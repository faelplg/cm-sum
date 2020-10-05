package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(5, 5)
	if got != 10 {
		t.Errorf("sum(5 + 5) == %d, want %d", got, 10)
	}
}
