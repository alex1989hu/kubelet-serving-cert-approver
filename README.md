# Kubelet Serving Certificate Approver

![CI](https://github.com/alex1989hu/kubelet-serving-cert-approver/workflows/build-and-test/badge.svg)
![e2e-test](https://github.com/alex1989hu/kubelet-serving-cert-approver/workflows/e2e-test/badge.svg)
[![codecov](https://codecov.io/gh/alex1989hu/kubelet-serving-cert-approver/branch/main/graph/badge.svg)](https://codecov.io/gh/alex1989hu/kubelet-serving-cert-approver)

Kubelet Serving Certificate Approver is a custom approving controller which approves `kubernetes.io/kubelet-serving` Certificate Signing Request that kubelet use to serve TLS endpoints.

## Why should I use Kubelet Serving Certificate Approver?

* You want to securely - in terms of trusted Certificate Authoritity (CA) - reach kubelet endpoint

* Signed serving certificates are honored as a valid kubelet serving certificate by the API server

* Don't want to use `--kubelet-insecure-tls` flag during installation of [metrics-server](https://github.com/kubernetes-sigs/metrics-server/)

## Do I need to have a commercial certificate?

No. Every Kubernetes cluster has a Cluster Root Certificate Authority (CA).

## How do I use Kubelet Serving Certificate Approver?

To install into your Kubernetes cluster, please navigate to [deploy](deploy) directory.

*Note: your Kubernetes cluster must be configured with enabled TLS Bootstrapping and provided `rotate-server-certificates: true` kubelet argument.*

## Kubernetes Compatibility Matrix

For older Kubernetes versions (`v1.19`, `v1.20`, `v1.21`) please see [older releases](https://github.com/alex1989hu/kubelet-serving-cert-approver/releases).

| Version        | Compatible |
| -------------- | ---------- |
| `v1.22`        | &check;    |
| `v1.23`        | &check;    |
| `v1.24`        | &check;    |
| `v1.25`        | &check;    |
| `v1.26`        | &check;    |
| `v1.27`        | &check;    |
| `v1.28`        | &check;    |

## Prometheus Metrics

You can download Prometheus metrics `/metrics` endpoint.

### Custom Metrics

| Metric                                                                     | Description                                        |
| -------------------------------------------------------------------------- | -------------------------------------------------- |
| `kubelet_serving_cert_approver_approved_certificate_signing_request_count` | The number of approved Certificate Signing Request |
| `kubelet_serving_cert_approver_invalid_certificate_signing_request_count`  | The number of invalid Certificate Signing Request  |

## Reference

* Original idea: <https://github.com/kontena/kubelet-rubber-stamp> which is unfortunately not maintained.
* Kubernetes TLS bootstrapping: <https://kubernetes.io/docs/reference/access-authn-authz/kubelet-tls-bootstrapping/>
* Conformant Rules: <https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers>

## License

Apache License, Version 2.0, see [LICENSE](LICENSE).
