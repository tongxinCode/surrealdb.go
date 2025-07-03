package models

import (
	"reflect"
	"time"
)

func replacerBeforeEncode(value interface{}) interface{} {
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()

	// TODO:when to use this?
	if valueType == reflect.TypeOf(time.Duration(0)) {
		oldVal := value.(time.Duration)
		newValue := CustomDuration{oldVal}
		return newValue
	}

	// TODO:when to use this?
	if valueType == reflect.TypeOf(time.Time{}) {
		oldVal := value.(time.Time)
		newValue := CustomDateTime{oldVal}
		return newValue
	}

	// TODO:when to use this?
	if valueKind == reflect.Map {
		oldValue := value.(map[string]interface{})
		newValue := make(map[interface{}]interface{})
		for k, v := range oldValue {
			newKey := replacerBeforeEncode(k)
			newVal := replacerBeforeEncode(v)
			newValue[newKey] = newVal
		}

		return newValue
	}

	// todo: handle slices

	// todo: handle structs

	return value
}

func replacerAfterDecode(value interface{}) interface{} {
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()

	// TODO:when to use this?
	if valueType == reflect.TypeOf(CustomDuration{}) {
		oldVal := value.(CustomDuration)
		newValue := oldVal.Duration
		return newValue
	}

	// TODO:when to use this?
	if valueType == reflect.TypeOf(CustomDateTime{}) {
		oldVal := value.(CustomDateTime)
		newValue := oldVal.Time
		return newValue
	}

	// TODO:when to use this?
	if valueKind == reflect.Map {
		oldValue := value.(map[string]interface{})
		newValue := make(map[interface{}]interface{})
		for k, v := range oldValue {
			newKey := replacerAfterDecode(k)
			newVal := replacerAfterDecode(v)
			newValue[newKey] = newVal
		}

		return newValue
	}

	// todo: handle slices

	// todo: handle structs

	return value
}
