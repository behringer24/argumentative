package argumentative

import "testing"

func TestNewStringFlag(t *testing.T) {
	flag := NewStringFlag("longname", "s", true, "default", "description")

	if flag.Longflag != "longname" {
		t.Errorf("Longflag assignment wrong, got [%s], want [%s]", flag.Longflag, "longname")
	}

	if flag.Shortflag != "s" {
		t.Errorf("Shortflag assignment wrong, got [%s], want [%s]", flag.Shortflag, "s")
	}

	if flag.Default != "default" {
		t.Errorf("Default assignment wrong, got [%s], want [%s]", flag.Default, "default")
	}

	if flag.Description != "description" {
		t.Errorf("Description assignment wrong, got [%s], want [%s]", flag.Description, "description")
	}

	if *flag.Value != "default" {
		t.Errorf("Assignment of default to value wrong, got [%s], want [%s]", *flag.Value, flag.Default)
	}
}

func TestGetLongDescription(t *testing.T) {
	flag := NewStringFlag("longname", "s", true, "default", "description")
	result := flag.GetLongDescription()
	await := "-s, --longname           description (Default: default)"

	if result != await {
		t.Errorf("Generation of long description failed, got [%s], want [%s]", result, await)
	}

	flag = NewStringFlag("longname", "s", true, "", "description")
	result = flag.GetLongDescription()
	await = "-s, --longname           description"

	if result != await {
		t.Errorf("Generation of long description without default failed, got [%s], want [%s]", result, await)
	}

	flag = NewStringFlag("longname", "s", false, "", "description")
	result = flag.GetLongDescription()
	await = "-s, --longname           description"

	if result != await {
		t.Errorf("Generation of long description without default failed, got [%s], want [%s]", result, await)
	}

	flag = NewStringFlag("longname", "", true, "", "description")
	result = flag.GetLongDescription()
	await = "--longname               description"

	if result != await {
		t.Errorf("Generation of long description without shortname failed, got [%s], want [%s]", result, await)
	}
}

func TestGetShortDescription(t *testing.T) {
	flag := NewStringFlag("longname", "s", true, "default", "description")
	result := flag.GetShortDescription()
	await := " -s" // @todo: remove space

	if result != await {
		t.Errorf("Generation of short description failed, got [%s], want [%s]", result, await)
	}

	flag = NewStringFlag("longname", "s", false, "", "description")
	result = flag.GetShortDescription()
	await = " [-s]" // @todo remove space

	if result != await {
		t.Errorf("Generation of short description no required failed, got [%s], want [%s]", result, await)
	}

	flag = NewStringFlag("longname", "", false, "", "description")
	result = flag.GetShortDescription()
	await = " [--longname]" // @todo remove space

	if result != await {
		t.Errorf("Generation of short description no short name failed, got [%s], want [%s]", result, await)
	}
}
