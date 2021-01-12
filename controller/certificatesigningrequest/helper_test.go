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

//nolint: testpackage // Need to reach functions.
package certificatesigningrequest

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"net"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	certificatesv1 "k8s.io/api/certificates/v1"
)

const (
	validOrganization = "system:nodes"
	validUsername     = "system:node:node-01"
)

//nolint: gochecknoglobals
var (
	validDNSNames   = []string{"foo.bar"}
	valdIPAddresses = []net.IP{net.ParseIP("1.2.3.4")}
	validUsages     = []certificatesv1.KeyUsage{
		certificatesv1.UsageKeyEncipherment,
		certificatesv1.UsageDigitalSignature,
		certificatesv1.UsageServerAuth,
	}
)

func TestParseCSR(t *testing.T) {
	t.Parallel()

	csr, err := parseCSR(nil)

	assert.NotNil(t, err)
	assert.Nil(t, csr)
	assert.True(t, errors.Is(err, errNotCertificateRequest))
}

func TestParseCSRMissingBlock(t *testing.T) {
	t.Parallel()

	pemCSR := []byte(`
-----BEGIN CERTIFICATE REQUEST-----
-----END CERTIFICATE REQUEST-----
`)
	csr, err := parseCSR(pemCSR)

	assert.NotNil(t, err)
	assert.Nil(t, csr)
	assert.Contains(t, err.Error(), "during parsing of Certificate Signing Request")
}

func TestParseCSRValidInput(t *testing.T) {
	t.Parallel()

	csr, err := parseCSR(generatePEMEncodedCSR(t))

	assert.Nil(t, err)
	assert.NotNil(t, csr)
}

func TestIsRequestConformInvalidSigningRequest(t *testing.T) {
	t.Parallel()

	tables := []struct {
		csr           certificatesv1.CertificateSigningRequest
		x509cr        x509.CertificateRequest
		expectedError error
	}{
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Usages:     validUsages,
					Username:   validUsername,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{"invalid"},
				},
				DNSNames:    validDNSNames,
				IPAddresses: valdIPAddresses,
			},
			expectedError: errOrganizationMismatch,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Usages:     validUsages,
					Username:   validUsername,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   "bad:prefix:node:node-01",
					Organization: []string{validOrganization},
				},
				DNSNames:    validDNSNames,
				IPAddresses: valdIPAddresses,
			},
			expectedError: errX509CommonNamePrefixMismatch,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Usages:     validUsages,
					Username:   validUsername,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername + "mismatch",
					Organization: []string{validOrganization},
				},
				DNSNames:    validDNSNames,
				IPAddresses: valdIPAddresses,
			},
			expectedError: errX509CommonNameMismatch,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username:   validUsername,
					Usages:     validUsages,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				EmailAddresses: []string{"foo@no-reply.bar"},
			},
			expectedError: errExtraExtensionsPresent,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username:   validUsername,
					Usages:     validUsages,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				URIs: []*url.URL{{Host: "foo.bar.acme"}},
			},
			expectedError: errExtraExtensionsPresent,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username:   validUsername,
					Usages:     validUsages,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				EmailAddresses: []string{"foo@no-reply.bar"},
				URIs:           []*url.URL{{Host: "foo.bar.acme"}},
			},
			expectedError: errExtraExtensionsPresent,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username:   validUsername,
					Usages:     validUsages,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
			},
			expectedError: errDNSOrIPMissing,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username: validUsername,
					Usages: []certificatesv1.KeyUsage{
						certificatesv1.UsageServerAuth,
						certificatesv1.UsageDigitalSignature,
						certificatesv1.UsageKeyEncipherment,
						certificatesv1.UsageCodeSigning,
					},
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				DNSNames:    validDNSNames,
				IPAddresses: valdIPAddresses,
			},
			expectedError: errKeyUsageMismatch,
		},
		{
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Username: validUsername,
					Usages: []certificatesv1.KeyUsage{
						certificatesv1.UsageKeyEncipherment,
						certificatesv1.UsageDigitalSignature,
						certificatesv1.UsageCodeSigning,
					},
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				DNSNames:    validDNSNames,
				IPAddresses: valdIPAddresses,
			},
			expectedError: errKeyUsageMismatch,
		},
	}

	for _, table := range tables { //nolint: paralleltest // Disable false-positive finding due to linter bug.
		table := table // scopelint, pin!

		t.Run(fmt.Sprint(table.expectedError), func(t *testing.T) {
			t.Parallel()
			assert.True(t, errors.Is(isRequestConform(table.csr, &table.x509cr), table.expectedError))
		})
	}
}

func TestConformantKubeletServingCertificateSigningRequest(t *testing.T) {
	t.Parallel()

	tables := []struct {
		goal   string
		csr    certificatesv1.CertificateSigningRequest
		x509cr x509.CertificateRequest
	}{
		{
			goal: "Only DNSNames present",
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Usages:     validUsages,
					Username:   validUsername,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				DNSNames: validDNSNames,
			},
		},
		{
			goal: "Only IPAddresses present",
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					Usages:     validUsages,
					Username:   validUsername,
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
			x509cr: x509.CertificateRequest{
				Subject: pkix.Name{
					CommonName:   validUsername,
					Organization: []string{validOrganization},
				},
				IPAddresses: valdIPAddresses,
			},
		},
	}

	for _, table := range tables { //nolint: paralleltest // Disable false-positive finding due to linter bug.
		table := table // scopelint, pin!

		t.Run(fmt.Sprint(table.goal), func(t *testing.T) {
			t.Parallel()
			assert.NoError(t, isRequestConform(table.csr, &table.x509cr))
		})
	}
}

// TestMain is needed due to t.Parallel() incompatibility of goleak.
// https://github.com/uber-go/goleak/issues/16
func TestMain(m *testing.M) { //nolint: interfacer
	// flushDaemon leaks: https://github.com/kubernetes/client-go/issues/900
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("k8s.io/klog/v2.(*loggingT).flushDaemon"))
}
