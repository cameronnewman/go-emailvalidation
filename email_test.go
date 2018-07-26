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
			NewValidation(),
			args{
				email: "john.snow@gmail.com",
			},
			false,
		},
		{
			"InvalidEmailAddress/Format",
			NewValidation(),
			args{
				email: "john.snow()()()@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Format",
			NewValidation(),
			args{
				email: "john.snow(*&(&*))@gmail.com",
			},
			true,
		},
		{
			"InvalidEmailAddress/Domain",
			NewValidation(),
			args{
				email: "john.snow&@gggggggmmmmaaaiiilllll.com",
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
			NewValidation(),
			args{
				domain: "gmail.com",
			},
			false,
		},
		{
			"InvalidDomain/NonExistantDomain",
			NewValidation(),
			args{
				domain: "gggggggmmmmaaaiiilllll.com",
			},
			true,
		},
		{
			"InvalidDomain/NoMailRecords",
			NewValidation(),
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
			NewValidation(),
			args{
				email: "john.snow@gmail.com",
			},
			"john.snow",
			"gmail.com",
		},
		{
			"InValidEmailAddress",
			NewValidation(),
			args{
				email: "john.snow_gmail.com",
			},
			"",
			"",
		},
		{
			"InValidEmailAddress",
			NewValidation(),
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
