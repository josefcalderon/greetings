package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Josef"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Josef")
	if !want.MatchString(msg) || err != nil { // Analisa el mensaje si coincide
		t.Fatalf(`Hello("Josef") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
