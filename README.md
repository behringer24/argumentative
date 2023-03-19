# argumentative
Go argument parser fast and simple

## Why
I experimented with a lot of Go command line argument parsers to build simple but powerful cli applications. I like thge way that pythons argparse module works but did not find a suitable replacemant for my current Go projects. Some seem to be abandoned and even the big framework like tools did not handle positional arguments like I would like them to work. So this is because and because: Why not.

## Installation
```
go get -u -v https://github.com/behringer24/argumentative
```

### Dependencies
There are no dependencies beside the "fmt", so thats a "no dependencies".

## Usage
Example 'argtest.go' file
``` Golang
package main

import (
	"fmt"
	"os"

	"github.com/behringer24/argumentative"
)

const (
	title       = "argtest"
	description = "A small demonstration"
	version     = "v0.0.1"
)

var (
	test1       *string
	test2       *string
	inFileName  *string
	outFileName *string
	showHelp    *bool
	showVer     *bool
)

func parseArgs() {
	flags := &argumentative.Flags{}
	test1 = flags.Flags().AddString("test", "t", true, "", "A small description of the test string flag")
	test2 = flags.Flags().AddString("norequired", "n", false, "norequireddefault", "Another description")
	showHelp = flags.Flags().AddBool("help", "h", "Show this help text")
	showVer = flags.Flags().AddBool("version", "", "Show version information")
	inFileName = flags.Flags().AddPositional("infile", true, "", "File to read from")
	outFileName = flags.Flags().AddPositional("outfile", false, "", "File to write to")

	err := flags.Parse(os.Args)
	if *showHelp {
		flags.Usage(title, description, nil)
		os.Exit(0)
	} else if *showVer {
		fmt.Print("argtest version " + version)
		os.Exit(0)
	} else if err != nil {
		flags.Usage("argtest", "A small demonstartion", err)
		os.Exit(1)
	}
}

func main() {
	parseArgs()

	fmt.Println("\nResult:", *test1, *test2, *inFileName, *outFileName, *showHelp, *showVer)
}
```

In the var block all the parameter holders are defined. argumentative returns pointers to the internal value fields and fills them after the ``flags.Parse(os.Args)`` call.

Parsing of the parameters has been move out of ``main()`` into its own function.

There are also the definitions of all available parameters and arguments.

```
go run .\argtest.go -h
```
result in the output of
```
argtest
A small demonstration

Usage: argtest [-h] [--version] [-n] -t infile [outfile]

Flags:
-h, --help               Show this help text
--version                Show version information

Options:
-t, --test               A small description of the test string flag
-n, --norequired         Another description (Default: norequireddefault)

Positional arguments:
infile                   File to read from
outfile                  File to write to
```