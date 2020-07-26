package prob6

import (
	"testing"
)

func TestProblemSix(t *testing.T) {
	inputOutputs := map[int]int{
		10:   2640,
		100:  25164150,
		1000: 250166416500,
	}

	solutions := map[string]func(int) int{
		"brute force": ProblemSixA,
		"arithmetic":  ProblemSixB,
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

func BenchmarkProblemFiveA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemSixA(100)
	}
}

func BenchmarkProblemFiveB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemSixB(100)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_6
//BenchmarkProblemFiveA
//BenchmarkProblemFiveA-8   	20764658	        52.2 ns/op
//BenchmarkProblemFiveB
//BenchmarkProblemFiveB-8   	1000000000	         0.500 ns/op
//PASS
