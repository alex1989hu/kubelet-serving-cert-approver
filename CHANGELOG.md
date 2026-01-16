
<a name="v0.10.2"></a>
## [v0.10.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.10.1...v0.10.2) (2026-01-16)

### Chore

* upgrade golangci-lint v2.8.0
* upgrade go 1.25.6
* enable modernize golangci-lint linter
* bump docker/setup-buildx-action from 3.11.1 to 3.12.0

### Ci

* add kubernetes 1.35.0 e2e image
* upgrade e2e kind images

### Feat

* reduce memory usage


<a name="v0.10.1"></a>
## [v0.10.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.10.0...v0.10.1) (2025-12-06)

### Chore

* bump docker/metadata-action from 5.9.0 to 5.10.0
* upgrade cobra v1.10.2
* upgrade go 1.25.5


<a name="v0.10.0"></a>
## [v0.10.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.9.3...v0.10.0) (2025-11-22)

### Chore

* upgrade controller-runtime v0.22.4
* use same kubernetes version for worker nodes
* add namespace of the csr as a log field
* bump actions/checkout from 5 to 6
* bump sigstore/cosign-installer from 3.10.0 to 4.0.0
* bump github/codeql-action from 3 to 4
* bump docker/metadata-action from 5.8.0 to 5.9.0
* bump docker/setup-qemu-action from 3.6.0 to 3.7.0
* bump golangci/golangci-lint-action from 8 to 9
* upgrade golangci-lint v2.6.2
* upgrade golang.org/x/net v0.47.0
* upgrade go 1.25.4
* bump docker/login-action from 3.5.0 to 3.6.0
* bump anchore/scan-action from 6 to 7
* bump sigstore/cosign-installer from 3.9.2 to 3.10.0
* bump actions/setup-go from 5 to 6
* bump aquasecurity/trivy-action from 0.32.0 to 0.33.1
* bump actions/checkout from 4 to 5
* upgrade go 1.24.6
* bump docker/metadata-action from 5.7.0 to 5.8.0
* bump docker/login-action from 3.4.0 to 3.5.0
* bump sigstore/cosign-installer from 3.9.1 to 3.9.2
* upgrade go 1.24.5
* bump aquasecurity/trivy-action from 0.31.0 to 0.32.0
* bump sigstore/cosign-installer from 3.9.0 to 3.9.1
* bump docker/setup-buildx-action from 3.11.0 to 3.11.1
* bump sigstore/cosign-installer from 3.8.2 to 3.9.0
* bump docker/setup-buildx-action from 3.10.0 to 3.11.0

### Ci

* use lower codecov target due to atomic coverage change
* do not run nancy on pull requests
* authenticate against oss index
* remove obsolete GOEXPERIMENT


<a name="v0.9.3"></a>
## [v0.9.3](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.9.2...v0.9.3) (2025-06-15)

### Chore

* upgrade golang.org/x/net v0.41.0
* upgrade controller-runtime v0.21.0
* upgrade go 1.24.4
* bump aquasecurity/trivy-action from 0.30.0 to 0.31.0
* bump docker/build-push-action from 6.17.0 to 6.18.0

### Ci

* upgrade kind v0.29.0
* fix codecov-action file warning
* upgrade e2e kind images
* upgrade github runner to ubuntu-24.04


<a name="v0.9.2"></a>
## [v0.9.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.9.1...v0.9.2) (2025-05-17)

### Chore

* upgrade go 1.24.3
* add node.cloudprovider.kubernetes.io/uninitialized toleration
* bump docker/build-push-action from 6.16.0 to 6.17.0
* restore gci formatter configuration
* upgrade golangci-lint v2.1.6
* bump golangci/golangci-lint-action from 7 to 8
* bump docker/build-push-action from 6.15.0 to 6.16.0
* bump sigstore/cosign-installer from 3.8.1 to 3.8.2


<a name="v0.9.1"></a>
## [v0.9.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.9.0...v0.9.1) (2025-04-22)

### Chore

* upgrade golangci-lint v2.1.2
* upgrade golang.org/x/net v0.39.0
* upgrade go 1.24.2
* bump docker/login-action from 3.3.0 to 3.4.0
* bump aquasecurity/trivy-action from 0.29.0 to 0.30.0
* upgrade golangci-lint v1.64.6
* upgrade golang.org/x/net v0.37.0
* upgrade go 1.24.1
* bump docker/setup-qemu-action from 3.5.0 to 3.6.0
* bump docker/build-push-action from 6.13.0 to 6.15.0
* bump docker/setup-buildx-action from 3.8.0 to 3.10.0
* bump docker/setup-qemu-action from 3.4.0 to 3.5.0
* bump sigstore/cosign-installer from 3.7.0 to 3.8.1
* bump docker/metadata-action from 5.6.1 to 5.7.0
* upgrade go 1.23.6
* bump docker/setup-qemu-action from 3.3.0 to 3.4.0
* upgrade controller-runtime v0.20.1
* bump docker/build-push-action from 6.12.0 to 6.13.0

### Ci

* update nancy ignore
* update nancy ignore


<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.7...v0.9.0) (2025-01-19)

### Chore

