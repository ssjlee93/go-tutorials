package greetings

import (
	"regexp"
	"testing"
)

// TestHello calls greetings.Hello with a name,
// checking for a valid return value
func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellos calls greetings.Hellos with a slice of names,
// checking for a valid return value
func TestHellos(t *testing.T) {
	names := []string{"Gladys"}
	want := regexp.MustCompile(`\b` + names[0] + `\b`)
	msgs, err := Hellos(names)
	if !want.MatchString(msgs[names[0]]) {
		t.Fatalf(`Hellos("Gladys") = %q, %v, want match for %#q, nil`, msgs[names[0]], err, want)
	}
}

// TestHellosEmpty calls greetings.Hellos with an empty string,
// checking for an error.
func TestHellosEmpty(t *testing.T) {
	names := []string{""}
	msg, err := Hellos(names)
	if msg != nil || err == nil {
		t.Fatalf(`Hellos([""]) = %q, %v, want "", error`, msg, err)
	}
}
