package standards

import (
	"testing"
)

// "regexp"

// TestCommand calls standards.FindCommand with a command,
// and checks for no errors
func TestCommand(t *testing.T) {
	// command := "help"
	// want := regexp.MustCompile(`\b` + command + `\b`)
	err := FindCommand("help")
	if err != "nil" {
		t.Fatalf(`FindCommand("help") = %v, want nil`, err)
	}
}

// TestInvalid calls standards.FindCommand with an invalid
// command, checking for an error.
func TestInvalid(t *testing.T) {
	err := FindCommand("DNE")
	if err == "nil" {
		t.Fatalf(`FindCommand("") = %v, want error`, err)
	}
}
