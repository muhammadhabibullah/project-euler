package prob5

import "math"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func ProblemFiveA(a, b int) (sm int) {
	sm = a
	for a <= b {
		if (sm % a) > 0 {
			ab := lcm(a, b)
			sm = lcm(sm, ab)
		}
		a++
	}
	return
}

func ProblemFiveB(a, b int) int {
	var i = 1
	for k := a; k <= b; k++ {
		if (i % k) > 0 {
			for j := a; j <= b; j++ {
				if ((i * j) % k) == 0 {
					i *= j
					break
				}
			}
		}
	}

	return i
}

// https://gist.github.com/lakshanwd/44c79820a6b9399afa3f56341c2825ef
func nthPrime(limit int) int {
	primes := make([]int, limit)
	count := 0
	isPrimeDivisible := func(candidate int) bool {
		for j := 0; j < count; j++ {
			if math.Sqrt(float64(candidate)) < float64(primes[j]) {
				return false
			}
			if (candidate % primes[j]) == 0 {
				return true
			}
		}
		return false
	}
	for candidate := 2; ; {
		if count < limit {
			if !isPrimeDivisible(candidate) {
				primes[count] = candidate
				count++
			}
			candidate++
		} else {
			break
		}
	}
	return primes[limit-1]
}

func ProblemFiveC(i, k int) (n int) {
	n = 1
	var check = true
	limit := int(math.Sqrt(float64(k)))
	pi := nthPrime(i)

	for pi <= k {
		ai := 1
		if check {
			if pi <= limit {
				ai = int(math.Floor(math.Log(float64(k)) / math.Log(float64(pi))))
			} else {
				check = false
			}
		}
		n *= int(math.Pow(float64(pi), float64(ai)))
		i++
		pi = nthPrime(i)
	}

	return
}
