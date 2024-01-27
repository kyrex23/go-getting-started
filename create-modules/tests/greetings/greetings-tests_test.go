package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking for a valid return type
func TestHelloName(t *testing.T) {
	name := "Kyrex"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, message, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string, checking for an error
func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
