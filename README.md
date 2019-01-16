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
[5]: https://travis-ci.org/cameronnewman/go-emailvalidation.svg?branch=master
[6]: https://travis-ci.org/cameronnewman/go-emailvalidation
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

	"github.com/cameronnewman/go-emailvalidation"
)

func main() {

	emailValidate := email.New()

	emailaddress := "john.snow@gmaiiiiiiillllll.com"

	// Run all checks
	err := emailValidate.ValidateEmailAddress(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

	// Just check the format
	err = emailValidate.ValidateFormat(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

	// Just check domain MX records exist
	err = emailValidate.ValidateDomainMailRecords(emailaddress)
	if err != nil {
		fmt.Println(err)
	}
}
```


## Issues
* None

## License
MIT Licensed