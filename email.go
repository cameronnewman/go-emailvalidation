package email

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const (
	emptyString string = ""
)

var (
	// Matches the local part as an RFC 5322 dot-atom (no leading, trailing
	// or consecutive dots) followed by a hostname of dot-separated labels
	// of up to 63 characters each. https://tools.ietf.org/html/rfc5322
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Resolver is the subset of *net.Resolver used for domain record lookups.
// It can be replaced with a custom implementation, e.g. for testing or to
// use a specific DNS server.
type Resolver interface {
	LookupNS(ctx context.Context, name string) ([]*net.NS, error)
	LookupMX(ctx context.Context, name string) ([]*net.MX, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
}

// Validator validates email addresses. The zero value is ready to use and
// performs DNS lookups with net.DefaultResolver.
type Validator struct {
	// Resolver used for DNS lookups. If nil, net.DefaultResolver is used.
	Resolver Resolver
}

var defaultValidator = &Validator{}

// Validate - validates an email address via all options using
// net.DefaultResolver without a timeout. Use ValidateContext to bound the
// DNS lookups.
func Validate(email string) error {
	return defaultValidator.Validate(context.Background(), email)
}

// ValidateContext - validates an email address via all options. The
// context bounds the DNS lookups, allowing timeouts and cancellation.
func ValidateContext(ctx context.Context, email string) error {
	return defaultValidator.Validate(ctx, email)
}

// ValidateFormat - validates an email address meets rfc 5322 format via a regex
func ValidateFormat(email string) error {

	_, _, err := validateFormatAndSplit(email)
	if err != nil {
		return err
	}

	return nil
}

// ValidateDomainRecords - validates an email address domain's NS and MX
// records via a DNS lookup using net.DefaultResolver without a timeout.
// Use ValidateDomainRecordsContext to bound the DNS lookups.
func ValidateDomainRecords(email string) error {
	return defaultValidator.ValidateDomainRecords(context.Background(), email)
}

// ValidateDomainRecordsContext - validates an email address domain's NS and
// MX records via a DNS lookup. The context bounds the DNS lookups, allowing
// timeouts and cancellation.
func ValidateDomainRecordsContext(ctx context.Context, email string) error {
	return defaultValidator.ValidateDomainRecords(ctx, email)
}

// Validate - validates an email address via all options
func (v *Validator) Validate(ctx context.Context, email string) error {
	return v.ValidateDomainRecords(ctx, email)
}

// ValidateDomainRecords - validates an email address domain's NS and MX
// records via a DNS lookup.
//
// Failed lookups are wrapped: use errors.Is with ErrInvalidDomainNoNameServers
// or ErrInvalidDomainNoMXRecords to classify the failure, and errors.As with
// *net.DNSError (or errors.Is with context.DeadlineExceeded / context.Canceled)
// to distinguish a transient lookup error from a domain that genuinely has no
// records.
func (v *Validator) ValidateDomainRecords(ctx context.Context, email string) error {

	_, domain, err := validateFormatAndSplit(email)
	if err != nil {
		return err
	}

	effectiveTLDPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return ErrInvalidDomainNoNameServers
	}

	r := v.resolver()

	// Added NS check as some ISPs hijack the MX record lookup :(
	nsRecords, err := r.LookupNS(ctx, effectiveTLDPlusOne)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidDomainNoNameServers, err)
	}
	if len(nsRecords) == 0 {
		return ErrInvalidDomainNoNameServers
	}

	mxRecords, err := r.LookupMX(ctx, domain)
	if err == nil && len(mxRecords) > 0 {
		return nil
	}

	// No MX records; fall back to an address lookup (RFC 5321 section 5.1)
	ipRecords, err := r.LookupIP(ctx, "ip", domain)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidDomainNoMXRecords, err)
	}
	if len(ipRecords) == 0 {
		return ErrInvalidDomainNoMXRecords
	}

	return nil
}

func (v *Validator) resolver() Resolver {
	if v.Resolver != nil {
		return v.Resolver
	}
	return net.DefaultResolver
}

// Split - Splits an email address into a username prefix and domain
func Split(email string) (username string, domain string) {

	username, domain, err := validateFormatAndSplit(email)
	if err != nil {
		return emptyString, emptyString
	}

	return username, domain
}

// Normalize - Trim whitespaces, extra dots in the hostname and converts to Lowercase.
//
// Note: RFC 5321 only requires the domain to be case-insensitive; the local
// part is technically case-sensitive. Lowercasing the whole address is a
// deliberate practical choice for storage and comparison, as virtually all
// mail providers treat the local part case-insensitively.
func Normalize(email string) string {

	email = strings.TrimSpace(email)
	email = strings.TrimRight(email, ".")
	email = strings.ToLower(email)

	return email
}

func validateFormatAndSplit(email string) (username string, domain string, err error) {
	// RFC 5321 limits a path to 254 octets; the shortest plausible
	// address ("a@b.co") is 6 characters.
	if len(email) < 6 || len(email) > 254 {
		return emptyString, emptyString, ErrInvalidFormat
	}

	// Regex matches as per rfc 5322 https://tools.ietf.org/html/rfc5322
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
