/*
Package email is a simple email validation package. Package valids email address string to RFC requirements and performs a DNS lookup for the MX records using the local DNS settings.
*/
package email

// Simple example:
//
// 	fmt.Println("Hello")
//	emailaddress := "john.snow@gmaiiiiiiillllll.com"
//
//	err := email.New().ValidateEmailAddress(emailaddress)
//	if err != nil {
//		fmt.Println(err)
//	}
//
