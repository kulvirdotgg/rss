package validator

import (
	"regexp"
	"slices"
)

type Validator struct {
	Errors map[string]string
}

var (
	EmailRGX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) In(value string, list ...string) bool {
	if slices.Contains(list, value) {
		return true
	}
	return false
}

func (v *Validator) Matches(value string, rgx regexp.Regexp) bool {
	return rgx.MatchString(value)
}

func (v *Validator) Unique(values []string) bool {
	unique := make(map[string]struct{})

	for _, value := range values {
		unique[value] = struct{}{}
	}

	return len(unique) == len(values)
}
