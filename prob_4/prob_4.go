package prob4

import "unicode/utf8"

func largestNumber(digit int) int {
	var num = 1
	for digit != 0 {
		num *= 10
		digit--
	}

	return num - 1
}

func ProblemFourA(digit int) int {
	isPalindrome := func(num int) bool {
		if num < 10 {
			return true
		}

		div := 1
		for div <= num {
			div *= 10
		}

		var arrNum []int
		tmpNum := num
		for tmpNum > 0 {
			div /= 10
			if div > tmpNum {
				arrNum = append(arrNum, 0)
			} else {
				divNum := tmpNum / div
				arrNum = append(arrNum, divNum)
				tmpNum -= divNum * div
			}
		}

		var pal int
		mtp := 1
		for _, n := range arrNum {
			pal += mtp * n
			mtp *= 10
		}

		return pal == num
	}

	a, b := 1, 1
	limit := largestNumber(digit)
	var pal int

	for a <= limit {
		for b <= limit {
			ab := a * b
			if isPalindrome(ab) && ab > pal {
				pal = ab
				//fmt.Println(a, b, pal)
			}
			b++
		}
		a++
		b = a
	}

	return pal
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

func ProblemFourB(digit int) int {
	return 0
}
