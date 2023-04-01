// Copyright 2023 The Argumentative Authors
//
// Licensed under the GNU GENERAL PUBLIC LICENSE, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://fsf.org/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package argumentative

import (
	"fmt"
)

// struct with all maps that hold the different flag types
type Flags struct {
	boolflags   map[string]BoolFlag
	stringflags map[string]StringFlag
	positionals []Positional

	shortflags map[byte]string
}

// constructor like chain command to init all maps
func (f *Flags) Flags() *Flags {
	if f.stringflags == nil {
		f.boolflags = make(map[string]BoolFlag)
		f.stringflags = make(map[string]StringFlag)
		f.shortflags = make(map[byte]string)
	}

	return f
}

// Add string type flag to map and return pointer to value
func (f *Flags) AddString(longflag string, shortflag string, required bool, defaultvalue string, description string) *string {
	f.stringflags[longflag] = NewStringFlag(longflag, shortflag, required, defaultvalue, description)
	if shortflag != "" {
		f.shortflags[shortflag[0]] = longflag
	}
	return f.stringflags[longflag].Value
}

// Add boolean type flag to map and return pointer to value
func (f *Flags) AddBool(longflag string, shortflag string, description string) *bool {
	f.boolflags[longflag] = NewBoolFlag(longflag, shortflag, description)
	if shortflag != "" {
		f.shortflags[shortflag[0]] = longflag
	}
	return f.boolflags[longflag].Value
}

// Add positional argument to map and return pointer to value
func (f *Flags) AddPositional(longflag string, required bool, defaultvalue string, description string) *string {
	f.positionals = append(f.positionals, NewPositional(longflag, required, defaultvalue, description))
	return f.positionals[len(f.positionals)-1].Value
}

// Check if argument is a flag or positional argument
func (f *Flags) isFlag(name string) bool {
	return name[0] == '-'
}

// Check if argument is a long flag
func (f *Flags) isLongFlag(name string) bool {
	return len(name) > 1 && name[0] == '-' && name[1] == '-'
}

// Get name of flag and translate short flags to their long equivalent
func (f *Flags) GetFlagName(name string, pos int) string {
	var longname string

	if len(name) > 1 && name[0] == '-' {
		if len(name) > 2 && name[1] == '-' {
			longname = name[2:]
		} else {
			if pos < len(name) {
				longname = f.shortflags[name[pos]]
			} else {
				return ""
			}
		}
		return longname
	}
	return ""
}

// Validate the parameters and check if all required parameters have a value
func (f *Flags) Validate() (err error) {
	for _, flag := range f.stringflags {
		if flag.Required && *flag.Value == "" {
			return fmt.Errorf("required flag --%s missing", flag.Longflag)
		}
	}
	for _, positional := range f.positionals {
		if positional.Required && *positional.Value == "" {
			return fmt.Errorf("required positional argument [%s] missing", positional.Longflag)
		}
	}
	return nil
}

// Parse arguments
func (f *Flags) Parse(args []string) (err error) {
	positional := 0
	i := 1 // leave out the first one as this is usually the (cli-) command itself
	for i < len(args) {
		if f.isFlag(args[i]) {
			// Parse flags with string values
			if _, ok := f.stringflags[f.GetFlagName(args[i], 1)]; ok {
				if !f.Flags().isLongFlag(args[i]) && len(args[i]) > 2 {
					return fmt.Errorf("options with parameters can not be combined %s", args[i])
				}
				*f.stringflags[f.GetFlagName(args[i], 1)].Value = args[i+1]
				i += 1
			} else {
				// Parse flags the switch to true if exists, allow "-xvzf" as combinations
				if f.isLongFlag(args[i]) {
					if _, ok := f.boolflags[f.GetFlagName(args[i], 1)]; ok {
						*f.boolflags[f.GetFlagName(args[i], 1)].Value = true
					} else {
						return fmt.Errorf("unknown flag %s", args[i])
					}
				} else {
					for j := 1; j < len(args[i]); j++ {
						if _, ok := f.boolflags[f.GetFlagName(args[i], j)]; ok {
							*f.boolflags[f.GetFlagName(args[i], j)].Value = true
						} else {
							if _, ok := f.stringflags[f.GetFlagName(args[i], j)]; ok {
								return fmt.Errorf("options with parameters can not be combined: %c in %s", args[i][j], args[i])
							} else {
								return fmt.Errorf("unknown flag -%c", args[i][j])
							}
						}
					}
				}
			}
			// Parse positional arguments sequentially while there are unset ones
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

// Print usage instructions
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