* upgrade controller-runtime v0.20.0
* set go patch version in go.mod
* bump docker/build-push-action from 6.11.0 to 6.12.0
* upgrade go 1.23.5
* bump docker/setup-qemu-action from 3.2.0 to 3.3.0
* bump docker/build-push-action from 6.10.0 to 6.11.0
* upgrade golangci-lint v1.63.4
* upgrade golangci-lint v1.63.3
* upgrade golangci-lint v1.63.0
* upgrade golang.org/x/net v0.33.0
* bump docker/setup-buildx-action from 3.7.1 to 3.8.0
* bump anchore/scan-action from 5 to 6
* upgrade golangci-lint v1.62.2
* upgrade go 1.23.4
* upgrade controller-runtime v0.19.3
* upgrade stretchr/testify v1.10.0
* bump docker/build-push-action from 6.9.0 to 6.10.0
* upgrade controller-runtime v0.19.2
* bump codecov/codecov-action from 4 to 5
* bump aquasecurity/trivy-action from 0.28.0 to 0.29.0
* bump docker/metadata-action from 5.5.1 to 5.6.1
* bump engineerd/setup-kind from 0.5.0 to 0.6.2
* upgrade go 1.23.2
* upgrade controller-runtime v0.19.1
* upgrade godog v0.15.0
* bump aquasecurity/trivy-action from 0.27.0 to 0.28.0
* bump anchore/scan-action from 4 to 5

### Ci

* upgrade kind v0.26.0
* update nancy ignore
* replace invalid comment opiton of codecov
* disable kind cluster log export
* upgrade kind v0.24.0


<a name="v0.8.7"></a>
## [v0.8.7](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.6...v0.8.7) (2024-10-13)

### Chore

* bump aquasecurity/trivy-action from 0.26.0 to 0.27.0
* bump aquasecurity/trivy-action from 0.25.0 to 0.26.0
* bump aquasecurity/trivy-action from 0.24.0 to 0.25.0
* as keyword should match the case of the from keyword in Dockerfile
* bump sigstore/cosign-installer from 3.6.0 to 3.7.0
* bump docker/setup-buildx-action from 3.7.0 to 3.7.1
* bump docker/setup-buildx-action from 3.6.1 to 3.7.0
* bump docker/build-push-action from 6.7.0 to 6.9.0

### Fix

* set logger to controller-runtime


<a name="v0.8.6"></a>
## [v0.8.6](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.5...v0.8.6) (2024-09-17)

### Chore

* upgrade godog v0.14.1
* upgrade viper v1.19.0
* upgrade zap v1.27.0
* upgrade golangci-lint v1.61.0
* upgrade controller-runtime v0.19.0
* upgrade go 1.23.1
* bump docker/build-push-action from 6.6.1 to 6.7.0
* bump docker/build-push-action from 6.5.0 to 6.6.1
* bump sigstore/cosign-installer from 3.5.0 to 3.6.0
* bump docker/setup-buildx-action from 3.5.0 to 3.6.1
* bump docker/login-action from 3.2.0 to 3.3.0
* bump docker/build-push-action from 6.4.1 to 6.5.0
* bump docker/setup-buildx-action from 3.4.0 to 3.5.0
* bump docker/setup-qemu-action from 3.1.0 to 3.2.0
* bump docker/build-push-action from 6.4.0 to 6.4.1
* bump anchore/scan-action from 3 to 4
* bump docker/build-push-action from 6.3.0 to 6.4.0
* bump aquasecurity/trivy-action from 0.23.0 to 0.24.0
* bump docker/setup-buildx-action from 3.3.0 to 3.4.0

### Ci

* update nancy ignore


<a name="v0.8.5"></a>
## [v0.8.5](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.4...v0.8.5) (2024-07-03)

### Chore

* bump docker/build-push-action from 6.1.0 to 6.3.0
* bump docker/setup-qemu-action from 3.0.0 to 3.1.0
* upgrade go 1.22.5
* bump docker/build-push-action from 6.0.2 to 6.1.0
* bump docker/build-push-action from 6.0.1 to 6.0.2
* bump docker/build-push-action from 5.4.0 to 6.0.1
* bump aquasecurity/trivy-action from 0.22.0 to 0.23.0
* bump docker/build-push-action from 5.3.0 to 5.4.0
* bump aquasecurity/trivy-action from 0.21.0 to 0.22.0


<a name="v0.8.4"></a>
## [v0.8.4](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.3...v0.8.4) (2024-06-05)

### Chore

* upgrade go 1.22.4
* upgrade kubernetes v1.30.0
* bump docker/login-action from 3.1.0 to 3.2.0
* bump aquasecurity/trivy-action from 0.20.0 to 0.21.0

### Ci

* use digest of kind images


<a name="v0.8.3"></a>
## [v0.8.3](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.2...v0.8.3) (2024-05-08)

### Chore

* enable testifylint golangci-lint linter
* bump golangci/golangci-lint-action from 5 to 6
* bump aquasecurity/trivy-action from 0.19.0 to 0.20.0
* upgrade golangci-lint v1.58.0
* upgrade go 1.22.3
* bump golangci/golangci-lint-action from 4 to 5
* bump sigstore/cosign-installer from 3.4.0 to 3.5.0
* upgrade controller-runtime v0.17.3
* bump docker/setup-buildx-action from 3.2.0 to 3.3.0

### Ci

* use verbose flag of govulncheck


<a name="v0.8.2"></a>
## [v0.8.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.1...v0.8.2) (2024-04-05)

### Chore

* upgrade go 1.22.2
* upgrade golang.org/x/net v0.24.0
* bump aquasecurity/trivy-action from 0.18.0 to 0.19.0
* bump docker/build-push-action from 5.2.0 to 5.3.0
* bump docker/setup-buildx-action from 3.1.0 to 3.2.0
* upgrade google.golang.org/protobuf v1.33.0
* bump docker/build-push-action from 5.1.0 to 5.2.0
* bump docker/login-action from 3.0.0 to 3.1.0
* upgrade golang.org/x/net v0.17.0
* upgrade go 1.21.8

