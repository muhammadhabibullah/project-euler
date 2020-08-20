package prob10

import (
	"math"
)

func ProblemTenA(max int) (sum int) {
	isPrime := make([]bool, max)
	for i := range isPrime {
		isPrime[i] = true
	}

	for p := 2; p*p < max; p++ {
		if isPrime[p] {
			for i := p * p; i < max; i += p {
				isPrime[i] = false
			}
		}
	}

	for p := 2; p < max; p++ {
		if isPrime[p] {
			sum += p
		}
	}
	return
}

func ProblemTenB(max int) (sum int) {
	isPrime := func(n int) bool {
		if (n % 2) == 0 {
			return false
		}

		i := 3
		primeRange := int(math.Sqrt(float64(n))) + 1
		for i < primeRange {
			if (n % i) == 0 {
				return false
			}
			i++
		}
		return true
	}

	if max > 2 {
		sum += 2
	}

	t := 3
	for t < max {
		if isPrime(t) {
			sum += t
		}
		t += 2
	}

	return
}
