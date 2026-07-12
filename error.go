package email

import (
	"errors"
)

var (
	//ErrInvalidFormat is an error generated when the format is incorrect
	ErrInvalidFormat = errors.New("invalid email format")

	//ErrInvalidDomainNoNameServers is an error generated when the domain
	//returns no Name Server records. When the DNS lookup itself failed the
	//underlying error is wrapped and can be inspected with errors.As, e.g.
	//as a *net.DNSError to detect timeouts.
	ErrInvalidDomainNoNameServers = errors.New("invalid email domain, unable to find name servers records")

	//ErrInvalidDomainNoMXRecords is an error generated when the domain
	//returns no MX records. When the fallback address lookup failed the
	//underlying error is wrapped and can be inspected with errors.As.
	ErrInvalidDomainNoMXRecords = errors.New("invalid email domain no mx records found")
)
