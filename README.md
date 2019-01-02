# go-emailvalidation

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build][5]][6]
[![codecov][7]][8]

[1]: https://godoc.org/github.com/cameronnewman/go-emailvalidation?status.svg
[2]: https://godoc.org/github.com/cameronnewman/go-emailvalidation
[3]: https://goreportcard.com/badge/github.com/cameronnewman/go-emailvalidation
[4]: https://goreportcard.com/report/github.com/cameronnewman/go-emailvalidation
[5]: https://travis-ci.org/cameronnewman/go-emailvalidation.svg?branch=master
[6]: https://travis-ci.org/cameronnewman/go-emailvalidation
[7]: https://codecov.io/gh/cameronnewman/go-emailvalidation/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/cameronnewman/go-emailvalidation


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

	emailaddress := "john.snow@gmaiiiiiiillllll.com"

	err := email.New().ValidateEmailAddress(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

}
```


## Issues
* None