### Ci

* upgrade azure/setup-kubectl to 4
* update nancy ignore


<a name="v0.8.1"></a>
## [v0.8.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.8.0...v0.8.1) (2024-03-04)

### Chore

* upgrade golangci-lint v1.56.2
* use meaningful log field names
* bump aquasecurity/trivy-action from 0.17.0 to 0.18.0
* bump docker/setup-buildx-action from 3.0.0 to 3.1.0
* upgrade controller-runtime v0.17.2
* upgrade controller-runtime v0.17.1
* upgrade golangci-lint v1.56.1
* upgrade go 1.21.7
* bump golangci/golangci-lint-action from 3 to 4
* bump aquasecurity/trivy-action from 0.16.1 to 0.17.0
* upgrade godog v0.14.0
* upgrade controller-runtime v0.17.0
* bump codecov/codecov-action from 3 to 4
* upgrade go 1.21.6
* bump sigstore/cosign-installer from 3.3.0 to 3.4.0
* bump docker/metadata-action from 5.5.0 to 5.5.1

### Ci

* get digest from docker_build step
* specify CODECOV_TOKEN due to codecov-action[@4](https://github.com/4) change
* call cosign to sign multiple tags

### Fix

* set v1.29.0 image for 2nd kind worker node as well


<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.7...v0.8.0) (2024-01-05)

### Chore

* upgrade prometheus/client_golang v1.18.0
* bump docker/metadata-action from 5.4.0 to 5.5.0
* bump aquasecurity/trivy-action from 0.16.0 to 0.16.1
* upgrade viper v1.18.2
* upgrade kubernetes v1.29.0
* bump github/codeql-action from 2 to 3
* bump docker/metadata-action from 5.3.0 to 5.4.0
* bump aquasecurity/trivy-action from 0.15.0 to 0.16.0
* bump sigstore/cosign-installer from 3.2.0 to 3.3.0
* upgrade viper v1.18.1

### Ci

* update nancy ignore
* add setup-go to codeql scan


<a name="v0.7.7"></a>
## [v0.7.7](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.6...v0.7.7) (2023-12-06)

### Chore

* bump actions/setup-go from 4 to 5
* upgrade viper v1.18.0
* upgrade go 1.21.5
* bump aquasecurity/trivy-action from 0.14.0 to 0.15.0
* bump docker/metadata-action from 5.2.0 to 5.3.0
* bump docker/metadata-action from 5.1.0 to 5.2.0
* bump docker/metadata-action from 5.0.0 to 5.1.0
* bump docker/build-push-action from 5.0.0 to 5.1.0
* upgrade kubernetes v1.28.4

### Ci

* update nancy ignore

### Docs

* remove trailing dot


<a name="v0.7.6"></a>
## [v0.7.6](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.5...v0.7.6) (2023-11-07)

### Chore

* bump aquasecurity/trivy-action from 0.13.1 to 0.14.0
* upgrade go 1.21.4
* bump sigstore/cosign-installer from 3.1.2 to 3.2.0
* bump aquasecurity/trivy-action from 0.13.0 to 0.13.1
* upgrade golangci-lint v1.55.2
* upgrade cobra v1.8.0
* upgrade goleak v1.3.0
* upgrade golangci-lint v1.55.1
* bump aquasecurity/trivy-action from 0.12.0 to 0.13.0
* upgrade controller-runtime v0.16.3

### Ci

* update nancy ignore


<a name="v0.7.5"></a>
## [v0.7.5](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.4...v0.7.5) (2023-10-12)

### Chore

* upgrade golang.org/x/net v0.17.0
* upgrade go 1.21.3
* remove mkdir step from dockerfile
* upgrade kubernetes v1.28.2
* upgrade viper v1.17.0
* upgrade go 1.21.2

### Docs

* fix tls bootstrapping url in reference section


<a name="v0.7.4"></a>
## [v0.7.4](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.3...v0.7.4) (2023-09-20)

### Chore

* upgrade controller-runtime v0.16.2
* bump codecov/codecov-action from 3 to 4
* upgrade zap v1.26.0
* bump docker/setup-qemu-action from 2.2.0 to 3.0.0
* bump docker/setup-buildx-action from 2.10.0 to 3.0.0
* bump docker/login-action from 2.2.0 to 3.0.0
* bump docker/metadata-action from 4.6.0 to 5.0.0
* bump docker/build-push-action from 4.2.1 to 5.0.0
* upgrade go 1.21.1
* bump docker/build-push-action from 4.1.1 to 4.2.1
* bump sigstore/cosign-installer from 3.1.1 to 3.1.2
* bump actions/checkout from 3 to 4
* bump aquasecurity/trivy-action from 0.11.2 to 0.12.0
* upgrade kubernetes v1.28.1

### Reverts

* chore: bump codecov/codecov-action from 3 to 4


<a name="v0.7.3"></a>
## [v0.7.3](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.2...v0.7.3) (2023-09-03)

### Chore

* upgrade to latest distroless image
* remove toolchain constraint due to codeql limitation
* upgrade controller-runtime v0.15.2
* upgrade stretchr/testify v1.8.4
* upgrade godog v0.13.0
* upgrade go 1.21.0
* bump docker/setup-buildx-action from 2.9.1 to 2.10.0
* upgrade golang.org/x/net v0.14.0
* upgrade controller-runtime v0.15.1

### Ci

* upgrade kind v0.20.0
* update nancy ignore


<a name="v0.7.2"></a>
## [v0.7.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.1...v0.7.2) (2023-08-02)

### Chore

