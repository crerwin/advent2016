package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/crerwin/advent2016/util"
)

type Person struct {
	x         int
	y         int
	direction int // N - W 1 - 4
}

func (p *Person) Init() {
	p.x = 0
	p.y = 0
	p.direction = 1
}

func (p *Person) turn(direction string) int {
	if direction == "R" {
		p.direction += 1
	} else if direction == "L" {
		p.direction -= 1
	} else {
		fmt.Println("bad direction")
	}
	if p.direction > 4 {
		p.direction = 1
	}
	if p.direction < 1 {
		p.direction = 4
	}
	return p.direction
}

func (p *Person) walk(distance int) {
	if p.direction == 1 {
		p.y += distance
	} else if p.direction == 2 {
		p.x += distance
	} else if p.direction == 3 {
		p.y -= distance
	} else if p.direction == 4 {
		p.x -= distance
	}
}

func (p *Person) move(instruction string) {
	direction := instruction[0:1]
	distance, err := strconv.Atoi(instruction[1:len(instruction)])
	if err != nil {
		fmt.Println(err)
	}
	p.turn(direction)
	p.walk(distance)
}

func (p *Person) FollowDirections(directions string) {
	dirlist := strings.Split(directions, ", ")
	for _, dir := range dirlist {
		p.move(dir)
	}
}

func (p *Person) GetDistance() int {
	return util.Abs(p.x) + util.Abs(p.y)
}

func CreatePerson() Person {
	tempPerson := Person{x: 0, y: 0, direction: 1}
	return tempPerson
}
