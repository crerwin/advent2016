package day1

import "testing"

func TestGetDistance(t *testing.T) {
	testPerson := CreatePerson()
	cases := []struct {
		inx, iny, want int
	}{
		{1, 1, 2},
		{5, 5, 10},
		{-5, 5, 10},
		{0, 0, 0},
		{-1, 1, 2},
		{-1000000, 1000000, 2000000},
	}
	for _, c := range cases {
		testPerson.location.x, testPerson.location.y = c.inx, c.iny
		got := testPerson.GetDistance()
		if got != c.want {
			t.Errorf("abs(%v, %v) == %v, want %v", c.inx, c.iny, got, c.want)
		}
	}
}

func TestTurn(t *testing.T) {
	testPerson := CreatePerson()
	cases := []struct {
		in   string
		want int
	}{
		{"R", 2},
		{"R", 3},
		{"R", 4},
		{"R", 1},
		{"L", 4},
		{"L", 3},
	}
	for _, c := range cases {
		testPerson.turn(c.in)
		if testPerson.direction != c.want {
			t.Errorf("turn(%v, %v) == %v, want %v", testPerson, c.in, testPerson.direction, c.want)
		}
	}
}

func TestWalk(t *testing.T) {
	testPerson := CreatePerson()
	cases := []struct {
		inx         int
		iny         int
		indirection int
		indistance  int
		wantx       int
		wanty       int
	}{
		{0, 0, 1, 10, 0, 10},
		{0, 0, 2, 10, 10, 0},
		{0, 0, 3, 10, 0, -10},
		{0, 0, 4, 10, -10, 0},
		{-3, -3, 1, 10, -3, 7},
		{-3, -3, 2, 10, 7, -3},
	}
	for _, c := range cases {
		testPerson.location.x = c.inx
		testPerson.location.y = c.iny
		testPerson.direction = c.indirection
		testPerson.walk(c.indistance)
		if testPerson.location.x != c.wantx || testPerson.location.y != c.wanty {
			t.Errorf("walk(%v) == (%v, %v), want (%v, %v)", c.indistance, testPerson.location.x, testPerson.location.y, c.wantx, c.wanty)
		}
	}
}

func TestMove(t *testing.T) {
	testPerson := CreatePerson()
	cases := []struct {
		inx         int
		iny         int
		indirection int
		instruction string
		wantx       int
		wanty       int
	}{
		{0, 0, 2, "L10", 0, 10},
		{0, 0, 1, "R10", 10, 0},
	}
	for _, c := range cases {
		testPerson.location.x = c.inx
		testPerson.location.y = c.iny
		testPerson.direction = c.indirection
		testPerson.move(c.instruction)
		if testPerson.location.x != c.wantx || testPerson.location.y != c.wanty {
			t.Errorf("move(%v) == (%v, %v), want (%v, %v)", c.instruction, testPerson.location.x, testPerson.location.y, c.wantx, c.wanty)
		}
	}
}

func TestFollowDirections(t *testing.T) {
	testPerson := CreatePerson()
	cases := []struct {
		indirs string
		want   int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}
	for _, c := range cases {
		testPerson.Init()
		testPerson.FollowDirections(c.indirs, false)
		if testPerson.GetDistance() != c.want {
			t.Errorf("followDirections(%v) == %v, want %v.  Person.x = %v Person.y = %v", c.indirs, testPerson.GetDistance(), c.want, testPerson.location.x, testPerson.location.y)
		}
	}
}
