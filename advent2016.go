package main

import (
	"fmt"

	"github.com/crerwin/advent2016/day1"
	"github.com/crerwin/advent2016/input"
)

func main() {
	person := day1.CreatePerson()
	person.FollowDirections(input.Day1(1), false)
	fmt.Println(person.GetDistance())
}
