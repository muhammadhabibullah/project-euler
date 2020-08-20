package prob10

import (
	"testing"
)

func TestProblemTen(t *testing.T) {
	inputOutputs := map[int]int{
		1:       0,
		2:       0,
		3:       2,
		5:       5,
		7:       10,
		10:      17,
		2000000: 142913828922,
	}

	solutions := map[string]func(int) int{
		"Sieve of Eratosthenes": ProblemTenA,
		"Brute Force":           ProblemTenB,
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

func BenchmarkProblemTenA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemTenA(2000000)
	}
}

func BenchmarkProblemTenB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemTenB(2000000)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_10
//BenchmarkProblemTenA
//BenchmarkProblemTenA-8   	     214	   5748200 ns/op
//BenchmarkProblemTenB
//BenchmarkProblemTenB-8   	       2	 864060978 ns/op
//PASS
