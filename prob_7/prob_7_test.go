package prob7

import (
	"testing"
)

func TestProblemSeven(t *testing.T) {
	inputOutputs := map[int]int{
		1:     2,
		2:     3,
		3:     5,
		4:     7,
		5:     11,
		6:     13,
		10001: 104743,
	}

	solutions := map[string]func(int) int{
		"Sieve of Eratosthenes": ProblemSevenA,
		"Brute force":           ProblemSevenB,
	}

	for input, expected := range inputOutputs {
		for solutionName, solutionFunc := range solutions {
			got := solutionFunc(input)
			if got != expected {
				t.Errorf("%s solution function give wrong answer : %d; expected: %d", solutionName, got, expected)
			}
		}
	}
}

func BenchmarkProblemSevenA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemSevenA(10001)
	}
}

func BenchmarkProblemSevenB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemSevenB(10001)
	}
}

//goos: darwin
//goarch: amd64
//pkg: project-euler/prob_7
//BenchmarkProblemSevenA
//BenchmarkProblemSevenA-4   	     180	   6293823 ns/op
//BenchmarkProblemSevenB
//BenchmarkProblemSevenB-4   	      36	  30683898 ns/op
//PASS
