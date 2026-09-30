# Repository Guidelines

## Project Structure & Module Organization

This repository contains two independent Go `package main` AWS Lambda samples:

- `sample_source1/`: converts portal JSON from S3 into standard and API-oriented data formats.
- `sample_source2/`: generates XML, creates ZIP payloads, calls external APIs, and stores results in S3.
- Each directory contains its Lambda handler in `sample_sourceN.go` and data structures in `data.go`; the second sample also keeps form templates and conversion helpers there.
- `README.md` provides a Japanese workflow overview. Check implementation details before relying on that summary.

There are currently no separate asset directories, test files, or deployment definitions.

## Build, Test, and Development Commands

No `go.mod`, dependency lockfile, build script, or CI configuration is committed. Configure dependencies before building: both samples import `aws-lambda-go` and `aws-sdk-go`; the first also uses `github.com/zenwerk/jptel`. Document dependency versions and setup when introducing reproducible builds.

- `gofmt -w sample_source1/*.go sample_source2/*.go`: format Go sources; review the diff for unrelated changes.
- After root module setup, `go build -o /tmp/sample1 ./sample_source1`: compile the first sample without colliding with its directory name. Substitute `sample2` and `sample_source2` for the second.
- After dependency setup, `go test ./...` and `go vet ./...`: run tests and static checks from the configured module root.

Handlers require Lambda runtime context and S3 events; there is no standalone local development server.

## Coding Style & Naming Conventions

Use standard Go formatting and tab indentation. Use exported CamelCase names and unexported mixedCase helpers. Preserve schema field identifiers such as `K0001` and `D0001`, serialization tags, and Japanese domain comments. Keep changes scoped to the affected sample. No custom linter configuration exists.

## Testing Guidelines

Use Go's `testing` package, colocated `*_test.go` files, and `TestXxx` functions. Add table-driven cases for field mappings, date conversions, empty inputs, and XML/ZIP output. Isolate S3 and HTTP calls from unit tests. No coverage threshold is currently defined.

## Commit & Pull Request Guidelines

History uses short subjects such as `Update README.md` and `sample source update`; no Conventional Commits convention is established. Write concise, specific subjects. PRs should describe the affected sample, behavior changes, validation results or blockers, and configuration changes. Link related issues when available.

## Security & Configuration

Supply bucket paths and API settings through environment variables, including `BUCKET`, `OUTPUT_DIR`, and `API_PASS`. Keep credentials and real applicant data out of commits, fixtures, and logs. Use isolated S3 buckets and test API endpoints for integration checks.
