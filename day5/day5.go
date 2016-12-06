package day5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func getPassword(doorid string) {
	password := ""
	i := 0
	for len(password) < 8 {
		stringToHash := doorid + strconv.Itoa(i)
		hasher := md5.New()
		hasher.Write([]byte(stringToHash))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if hash[0:5] == "00000" {
			fmt.Print(hash[5:6])
			password += hash[5:6]
		}
		i++
	}
	fmt.Println()
}

func getPart2Password(doorid string) {
	password := make([]string, 8)
	charcount := 0
	i := 0
	for charcount < 8 {
		stringToHash := doorid + strconv.Itoa(i)
		hasher := md5.New()
		hasher.Write([]byte(stringToHash))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if hash[0:5] == "00000" {
			position, _ := strconv.Atoi(hash[5:6])
			if position < 8 && password[position] == "" {
				password[position] = hash[6:7]
				charcount++
			}
			fmt.Println(password)
		}
		i++
	}
	fmt.Println()
}

func Part1() {
	getPassword("abbhdwsy")
}

func Part2() {
	getPart2Password("abbhdwsy")
}
