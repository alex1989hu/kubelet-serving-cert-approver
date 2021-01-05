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

package metrics

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

//nolint: gochecknoglobals
var (
	NumberOfApprovedCertificateRequests = prom.NewCounter(prom.CounterOpts{
		Name: "kubelet_serving_cert_approver_approved_certificate_signing_request_count",
		Help: "The number of approved Certificate Signing Request",
	})

	NumberOfInvalidCertificateSigningRequests = prom.NewCounter(prom.CounterOpts{
		Name: "kubelet_serving_cert_approver_invalid_certificate_signing_request_count",
		Help: "The number of invalid Certificate Signing Request",
	})
)
