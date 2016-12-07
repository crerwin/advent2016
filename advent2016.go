package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/crerwin/advent2016/day6"
)

func main() {
	// person := day1.CreatePerson()
	// person.FollowDirections(input.Day1(1), false)
	// fmt.Println(person.GetDistance())

	if len(os.Args) != 3 {
		fmt.Println("Invalid arguments.  Usage: advent2016 n p - where n is the day number and p is the part number")
	} else {
		flag.Parse()
		day6.Part2()
	}

}
