package validation

import (
	"fmt"
	"reflect"
)

type Validator struct {
	rules map[string][]ValidationRule
}

type ValidationRule struct {
	Check   func(interface{}) bool
	Message string
}

func NewValidator() *Validator {
	return &Validator{
		rules: make(map[string][]ValidationRule),
	}
}

func (v *Validator) AddRule(field string, rule ValidationRule) {
	v.rules[field] = append(v.rules[field], rule)
}

func (v *Validator) Validate(entity interface{}) error {
	val := reflect.ValueOf(entity)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for field, rules := range v.rules {
		value := val.FieldByName(field)
		if !value.IsValid() {
			return fmt.Errorf("field %s not found in entity", field)
		}

		for _, rule := range rules {
			if !rule.Check(value.Interface()) {
				return fmt.Errorf("validation failed for %s: %s", field, rule.Message)
			}
		}
	}

	return nil
}
