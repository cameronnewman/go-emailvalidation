# go-emailvalidation

[![Build][1]][2]
[![GoDoc][3]][4]
[![Go Report Card][5]][6]
[![FOSSA Status][9]][10]

[1]: https://github.com/cameronnewman/go-emailvalidation/workflows/pipeline/badge.svg
[2]: https://github.com/cameronnewman/go-emailvalidation/actions/workflows/pipeline.yml/badge.svg
[3]: https://godoc.org/github.com/cameronnewman/go-emailvalidation?status.svg
[4]: https://godoc.org/github.com/cameronnewman/go-emailvalidation
[5]: https://goreportcard.com/badge/github.com/cameronnewman/go-emailvalidation
[6]: https://goreportcard.com/report/github.com/cameronnewman/go-emailvalidation
[9]: https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation.svg?type=shield
[10]: https://app.fossa.io/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation?ref=badge_shield

## Purpose

Simple email validation package. Package supports either fast email validation via
RFC compliant regex OR slower recursive DNS lookup. DNS validation checks for
valid NS & MX records using the local DNS settings.

## Supported versions

The current version is v3

Please import via go modules `github.com/cameronnewman/go-emailvalidation/v3`.

## Usage

```golang
package main

import (
    "fmt"

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
}
```

## Issues

* None

## License

MIT Licensed
