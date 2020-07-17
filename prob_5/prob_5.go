package prob5

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
		ab := lcm(a, b)
		sm = lcm(sm, ab)
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
