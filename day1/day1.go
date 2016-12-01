package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/crerwin/advent2016/util"
)

type Coord struct {
	x int
	y int
}

type Person struct {
	location  Coord
	visited   map[Coord]bool
	direction int // N - W 1 - 4
}

func (p *Person) Init() {
	p.location.x = 0
	p.location.y = 0
	p.direction = 1
	p.visited = make(map[Coord]bool)
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
		p.location.y += distance
	} else if p.direction == 2 {
		p.location.x += distance
	} else if p.direction == 3 {
		p.location.y -= distance
	} else if p.direction == 4 {
		p.location.x -= distance
	}
}

func (p *Person) move(instruction string) bool {
	direction := instruction[0:1]
	distance, err := strconv.Atoi(instruction[1:len(instruction)])
	if err != nil {
		fmt.Println(err)
	}
	p.turn(direction)
	p.walk(distance)
	if p.visited[p.location] {
		return true
	} else {
		p.visited[p.location] = true
		return false
	}
}

func (p *Person) FollowDirections(directions string, stopAtDuplicateLocation bool) {
	dirlist := strings.Split(directions, ", ")
	for _, dir := range dirlist {
		visitedTwice := p.move(dir)
		if visitedTwice && stopAtDuplicateLocation {
			break
		}
	}
}

func (p *Person) GetDistance() int {
	return util.Abs(p.location.x) + util.Abs(p.location.y)
}

func CreatePerson() Person {
	tempPerson := Person{}
	tempPerson.Init()
	return tempPerson
}
