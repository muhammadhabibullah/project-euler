package prob2

func ProblemTwoA(limit int) (sum int) {
	var i, next = 1, 2
	for i < limit {
		if (i % 2) == 0 {
			sum += i
		}

		next, i = next+i, next
	}
	return
}

func ProblemTwoA2(limit int) (sum int) {
	var i, next = 1, 1
	for sum < limit {
		total := i + next
		if (total % 2) == 0 {
			sum += total
		}
		i, next = next, total
	}
	return
}
