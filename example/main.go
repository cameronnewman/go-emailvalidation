package main

import (
	"fmt"

	"github.com/cameronnewman/go-emailvalidation"
)

func main() {

	emailaddress := "john.snow@gmaiiiiiiillllll.com"

	e := email.NewValidation()

	err := e.ValidateEmailAddress(emailaddress)
	if err != nil {
		fmt.Println(err)
	}

}
