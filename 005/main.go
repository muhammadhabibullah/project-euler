package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BruteForce(low, high int) (sm int) {
	sm = high
	div := low

	for div <= high {
		if sm%div != 0 {
			sm += high
			div = low
			continue
		}
		div++
	}

	return sm
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func LCMnGCD(low, high int) (sm int) {
	sm = low
	for low <= high {
		if (sm % low) > 0 {
			ab := lcm(low, high)
			sm = lcm(sm, ab)
		}
		low++
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
		fmt.Println(LCMnGCD(1, input[i]))
	}
}
