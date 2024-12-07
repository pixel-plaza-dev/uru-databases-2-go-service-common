package flag

import (
	"fmt"
	"strings"
)

// Flag is a custom flag type
type Flag struct {
	value   string
	allowed []string
}

// String returns the string representation of the flag value
func (f *Flag) String() string {
	return f.value
}

// Set validates and sets the flag value
func (f *Flag) Set(value string) error {
	for _, v := range f.allowed {
		if value == v {
			f.value = value
			return nil
		}
	}
	return fmt.Errorf(
		"invalid value %q, allowed values are: %s", value,
		strings.Join(f.allowed, ", "),
	)
}
