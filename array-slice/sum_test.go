package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	moreNums := []int{6, 7, 8}
	got := SumAll(nums, moreNums)
	want := []int{15, 21}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("sum two sets of number", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		moreNums := []int{6, 7, 8}
		got := SumAllTails(nums, moreNums)
		want := []int{14, 15}
		assertMessage(t, got, want)
	})
	t.Run("sum an empty slice safely", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		moreNums := []int{}
		got := SumAllTails(nums, moreNums)
		want := []int{14, 0}
		assertMessage(t, got, want)
	})
}
