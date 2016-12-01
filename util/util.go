package util

import (
	"io/ioutil"
	"log"
)

func Abs(n int) int {
	// why the math package only deals with floats I'll never get
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func ReadFileAsString(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
