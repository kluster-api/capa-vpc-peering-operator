# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.klusters.dev/capa-vpc-peering-operator` — a Kubernetes operator that automates AWS VPC peering between Cluster API (CAPA) managed EKS clusters and a designated "peer" VPC. It watches two upstream resources:

- `sigs.k8s.io/cluster-api-provider-aws/v2` `AWSManagedControlPlane` — discovers the workload cluster's VPC.
- `kubedb.dev/provider-aws/apis/ec2/v1alpha1.VPCPeeringConnection` — the Crossplane AWS provider's VPC peering resource.

When both are present, it reconciles the security-group / firewall rules needed for the peered VPCs to communicate.

The produced binary is `capa-vpc-peering-operator`.

## Architecture

- `cmd/capa-vpc-peering-operator/` — entry point.
- `pkg/cmds/` — Cobra `root` and `run` commands.
- `pkg/controllers/` — two reconcilers, both `controller-runtime`-based:
  - `awsmanagedcontrolplane_controller.go` — `AWSManagedControlPlaneReconciler`; gates work behind `DeletionTimestamp`, triggers `firewall` operations.
  - `vpcpeeringconnection_controller.go` — `VPCPeeringConnectionReconciler`; reacts to Crossplane `VPCPeeringConnection` and updates the matching firewall rules. Uses `handler` watches to requeue parents when peering state changes.
  - `suite_test.go` — `envtest`-based suite using Ginkgo/Gomega.
- `pkg/firewall/` — the actual security-group reconciliation:
  - `crossplane.go` — calls Crossplane AWS provider types (`ec2api`) via `kmodules.xyz/client-go/client` patch helpers.
  - `utils.go` — shared helpers.
- `Dockerfile.in` (PROD, distroless), `Dockerfile.dbg` (debian) — two image variants (no UBI for this one).
- `hack/`, `Makefile` — AppsCode build harness (everything runs inside `ghcr.io/appscode/golang-dev`).
- `vendor/` — checked-in deps.

The operator does **not** define its own CRDs — it composes upstream CAPA + Crossplane AWS provider types. Bumping either upstream dep affects the reconciler surface directly.

## Common commands

All build/test/lint targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make ci` — CI pipeline: `check-license lint build` (unit-tests and verify are commented out in CI; run `make test`/`make verify` locally before opening a PR).
- `make build` / `make all-build` — build host or all-platform binaries.
- `make fmt` — gofmt + goimports.
- `make lint` — golangci-lint.
- `make unit-tests` — Go unit tests.
- `make e2e-tests` / `make test` — runs both unit and e2e (Ginkgo, controller-runtime envtest under `pkg/controllers/suite_test.go`).
- `make verify` — `verify-gen verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make container` — build PROD and DBG images.
- `make push` — push both image variants; `make docker-manifest` writes multi-arch manifests; `make release` is the full publish flow.
- `make push-to-kind` / `make deploy-to-kind` — load into Kind and Helm-install.
- `make install` / `make uninstall` / `make purge` — Helm install lifecycle.
- `make add-license` / `make check-license` — manage license headers.
- `make run` — run the binary locally against the current kubeconfig.

Run a single Go test locally (requires a Go toolchain):

```
go test ./pkg/controllers/... -run TestName -v
```

## Conventions

- Module path is `go.klusters.dev/capa-vpc-peering-operator` (vanity URL); imports must use that, not the GitHub URL.
- License: **AppsCode Community License 1.0.0** (`LICENSE.md`); new files need the standard AppsCode header (`make add-license`).
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO` file).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- Two Dockerfiles, one binary (`Dockerfile.in` PROD, `Dockerfile.dbg` DBG) — no UBI variant; keep them in sync when changing build steps.
- The operator depends tightly on **two upstream APIs** (CAPA and KubeDB's Crossplane AWS provider). Pin both deliberately; when bumping, expect both reconcilers to compile against the new API shapes simultaneously.
