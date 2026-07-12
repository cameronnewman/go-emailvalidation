# CHANGELOG

## v3.4

- Added context-aware validation (ValidateContext,
  ValidateDomainRecordsContext) and a Validator type with a pluggable
  DNS Resolver
- DNS lookup failures are now wrapped, so errors.Is/errors.As can
  distinguish transient lookup errors (timeouts, cancellation) from
  domains with missing records
- Stricter RFC 5322 format check: consecutive, leading or trailing dots
  in the local part are now rejected
- Removed duplicated name server lookup logic
- Unit tests no longer depend on live DNS (go test -short skips the
  single remaining live-DNS integration test)

## v3.3

- Updated to golang 1.26
- Bumped golang.org/x/net from v0.23.0 to v0.57.0
- Added Trivy vulnerability/secret/misconfiguration scanning to CI,
  with scan policy in trivy.yaml and suppressions in .trivyignore
- Updated GitHub Actions (checkout v6, CodeQL v4) with least-privilege
  workflow permissions
- Updated linter images (golangci-lint v2.12.2, shellcheck v0.11.0)
- Updated SECURITY.md to use GitHub private vulnerability reporting

## v3.2

- Updated to golang 1.20
- Updated all linters

## v3.1

- Updated to golang 1.19
- Updated all linters

## v2.3

- Switched to Github Actions and removed TravisCI
- Fixed shell lint errors

## v2.1.144

- Removed double function call

## v2.0.131

- simplified all functions
- added Normalize function

- Massively improved all unit tests
- Updated the docs/readme/example

## v1.1.127

- Added FOSSA license check

## v1.1.122

- Added changelog

## v1.1.117

- Updated the readme to include license details

## v1.1.113

### New features

- Added new function to only validate an email address
- Added further tests for email address domains

## v1.0.105

- Updated readme layout

## v1.0.101

- Added Coverage report via codecov.io

## v1.0.95

- Add Github token check for build tags

## v1.0.91

- Fixed typo in doc.go

## v1.0.87

- Updated doc.go and README with clearer details of the project

## v1.0.83

- Add more logic for PR check
