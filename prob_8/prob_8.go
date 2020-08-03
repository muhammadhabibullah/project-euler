package prob8

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readDigits() string {
	file, _ := os.Open("digits.txt")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return strings.Replace(string(bytes), "\n", "", -1)
}

func parseNth(str uint8) int {
	n, _ := strconv.Atoi(string(str))
	return n
}

func ProblemEightA(digit int) (maxProduct int) {
	digits := readDigits()
	limit := len(digits) - digit + 1

next:
	for i := 0; i < limit; i++ {
		product := 1
		for j := i; j < (i + digit); j++ {
			multiply := parseNth(digits[j])
			if multiply == 0 {
				continue next
			}
			product *= multiply
		}

		if maxProduct < product {
			maxProduct = product
		}
	}
	return
}

func ProblemEightA2(digit int) (maxProduct int) {
	digits := readDigits()
	limit := len(digits) - digit + 1
	var lastProduct int

next:
	for i := 0; i < limit; i++ {
		product := 1

		if lastProduct != 0 {
			div := parseNth(digits[i-1])
			mul := parseNth(digits[i+digit-1])
			product = lastProduct * mul / div
		} else {
			for j := i; j < (i + digit); j++ {
				multiply := parseNth(digits[j])
				if multiply == 0 {
					lastProduct = 0
					continue next
				}
				product *= multiply
			}
		}

		if maxProduct < product {
			maxProduct = product
		}
		lastProduct = product
	}
	return
}