* upgrade go 1.20.7
* upgrade zap v1.25.0
* bump docker/setup-buildx-action from 2.9.0 to 2.9.1
* upgrade go v1.20.6
* bump docker/setup-buildx-action from 2.8.0 to 2.9.0
* bump docker/setup-buildx-action from 2.7.0 to 2.8.0
* bump sigstore/cosign-installer from 3.1.0 to 3.1.1
* bump docker/build-push-action from 4.1.0 to 4.1.1
* bump docker/setup-buildx-action from 2.6.0 to 2.7.0
* bump sigstore/cosign-installer from 3.0.5 to 3.1.0
* bump docker/metadata-action from 4.5.0 to 4.6.0
* bump aquasecurity/trivy-action from 0.11.0 to 0.11.2
* bump docker/login-action from 2.1.0 to 2.2.0
* bump docker/build-push-action from 4.0.0 to 4.1.0
* bump docker/setup-qemu-action from 2.1.0 to 2.2.0
* upgrade go 1.20.5
* bump docker/metadata-action from 4.4.0 to 4.5.0
* bump docker/setup-buildx-action from 2.5.0 to 2.6.0
* bump aquasecurity/trivy-action from 0.10.0 to 0.11.0

### Ci

* update govulncheck call


<a name="v0.7.1"></a>
## [v0.7.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.7.0...v0.7.1) (2023-06-04)

### Chore

* upgrade cobra v1.7.0
* upgrade viper v1.16.0
* upgrade golangci-lint v1.53.2
* disable caching of CertificateSigningRequest
* enable pprof in debug mode


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.10...v0.7.0) (2023-05-28)

### Chore

* upgrade kubernetes v1.27.2
* bump sigstore/cosign-installer from 3.0.4 to 3.0.5
* bump sigstore/cosign-installer from 3.0.3 to 3.0.4
* upgrade go 1.20.4
* bump sigstore/cosign-installer from 3.0.2 to 3.0.3
* bump aquasecurity/trivy-action from 0.9.2 to 0.10.0
* bump docker/metadata-action from 4.3.0 to 4.4.0


<a name="v0.6.10"></a>
## [v0.6.10](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.9...v0.6.10) (2023-04-10)

### Chore

* bump sigstore/cosign-installer from 3.0.1 to 3.0.2
* upgrade golangci-lint v1.52.2
* upgrade go 1.20.3
* upgrade kubernetes v1.26.3
* upgrade controller-runtime v0.14.6

### Ci

* upgrade kind v0.18.0
* update nancy ignore


<a name="v0.6.9"></a>
## [v0.6.9](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.8...v0.6.9) (2023-03-18)

### Ci

* add yes argument to cosign call


<a name="v0.6.8"></a>
## [v0.6.8](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.7...v0.6.8) (2023-03-18)

### Chore

* bump sigstore/cosign-installer from 2.8.1 to 3.0.1
* bump actions/setup-go from 3 to 4
* bump docker/setup-buildx-action from 2.4.1 to 2.5.0
* upgrade go 1.20.2
* bump aquasecurity/trivy-action from 0.9.1 to 0.9.2
* bump sonatype-nexus-community/nancy-github-action

### Ci

* set GitHub token for nancy-github-action


<a name="v0.6.7"></a>
## [v0.6.7](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.6...v0.6.7) (2023-02-18)

### Chore

* upgrade golang.org/x/net v0.7.0
* upgrade go v1.20.1
* set project to go1.20
* upgrade viper v1.15.0
* upgrade go v1.20.0
* bump aquasecurity/trivy-action from 0.9.0 to 0.9.1
* upgrade golangci-lint v1.51.1
* upgrade controller-runtime v0.14.4
* bump docker/setup-buildx-action from 2.4.0 to 2.4.1
* upgrade golangci-lint v1.51.0
* upgrade controller-runtime v0.14.3
* bump aquasecurity/trivy-action from 0.8.0 to 0.9.0
* upgrade controller-runtime v0.14.2
* bump docker/setup-buildx-action from 2.2.1 to 2.4.0
* bump docker/build-push-action from 3.2.0 to 3.3.0
* bump docker/metadata-action from 4.2.0 to 4.3.0
* bump docker/metadata-action from 4.1.1 to 4.2.0
* upgrade godog v0.12.6
* upgrade go v1.19.5
* bump azure/setup-kubectl from 3.1 to 3.2

### Ci

* add buildcsv buildflag for snyk job
* upgra build-push-action v4.0.0
* add grype ignore
* do not use fail-build with grype
* add grype


<a name="v0.6.6"></a>
## [v0.6.6](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.5...v0.6.6) (2022-12-20)

### Chore

* upgrade controller-runtime v0.14.1


<a name="v0.6.5"></a>
## [v0.6.5](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.4...v0.6.5) (2022-12-18)

### Chore

* godocs typo in new mocked function
* drop old build tag format in e2e tests
* upgrade kubernetes v1.26.0

### Ci

* update nancy ignore


<a name="v0.6.4"></a>
## [v0.6.4](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.3...v0.6.4) (2022-12-10)

### Chore

* upgrade golang.org/x/net
* upgrade go v1.19.4
* bump azure/setup-kubectl from 3.0 to 3.1

### Ci

* upgrade github runner to ubuntu-22.04


<a name="v0.6.3"></a>
## [v0.6.3](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.2...v0.6.3) (2022-11-06)

### Chore

