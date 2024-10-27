package flags

import (
	"flag"
	"fmt"
	"strings"
)

// ModeFlag is a custom flag type for mode
type ModeFlag struct {
	value   string
	allowed []string
}

// String returns the string representation of the flag value
func (m *ModeFlag) String() string {
	return m.value
}

// Set validates and sets the flag value
func (m *ModeFlag) Set(value string) error {
	for _, v := range m.allowed {
		if value == v {
			m.value = value
			return nil
		}
	}
	return fmt.Errorf("invalid value %q, allowed values are: %s", value, strings.Join(m.allowed, ", "))
}

// NewModeFlag creates a new ModeFlag with allowed values
func NewModeFlag(defaultValue string, allowed []string) *ModeFlag {
	return &ModeFlag{
		value:   defaultValue,
		allowed: allowed,
	}
}

const (
	// ModeDev is the development mode
	ModeDev = "dev"

	// ModeProd is the production mode
	ModeProd = "prod"
)

// EnvironmentMode is the environment mode
var EnvironmentMode = NewModeFlag(ModeDev, []string{ModeDev, ModeProd})

// SetModeFlag sets the mode flag
func SetModeFlag() {
	flag.Var(EnvironmentMode, "m", "Specify mode. Allowed values are: dev, prod. Default is the development mode")
}
