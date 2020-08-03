package prob7

import (
	"math"
)

func ProblemSevenA(n int) (prime int) {
	const MaxCheckedPrime = 1000000
	var primes []int

	isPrime := make([]bool, MaxCheckedPrime)
	for i := range isPrime {
		isPrime[i] = true
	}

	for p := 2; p*p < MaxCheckedPrime; p++ {
		if isPrime[p] {
			for i := p * p; i < MaxCheckedPrime; i += p {
				isPrime[i] = false
			}
		}
	}

	for p := 2; p < MaxCheckedPrime; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	if primes != nil {
		return primes[n-1]
	}
	return
}

func ProblemSevenB(nth int) int {

	isPrime := func(n int) bool {
		if n == 2 {
			return true
		}
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

	if nth == 1 {
		return 2
	}

	n, t := 1, 1
	for n < nth {
		t += 2
		if isPrime(t) {
			n++
		}
	}

	return t
}
