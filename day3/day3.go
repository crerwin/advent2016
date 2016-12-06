package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type triangle struct {
	a, b, c int
}

func (t *triangle) isValid() bool {
	return checkTriangle(t.a, t.b, t.c)
}

func createTriangle(a, b, c int) triangle {
	return triangle{a: a, b: b, c: c}
}

func createTriangleFromString(s string) triangle {
	splits := strings.Fields(s)
	a, aerr := strconv.Atoi(splits[0])
	b, berr := strconv.Atoi(splits[1])
	c, cerr := strconv.Atoi(splits[2])
	if aerr != nil || berr != nil || cerr != nil {
		fmt.Println(aerr, berr, cerr)
	}
	return createTriangle(a, b, c)
}

func checkTriangle(a int, b int, c int) bool {
	if a+b > c && b+c > a && a+c > b {
		return true
	} else {
		return false
	}
}

func CountPossibleTriangles(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		t := createTriangleFromString(scanner.Text())
		if t.isValid() {
			count++
		}
	}
	return count
}
