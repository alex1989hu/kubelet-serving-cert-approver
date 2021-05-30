<a name="0.1.1"></a>
## [0.1.1](https://github.com/alex1989hu/kubelet-serving-cert-approver/compare/0.1.0...v0.1.1) (2021-05-30)

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
