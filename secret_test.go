package secret_test

import (
	"github.com/hill-daniel/secret"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidator_Validate_should_be_valid_if_code_equals_otp_and_is_not_expired(t *testing.T) {
	fixedTime := time.Date(2022, 12, 24, 13, 37, 0, 0, time.UTC)
	clock := func() time.Time {
		return fixedTime
	}
	otp := secret.Otp{
		Expiry: fixedTime.Add(30 * time.Second).Unix(),
		Code:   "given secret",
	}
	validator := secret.NewValidator(clock)

	result := validator.Validate("given secret", otp)

	assert.Equal(t, secret.Valid, result)
}

func TestValidator_Validate_should_be_expired_if_expiry_is_exceeded(t *testing.T) {
	fixedTime := time.Date(2022, 12, 24, 13, 37, 0, 0, time.UTC)
	clock := func() time.Time {
		return fixedTime
	}
	otp := secret.Otp{
		Code:   "code",
		Expiry: fixedTime.Add(-1 * time.Second).Unix(),
	}
	validator := secret.NewValidator(clock)

	result := validator.Validate("does not matter", otp)

	assert.Equal(t, secret.Expired, result)
}

func TestValidator_Validate_should_be_invalid_if_given_code_does_not_match_otp(t *testing.T) {
	fixedTime := time.Date(2022, 12, 24, 13, 37, 0, 0, time.UTC)
	clock := func() time.Time {
		return fixedTime
	}
	otp := secret.Otp{
		Code:   "the secret code",
		Expiry: fixedTime.Add(1 * time.Minute).Unix(),
	}
	validator := secret.NewValidator(clock)

	result := validator.Validate("does not match", otp)

	assert.Equal(t, secret.Invalid, result)
}

func TestValidator_Validate_should_be_invalid_if_code_is_empty(t *testing.T) {
	fixedTime := time.Date(2022, 12, 24, 13, 37, 0, 0, time.UTC)
	clock := func() time.Time {
		return fixedTime
	}
	otp := secret.Otp{
		Code:   "",
		Expiry: fixedTime.Add(1 * time.Minute).Unix(),
	}
	validator := secret.NewValidator(clock)

	result := validator.Validate("", otp)

	assert.Equal(t, secret.Invalid, result)
}

func TestValidator_Validate_should_be_missing_if_otp_is_empty(t *testing.T) {
	fixedTime := time.Date(2022, 12, 24, 13, 37, 0, 0, time.UTC)
	clock := func() time.Time {
		return fixedTime
	}
	validator := secret.NewValidator(clock)

	result := validator.Validate("does not matter", secret.Otp{})

	assert.Equal(t, secret.Missing, result)
}
