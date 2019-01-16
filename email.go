package email

import (
	"net"
	"regexp"
	"strings"
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//Validation struct
type Validation struct {
}

//New returns a new Validation struct
func New() *Validation {
	return &Validation{}
}

//ValidateEmailAddress - validates an email address via all methods
func (e *Validation) ValidateEmailAddress(email string) error {

	err := e.ValidateFormat(email)
	if err != nil {
		return err
	}
	_, domain := e.SplitEmailAddress(email)

	err = e.ValidateDomainMailRecords(domain)
	if err != nil {
		return err
	}

	return nil
}

//ValidateFormat - validates an email address meets rfc 822 format via a regex
func (e *Validation) ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrEmailInvalidFormat
	}
	return nil
}

//ValidateDomainMailRecords - validates a domain MX records via a DNS lookup
func (e *Validation) ValidateDomainMailRecords(domain string) error {

	_, err := net.LookupMX(domain)
	if err != nil {
		return ErrEmailInvalidDomain
	}
	return nil
}

//SplitEmailAddress - Splits an email address into a prefix and domain
func (e *Validation) SplitEmailAddress(email string) (username, domain string) {

	components := strings.Split(email, "@")
	if len(components) == 2 {
		username, domain := components[0], components[1]
		return username, domain
	}

	return
}
