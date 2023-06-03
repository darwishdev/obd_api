package validator

import (
	"math"
	"net/mail"
	"regexp"
	"time"
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

func ValidateUnsignedInt(value int64) error {
	if value < 0 {
		return ErrorInvalid()
	}
	return nil
}

func ValidateUnsignedFloat(float float32) error {
	if float < 0 {
		return ErrorInvalid()
	}
	return nil
}

func ValidateInt(value int64, minLength int64, maxLength int64) error {
	if value > maxLength || value < minLength {
		return ErrorMinMax(float32(minLength), float32(maxLength))
	}
	return nil
}
func ValidateDate(date time.Time, minYear int, maxYear int) bool {
	minDate := time.Date(minYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(maxYear, time.December, 31, 23, 59, 59, 0, time.UTC)
	return !date.Before(minDate) && !date.After(maxDate)
}

func ValidateLatLng(lat, lng float32) error {
	if lat < -90 || lat > 90 {
		return ErrorInvalidLocation()
	}
	if lng < -180 || lng > 180 {
		return ErrorInvalidLocation()
	}
	if math.IsNaN(float64(lat)) || math.IsNaN(float64(lng)) {
		return ErrorInvalidLocation()
	}
	return nil
}
