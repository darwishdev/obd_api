package util

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	// Get the struct value from the interface
	structValue := reflect.ValueOf(obj).Elem()

	// Get the field value by name
	fieldValue := structValue.FieldByName(name)

	// Check if the field exists
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	// Get the type of the field
	fieldType := fieldValue.Type()

	// Get the value to be set as a reflect.Value
	newValue := reflect.ValueOf(value)

	// Check if the type of the value to be set matches the type of the field
	if !newValue.Type().AssignableTo(fieldType) {
		return fmt.Errorf("value type %v is not assignable to field type %v", newValue.Type(), fieldType)
	}

	// Set the field value to the new value
	fieldValue.Set(newValue)

	return nil
}

func ConvertStruct(src interface{}, dst interface{}) error {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Struct {
		return errors.New("src must be a struct")
	}

	if dstValue.Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		return errors.New("dst must be a pointer to a struct")
	}

	srcType := srcValue.Type()
	dstType := dstValue.Type().Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcFieldName := srcType.Field(i).Name
		dstField, ok := dstType.FieldByName(srcFieldName)

		if ok {
			dstFieldType := dstField.Type
			srcFieldValue := srcValue.Field(i)
			dstFieldValue := dstValue.Elem().FieldByName(srcFieldName)

			if !dstFieldValue.IsValid() {
				return fmt.Errorf("dst does not have a field named %s", srcFieldName)
			}

			if dstFieldType == srcFieldValue.Type() {
				dstFieldValue.Set(srcFieldValue)
			} else {
				return fmt.Errorf("field %s: cannot convert from %s to %s", srcFieldName, srcFieldValue.Type(), dstFieldType)
			}
		}
	}

	return nil
}
