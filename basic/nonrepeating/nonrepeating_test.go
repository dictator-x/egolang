package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		a string
		r int
	}{
		{"abcabcbb", 3},
	}

	for _, tt := range tests {
		if actual := lengthOfNoRepeatingSubStr(tt.a); actual != tt.r {
			t.Errorf("got %d for input %s; expect %d", actual, tt.a, tt.r)
		}
	}
}
