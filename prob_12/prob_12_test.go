package prob12

import (
	"testing"
)

func TestProblemTwelve(t *testing.T) {
	inputOutputs := map[int]int{
		3:   6,
		4:   28,
		5:   28,
		500: 76576500,
	}

	divisors := map[string]countDivFn{
		"O(n^1/3)": countDivs,
	}

	solutions := map[string]func(int, countDivFn) int{
		"Standard": ProblemTwelveA,
	}

	for input, expected := range inputOutputs {
		for solutionName, solutionFunc := range solutions {
			for divisorName, divisorFunc := range divisors {
				got := solutionFunc(input, divisorFunc)
				if got != expected {
					t.Errorf("%s solution with %s divisor function give wrong answer : %d; expected: %d",
						solutionName, divisorName, got, expected)
				}
			}
		}
	}
}
