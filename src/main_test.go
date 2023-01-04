package main

import "testing"

func Test_CreateNumList(t *testing.T) {
	args := 3

	got := createNumList(args)
	want := []int{1, 2, 3}

	if len(got) != len(want) {
		t.Errorf("got %d want %d given, %v", got, want, args)
	}
}
