package day2

import "testing"

func TestGet(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 1},
		{0, 1, 4},
		{0, 2, 7},
		{1, 0, 2},
		{1, 1, 5},
		{1, 2, 8},
		{2, 0, 3},
		{2, 1, 6},
		{2, 2, 9},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		got := k.get()
		if got != c.want {
			t.Errorf("k.get() at %v, %v == %v, want %v", c.x, c.y, got, c.want)
		}
	}
}

func TestMoveUp(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 1},
		{1, 0, 2},
		{0, 1, 1},
		{0, 2, 4},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		start := k.get()
		k.moveUp()
		got := k.get()
		if got != c.want {
			t.Errorf("k.moveUp() starting at %v, %v (%v) == %v, want %v", c.x, c.y, start, got, c.want)
		}
	}
}

func TestMoveDown(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 4},
		{2, 2, 9},
		{0, 1, 7},
		{0, 2, 7},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		start := k.get()
		k.moveDown()
		got := k.get()
		if got != c.want {
			t.Errorf("k.moveDown() starting at %v, %v (%v) == %v, want %v", c.x, c.y, start, got, c.want)
		}
	}
}

func TestMoveRight(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 2},
		{1, 0, 3},
		{0, 1, 5},
		{0, 2, 8},
		{2, 0, 3},
		{2, 1, 6},
		{2, 2, 9},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		start := k.get()
		k.moveRight()
		got := k.get()
		if got != c.want {
			t.Errorf("k.moveRight() starting at %v, %v (%v) == %v, want %v", c.x, c.y, start, got, c.want)
		}
	}
}

func TestMoveLeft(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 4},
		{0, 2, 7},
		{2, 1, 5},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		start := k.get()
		k.moveLeft()
		got := k.get()
		if got != c.want {
			t.Errorf("k.moveLeft() starting at %v, %v (%v) == %v, want %v", c.x, c.y, start, got, c.want)
		}
	}
}

func TestMove(t *testing.T) {
	cases := []struct {
		x         int
		y         int
		direction string
		want      int
	}{
		{0, 0, "U", 1},
		{1, 0, "D", 5},
		{0, 1, "L", 4},
		{0, 2, "R", 8},
		{2, 1, "L", 5},
	}
	k := createKeypad()
	for _, c := range cases {
		k.setLocation(c.x, c.y)
		start := k.get()
		k.move(c.direction)
		got := k.get()
		if got != c.want {
			t.Errorf("k.move(%v) starting at %v, %v (%v) == %v, want %v", c.direction, c.x, c.y, start, got, c.want)
		}
	}
}

func TestMoveToNextNumber(t *testing.T) {
	cases := []struct {
		directions string
		want       int
	}{
		{"ULL", 1},
		{"RRDDD", 9},
		{"LURDL", 8},
		{"UUUUD", 5},
	}
	k := createKeypad()
	for _, c := range cases {
		start := k.get()
		got := k.moveToNextNumber(c.directions)
		if got != c.want {
			t.Errorf("k.moveToNextNumber(%v) starting at %v == %v, want %v", c.directions, start, got, c.want)
		}
	}
}
