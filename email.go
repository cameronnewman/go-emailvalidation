package email

import (
	"net"
	"regexp"
	"strings"
)

const (
	emptyString string = ""
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//Validate - validates an email address via all options
func Validate(email string) error {

	err := ValidateFormat(email)
	if err != nil {
		return err
	}

	err = ValidateDomainRecords(email)
	if err != nil {
		return err
	}

	return nil
}

//ValidateFormat - validates an email address meets rfc 822 format via a regex
func ValidateFormat(email string) error {

	_, _, err := validateFormatAndSplit(email)
	if err != nil {
		return err
	}

	return nil
}

//ValidateDomainRecords - validates an email address domain's NS and MX records via a DNS lookup
func ValidateDomainRecords(email string) error {

	_, domain, err := validateFormatAndSplit(email)
	if err != nil {
		return ErrInvalidFormat
	}

	// Added NS check as some ISPs hijack the MX record lookup :(
	nsRecords, err := net.LookupNS(domain)
	if err != nil || len(nsRecords) == 0 {
		return ErrInvalidDomainNoNameServers
	}

	if _, err := net.LookupMX(domain); err != nil {
		if _, err := net.LookupIP(domain); err != nil {
			return ErrInvalidDomainNoMXRecords
		}
	}

	return nil
}

//Split - Splits an email address into a username prefix and domain
func Split(email string) (username string, domain string) {

	username, domain, err := validateFormatAndSplit(email)
	if err != nil {
		return emptyString, emptyString
	}

	return username, domain
}

// Normalize - Trim whitespaces, extra dots in the hostname and converts to Lowercase.
func Normalize(email string) string {

	email = strings.TrimSpace(email)
	email = strings.TrimRight(email, ".")
	email = strings.ToLower(email)

	return email
}

func validateFormatAndSplit(email string) (username string, domain string, err error) {
	if len(email) < 6 || len(email) > 254 {
		return emptyString, emptyString, ErrInvalidFormat
	}

	if !emailRegexp.MatchString(email) {
		return emptyString, emptyString, ErrInvalidFormat
	}

	i := strings.LastIndexByte(email, '@')
	username = email[:i]
	domain = email[i+1:]

	if len(username) > 64 {
		return emptyString, emptyString, ErrInvalidFormat
	}

	return username, domain, nil
}
