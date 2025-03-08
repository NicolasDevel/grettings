package grettings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Nicolas"
	wan := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !wan.MatchString(msg) || err != nil {
		t.Fatalf(`Hello ("juan") = %q, %v, quiere como coincidencia para %#q, nil`, msg, err, wan)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""

	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Fatalf(`Hello ("") = %q, %v, quiere "", error`, msg, err)
	}

}
