package util

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// checks if a given string str is present in an array of strings arr.
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(a, str) {
			return true
		}
	}
	return false
}

func ToNullString(input string) sql.NullString {
	if input == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: input, Valid: true}
}
func SetStructOverrides(obj interface{}, overrides map[string]interface{}) error {
	reflectObj := reflect.ValueOf(obj).Elem()

	for overrideFieldName, overrideValue := range overrides {
		field := reflectObj.FieldByName(overrideFieldName)

		if !field.IsValid() {
			return fmt.Errorf("field '%s' not found in struct", overrideFieldName)
		}

		if !field.CanSet() {
			return fmt.Errorf("field '%s' cannot be set", overrideFieldName)
		}

		overrideValueReflect := reflect.ValueOf(overrideValue)

		if field.Type() != overrideValueReflect.Type() {
			return fmt.Errorf("field '%s' has type '%s' but override value has type '%s'", overrideFieldName, field.Type(), overrideValueReflect.Type())
		}

		field.Set(overrideValueReflect)
	}

	return nil
}
func ToNullInt64(input int64) sql.NullInt64 {
	if input == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: input, Valid: true}
}
