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

	// Checks the format - this function performs no network operations and is very fast
	err = email.ValidateFormat(emailAddress)
	if err != nil {
		fmt.Println(err)
	}

	// Checks domain NS & MX records exist
	err = email.ValidateDomainRecords(emailAddress)
	if err != nil {
		fmt.Println(err)
	}

	// Normalize email address for storage
	emailNormal := email.Normalize(emailAddress)
	fmt.Println(emailNormal)
}
