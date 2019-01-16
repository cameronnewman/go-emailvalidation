package main

import (
	"fmt"

	email "github.com/cameronnewman/go-emailvalidation"
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
