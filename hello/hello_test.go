package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tim")
	want := "Hello, Tim"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
