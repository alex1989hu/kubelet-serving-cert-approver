# How to Contribute

The Contribution Guideline helps you to understand how to contribute to **Kubelet Serving Certificate Approver**.

## Development

To ease development there is an option to use [Tilt](https://tilt.dev/) and [KinD (Kubernetes in Docker)](https://kind.sigs.k8s.io/).

Once you already installed all of them, you can start the development.

Files being used here:

* [hack/kind-with-registry.sh](hack/kind-with-registry.sh)
* [hack/teardown-kind-with-registry.sh](hack/teardown-kind-with-registry.sh)
* [Tiltfile](Tiltfile)

```bash
# Start Kubernetes Cluster with local Image Registry
hack/kind-with-registry.sh
# Start Tilt
tilt up

# Navigate back to console to remove your Tilt development
tilt down --delete-namespaces
# Stop and remove Kubernetes Cluster with local Image Registry
hack/teardown-kind-with-registry.sh
```

## Pull Requests

Use the [GitHub flow](https://guides.github.com/introduction/flow/) as main versioning workflow.

## Style Guide

All pull requests shall adhere to the [Conventional Commits specification](https://conventionalcommits.org/).
