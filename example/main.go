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
