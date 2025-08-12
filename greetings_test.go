package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Kevin"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Kevin")

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Kevin") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
