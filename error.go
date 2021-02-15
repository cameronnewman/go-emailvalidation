package email

import (
	"errors"
)

var (
	//ErrInvalidFormat is an error generated when the format is incorrect
	ErrInvalidFormat = errors.New("invalid email format")

	//ErrInvalidDomainNoNameServers is an error generated when the domain returns no Name Server records
	ErrInvalidDomainNoNameServers = errors.New("invalid email domain, unable to find name servers records")

	//ErrInvalidDomainNoMXRecords is an error generated when the domain returns no MX records
	ErrInvalidDomainNoMXRecords = errors.New("invalid email domain no mx records found")
)
