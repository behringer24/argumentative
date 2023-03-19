package argumentative

import (
	"fmt"
)

type BoolFlag struct {
	Longflag    string
	Shortflag   string
	Description string
	Required    bool
	Value       *bool
}

func NewBoolFlag(longflag string, shortflag string, description string) BoolFlag {
	flag := BoolFlag{
		Longflag:    longflag,
		Shortflag:   shortflag,
		Description: description,
		Value:       new(bool),
	}
	*flag.Value = false

	return flag
}

func (f *BoolFlag) GetLongDescription() string {
	flagnames := ""
	if f.Shortflag != "" {
		flagnames += "-" + f.Shortflag + ", "
	}
	flagnames += "--" + f.Longflag
	output := fmt.Sprintf("%-25s", flagnames)
	if f.Description != "" {
		output += f.Description
	}
	return output
}

func (f *BoolFlag) GetShortDescription() string {
	output := " ["
	if f.Shortflag != "" {
		output += "-" + f.Shortflag
	} else {
		output += "--" + f.Longflag
	}
	output += "]"
	return output
}
