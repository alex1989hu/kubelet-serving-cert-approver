// Copyright 2021 Alex Szakaly
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package certificatesigningrequest

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/zap"
	certificatesv1 "k8s.io/api/certificates/v1"
)

//nolint:lll
var (
	errDNSOrIPMissing               = errors.New("either DNS Names or IP Addresses is missing")
	errExtraExtensionsPresent       = errors.New("emailAddress and URI subjectAltName extensions are forbidden")
	errKeyUsageMismatch             = errors.New("key usage does not match")
	errNotCertificateRequest        = errors.New("PEM Block Type must be CERTIFICATE REQUEST")
	errOrganizationMismatch         = errors.New("organization does not match")
	errX509CommonNameMismatch       = errors.New("x509 Common Name does not match with Certificate Signing Request Username")
	errX509CommonNamePrefixMismatch = errors.New("x509 Common Name does not start with 'system:node'")
)

// parseCSR decodes a PEM encoded Certificate Signing Request.
// https://github.com/kubernetes/kubernetes/blob/v1.20.1/pkg/apis/certificates/v1/helpers.go#L26
func parseCSR(pemBytes []byte) (*x509.CertificateRequest, error) {
	block, _ := pem.Decode(pemBytes)

	if block == nil || block.Type != "CERTIFICATE REQUEST" {
		return nil, errNotCertificateRequest
	}

	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error occurred during parsing of Certificate Signing Request: %w", err)
	}

	return csr, nil
}

// hasExactUsages check the permitted key usages - exactly ["key encipherment", "digital signature", "server auth"]
// for RSA and ["digital signature", "server auth"] for non-RSA certificates.
func hasExactUsages(log *zap.Logger, csr certificatesv1.CertificateSigningRequest) bool {
	permittedUsages := [3]certificatesv1.KeyUsage{
		certificatesv1.UsageDigitalSignature,
		certificatesv1.UsageServerAuth,
		// Optional since Kubernetes v1.27 https://github.com/kubernetes/kubernetes/pull/111660
		certificatesv1.UsageKeyEncipherment,
	}

	permittedUsagesMap := map[certificatesv1.KeyUsage]struct{}{}
	for _, u := range permittedUsages {
		permittedUsagesMap[u] = struct{}{}
	}

	for _, u := range csr.Spec.Usages {
		if _, ok := permittedUsagesMap[u]; !ok {
			log.Warn("Found disallowed certificate usage(s)", zap.String("usage", string(u)))

			return false
		}
	}

	return true
}

// isRequestConform returns error if the input does not conform with Kubernetes rules.
// Reference: https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
func isRequestConform(log *zap.Logger, csr certificatesv1.CertificateSigningRequest,
	x509cr *x509.CertificateRequest,
) error {
	expectedOrg := "system:nodes"
	expectedPrefix := "system:node:"

	if !reflect.DeepEqual([]string{expectedOrg}, x509cr.Subject.Organization) {
		log.Warn("X509 Organization does not match",
			zap.Strings("actual", x509cr.Subject.Organization),
			zap.String("expected", expectedOrg))

		return errOrganizationMismatch
	}

	if !strings.HasPrefix(x509cr.Subject.CommonName, expectedPrefix) {
		log.Warn("X509 Common Name does not start with expected prefix",
			zap.String("actual", x509cr.Subject.CommonName),
			zap.String("expected", expectedPrefix))

		return errX509CommonNamePrefixMismatch
	}

	if csr.Spec.Username != x509cr.Subject.CommonName {
		log.Warn("X509 Common Name does not match with Certificate Signing Request Username",
			zap.String("expected", csr.Spec.Username),
			zap.String("actual", x509cr.Subject.CommonName))

		return errX509CommonNameMismatch
	}

	if (len(x509cr.EmailAddresses) != 0) || (len(x509cr.URIs) != 0) {
		log.Warn("Forbidden EmailAddress and URI subjectAltName extensions are found")

		return errExtraExtensionsPresent
	}

	if (len(x509cr.DNSNames) < 1) && (len(x509cr.IPAddresses) < 1) {
		log.Warn("DNSNames or IP Addresses must be present")

		return errDNSOrIPMissing
	}

	if !hasExactUsages(log, csr) {
		log.Warn("Certificate Signing Request Usages do not match",
			zap.Any("usages", csr.Spec.Usages),
		)

		return errKeyUsageMismatch
	}

	return nil
}
