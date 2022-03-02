package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want) // %q wraps value in double quotes
	}

}
