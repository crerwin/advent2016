package util

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{5, 5},
		{-5, 5},
		{0, 0},
		{-1, 1},
		{-1000000, 1000000},
	}
	for _, c := range cases {
		got := Abs(c.in)
		if got != c.want {
			t.Errorf("abs(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
