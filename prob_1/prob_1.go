package prob1

func ProblemOneA(below int) (sum int) {
	var i = 1
	for i < below {
		if (i%3) == 0 || (i%5) == 0 {
			sum += i
		}
		i++
	}
	return
}

func ProblemOneB(below int) int {
	addDivisible := func(div, above, below int) (sum int) {
		i := above + 1
		for i < below {
			if (i % div) == 0 {
				sum += i
			}
			i++
		}
		return
	}
	countDivisible := func(div, below int) int {
		return (below / div) * (below + div) / 2
	}

	greatestDivisorBelow := below - (below % (3 * 5))

	aboveGDBSum := addDivisible(3, greatestDivisorBelow, below) +
		addDivisible(5, greatestDivisorBelow, below)
	divisibleBy3 := countDivisible(3, greatestDivisorBelow)
	divisibleBy5 := countDivisible(5, greatestDivisorBelow)
	divisibleBy15 := countDivisible(15, greatestDivisorBelow)

	return aboveGDBSum + divisibleBy3 + divisibleBy5 - divisibleBy15
}
