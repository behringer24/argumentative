package argumentative

import (
	"fmt"
)

type Positional struct {
	Longflag    string
	Description string
	Required    bool
	Default     string
	Value       *string
}

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
