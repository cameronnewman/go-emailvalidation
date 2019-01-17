package email

import (
	"errors"
)

var (
	//ErrInvalidFormat is an error generatd when the format is incorrect
	ErrInvalidFormat = errors.New("Invalid email format")

	//ErrInvalidDomainNoNameServers is an error generatd when the domain returns no Name Server reocrds
	ErrInvalidDomainNoNameServers = errors.New("Invalid email domain, unable to find Name Servers records")

	//ErrInvalidDomainNoMXRecords is an error generatd when the domain returns no MX reocrds
	ErrInvalidDomainNoMXRecords = errors.New("Invalid email domain no MX records found")
)
