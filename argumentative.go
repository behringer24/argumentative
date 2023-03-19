package argumentative

import (
	"fmt"
)

type Flags struct {
	boolflags   map[string]BoolFlag
	stringflags map[string]StringFlag
	positionals []Positional

	shortflags map[byte]string
}

func (f *Flags) Flags() *Flags {
	if f.stringflags == nil {
		f.boolflags = make(map[string]BoolFlag)
		f.stringflags = make(map[string]StringFlag)
		f.shortflags = make(map[byte]string)
	}

	return f
}

func (f *Flags) AddString(longflag string, shortflag string, required bool, defaultvalue string, description string) *string {
	f.stringflags[longflag] = NewStringFlag(longflag, shortflag, required, defaultvalue, description)
	if shortflag != "" {
		f.shortflags[shortflag[0]] = longflag
	}
	return f.stringflags[longflag].Value
}

func (f *Flags) AddBool(longflag string, shortflag string, description string) *bool {
	f.boolflags[longflag] = NewBoolFlag(longflag, shortflag, description)
	if shortflag != "" {
		f.shortflags[shortflag[0]] = longflag
	}
	return f.boolflags[longflag].Value
}

func (f *Flags) AddPositional(longflag string, required bool, defaultvalue string, description string) *string {
	f.positionals = append(f.positionals, NewPositional(longflag, required, defaultvalue, description))
	return f.positionals[len(f.positionals)-1].Value
}

func (f *Flags) isFlag(name string) bool {
	return name[0] == '-'
}

func (f *Flags) GetFlagName(name string) string {
	var longname string
	if len(name) > 1 && name[0] == '-' {
		if len(name) > 2 && name[1] == '-' {
			longname = name[2:]
		} else {
			longname = f.shortflags[name[1]]
		}
		return longname
	}
	return ""
}

func (f *Flags) Validate() (err error) {
	for _, flag := range f.stringflags {
		if flag.Required && *flag.Value == "" {
			return fmt.Errorf("flag %s missing", flag.Longflag)
		}
	}
	for _, positional := range f.positionals {
		if positional.Required && *positional.Value == "" {
			return fmt.Errorf("positional argument %s missing", positional.Longflag)
		}
	}
	return nil
}

func (f *Flags) Parse(args []string) (err error) {
	positional := 0
	i := 1
	for i < len(args) {
		if f.isFlag(args[i]) {
			// Parse flags with string values
			if _, ok := f.stringflags[f.GetFlagName(args[i])]; ok {
				*f.stringflags[f.GetFlagName(args[i])].Value = args[i+1]
				i += 1
			} else
			// Parse flags the switch to true if exists
			if _, ok := f.boolflags[f.GetFlagName(args[i])]; ok {
				*f.boolflags[f.GetFlagName(args[i])].Value = true
			} else {
				return fmt.Errorf("unknown flag %s", args[i])
			}
		} else if positional < len(f.positionals) {
			*f.positionals[positional].Value = args[i]
			positional += 1
		} else {
			return fmt.Errorf("unknown positional argument %s", args[i])
		}
		i += 1
	}
	return f.Validate()
}

func (f *Flags) Usage(name string, description string, err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(name)
		fmt.Println(description)
	}
	output := "\nUsage: " + name
	if len(f.stringflags) > 0 || len(f.boolflags) > 0 {
		for _, flag := range f.boolflags {
			output += flag.GetShortDescription()
		}
		for _, flag := range f.stringflags {
			output += flag.GetShortDescription()
		}
	}
	if len(f.positionals) > 0 {
		for _, positional := range f.positionals {
			output += positional.GetShortDescription()
		}
	}

	fmt.Println(output)

	if len(f.boolflags) > 0 {
		fmt.Println("\nFlags:")
		for _, flag := range f.boolflags {
			fmt.Println(flag.GetLongDescription())
		}
	}

	if len(f.stringflags) > 0 {
		fmt.Println("\nOptions:")
		for _, flag := range f.stringflags {
			fmt.Println(flag.GetLongDescription())
		}
	}

	if len(f.positionals) > 0 {
		fmt.Println("\nPositional arguments:")
		for _, positional := range f.positionals {
			fmt.Println(positional.GetLongDescription())
		}
	}

}
