package email

import (
	"testing"
)

func TestEmailValidation_ValidateEmailAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		e       *Validation
		args    args
		wantErr bool
	}{
		{
			"ValidEmailAddress",
			New(),
			args{
				email: "john.snow@gmail.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/Format",
			New(),
			args{
				email: "john.snow()()()@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Format",
			New(),
			args{
				email: "john.snow(*&(&*))@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Domain",
			New(),
			args{
				email: "john.snow&@gggggggmmmmaaaiizilllll.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Validation{}
			if err := e.ValidateEmailAddress(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Validation.ValidateEmailAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailValidation_ValidateFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		e       *Validation
		args    args
		wantErr bool
	}{
		{
			"ValidEmailAddressDot",
			New(),
			args{
				email: "john.snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddressUnderscore",
			New(),
			args{
				email: "john_snow@gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddressCase",
			New(),
			args{
				email: "JohnSnow@Gmail.com",
			},
			false,
		},
		{
			"ValidEmailAddressDomainDash",
			New(),
			args{
				email: "JohnSnow@g-mail.com",
			},
			false,
		},
		{
			"InvalidEmailAddressDomainUnderScore",
			New(),
			args{
				email: "JohnSnow@g_mail.com",
			},
			true,
		},
		{
			"InvalidEmailAddressDomainUnderScore",
			New(),
			args{
				email: "JohnSnow@G_mAil.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Format",
			New(),
			args{
				email: "john.snow()()()@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Format",
			New(),
			args{
				email: "john.snow(*&(&*))@gmail.com",
			},
			true,
		},
		{
			"ValidEmailAddress/InvalidDomain",
			New(),
			args{
				email: "john.snow@gggggggmmmmaaaiizilllll.com",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Validation{}
			if err := e.ValidateFormat(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Validation.ValidateFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailValidation_ValidateDomainMailRecords(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		e       *Validation
		args    args
		wantErr bool
	}{
		{
			"ValidDomain",
			New(),
			args{
				domain: "gmail.com",
			},
			false,
		},
		{
			"InvalidDomain/NonExistantDomain",
			New(),
			args{
				domain: "gggggggmmmmaaaiiilllll.com",
			},
			true,
		},
		{
			"InvalidDomain/NoMailRecords",
			New(),
			args{
				domain: "support.google.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Validation{}
			if err := e.ValidateDomainMailRecords(tt.args.domain); (err != nil) != tt.wantErr {
				t.Errorf("Validation.ValidateDomainMailRecords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailValidation_SplitEmailAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name         string
		e            *Validation
		args         args
		wantUsername string
		wantDomain   string
	}{
		{
			"ValidEmailAddress",
			New(),
			args{
				email: "john.snow@gmail.com",
			},
			"john.snow",
			"gmail.com",
		},
		{
			"InValidEmailAddress",
			New(),
			args{
				email: "john.snow_gmail.com",
			},
			"",
			"",
		},
		{
			"InValidEmailAddress",
			New(),
			args{
				email: "john.snow@@gmail.com",
			},
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Validation{}
			gotUsername, gotDomain := e.SplitEmailAddress(tt.args.email)
			if gotUsername != tt.wantUsername {
				t.Errorf("Validation.SplitEmailAddress() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("Validation.SplitEmailAddress() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
		})
	}
}
