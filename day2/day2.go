package day2

import (
	"fmt"
	"strconv"

	"github.com/crerwin/advent2016/input"
)

type keypad struct {
	keys [3][3]int
	x    int
	y    int
}

func (k *keypad) init() {
	k.keys = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	k.x = 1
	k.y = 1
}

func createKeypad() keypad {
	k := keypad{}
	k.init()
	return k
}

func (k *keypad) setLocation(x int, y int) {
	k.x = x
	k.y = y
}

func (k *keypad) moveLeft() {
	if k.x > 0 {
		k.x--
	}
}

func (k *keypad) moveRight() {
	if k.x < 2 {
		k.x++
	}
}

func (k *keypad) moveUp() {
	if k.y > 0 {
		k.y--
	}
}

func (k *keypad) moveDown() {
	if k.y < 2 {
		k.y++
	}
}

func (k *keypad) move(direction string) {
	if direction == "U" {
		k.moveUp()
	} else if direction == "D" {
		k.moveDown()
	} else if direction == "R" {
		k.moveRight()
	} else if direction == "L" {
		k.moveLeft()
	} else {
		fmt.Println("bad direction")
	}
}

func (k *keypad) get() int {
	return k.keys[k.y][k.x]
}

func (k *keypad) moveToNextNumber(directions string) int {
	for _, r := range directions {
		k.move(string(r))
	}
	return k.get()
}

func Part1() {
	input := input.Day2(1)
	code := ""
	k := createKeypad()
	for _, directions := range input {
		button := k.moveToNextNumber(directions)
		code += strconv.Itoa(button)
	}
	fmt.Println(code)
}
