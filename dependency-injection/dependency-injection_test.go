package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Prajjwol")

	got := buffer.String()
	want := "Hello, Prajjwol"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
