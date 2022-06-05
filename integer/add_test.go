package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func ExampleAdd() {
	sum := Add(0, 2)
	fmt.Println(sum)
	// Output: 2
}
