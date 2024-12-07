package flag

import (
	"flag"
)

// ModeFlag is a custom flag type for mode
type ModeFlag struct {
	Flag
}

// NewModeFlag creates a new ModeFlag with allowed values
func NewModeFlag(defaultValue string, allowed []string) *ModeFlag {
	return &ModeFlag{
		Flag: Flag{
			value:   defaultValue,
			allowed: allowed,
		},
	}
}

// IsDev returns true if the mode is development
func (m *ModeFlag) IsDev() bool {
	return m.value == ModeDev
}

// IsProd returns true if the mode is production
func (m *ModeFlag) IsProd() bool {
	return m.value == ModeProd
}

const (
	// ModeDev is the development mode
	ModeDev = "dev"

	// ModeProd is the production mode
	ModeProd = "prod"
)

// Mode is the environment mode
var Mode = NewModeFlag(ModeDev, []string{ModeDev, ModeProd})

// SetModeFlag sets the mode flag
func SetModeFlag() {
	flag.Var(
		Mode,
		"m",
		"Specify mode. Allowed values are: dev, prod. Default is the development mode",
	)
}
