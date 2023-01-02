package secret

import (
	"time"
)

const (
	Valid = iota
	Expired
	Invalid
	Missing
)

type Otp struct {
	Expiry int64
	Code   string
}

type Validator struct {
	clock func() time.Time
}

func NewValidator(clock func() time.Time) *Validator {
	return &Validator{clock: clock}
}

func (v *Validator) Validate(code string, otp Otp) int {
	empty := Otp{}
	if otp == empty {
		return Missing
	}

	now := v.clock().Unix()
	if now > otp.Expiry {
		return Expired
	}

	if code == "" || code != otp.Code {
		return Invalid
	}

	return Valid
}
