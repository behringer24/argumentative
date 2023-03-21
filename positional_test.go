package argumentative

import "testing"

func TestNewPositional(t *testing.T) {
	flag := NewPositional("longname", true, "default", "description")

	if flag.Longflag != "longname" {
		t.Errorf("Longflag assignment wrong, got [%s], want [%s]", flag.Longflag, "longname")
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
func TestGetLongPositionalDescription(t *testing.T) {
	flag := NewPositional("longname", true, "default", "description")
	result := flag.GetLongDescription()
	await := "longname                 description (Default: default)"

	if result != await {
		t.Errorf("Generation of long description failed, got [%s], want [%s]", result, await)
	}

	flag = NewPositional("longname", false, "default", "description")
	result = flag.GetLongDescription()
	await = "longname                 description (Default: default)"

	if result != await {
		t.Errorf("Generation of long description not required failed, got [%s], want [%s]", result, await)
	}

	flag = NewPositional("longname", false, "", "description")
	result = flag.GetLongDescription()
	await = "longname                 description"

	if result != await {
		t.Errorf("Generation of long description no default failed, got [%s], want [%s]", result, await)
	}
}

func TestGetShortPositionalDescription(t *testing.T) {
	flag := NewPositional("longname", true, "default", "description")
	result := flag.GetShortDescription()
	await := " longname" // @todo: remove space

	if result != await {
		t.Errorf("Generation of short description failed, got [%s], want [%s]", result, await)
	}

	flag = NewPositional("longname", false, "", "description")
	result = flag.GetShortDescription()
	await = " [longname]" // @todo remove space

	if result != await {
		t.Errorf("Generation of short description no required failed, got [%s], want [%s]", result, await)
	}

	flag = NewPositional("longname", false, "default", "description")
	result = flag.GetShortDescription()
	await = " [longname]" // @todo remove space

	if result != await {
		t.Errorf("Generation of short description no required but default failed, got [%s], want [%s]", result, await)
	}
}
