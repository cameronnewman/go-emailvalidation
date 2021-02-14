package email

import (
	"testing"
)

func TestValidate(t *testing.T) {
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
			"InvalidDomain/NoMailRecords",
			args{
				email: "justin@support.google.com",
			},
			true,
		},
		{
			"InvalidDomain/InvalidCharacters",
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
			"ValidEmailAddress/DomainNoMXbutA",
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
			"ValidEmailAddress/DomainNoMXbutA",
			args{
				email: "JohnSnow@support.google.com",
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
			true,
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
