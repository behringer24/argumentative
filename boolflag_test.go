package argumentative

import "testing"

func TestNewBoolFlag(t *testing.T) {
	flag := NewBoolFlag("longname", "s", "description")

	if flag.Longflag != "longname" {
		t.Errorf("Longflag assignment wrong, got [%s], want [%s]", flag.Longflag, "longname")
	}

	if flag.Shortflag != "s" {
		t.Errorf("Shortflag assignment wrong, got [%s], want [%s]", flag.Shortflag, "s")
	}

	if flag.Description != "description" {
		t.Errorf("Description assignment wrong, got [%s], want [%s]", flag.Description, "description")
	}

	if *flag.Value != false {
		t.Errorf("Assignment of default to value wrong, got [%t], want [%t]", *flag.Value, false)
	}
}

func TestGetLongBoolDescription(t *testing.T) {
	flag := NewBoolFlag("longname", "s", "description")
	result := flag.GetLongDescription()
	await := "-s, --longname           description"

	if result != await {
		t.Errorf("Generation of long description failed, got [%s], want [%s]", result, await)
	}

	flag = NewBoolFlag("longname", "", "description")
	result = flag.GetLongDescription()
	await = "--longname               description"

	if result != await {
		t.Errorf("Generation of long description without default failed, got [%s], want [%s]", result, await)
	}
}

func TestGetShortBoolDescription(t *testing.T) {
	flag := NewBoolFlag("longname", "s", "description")
	result := flag.GetShortDescription()
	await := " [-s]" // @todo: remove space

	if result != await {
		t.Errorf("Generation of short description failed, got [%s], want [%s]", result, await)
	}

	flag = NewBoolFlag("longname", "", "description")
	result = flag.GetShortDescription()
	await = " [--longname]" // @todo remove space

	if result != await {
		t.Errorf("Generation of short description no required failed, got [%s], want [%s]", result, await)
	}
}
