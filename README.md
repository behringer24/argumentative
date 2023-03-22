# argumentative
Go argument parser fast and simple

[![Build + Test](https://github.com/behringer24/argumentative/actions/workflows/go.yml/badge.svg)](https://github.com/behringer24/argumentative/actions/workflows/go.yml)

## Why
I experimented with a lot of Go command line argument parsers to build simple but powerful cli applications. I like the way that pythons _argparse_ module works, but did not find a suitable replacemant for my current Go projects. Some seem to be abandoned and even the big framework-like tools did not handle positional arguments like I would like them to work. So this is because and because: Why not?

## Installation
```
go get -u -v github.com/behringer24/argumentative
```

### Dependencies
There are no dependencies beside the "fmt", so thats a "no dependencies".

## Usage
The simple structure of the argumentative package, that is required to get your parameters is:

* init a flags object from the package
* define you flags
* call the Parse method to get the values
* handle the errors or override flags and display the usage instructions

_Overriding flags_ are flags that should work even when other errors where happening. The most popular examples are `--help|-h` or `--version|-v` that should display results and make a clean exit even if required arguments are missing (The user wants to learn about these by calling `-h` for example)

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
		fmt.Print(title, "version", version)
		os.Exit(0)
	} else if err != nil {
		flags.Usage(title, description, err)
		os.Exit(1)
	}
}

func main() {
	parseArgs()

	fmt.Println("\nResult:", *test1, *test2, *inFileName, *outFileName, *showHelp, *showVer)
}
```

In the var block all the parameter holders are defined. argumentative returns pointers to the internal value fields and fills them after the ``flags.Parse(os.Args)`` call.

Parsing of the parameters has been move out of ``main()`` into its own 'parseArgs' function.

There are also the definitions of all available parameters and arguments for this demo app.

```
go run .\argtest.go -h
```
result in the output:
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
omitting a required parameter results in an error message. In this case as an example we omit the required `-t` parameter:
```
go run .\argtest.go foo
```
results in the output:
```
Error: required flag --test missing

Usage: argtest [-h] [--version] -t [-n] infile [outfile]

Flags:
-h, --help               Show this help text
--version                Show version information

Options:
-t, --test               A small description of the test string flag
-n, --norequired         Another description (Default: norequireddefault)

Positional arguments:
infile                   File to read from
outfile                  File to write to
exit status 1
```
Notice that we satisfied the required parameter `infile` with `foo` in the above example.

If we add a parameter `--version` or `--help` the output of the errors of the missing parameters is supressed. This behavior has to be handled in the application code and is _not_ part of the argumentative lib.

## Add Parameters to your cli app
### Add boolean parameter
Boolean parameters are simple switches that return true if they are present and false if they are omitted. They do not support a default value or a required flag.

``` Golang
var result *bool

flags := &argumentative.Flags{}
result = flags.Flags().AddBool("longname", "short", "Descriptive help text")
```

`longname` is the long version of the parameter and must not contain space chars. On the commandline the long parameter will be preceeded by two dashes '--'

`short` the one character short version of the parameter, like the `-h` for `--help`. This can only be one character. The short version is optional (see the above example vor `--version`). The short version is preceeded with one dash '-'.

`Descriptive help text` A brief description of the parameter. Keep it short and simple, may be omitted but why should you?

`.Flags{}.` in the above call makes sure, that the internal flags storages are initialized and is simply chained.

### Add string parameter
String parameters are like boolean parameters but always contain an additional value after the parameter seperated with a space char from the short or long parameter name. They return a string, can be optional or required and may contain a default value (which makes switching to 'required' obsolete).

``` Golang
var result *string

flags := &argumentative.Flags{}
result = flags.Flags().AddString("longname", "short", required, "default", "Descriptive help text")
```

`longname` is the long version of the parameter and must not contain space chars. On the commandline the long parameter will be preceeded by two dashes '--'

`short` the one character short version of the parameter, like the `-h` for `--help`. This can only be one character. The short version is optional (see the above example vor `--version`). The short version is preceeded with one dash '-'.

`required` a boolean _true_ or _false_ if this parameter is required (true). If this parameter is omitted it will stop execution and display an error text and the usage instructions.

`default` The default value of the parameter (if not present in the cli call). If there is a default value it makes no sense to set the parameter to required, too. 

`Descriptive help text` A brief description of the parameter. Keep it short and simple, may be omitted but why should you?

### Positional arguments
Positional arguments are parameters without a short or long name and come in a specific order, the order you defined them in. They can be required or have a default value. They return a string. For displaying the arguments in the help text, they require a `longname`.

``` Golang
var result *string

flags := &argumentative.Flags{}
result = flags.Flags().AddPositional("longname", required, "default", "Descriptive help text")
```

`longname` is the name of the parameter and must not contain space chars. The name will be shown on the Usage instruction line and in the brief description section of the help text.

`required` a boolean _true_ or _false_ if this parameter is required (true). If this parameter is omitted it will stop execution and display an error text and the usage instructions. 

`default` The default value of the parameter (if not present in the cli call). If there is a default value it makes no sense to set the parameter to required, too.

`Descriptive help text` A brief description of the parameter. Keep it short and simple, may be omitted but why should you?

Consider the order of positional arguments in your command line. Optional arguments must come last as they would be confused with other arguments. Required arguments must come first. If you are struggling consider to use named string flags.
