package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SolutionOf001 func(x, y, limit int) int

func BruteForce(x, y, limit int) (sum int) {
	for i := 1; i < limit; i++ {
		if (i%x) == 0 || (i%y) == 0 {
			sum += i
		}
	}

	return sum
}

func Arithmetic(x, y, limit int) int {
	max := limit - 1
	sumDivisibleBy := func(n, max int) int {
		N := max / n
		return n * N * (N + 1) / 2
	}

	return sumDivisibleBy(x, max) + sumDivisibleBy(y, max) - sumDivisibleBy(x*y, max)
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
		fmt.Println(Arithmetic(3, 5, input[i]))
	}
}
