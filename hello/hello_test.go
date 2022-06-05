package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, want, got interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("say hello to Tim", func(t *testing.T) {
		got := Hello("Tim")
		want := "Hello, Tim"
		assertMessage(t, want, got)
	})
	t.Run("say hello world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertMessage(t, want, got)
	})
}