* bump aquasecurity/trivy-action from 0.7.1 to 0.8.0
* upgrade prometheus/client_golang v1.13.1
* upgrade controller-runtime v0.13.1
* upgrade viper v1.14.0
* upgrade cobra v1.6.1
* upgrade go v1.19.3
* upgrade golangci-lint v1.50.1
* bump docker/setup-buildx-action from 2.2.0 to 2.2.1
* bump docker/metadata-action from 4.1.0 to 4.1.1
* bump sigstore/cosign-installer from 2.8.0 to 2.8.1
* bump docker/setup-buildx-action from 2.1.0 to 2.2.0
* upgrade golang.org/x/text
* upgrade kubernetes v1.25.3
* upgrade cobra v1.6.0
* bump docker/setup-buildx-action from 2.0.0 to 2.1.0
* bump docker/login-action from 2.0.0 to 2.1.0
* bump docker/metadata-action from 4.0.1 to 4.1.0
* bump docker/build-push-action from 3.1.1 to 3.2.0
* bump docker/setup-qemu-action from 2.0.0 to 2.1.0
* bump sigstore/cosign-installer from 2.7.0 to 2.8.0
* upgrade to latest distroless image
* upgrade golangci-lint v1.50.0
* upgrade go v1.19.2

### Ci

* update nancy ignore
* limit image-publish token scope
* limit build-and-test token scope
* upgrade kind v0.17.0
* limit e2e token scope
* limit snyk token scope
* limit govulncheck token scope
* limit codeql token scope
* limit trivy token scope
* limit nancy token scope
* limit golangci-lint token scope
* enable cache for setup-go
* update nancy ignore


<a name="v0.6.2"></a>
## [v0.6.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.1...v0.6.2) (2022-09-28)

### Chore

* upgrade golang.org/x/net/http2
* upgrade goleak v1.2.0
* upgrade kubernetes v1.25.2
* upgrade kind v0.16.0
* bump sigstore/cosign-installer from 2.6.0 to 2.7.0
* bump sigstore/cosign-installer from 2.5.1 to 2.6.0

### Ci

* add govulncheck


<a name="v0.6.1"></a>
## [v0.6.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.6.0...v0.6.1) (2022-09-08)

### Chore

* upgrade go v1.19.1


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.5.1...v0.6.0) (2022-09-06)

### Chore

* upgrade golangci-lint v1.49.0
* upgrade viper v1.13.0
* upgrade zap v1.23.0
* upgrade e2e kubernetes versions to recent patch level
* upgrade kubernetes v1.25.0
* bump aquasecurity/trivy-action from 0.6.2 to 0.7.1
* bump sigstore/cosign-installer from 2.5.0 to 2.5.1
* upgrade prometheus/client_golang v1.13.0
* upgrade kubernetes dependencies v1.24.3
* upgrade zap v1.22.0
* upgrade go v1.19.0
* bump docker/build-push-action from 3.1.0 to 3.1.1
* bump aquasecurity/trivy-action from 0.6.1 to 0.6.2
* upgrade go v1.18.5
* upgrade golangci-lint v1.47.3
* bump sigstore/cosign-installer from 2.4.1 to 2.5.0
* bump aquasecurity/trivy-action from 0.6.0 to 0.6.1
* bump aquasecurity/trivy-action from 0.5.1 to 0.6.0
* bump docker/build-push-action from 3.0.0 to 3.1.0
* upgrade go v1.18.4
* upgrade stretchr/testify v1.8.0
* upgrade controller-runtime v0.12.3
* bump sigstore/cosign-installer from 2.4.0 to 2.4.1
* bump aquasecurity/trivy-action from 0.5.0 to 0.5.1
* upgrade e2e kubernetes version to v1.24.2
* upgrade cobra v1.5.0
* upgrade stretchr/testify v1.7.5
* upgrade controller-runtime v0.12.2
* bump aquasecurity/trivy-action from 0.4.0 to 0.5.0
* bump azure/setup-kubectl from 2.1 to 3.0
* upgrade kubernetes dependencies v1.24.2
* upgrade stretchr/testify v1.7.4
* bump aquasecurity/trivy-action from 0.3.0 to 0.4.0
* upgrade golangci-lint v1.46.2
* upgrade kind v0.14.0
* upgrade kubernetes dependencies v1.24.1
* upgrade go v1.18.3
* bump sigstore/cosign-installer from 2.3.0 to 2.4.0
* upgrade viper v1.12.0
* upgrade controller-runtime v0.12.1

### Ci

* remove no longer supported kubernetes versions
* upgrade kubernetes image to v1.24.2
* update nancy ignore

### Fix

* logger debug flag not being set


<a name="v0.5.1"></a>
## [v0.5.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.5.0...v0.5.1) (2022-06-24)

### Chore

* upgrade e2e kubernetes version to v1.24.2
* upgrade cobra v1.5.0
* upgrade stretchr/testify v1.7.5
* upgrade controller-runtime v0.12.2
* bump aquasecurity/trivy-action from 0.4.0 to 0.5.0
* bump azure/setup-kubectl from 2.1 to 3.0
* upgrade kubernetes dependencies v1.24.2
* upgrade stretchr/testify v1.7.4
* bump aquasecurity/trivy-action from 0.3.0 to 0.4.0
* upgrade golangci-lint v1.46.2
* upgrade kind v0.14.0
* upgrade kubernetes dependencies v1.24.1
* upgrade go v1.18.3
* bump sigstore/cosign-installer from 2.3.0 to 2.4.0
* upgrade viper v1.12.0
* upgrade controller-runtime v0.12.1

### Ci

* upgrade kubernetes image to v1.24.2
* update nancy ignore

### Fix

* logger debug flag not being set


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.4.4...v0.5.0) (2022-05-16)

### Chore

