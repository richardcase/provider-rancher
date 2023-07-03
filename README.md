# Provider Rancher

`provider-rancher` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Rancher API.

```bash
export TERRAFORM_PROVIDER_SOURCE=rancher/rancher2
export TERRAFORM_PROVIDER_REPO=https://github.com/rancher/terraform-provider-rancher2
export TERRAFORM_PROVIDER_VERSION=3.0.2
export TERRAFORM_PROVIDER_DOWNLOAD_NAME=terraform-provider-rancher2
export TERRAFORM_NATIVE_PROVIDER_BINARY=terraform-provider-rancher2_3.0.2_x5
export TERRAFORM_DOCS_PATH=docs/resources
```

Or for a fork:

```bash
export TERRAFORM_PROVIDER_SOURCE=rancher/rancher2
export TERRAFORM_PROVIDER_REPO=https://github.com/richardcase/terraform-provider-rancher2
export TERRAFORM_PROVIDER_VERSION=3.0.2
export TERRAFORM_PROVIDER_DOWNLOAD_NAME=terraform-provider-rancher2
export TERRAFORM_NATIVE_PROVIDER_BINARY=terraform-provider-rancher2_3.0.2_x5
export TERRAFORM_DOCS_PATH=docs/resources
```

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/richardcase/provider-rancher):
```
up ctp provider install richardcase/provider-rancher:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-rancher
spec:
  package: richardcase/provider-rancher:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/richardcase/provider-rancher).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/richardcase/provider-rancher/issues).
