# Repository Guidelines

## Project Structure & Module Organization
- Source: package `underscore` in repo root (`*.go`).
- Tests: co-located `*_test.go` files in root; examples in `examples/`.
- Module: `github.com/rjNemo/underscore` (Go 1.24+).
- Docs: Hugo site under `docs/` (content, themes, public).
- CI/Security assets: Dockerfile(s), `.golangci.yml`, `.trivycache/`.

## Build, Test, and Development Commands
- `go mod download`: fetch dependencies for local dev.
- `go test ./...`: run unit tests locally.
- `make build`: build Docker image `underscore:latest`.
- `make test`: run tests in Docker with coverage summary.
- `make docs`: serve docs locally (`cd docs && hugo server -D`).
- `make build-docs`: build static docs site.
- `make scan` / `make scan-config`: security scan with Trivy (image/config).

## Coding Style & Naming Conventions
- Formatting: `gofmt`/`goimports` (enforced via `golangci-lint`).
- Lint: `golangci-lint run` (uses `.golangci.yml`).
- Indentation: tabs; idiomatic Go style.
- Naming: exported APIs use PascalCase (e.g., `Filter`, `Map`); files are lower_snake (`filter.go`).
- Imports: group stdlib/external/local; local prefix `github.com/rjNemo/underscore`.

## Testing Guidelines
- Frameworks: Go `testing` with `testify` assertions.
- Conventions: `TestXxx` functions; prefer table-driven tests.
- Coverage: keep or increase overall coverage; `make test` prints summary.
- Run subset: `go test ./... -run TestFilter` for focused runs.

## Commit & Pull Request Guidelines
- Commits: imperative, concise subject; include scope when helpful.
- Before PR: `golangci-lint run` and `make test` must pass.
- PRs: clear description, linked issues, tests for new/changed behavior, update README/docs when API changes.
- Reviews: follow CONTRIBUTING; require two approvals before merge.

## Security & Configuration Tips
- Go 1.24+ recommended; Docker image pins Go 1.24.
- Do not commit secrets; prefer environment variables for local runs.
- Use `make scan` regularly; fix CRITICAL findings before release.
