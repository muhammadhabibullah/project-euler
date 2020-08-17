package prob9

import (
	"math"
)

func ProblemNineA(n int) int {
	var a, b, c int
	var found bool
	for a = 1; a < (n / 3); a++ {
		for b = a; b < (n / 2); b++ {
			c = n - a - b
			if ((a * a) + (b * b)) == (c * c) {
				found = true
				break
			}
		}

		if found {
			return a * b * c
		}
	}

	return 0
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func ProblemNineB(n int) int {
	var a, b, c, m, k, d int
	var found bool
	var mLimit = int(math.Sqrt(float64(n / 2)))

	for m = 2; m <= mLimit; m++ {
		if ((n / 2) % m) == 0 {
			if (m % 2) == 0 {
				k = m + 1
			} else {
				k = m + 2
			}

			for (k < (m * 2)) && (k <= (n / (m * 2))) {
				if (((n / (m * 2)) % k) == 0) && (gcd(k, m) == 1) {
					d = n / 2 / (k * m)
					n = k - m
					a = d * ((m * m) - (n * n))
					b = 2 * d * n * m
					c = d * ((m * m) + (n * n))
					found = true
					break
				}
				k += 2
			}
		}
		if found {
			return a * b * c
		}
	}

	return 0
}
