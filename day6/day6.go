package day6

import (
	"fmt"

	"github.com/crerwin/advent2016/input"
)

type letterCount struct {
	counts map[string]int
}

type positions struct {
	countsPerPosition []letterCount
}

func newLetterCount() letterCount {
	var lc letterCount
	counts := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	lc.counts = counts
	return lc
}

func newPositions(length int) positions {
	var p positions
	for i := 0; i < length; i++ {
		p.countsPerPosition = append(p.countsPerPosition, newLetterCount())
	}
	return p
}

func (lc *letterCount) addCount(letter string) {
	lc.counts[letter]++
}

func (lc *letterCount) getMostFrequent() string {
	max := 0
	mostFrequent := ""
	for key, value := range lc.counts {
		if value > max {
			mostFrequent = key
			max = value
		}
	}
	return mostFrequent
}

func (lc *letterCount) getLeastFrequent() string {
	min := 1000000
	leastFrequent := ""
	for key, value := range lc.counts {
		if value < min {
			leastFrequent = key
			min = value
		}
	}
	return leastFrequent
}

func (p *positions) addCountAtPosition(letter string, position int) {
	p.countsPerPosition[position].addCount(letter)
}

func (p *positions) parseLine(line string) {
	for i, r := range line {
		p.addCountAtPosition(string(r), i)
	}
}

func (p *positions) getMostFrequents() string {
	word := ""
	for _, c := range p.countsPerPosition {
		word += c.getMostFrequent()
	}
	return word
}

func (p *positions) getLeastFrequents() string {
	word := ""
	for _, c := range p.countsPerPosition {
		word += c.getLeastFrequent()
	}
	return word
}

func findPassword(transmissions []string, part int) string {
	length := len(transmissions[0])
	fmt.Println("password length will be", length, "characters.")
	counts := newPositions(length)
	for _, line := range transmissions {
		counts.parseLine(line)
	}
	if part == 1 {
		return counts.getMostFrequents()
	} else if part == 2 {
		return counts.getLeastFrequents()
	} else {
		return "wrong part specified"
	}

}

func Part1() {
	transmissions := input.Day6(1)
	password := findPassword(transmissions, 1)
	fmt.Println(password)
}

func Part2() {
	transmissions := input.Day6(1)
	password := findPassword(transmissions, 2)
	fmt.Println(password)
}
