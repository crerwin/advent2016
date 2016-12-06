package day3

import "testing"

func TestCheckTriangle(t *testing.T) {
	cases := []struct {
		a, b, c int
		want    bool
	}{
		{1, 3, 4, false},
		{3, 4, 5, true},
		{1, 2, 3, false},
		{5, 10, 25, false},
	}
	for _, c := range cases {
		got := checkTriangle(c.a, c.b, c.c)
		if got != c.want {
			t.Errorf("checkTriangle(%v, %v, %v) == %v, want %v", c.a, c.b, c.c, got, c.want)
		}
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		t    triangle
		want bool
	}{
		{triangle{a: 1, b: 3, c: 4}, false},
		{triangle{a: 3, b: 4, c: 5}, true},
		{triangle{a: 1, b: 2, c: 3}, false},
		{triangle{a: 5, b: 10, c: 25}, false},
	}
	for _, c := range cases {
		got := c.t.isValid()
		if got != c.want {
			t.Errorf("%v.isValid() == %v, want %v", c.t, got, c.want)
		}
	}
}

func TestCreateTriangle(t *testing.T) {
	cases := []struct {
		a, b, c int
		want    triangle
	}{
		{1, 2, 3, triangle{a: 1, b: 2, c: 3}},
	}
	for _, c := range cases {
		got := createTriangle(c.a, c.b, c.c)
		if got != c.want {
			t.Errorf("createTriangle(%v, %v, %v) == %v, want %v", c.a, c.b, c.c, got, c.want)
		}
	}
}

func TestCreateTriangleFromString(t *testing.T) {
	cases := []struct {
		s    string
		want triangle
	}{
		{"1 2 3", triangle{a: 1, b: 2, c: 3}},
	}
	for _, c := range cases {
		got := createTriangleFromString(c.s)
		if got != c.want {
			t.Errorf("createTriangleFromString(%v) == %v, want %v", c.s, got, c.want)
		}
	}
}
