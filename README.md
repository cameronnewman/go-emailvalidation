# go-emailvalidation

## Purpose ##

Simple Email Validation package.

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build][5]][6]

[1]: https://godoc.org/github.com/cameronnewman/go-emailvalidation?status.svg
[2]: https://godoc.org/github.com/cameronnewman/go-emailvalidation
[3]: https://goreportcard.com/badge/github.com/cameronnewman/go-emailvalidation
[4]: https://goreportcard.com/report/github.com/cameronnewman/go-emailvalidation
[5]: https://travis-ci.org/cameronnewman/go-emailvalidation.svg?branch=master
[6]: https://travis-ci.org/cameronnewman/go-emailvalidation

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
