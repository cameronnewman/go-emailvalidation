package email

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"
)

// fakeResolver serves DNS records from in-memory maps so tests do not
// depend on live DNS. Names are matched case-insensitively, like real DNS.
type fakeResolver struct {
	ns map[string][]*net.NS
	mx map[string][]*net.MX
	ip map[string][]net.IP
}

func (f *fakeResolver) LookupNS(ctx context.Context, name string) ([]*net.NS, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if records, ok := f.ns[strings.ToLower(name)]; ok {
		return records, nil
	}
	return nil, &net.DNSError{Err: "no such host", Name: name, IsNotFound: true}
}

func (f *fakeResolver) LookupMX(ctx context.Context, name string) ([]*net.MX, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if records, ok := f.mx[strings.ToLower(name)]; ok {
		return records, nil
	}
	return nil, &net.DNSError{Err: "no such host", Name: name, IsNotFound: true}
}

func (f *fakeResolver) LookupIP(ctx context.Context, network, host string) ([]net.IP, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if records, ok := f.ip[strings.ToLower(host)]; ok {
		return records, nil
	}
	return nil, &net.DNSError{Err: "no such host", Name: host, IsNotFound: true}
}

func newFakeResolver() *fakeResolver {
	return &fakeResolver{
		ns: map[string][]*net.NS{
			"gmail.com":      {{Host: "ns1.google.com."}},
			"google.com":     {{Host: "ns1.google.com."}},
			"sparkpost.com":  {{Host: "ns1.sparkpost.com."}},
			"cloudfront.net": {{Host: "ns1.cloudfront.net."}},
		},
		mx: map[string][]*net.MX{
			"gmail.com":               {{Host: "gmail-smtp-in.l.google.com.", Pref: 5}},
			"validator.sparkpost.com": {{Host: "mx.sparkpost.com.", Pref: 10}},
		},
		ip: map[string][]net.IP{
			"support.google.com": {net.ParseIP("142.250.70.142")},
		},
	}
}

// useFakeResolver points the package-level functions at the fake resolver
// for the duration of a test.
func useFakeResolver(t *testing.T) {
	t.Helper()
	orig := defaultValidator.Resolver
	defaultValidator.Resolver = newFakeResolver()
	t.Cleanup(func() { defaultValidator.Resolver = orig })
}

