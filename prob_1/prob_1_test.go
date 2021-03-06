package prob1

import "testing"

func TestProblemOne(t *testing.T) {
	solutions := map[string]func(x, y, below int) int{
		"brute force": ProblemOneA,
		"arithmetic":  ProblemOneB,
	}

	for solutionName, solution := range solutions {
		got := solution(3, 5, 1000)
		if got != 233168 {
			t.Errorf("%s solution give wrong answer : %d", solutionName, got)
			return
		}
	}
}

func BenchmarkProblemOneA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemOneA(3, 5, 1000)
	}
}

func BenchmarkProblemOneB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemOneB(3, 5, 1000)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_1
//BenchmarkProblemOneA
//BenchmarkProblemOneA-8   	  170623	      7107 ns/op
//BenchmarkProblemOneB
//BenchmarkProblemOneB-8   	12228836	        97.1 ns/op
//PASS
