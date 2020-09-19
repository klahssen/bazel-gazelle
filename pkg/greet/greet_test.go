package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buf bytes.Buffer

	Greet(&buf, "James")

	got := buf.String()
	want := "Hello James !\n"

	if got != want {
		t.Errorf("expected '%s' received '%s'", want, got)
	}
}
