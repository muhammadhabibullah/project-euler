package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SolutionOf002 func(limit int) int

func BruteForce(limit int) (sum int) {
	var i, next = 1, 2
	for i < limit {
		if (i % 2) == 0 {
			sum += i
		}

		next, i = next+i, next
	}
	return
}

func EvenNumberBruteForce(limit int) (sum int) {
	var i, next = 2, 0
	for i < limit {
		sum += i
		i, next = (4*i)+next, i
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
		fmt.Println(BruteForce(input[i]))
	}
}
