package prob3

import "math"

func ProblemThreeA(num int) (prime int) {
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

func ProblemThreeB(num int) (prime int) {
	var i = 2
	for (i * i) <= num {
		for (num % i) == 0 {
			prime = i
			num /= i
		}
		i++
	}
	if num > prime {
		prime = num
	}

	return
}
