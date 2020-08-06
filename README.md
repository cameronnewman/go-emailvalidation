# go-emailvalidation

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build][5]][6]
[![codecov][7]][8]
[![FOSSA Status][9]][10]

[1]: https://godoc.org/github.com/cameronnewman/go-emailvalidation?status.svg
[2]: https://godoc.org/github.com/cameronnewman/go-emailvalidation
[3]: https://goreportcard.com/badge/github.com/cameronnewman/go-emailvalidation
[4]: https://goreportcard.com/report/github.com/cameronnewman/go-emailvalidation
[5]: https://github.com/cameronnewman/go-emailvalidation/workflows/pipeline/badge.svg
[6]: https://github.com/cameronnewman/go-emailvalidation/actions
[7]: https://codecov.io/gh/cameronnewman/go-emailvalidation/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/cameronnewman/go-emailvalidation
[9]: https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation.svg?type=shield
[10]: https://app.fossa.io/projects/git%2Bgithub.com%2Fcameronnewman%2Fgo-emailvalidation?ref=badge_shield


## Purpose ##

Simple email validation package. Package valids email address string to RFC requirements and performs a DNS lookup for the MX records using the local DNS settings.

## Usage

```
package main

import (
	"fmt"

	email "github.com/cameronnewman/go-emailvalidation"
)

func main() {

	emailaddress := "John.Snow@gmaiiiiiiillllll.com"

	// Run all checks, including validating the format along with DNS lookups which 
	// may be slower depending on your DNS server performance
	err := email.Validate(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

	// Checks the format - this function performs no network operations and is very fast
	err = email.ValidateFormat(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

	// Checks domain NS & MX records exist
	err = email.ValidateDomainRecords(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

	// Normalize email address for storage
	emailNormal := email.Normalize(emailaddress)
	fmt.Println(emailNormal)
}
```


## Issues
* None

## License
MIT Licensed