# go-emailvalidation

[![Build][1]][2]
[![GoDoc][3]][4]
[![Go Report Card][5]][6]
[![FOSSA Status][9]][10]

[1]: https://github.com/cameronnewman/go-emailvalidation/actions/workflows/pipeline.yml/badge.svg
[2]: https://github.com/cameronnewman/go-emailvalidation/actions/workflows/pipeline.yml
[3]: https://pkg.go.dev/badge/github.com/cameronnewman/go-emailvalidation/v3.svg
[4]: https://pkg.go.dev/github.com/cameronnewman/go-emailvalidation/v3
[5]: https://goreportcard.com/badge/github.com/cameronnewman/go-emailvalidation
[6]: https://goreportcard.com/report/github.com/cameronnewman/go-emailvalidation
[9]: https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation.svg?type=shield
[10]: https://app.fossa.io/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation?ref=badge_shield

## Purpose

Simple email validation package. Package supports either fast email validation via
RFC compliant regex OR slower recursive DNS lookup. DNS validation checks for
valid NS & MX records using the local DNS settings.

## Supported versions

The current version is v3 and requires Go 1.26 or later.

Please import via go modules `github.com/cameronnewman/go-emailvalidation/v3`.

## Usage

```golang
package main

import (
    "context"
    "fmt"
    "time"

    email "github.com/cameronnewman/go-emailvalidation/v3"
)

func main() {

    emailAddress := "John.Snow@gmaiiiiiiillllll.com"

    // Run all checks, including validating the format along with DNS lookups which
    // may be slower depending on your DNS server performance
    err := email.Validate(emailAddress)
    if err != nil {
        fmt.Println(err)
    }

    // Same as above, with a timeout on the DNS lookups
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = email.ValidateContext(ctx, emailAddress)
    if err != nil {
        fmt.Println(err)
    }

    // Checks the format - this function performs no network
    // operations and is very fast
    err = email.ValidateFormat(emailAddress)
    if err != nil {
        fmt.Println(err)
    }

    // Checks domain NS & MX, along with format validation
    err = email.ValidateDomainRecords(emailAddress)
    if err != nil {
        fmt.Println(err)
    }

    // Normalize email address for storage
    address := email.Normalize(emailAddress)
    fmt.Println(address)

    // Split an email address into the username and domain parts
    username, domain := email.Split(emailAddress)
    fmt.Println(username, domain)
}
```

## Security

See [SECURITY.md](SECURITY.md) for the supported versions and how to
report a vulnerability. Dependencies and code are scanned in CI with
Trivy and CodeQL. The Trivy scan policy is defined in `trivy.yaml`,
and suppressed findings are tracked in `.trivyignore`.

## License

MIT Licensed