* rerun changelog generator
* remove permission bit from changelog generator
* upgrade controller-runtime v0.12.0
* upgrade golangci-lint v1.46.1
* upgrade golangci-lint v1.46.0
* upgrade go v1.18.2
* bump aquasecurity/trivy-action from 0.2.5 to 0.3.0
* bump docker/build-push-action from 2.10.0 to 3.0.0
* bump docker/setup-qemu-action from 1.2.0 to 2.0.0
* bump docker/metadata-action from 3.8.0 to 4.0.1
* bump docker/setup-buildx-action from 1.7.0 to 2.0.0
* bump docker/login-action from 1.14.1 to 2.0.0
* bump docker/metadata-action from 3.7.0 to 3.8.0
* bump docker/setup-buildx-action from 1.6.0 to 1.7.0
* bump sigstore/cosign-installer from 2.2.1 to 2.3.0
* upgrade kubernetes dependencies v1.23.6
* upgrade go v1.18.1
* upgrade viper v1.11.0
* bump github/codeql-action from 1 to 2
* bump aquasecurity/trivy-action from 0.2.4 to 0.2.5
* bump aquasecurity/trivy-action from 0.2.3 to 0.2.4
* bump sigstore/cosign-installer from 2.2.0 to 2.2.1
* bump aquasecurity/trivy-action from 0.2.2 to 0.2.3
* upgrade controller-runtime v0.11.2
* bump sigstore/cosign-installer from 2.1.0 to 2.2.0
* bump actions/setup-go from 2 to 3
* bump docker/metadata-action from 3.6.2 to 3.7.0
* bump codecov/codecov-action from 2.1.0 to 3
* upgrade godog v0.12.5
* upgrade golangci-lint v1.45.2
* upgrade kubernetes dependencies v1.23.5
* upgrade go v1.18.0
* bump docker/build-push-action from 2.9.0 to 2.10.0
* upgrade go v1.17.8
* upgrade kind v0.12.0
* upgrade kubernetes dependencies v1.23.4
* upgrade cobra v1.4.0
* bump azure/setup-kubectl from 2.0 to 2.1
* bump docker/login-action from 1.14.0 to 1.14.1
* bump actions/checkout from 2.4.0 to 3
* bump docker/login-action from 1.13.0 to 1.14.0
* upgrade prometheus/client_golang v1.11.1
* upgrade golangci-lint v1.44.1
* upgrade go v1.17.7
* bump docker/login-action from 1.12.0 to 1.13.0
* upgrade controller-runtime v0.11.1
* upgrade zap v1.21.0
* bump aquasecurity/trivy-action from 0.2.1 to 0.2.2
* upgrade golangci-lint v1.44.0
* upgrade kubernetes dependencies v1.23.3
* upgrade go v1.17.6
* bump docker/build-push-action from 2.8.0 to 2.9.0
* bump docker/build-push-action from 2.7.0 to 2.8.0
* upgrade e2e kubernetes version to v1.23.1
* upgrade godog v0.12.3
* upgrade zap v1.20.0
* bump azure/setup-kubectl from 1 to 2.0
* bump aquasecurity/trivy-action from 0.2.0 to 0.2.1
* bump docker/login-action from 1.10.0 to 1.12.0
* upgrade kubernetes dependencies v1.23.1
* upgrade go v1.17.5
* upgrade viper v1.10.1

### Ci

