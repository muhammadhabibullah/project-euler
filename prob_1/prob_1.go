package prob1

func ProblemOneA(x, y, limit int) (sum int) {
	var i = 1
	for i < limit {
		if (i%x) == 0 || (i%y) == 0 {
			sum += i
		}
		i++
	}
	return
}

func ProblemOneB(x, y, limit int) int {
	xy := x * y
	countAllDivisibleSum := func(div, bottomLimit, upLimit int) (sum int) {
		i := bottomLimit + 1
		for i < upLimit {
			if (i % div) == 0 {
				sum += i
			}
			i++
		}
		return
	}
	calculateAllDivisibleSum := func(div, limit int) int {
		return (limit / div) * (limit + div) / 2
	}

	greatestDivisible := limit - (limit % (xy))
	aboveGDDivisibleSum := countAllDivisibleSum(x, greatestDivisible, limit) +
		countAllDivisibleSum(y, greatestDivisible, limit)
	divisibleByX := calculateAllDivisibleSum(x, greatestDivisible)
	divisibleByY := calculateAllDivisibleSum(y, greatestDivisible)
	divisibleByXY := calculateAllDivisibleSum(xy, greatestDivisible)

	return aboveGDDivisibleSum + divisibleByX + divisibleByY - divisibleByXY
}
