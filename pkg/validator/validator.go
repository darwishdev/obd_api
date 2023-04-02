package validator

import (
	"net/mail"
	"regexp"
)

var (
	isValidName  = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidPhone = regexp.MustCompile(`^[+]?[0-9]{8,}$`).MatchString
)

func ValidatePrice(value float32) error {
	if value < 0.1 || value > 50000 {
		return ErrorMinMax(float32(0.1), float32(50000))
	}

	return nil
}
func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return ErrorMinMax(float32(minLength), float32(maxLength))
	}
	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return ErrorInvalid()
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 200)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return ErrorInvalid()
	}
	return nil
}

func ValidatePhone(value string) error {

	if !isValidPhone(value) {
		return ErrorInvalid()
	}
	return nil

}
