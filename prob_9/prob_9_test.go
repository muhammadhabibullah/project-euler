package prob9

import (
	"testing"
)

func TestProblemNine(t *testing.T) {
	inputOutputs := map[int]int{
		12:   60,       // 3,4,5
		220:  199980,   // 20, 99, 101
		1000: 31875000, // 200, 375, 425
	}

	solutions := map[string]func(int) int{
		"Brute force": ProblemNineA,
		"Arithmetic":  ProblemNineB,
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

func BenchmarkProblemNineA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemNineA(1000)
	}
}

func BenchmarkProblemNineA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemNineB(1000)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_9
//BenchmarkProblemNineA
//BenchmarkProblemNineA-8    	   13738	     85880 ns/op
//BenchmarkProblemNineA2
//BenchmarkProblemNineA2-8   	27349213	        43.9 ns/op
//PASS
