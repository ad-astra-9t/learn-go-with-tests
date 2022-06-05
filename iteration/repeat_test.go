package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat_character() {
	repeated := Repeat("x", 7)
	fmt.Println(repeated)
	// Output: xxxxxxx
}

func ExampleRepeat_string() {
	repeated := Repeat("xyz", 7)
	fmt.Println(repeated)
	// Output: xyzxyzxyzxyzxyzxyzxyz
}
