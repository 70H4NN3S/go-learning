package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Johannes")
	want := "Hello, Johannes"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