* switch trivy report sarif option
* upgrade kubernetes image to v1.23.5
* sign container images with cosign
* upgrade golangci-lint-action[@v3](https://github.com/v3)


<a name="v0.4.4"></a>
## [v0.4.4](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.4.3...v0.4.4) (2022-05-08)

### Chore

* bump docker/build-push-action from 2.10.0 to 3.0.0
* bump docker/setup-qemu-action from 1.2.0 to 2.0.0
* bump docker/metadata-action from 3.8.0 to 4.0.1
* bump docker/setup-buildx-action from 1.7.0 to 2.0.0
* bump docker/login-action from 1.14.1 to 2.0.0
* bump docker/metadata-action from 3.7.0 to 3.8.0
* bump docker/setup-buildx-action from 1.6.0 to 1.7.0
* bump sigstore/cosign-installer from 2.2.1 to 2.3.0
* upgrade kubernetes dependencies v1.23.6
* upgrade go v1.18.1
* upgrade viper v1.11.0
* bump github/codeql-action from 1 to 2
* bump aquasecurity/trivy-action from 0.2.4 to 0.2.5
* bump aquasecurity/trivy-action from 0.2.3 to 0.2.4
* bump sigstore/cosign-installer from 2.2.0 to 2.2.1
* bump aquasecurity/trivy-action from 0.2.2 to 0.2.3
* upgrade controller-runtime v0.11.2
* bump sigstore/cosign-installer from 2.1.0 to 2.2.0
* bump actions/setup-go from 2 to 3
* bump docker/metadata-action from 3.6.2 to 3.7.0
* bump codecov/codecov-action from 2.1.0 to 3
* upgrade godog v0.12.5
* upgrade golangci-lint v1.45.2
* upgrade kubernetes dependencies v1.23.5
* upgrade go v1.18.0
* bump docker/build-push-action from 2.9.0 to 2.10.0
* upgrade go v1.17.8
* upgrade kind v0.12.0
* upgrade kubernetes dependencies v1.23.4
* upgrade cobra v1.4.0
* bump azure/setup-kubectl from 2.0 to 2.1
* bump docker/login-action from 1.14.0 to 1.14.1
* bump actions/checkout from 2.4.0 to 3
* bump docker/login-action from 1.13.0 to 1.14.0

### Ci

* switch trivy report sarif option
* upgrade kubernetes image to v1.23.5
* sign container images with cosign
* upgrade golangci-lint-action[@v3](https://github.com/v3)


<a name="v0.4.3"></a>
## [v0.4.3](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.4.2...v0.4.3) (2022-02-17)

### Chore

* upgrade prometheus/client_golang v1.11.1
* upgrade golangci-lint v1.44.1
* upgrade go v1.17.7
* bump docker/login-action from 1.12.0 to 1.13.0
* upgrade controller-runtime v0.11.1
* upgrade zap v1.21.0


<a name="v0.4.2"></a>
## [v0.4.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.4.1...v0.4.2) (2022-02-03)

### Chore

* bump aquasecurity/trivy-action from 0.2.1 to 0.2.2
* upgrade golangci-lint v1.44.0
* upgrade kubernetes dependencies v1.23.3
* upgrade go v1.17.6
* bump docker/build-push-action from 2.8.0 to 2.9.0
* bump docker/build-push-action from 2.7.0 to 2.8.0
* upgrade e2e kubernetes version to v1.23.1
* upgrade godog v0.12.3
* upgrade zap v1.20.0
* bump azure/setup-kubectl from 1 to 2.0
* bump aquasecurity/trivy-action from 0.2.0 to 0.2.1
* bump docker/login-action from 1.10.0 to 1.12.0
* upgrade kubernetes dependencies v1.23.1
* upgrade go v1.17.5


<a name="v0.4.1"></a>
## [v0.4.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.4.0...v0.4.1) (2021-12-15)

### Chore

* upgrade viper v1.10.1


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.3.0...v0.4.0) (2021-12-15)

### Chore

* upgrade cobra v1.3.0
* upgrade controller-runtime v0.11.0
* upgrade go v1.17.4
* explicitly set user group in Dockerfile
* bump docker/metadata-action from 3.6.1 to 3.6.2
* upgrade kubernetes dependencies v1.22.4
* bump docker/metadata-action from 3.6.0 to 3.6.1
* bump aquasecurity/trivy-action from 0.1.0 to 0.2.0
* upgrade golangci-lint v1.43.0
* bump aquasecurity/trivy-action from 0.0.22 to 0.1.0
* upgrade kubernetes dependencies v1.22.3
* upgrade go v1.17.3
* upgrade controller-runtime v0.10.3
* bump actions/checkout from 2.3.5 to 2.4.0
* bump aquasecurity/trivy-action from 0.0.20 to 0.0.22
* bump docker/metadata-action from 3.5.0 to 3.6.0

### Ci

* add nancy scan action
* add snyk scan action

### Docs

* add kubernetes 1.23 to compatibility matrix


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.2.2...v0.3.0) (2021-10-18)

### Chore

* upgrade godog v0.12.2
* bump actions/checkout from 2.3.4 to 2.3.5
* upgrade controller-runtime v0.10.2
* upgrade go v1.17.2
* upgrade to latest distroless image
* upgrade viper v1.9.0
* upgrade controller-runtime v0.10.1
* upgrade golangci-lint v1.42.1
* upgrade go v1.17.1
* upgrade kubernetes dependencies v1.22.2
* upgrade goleak v1.1.11
* upgrade zap v1.19.1
* upgrade controller-runtime v0.10.0
* bump codecov/codecov-action from 2.0.3 to 2.1.0
* bump docker/setup-buildx-action from 1.5.1 to 1.6.0
* bump codecov/codecov-action from 2.0.2 to 2.0.3
* bump docker/metadata-action from 3.4.1 to 3.5.0
* bump docker/build-push-action from 2.6.1 to 2.7.0
* upgrade go v1.17.0
* upgrade godog v0.12.0
* bump aquasecurity/trivy-action from 0.0.19 to 0.0.20
* upgrade controller-runtime v0.9.6
* upgrade golangci-lint v1.42.0
* upgrade go v1.16.7
* increase memory request resource
* upgrade zap v1.19.0
* upgrade controller-runtime v0.9.5
* bump aquasecurity/trivy-action from 0.0.18 to 0.0.19
* bump codecov/codecov-action from 2.0.1 to 2.0.2
* upgrade go v1.16.6
* upgrade kubernetes dependencies v1.21.3
* upgrade controller-runtime v0.9.3
* bump codecov/codecov-action from 1 to 2.0.1
* bump docker/metadata-action from 3.4.0 to 3.4.1
* bump docker/setup-buildx-action from 1.5.0 to 1.5.1
* bump docker/metadata-action from 3.3.0 to 3.4.0
* upgrade cobra v1.2.1
* bump docker/setup-buildx-action from 1.4.1 to 1.5.0
* upgrade cobra v1.2.0
* bump docker/build-push-action from 2.5.0 to 2.6.1
* enhance golangci-lint revive configuration
* bump docker/setup-buildx-action from 1.3.0 to 1.4.1
* upgrade golangci-lint v1.41.1
* upgrade zap v1.18.1
* upgrade viper v1.8.1
* upgrade viper v1.8.0
* upgrade cobra v1.1.3
* upgrade kubernetes dependencies v1.21.2
* bump docker/login-action from 1.9.0 to 1.10.0
* add security hardening related e2e tests
* add health check e2e tests

### Ci

* use correct kubectl version
* add wait script for not yet created k8s resource


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.2.1...v0.2.2) (2021-08-03)

### Chore

* upgrade controller-runtime v0.9.5
* bump aquasecurity/trivy-action from 0.0.18 to 0.0.19
* bump codecov/codecov-action from 2.0.1 to 2.0.2


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.2.0...v0.2.1) (2021-07-23)

### Chore

