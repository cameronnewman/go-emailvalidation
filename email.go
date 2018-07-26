package email

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var (
	//ErrEmailInvalidFormat is an error generatd when the format is incorrect
	ErrEmailInvalidFormat = errors.New("Invalid email format")

	//ErrEmailInvalidDomain is an error generatd when the domain is invalid or no MX reocrds can be found
	ErrEmailInvalidDomain = errors.New("Invalid email domain OR MX records don't exist")

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//Validation struct
type Validation struct {
}

//NewValidation returns a new Validation struct
func NewValidation() *Validation {
	e := &Validation{}
	return e
}

//ValidateEmailAddress - validates and email address via a regix and then a DNS lookup for MX records
func (e *Validation) ValidateEmailAddress(email string) error {

	if !emailRegexp.MatchString(email) {
		return ErrEmailInvalidFormat
	}

	_, host := e.SplitEmailAddress(email)
	_, err := net.LookupMX(host)
	if err != nil {
		return ErrEmailInvalidDomain
	}
	return nil
}

//SplitEmailAddress - Splits an email address into a prefix and domains
func (e *Validation) SplitEmailAddress(email string) (username, domain string) {

	components := strings.Split(email, "@")
	if len(components) == 2 {
		username, domain := components[0], components[1]
		return username, domain
	}

	return
}
