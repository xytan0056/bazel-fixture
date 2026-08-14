# Design notes

Freeform notes about the fixture. This file is intentionally not referenced by
any BUILD target — it exists to exercise tango's handling of source-tree
changes that produce zero changed Bazel targets.

## Layers

- `pkg/*` — leaf utilities, no service-layer imports
- `service/*` — business logic; may import `pkg/*` and `proto/*`
- `cmd/*` — process entry points; import from `service/*` only

## Adding a new service

1. Create `service/<name>/{name}.go`, `{name}_test.go`, `BUILD.bazel`.
2. Wire it into `service/handlers` if it must be called on the request path.
3. Update `cmd/server/main.go` to construct it.
