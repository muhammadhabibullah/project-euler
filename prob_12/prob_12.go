package prob12

import (
	"math"
)

type countDivFn func(int) int

func countDivs(n int) (count int) {
	for i := 1; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if (n % i) == 0 {
			if (n / i) == i {
				count++
			} else {
				count += 2
			}
		}
	}
	return
}

// ProblemTwelveA count incrementally into get it's total divisors
// above the maximum input using count divisor function input
func ProblemTwelveA(maxDivs int, countDivFn countDivFn) int {
	var triangleNum int
	i := 1
	for {
		triangleNum += i
		i++
		if countDivFn(triangleNum) > maxDivs {
			break
		}
	}

	return triangleNum
}