func TestValidate(t *testing.T) {
	useFakeResolver(t)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"ValidEmailAddress/PrefixDot",
			args{
				email: "john.snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/PrefixUnderscore",
			args{
				email: "john_snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/AddressCase",
			args{
				email: "JohnSnow@Gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/DomainDash",
			args{
				email: "JohnSnow@g-mail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/DomainUnderscore",
			args{
				email: "JohnSnow@g_mail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/DomainCaseUnderscore",
			args{
				email: "JohnSnow@G_mAil.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/PrefixBrackets",
			args{
				email: "john.snow()()()@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/TooShort",
			args{
				email: "j@g.i",
			},
			true,
		},
		{
			"InvalidEmailAddress/TooLong",
			args{
				email: "jasdkjhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkjasdkjhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhahasdkhaskdhasdkjhasdkjhasdkhasdkhaskdha@gasdkjhaskdhaskdhaksdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdhkasdhkashdkasjhdkasjdhkashd.isadhaskdhaskdhaksjdh",
			},
			true,
		},
		{
			"InvalidEmailAddress/UsernameTooLongDomainOk",
			args{
				email: "sgafetdgetsgdhstegdtsheduejdiekdjsyehrjehdusjegdtsgetsfgedswedtgd@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/PrefixAmpersandStar",
			args{
				email: "john.snow(*&(&*))@gmail.com",
			},
			true,
		},
		{
			"ValidEmailAddress/InvalidDomain",
			args{
				email: "john.snow@gggggggmmmmaaaiizilllll.com",
			},
			true,
		},
		{
			"ValidEmailAddress/PrefixAmpersandInvalidDomain",
			args{
				email: "john.snow&@gggggggmmmmaaaiizilllll.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/NoPrefix",
			args{
				email: "@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/NoPrefixDoubleDomain",
			args{
				email: "@gmail.com@gmail.com",
			},
			true,
		},
		{
			"ValidDomain",
			args{
				email: "justin@gmail.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/NoAtSymbol",
			args{
				email: "john.snow_gmail.com",
			},
			true,
		},
		{
			"InvalidDomain/NonExistentDomain",
			args{
				email: "justin@gggggggmmmmaaaiiilllll.com",
			},
			true,
		},
		{
			"ValidEmailAddress/SubDomain",
			args{
				email: "kasdhaskdhaskdjhaksd@validator.sparkpost.com",
			},
			false,
		},
		{
			"InvalidDomain/SubDomain/NoMXORA",
			args{
				email: "kasdhaskdhaskdjhaksd@www.cloudfront.net",
			},
			true,
		},
		{
			"InvalidDomain/SubDomain/NoMXORA",
			args{
				email: "kasdhaskdhaskdjhaksd@bob.www.cloudfront.net",
			},
			true,
		},
		{
			"InvalidDomain/InvalidCharactersPrefixAmpersand",
			args{
				email: "john.snow&@gggggggmmmmaaaiizilllll.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("email.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"ValidEmailAddress/PrefixDot",
			args{
				email: "john.snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/PrefixUnderscore",
			args{
				email: "john_snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/AddressCase",
			args{
				email: "JohnSnow@Gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/DomainDash",
			args{
				email: "JohnSnow@g-mail.com",
			},
			false,
		},
		{
			"ValidEmailAddress/SubDomain/MissingMXSubDomainARecord",
			args{
				email: "JohnSnow@support.google.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/DomainUnderscore",
			args{
				email: "JohnSnow@g_mail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/DomainCaseUnderscore",
			args{
				email: "JohnSnow@G_mAil.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/PrefixBrackets",
			args{
				email: "john.snow()()()@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/TooShort",
			args{
				email: "j@g.i",
			},
			true,
		},
		{
			"InvalidEmailAddress/TooLong",
			args{
				email: "jasdkjhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkjasdkjhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhasdkjhasdkjhasdkhasdkhaskdhahasdkhaskdhasdkjhasdkjhasdkhasdkhaskdha@gasdkjhaskdhaskdhaksdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdgasdkjhaskdhaskdhaksdhkasdhkashdkasjhdkasjdhkashdhkasdhkashdkasjhdkasjdhkashd.isadhaskdhaskdhaksjdh",
			},
			true,
		},
		{
			"InvalidEmailAddress/PrefixAmpersandStar",
			args{
				email: "john.snow(*&(&*))@gmail.com",
			},
			true,
		},
		{
			"ValidEmailAddress/InvalidDomain",
			args{
				email: "john.snow@gggggggmmmmaaaiizilllll.com",
			},
			false,
		},
		{
			"ValidEmailAddress/PrefixAmpersandInvalidDomain",
			args{
				email: "john.snow&@gggggggmmmmaaaiizilllll.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/NoPrefix",
			args{
				email: "@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/NoPrefixDoubleDomain",
			args{
				email: "@gmail.com@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/ConsecutiveDotsInPrefix",
			args{
				email: "john..snow@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/LeadingDotInPrefix",
			args{
				email: ".john@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/TrailingDotInPrefix",
			args{
				email: "john.@gmail.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFormat(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("email.ValidateFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDomainRecords(t *testing.T) {
	useFakeResolver(t)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"ValidDomain",
			args{
				email: "justin@gmail.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/NoAt",
			args{
				email: "john.snow_gmail.com",
			},
			true,
		},
		{
			"InvalidDomain/NonExistentDomain",
			args{
				email: "justin@gggggggmmmmaaaiiilllll.com",
			},
			true,
		},
		{
			"InvalidDomain/NoMailRecords",
			args{
				email: "justin@support.google.com",
			},
			false,
		},
		{
			"InvalidDomain/NonExistentDomain",
			args{
				email: "john.snow&@gggggggmmmmaaaiizilllll.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDomainRecords(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("email.ValidateDomainMailRecords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatorCustomResolver(t *testing.T) {
	v := &Validator{Resolver: newFakeResolver()}

	if err := v.Validate(context.Background(), "john.snow@gmail.com"); err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}

	err := v.Validate(context.Background(), "john.snow@gggggggmmmmaaaiizilllll.com")
	if !errors.Is(err, ErrInvalidDomainNoNameServers) {
		t.Errorf("Validator.Validate() error = %v, want ErrInvalidDomainNoNameServers", err)
	}
}

func TestValidateDomainRecordsWrapsLookupErrors(t *testing.T) {
	useFakeResolver(t)

	err := ValidateDomainRecords("justin@gggggggmmmmaaaiiilllll.com")
	if !errors.Is(err, ErrInvalidDomainNoNameServers) {
		t.Fatalf("error = %v, want ErrInvalidDomainNoNameServers", err)
	}

	var dnsErr *net.DNSError
	if !errors.As(err, &dnsErr) {
		t.Fatalf("error = %v, want wrapped *net.DNSError", err)
	}
	if !dnsErr.IsNotFound {
		t.Errorf("dnsErr.IsNotFound = false, want true")
	}
}

func TestValidateContextCancellation(t *testing.T) {
	useFakeResolver(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ValidateContext(ctx, "john.snow@gmail.com")
	if !errors.Is(err, ErrInvalidDomainNoNameServers) {
		t.Fatalf("error = %v, want ErrInvalidDomainNoNameServers", err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("error = %v, want wrapped context.Canceled", err)
	}
}

// TestValidateLiveDNS is the only test that performs real DNS lookups.
// Run `go test -short` to skip it.
func TestValidateLiveDNS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live DNS test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := ValidateDomainRecordsContext(ctx, "john.snow@gmail.com"); err != nil {
		t.Errorf("email.ValidateDomainRecordsContext() error = %v, want nil", err)
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name         string
		args         args
		wantUsername string
		wantDomain   string
	}{
		{
			"ValidEmailAddress",
			args{
				email: "john.snow@gmail.com",
			},
			"john.snow",
			"gmail.com",
		},
		{
			"InvalidEmailAddress/NoAt",
			args{
				email: "john.snow_gmail.com",
			},
			"",
			"",
		},
		{
			"InvalidEmailAddress/TwoAtSymbols",
			args{
				email: "john.snow@@gmail.com",
			},
			"",
			"",
		},
		{
			"InvalidEmailAddress/NoPrefix",
			args{
				email: "@gmail.com",
			},
			"",
			"",
		},
		{
			"InvalidEmailAddress/NoPrefix",
			args{
				email: "@gmail.com@gmail.com",
			},
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsername, gotDomain := Split(tt.args.email)
			if gotUsername != tt.wantUsername {
				t.Errorf("email.SplitEmailAddress() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("email.SplitEmailAddress() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		args      args
		wantEmail string
	}{
		{
			"ValidEmailAddress/NoChanges",
			args{
				email: "john.snow@gmail.com",
			},
			"john.snow@gmail.com",
		},
		{
			"InvalidEmailAddress/NoChanges",
			args{
				email: "john.snow_gmail.com",
			},
			"john.snow_gmail.com",
		},
		{
			"ValidEmailAddress/RemoveWhiteSpace",
			args{
				email: " john.snow@gmail.com ",
			},
			"john.snow@gmail.com",
		},
		{
			"ValidEmailAddress/RemoveDomainDot",
			args{
				email: "john.snow@ns.gmail.com.",
			},
			"john.snow@ns.gmail.com",
		},
		{
			"ValidEmailAddress/SetToLowercase",
			args{
				email: "johN.Snow@Gmail.com.",
			},
			"john.snow@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmail := Normalize(tt.args.email)
			if gotEmail != tt.wantEmail {
				t.Errorf("email.Normalize() gotEmail = %v, want %v", gotEmail, tt.wantEmail)
			}
		})
	}
}
