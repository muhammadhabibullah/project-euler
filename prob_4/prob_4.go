package prob4

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func smallestNumber(digit int) int {
	return int(math.Pow(float64(10), float64(digit-1)))
}

func largestNumber(digit int) int {
	return int(math.Pow(float64(10), float64(digit))) - 1
}

func isIntPalindrome(num int) bool {
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

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func isStrPalindrome(num int) bool {
	str := strconv.Itoa(num)
	return str == reverse(str)
}

func ProblemFourA(digit int, isPalindrome func(int) bool) (pal int) {
	a, b := smallestNumber(digit), smallestNumber(digit)
	upperLimit := largestNumber(digit)

	for a <= upperLimit {
		for b <= upperLimit {
			ab := a * b
			if ab > pal {
				if isPalindrome(ab) {
					pal = ab
				}
			}
			b++
		}
		a++
		b = a
	}

	return
}

func ProblemFourB(digit int, isPalindrome func(int) bool) (pal int) {
	a := largestNumber(digit)
	bottomLimit := smallestNumber(digit)

	for a >= bottomLimit {
		b := largestNumber(digit)
		for b >= a {
			ab := a * b
			if ab <= pal {
				break
			}

			if isPalindrome(ab) {
				pal = ab
			}
			b--
		}
		a--
	}

	return
}

func ProblemFourC(digit int, isPalindrome func(int) bool) (pal int) {
	if digit == 1 {
		return 9
	}

	a := largestNumber(digit)
	bottomLimit := smallestNumber(digit)

	for a >= bottomLimit {
		var b, db int
		if (a % 11) == 0 {
			b = largestNumber(digit)
			db = 1
		} else {
			b = largestNumber(digit) - (largestNumber(digit) % 11)
			db = 11
		}
		for b >= a {
			ab := a * b
			if ab <= pal {
				break
			}

			if isPalindrome(ab) {
				pal = ab
			}
			b -= db
		}
		a--
	}

	return
}
