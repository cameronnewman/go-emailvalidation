package email

import (
	"errors"
)

var (
	//ErrEmailInvalidFormat is an error generatd when the format is incorrect
	ErrEmailInvalidFormat = errors.New("Invalid email format")

	//ErrEmailInvalidDomain is an error generatd when the domain is invalid or no MX reocrds can be found
	ErrEmailInvalidDomain = errors.New("Invalid email domain OR MX records don't exist")
)
