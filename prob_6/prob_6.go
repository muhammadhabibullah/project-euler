package prob6

func ProblemSixA(n int) (ssd int) {
	sqSum, sumSq := 0, 0
	for i := 1; i <= n; i++ {
		sumSq += i * i
		sqSum += i
	}
	ssd = sqSum*sqSum - sumSq
	return
}

func ProblemSixB(n int) (ssd int) {
	sqSum := ((n * n) * (n + 1) * (n + 1)) / 4
	sumSq := (n * (n + 1) * ((2 * n) + 1)) / 6
	ssd = sqSum - sumSq
	return
}
