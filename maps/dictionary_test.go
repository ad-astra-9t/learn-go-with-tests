package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search a word", func(t *testing.T) {
		word := "test"
		definition := "test dictionary"
		dictionary := Dictionary{word: definition}
		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatalf("unexpected error is returned: %#v", err)
		}
		want := "test dictionary"
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("search an unexisting word", func(t *testing.T) {
		word := "unexisting"
		dictionary := Dictionary{
			"test": "test dictionary",
		}
		_, err := dictionary.Search(word)
		if err == nil {
			t.Errorf("no error is returned")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		word := "test"
		definition := "test dictionary"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		if err != nil {
			t.Fatalf("unexpected error is returned: %#v", err)
		}
		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatalf("unexpected error is returned: %#v", err)
		}
		want := definition
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "test dictionary"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		if err == nil {
			t.Errorf("no error is returned")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update a word", func(t *testing.T) {
		word := "test"
		definition := "test dictionary"
		dictionary := Dictionary{word: definition}
		newDefinition := "new test dictionary"
		err := dictionary.Update("test", newDefinition)
		if err != nil {
			t.Fatalf("unexpected error is returned: %#v", err)
		}
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatalf("unexpected error is returned: %#v", err)
		}
		want := newDefinition
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("update an unexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "test dictionary")
		if err == nil {
			t.Errorf("no error is returned")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete a word", func(t *testing.T) {
		word := "test"
		definition := "test dictionary"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		if err == nil {
			t.Errorf("no error is returned")
		}
	})
}
