package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Moms")

	got := buffer.String()
	expected := "Hello, Moms"

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}
