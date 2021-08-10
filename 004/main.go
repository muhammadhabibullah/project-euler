package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SolutionOf004 struct {
	palindromeSearcher
	palindromeValidator
}

type palindromeSearcher func(num int, validator palindromeValidator) (palindrome int)

type palindromeValidator func(num int) bool

func iterativePalindromeValidator(num int) bool {
	var remainder int
	var reverse = 0
	tmpNum := num

	for {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10
		if num == 0 {
			break
		}
	}

	return tmpNum == reverse
}

func recursiveNumberReverse(n int, temp int) int {
	if n == 0 {
		return temp
	}
	temp = (temp * 10) + (n % 10)

	return recursiveNumberReverse(n/10, temp)
}

func recursivePalindromeValidator(num int) bool {
	return num == recursiveNumberReverse(num, 0)
}

const largestThreeDigitNumber = 999
const smallestThreeDigitNumber = 100

func BruteForce(num int, isPalindrome palindromeValidator) (palindrome int) {
	a := largestThreeDigitNumber
	bottomLimit := smallestThreeDigitNumber

	for a >= bottomLimit {
		b := largestThreeDigitNumber
		for b >= a {
			ab := a * b
			if ab <= palindrome {
				break
			}

			if isPalindrome(ab) && ab < num {
				palindrome = ab
			}
			b--
		}
		a--
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Replace(nStr, "\n", "", 1))

	input := make([]int, 0)
	for i := 0; i < n; i++ {
		xStr, _ := reader.ReadString('\n')
		x, _ := strconv.Atoi(strings.Replace(xStr, "\n", "", 1))
		input = append(input, x)
	}

	for i := 0; i < n; i++ {
		fmt.Println(BruteForce(input[i], iterativePalindromeValidator))
	}
}