* upgrade go v1.16.6
* upgrade kubernetes dependencies v1.21.3
* upgrade controller-runtime v0.9.3
* bump codecov/codecov-action from 1 to 2.0.1
* bump docker/metadata-action from 3.4.0 to 3.4.1
* bump docker/setup-buildx-action from 1.5.0 to 1.5.1
* bump docker/metadata-action from 3.3.0 to 3.4.0
* upgrade cobra v1.2.1
* bump docker/setup-buildx-action from 1.4.1 to 1.5.0
* upgrade cobra v1.2.0
* bump docker/build-push-action from 2.5.0 to 2.6.1
* enhance golangci-lint revive configuration
* bump docker/setup-buildx-action from 1.3.0 to 1.4.1
* upgrade golangci-lint v1.41.1
* upgrade zap v1.18.1
* upgrade viper v1.8.1
* upgrade viper v1.8.0
* upgrade cobra v1.1.3
* upgrade kubernetes dependencies v1.21.2
* bump docker/login-action from 1.9.0 to 1.10.0
* add security hardening related e2e tests
* add health check e2e tests

### Ci

* add wait script for not yet created k8s resource


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/v0.1.1...v0.2.0) (2021-06-10)

### Chore

* upgrade go v1.16.5
* upgrade kubernetes dependencies v1.21.1
* optimize debug logging with zap checkedentry
* upgrade golangci-lint v1.40.1
* upgrade zap v1.17.0
* upgrade kind v0.11.1
* generate manifests with v1.20.7 kubectl
* bump aquasecurity/trivy-action from 0.0.17 to 0.0.18
* bump docker/setup-qemu-action from 1.1.0 to 1.2.0
* bump docker/build-push-action from 2.4.0 to 2.5.0
* bump docker/metadata-action from 3.1.0 to 3.3.0

### Ci

* use v2 input for metadata-action to eliminate warning
* extend e2e build matrix with ha and standalone install


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/0.1.0...v0.1.1) (2021-05-30)

### Chore

* optimize debug logging with zap checkedentry
* upgrade golangci-lint v1.40.1
* upgrade zap v1.17.0
* upgrade kind v0.11.1
* generate manifests with v1.20.7 kubectl
* bump aquasecurity/trivy-action from 0.0.17 to 0.0.18
* bump docker/setup-qemu-action from 1.1.0 to 1.2.0
* bump docker/build-push-action from 2.4.0 to 2.5.0
* bump docker/metadata-action from 3.1.0 to 3.3.0

### Ci

* extend e2e build matrix with ha and standalone install


<a name="0.1.0"></a>
## 0.1.0 (2021-05-19)

### Chore

* add chglog configuration and template file
* upgrade kubernetes dependencies v1.20.7
* bump aquasecurity/trivy-action from 0.0.15 to 0.0.17
* bump docker/build-push-action from 2 to 2.4.0
* bump docker/setup-buildx-action from 1 to 1.3.0
* bump docker/setup-qemu-action from 1 to 1.1.0
* bump actions/checkout from 2 to 2.3.4
* bump docker/login-action from 1 to 1.9.0
* bump docker/metadata-action from 3 to 3.1.0
* upgrade golangci-lint v1.40.0
* upgrade go v1.16.4
* use new repository of ghaction-docker-meta
* bump crazy-max/ghaction-docker-meta from v2.5.0 to v3
* bump crazy-max/ghaction-docker-meta from v2.4.0 to v2.5.0
* bump aquasecurity/trivy-action from 0.0.14 to 0.0.15
* bump crazy-max/ghaction-docker-meta from v2.3.0 to v2.4.0
* bump crazy-max/ghaction-docker-meta from v2.2.1 to v2.3.0
* bump aquasecurity/trivy-action from 0.0.13 to 0.0.14
* upgrade distroless cd784033
* bump crazy-max/ghaction-docker-meta from v2.1.1 to v2.2.1
* bump golangci/golangci-lint-action from v2.5.1 to v2.5.2
* upgrade go v1.16.3
* upgrade go v1.16.2
* bump crazy-max/ghaction-docker-meta from v1 to v2.1.1
* add nolintlint configuration
* use the special comment form for nolint
* upgrade golangci-lint v1.39.0
* bump aquasecurity/trivy-action from 0.0.12 to 0.0.13
* bump aquasecurity/trivy-action from 0.0.11 to 0.0.12
* bump aquasecurity/trivy-action from v0.0.10 to 0.0.11
* upgrade controller-runtime v0.8.3
* bump aquasecurity/trivy-action from 0.0.9 to v0.0.10
* bump aquasecurity/trivy-action from 0.0.8 to 0.0.9
* bump golangci/golangci-lint-action from v2.4.0 to v2.5.1
* upgrade golangci-lint v1.37.1
* upgrade controller-runtime v0.8.2
* bump golangci/golangci-lint-action from v2 to v2.4.0
* upgrade go v1.15.8
* upgrade golangci-lint v1.36.0
* upgrade controller-runtime v0.8.1
* upgrade e2e kubernetes version to v1.19.7, v1.20.2
* upgrade kind v0.10.0
* upgrade go v1.15.7
* bind local registry to localhost
* upgrade cucumber/godog v0.11.0
* upgrade stretchr/testify v1.7.0
* upgrade kubernetes dependencies v1.20.2
* upgrade golangci-lint v1.35.2
* upgrade golangci-lint v1.35.0

### Ci

* exclude dependabot trivy codesql push due to recent changes
* exclude dependabot from codeql scan due to recent changes
* add trivy scan action

### Feat

* add e2e BDD tests
* add operator to approve kubelet certificate use to serve TLS endpoints

