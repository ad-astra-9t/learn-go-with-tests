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
		got := Hello("Tim", "en")
		want := "Hello, Tim"
		assertMessage(t, want, got)
	})
	t.Run("say hello world", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"
		assertMessage(t, want, got)
	})
	t.Run("say hello to Elodie in French", func(t *testing.T) {
		got := Hello("Elodie", "fr")
		want := "Bonjour, Elodie"
		assertMessage(t, want, got)
	})
	t.Run("say hello to Ronny in Chinese", func(t *testing.T) {
		got := Hello("Ronny", "zh")
		want := "你好，Ronny"
		assertMessage(t, want, got)
	})
}
