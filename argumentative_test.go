package argumentative

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestFlagsIntegration(t *testing.T) {
	flags := &Flags{}
	stringflag := flags.Flags().AddString("stringname", "s", true, "", "stringdescription")
	boolflag := flags.Flags().AddBool("boolname", "b", "booldescription")
	positional := flags.Flags().AddPositional("positionalname", false, "positionaldefault", "positionaldescription")
	optx := flags.Flags().AddBool("optx", "x", "Option X")
	opty := flags.Flags().AddBool("opty", "y", "Option Y")
	optz := flags.Flags().AddBool("optz", "z", "Option Z")

	var args []string

	args = append(args, "scriptname")
	err := flags.Parse(args)

	await := "required flag --stringname missing"

	if err == nil {
		t.Errorf("No error found, got [%p], want pointer", err)
	} else if err.Error() != await {
		t.Errorf("Wrong error message, got [%s], want [%s]", err, await)
	}

	if *stringflag != "" {
		t.Errorf("Wrong stringflag value, got [%s], want [%s]", *stringflag, "")
	}

	if *boolflag != false {
		t.Errorf("Wrong boolflag value, got [%t], want [%t]", *boolflag, false)
	}

	if *positional != "positionaldefault" {
		t.Errorf("Wrong positional value, got [%s], want [%s]", *positional, "positionaldefault")
	}

	args = append(args, "-s", "stringvalue", "-b", "positionalvalue")
	err = flags.Parse(args)

	await = "required flag --stringname missing"

	if err != nil {
		t.Errorf("Error found, got [%s], want nil", err.Error())
	}

	if *stringflag != "stringvalue" {
		t.Errorf("Wrong stringflag value, got [%s], want [%s]", *stringflag, "stringvalue")
	}

	if *boolflag != true {
		t.Errorf("Wrong boolflag value, got [%t], want [%t]", *boolflag, false)
	}

	if *positional != "positionalvalue" {
		t.Errorf("Wrong positional value, got [%s], want [%s]", *positional, "positionalvalue")
	}

	args = append(args, "-xy")
	err = flags.Parse(args)

	if err != nil {
		t.Errorf("Error found, got [%s], want nil", err.Error())
	}

	if !*optx {
		t.Errorf("Wrong combined boolflag value, got [%t], want [%t]", *optx, true)
	}

	if !*opty {
		t.Errorf("Wrong combined boolflag value, got [%t], want [%t]", *opty, true)
	}

	args = append(args, "-sz")
	err = flags.Parse(args)
	await = "options with parameters can not be combined -sz"

	if err == nil {
		t.Errorf("No error found, got [%p], want pointer", err)
	} else if err.Error() != await {
		t.Errorf("Wrong error message, got [%s], want [%s]", err, await)
	}

	if *optz {
		t.Errorf("Wrong combined boolflag value, got [%t], want [%t]", *optz, false)
	}
}

func TestUsage(t *testing.T) {
	flags := &Flags{}
	flags.Flags().AddString("stringname", "s", true, "", "stringdescription")
	flags.Flags().AddBool("boolname", "b", "booldescription")
	flags.Flags().AddPositional("positionalname", false, "positionaldefault", "positionaldescription")

	await := `title
description

Usage: title [-b] -s [positionalname]

Flags:
-b, --boolname           booldescription

Options:
-s, --stringname         stringdescription

Positional arguments:
positionalname           positionaldescription (Default: positionaldefault)
`

	result := captureOutput(func() {
		flags.Usage("title", "description", nil)
	})

	if result != await {
		t.Errorf("Wrong Usage output, got\n%s\n\nwant\n\n%s", result, await)
	}
}
