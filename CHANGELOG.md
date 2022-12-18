
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

### Ci

* remove no longer supported kubernetes versions


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

