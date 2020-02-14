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

func BenchmarkSubstr(b *testing.B) {
	s := "abcabcbb"
	for i := 0; i < 15; i++ {
		s = s + s
	}
	r := 3

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNoRepeatingSubStr(s); actual != r {
			b.Errorf("got %d for input %s; expect %d", actual, s, r)
		}
	}

}
