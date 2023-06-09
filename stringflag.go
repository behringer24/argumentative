package argumentative

import (
	"fmt"
)

// struct for a single configured flag
type StringFlag struct {
	Longflag    string
	Shortflag   string
	Description string
	Required    bool
	Default     string
	Value       *string
}

// Factory to generate a new flag
func NewStringFlag(longflag string, shortflag string, required bool, defaultvalue string, description string) StringFlag {
	flag := StringFlag{
		Longflag:    longflag,
		Shortflag:   shortflag,
		Description: description,
		Required:    required,
		Default:     defaultvalue,
		Value:       new(string),
	}
	if defaultvalue != "" {
		*flag.Value = defaultvalue
	}

	return flag
}

// Generate the string for the long description
func (f *StringFlag) GetLongDescription() string {
	flagnames := ""
	if f.Shortflag != "" {
		flagnames += "-" + f.Shortflag + ", "
	}
	flagnames += "--" + f.Longflag
	output := fmt.Sprintf("%-25s", flagnames)
	if f.Description != "" {
		output += f.Description
	}
	if f.Default != "" {
		output += " (Default: " + f.Default + ")"
	}

	return output
}

// Generate the string for a short description in the 'Usage:' line
func (f *StringFlag) GetShortDescription() string {
	output := " "
	if !f.Required {
		output += "["
	}
	if f.Shortflag != "" {
		output += "-" + f.Shortflag
	} else {
		output += "--" + f.Longflag
	}
	if !f.Required {
		output += "]"
	}
	return output
}
