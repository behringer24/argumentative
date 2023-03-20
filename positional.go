package argumentative

import (
	"fmt"
)

// struct for a single configured positional argument
type Positional struct {
	Longflag    string
	Description string
	Required    bool
	Default     string
	Value       *string
}

// Factory to generate a new positional argument
func NewPositional(longflag string, required bool, defaultvalue string, description string) Positional {
	positional := Positional{
		Longflag:    longflag,
		Description: description,
		Required:    required,
		Default:     defaultvalue,
		Value:       new(string),
	}
	if defaultvalue != "" {
		*positional.Value = defaultvalue
	}

	return positional
}

// Generate the string for the long description
func (f *Positional) GetLongDescription() string {
	output := fmt.Sprintf("%-25s", f.Longflag)
	if f.Description != "" {
		output += f.Description
	}
	if f.Default != "" {
		output += " (Default: " + f.Default + ")"
	}
	return output
}

// Generate the string for a short description in the 'Usage:' line
func (f *Positional) GetShortDescription() string {
	output := " "
	if !f.Required {
		output += "["
	}
	output += f.Longflag
	if !f.Required {
		output += "]"
	}
	return output
}
