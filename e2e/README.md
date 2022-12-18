# End-to-End (E2E)

The tests are bounded to `//go:build e2e` [build tag](https://golang.org/cmd/go/#hdr-Build_constraints) which is not activated by default.

## Which technology is being used?

The application [requirements](features) are being tested with [godog](https://github.com/cucumber/godog/) -  official **Cucumber BDD framework for Golang**.

## How do I execute End-to-End tests?

The are multiple options avaialable for test execution.

```bash
# in Git Repository root directory
go test -tags e2e -v ./e2e
# in the current directory where this README is also located
go test -tags e2e -v ./...
```

*Notice that the `go test` has an extra argument: `-tags e2e`.*

## Test Flakiness

Kubernetes removes `Events` and `Certificate Signing Requests` which has approval condition due to its `TTL` configuration *(e.g. `--event-ttl`, `--controllers: --csrcleaner`)*.

Meaning that execution of `e2e` tests can be flaky after a given period of time after approval of `Certificate Signing Request`.

To overcome this, the easiest way is to recreate/remove your [KinD (Kubernetes in Docker)](https://kind.sigs.k8s.io/) cluster.

The other option is to override previously mentioned configuration options before starting KiND:

```yaml
kubeadmConfigPatches:
- |-
  kind: ClusterConfiguration
  apiServer:
    extraArgs:
      # Increase duration of Event Time To Leave (TTL)
      "event-ttl": "8h0m0s"
  controllerManager:
    extraArgs:
      # Disable csrcleaner contoller to avoid removal of Certificate Signing Request; keep default KiND options
      "controllers": "*,bootstrapsigner,tokencleaner,-csrcleaner"
```

These options are already configured in resources being used by [Contribution Guideline](../CONTRIBUTING.md).

## Reference

* Kubernetes API Server Event TTL: <https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/>
* Kubernetes Controller Manager Controllers: <https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/>
