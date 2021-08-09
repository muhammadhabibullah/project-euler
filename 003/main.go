package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SolutionOf003 func(num int) (prime int)

func BruteForce(num int) (prime int) {
	for (num % 2) == 0 {
		prime = 2
		num /= 2
	}

	for i := 3; i < int(math.Sqrt(float64(num)))+1; i += 2 {
		for (num % i) == 0 {
			prime = i
			num /= i
		}
	}

	if num > prime {
		prime = num
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
