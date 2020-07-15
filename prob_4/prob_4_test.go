package prob4

import (
	"testing"
)

func TestProblemFourA(t *testing.T) {
	solutions := map[string]func(below int) int{
		"super brute force": ProblemFourA,
		//"smarter":     ProblemFourB,
	}
	outputs := map[int]int{
		1: 9,
		2: 9009,
		3: 906609,
		//4: 99000099,
	}

	for solutionName, solutionFunc := range solutions {
		for input, expected := range outputs {
			answer := solutionFunc(input)
			if answer != expected {
				t.Errorf("%s solution for %d-digit number give wrong answer: %d, expected: %d",
					solutionName, input, answer, expected)
			}
		}
	}
}

func BenchmarkProblemThreeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourA(3)
	}
}
