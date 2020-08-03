package prob8

import (
	"testing"
)

func TestProblemEight(t *testing.T) {
	inputOutputs := map[int]int{
		4:  5832,
		13: 23514624000,
	}

	solutions := map[string]func(int) int{
		"Standard":  ProblemEightA,
		"Optimized": ProblemEightA2,
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

func BenchmarkProblemEightA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemEightA(13)
	}
}

func BenchmarkProblemEightA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemEightA2(13)
	}
}

//goos: darwin
//goarch: amd64
//pkg: project-euler/prob_8
//BenchmarkProblemEightA
//BenchmarkProblemEightA-4    	    7818	    132416 ns/op
//BenchmarkProblemEightA2
//BenchmarkProblemEightA2-4   	   13060	     90673 ns/op
//PASS
