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
