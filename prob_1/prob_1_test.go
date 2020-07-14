package prob1

import "testing"

func TestProblemOne(t *testing.T) {
	solutions := []func(below int) int{
		ProblemOneA,
		ProblemOneB,
	}

	for i, solution := range solutions {
		answer := solution(1000)
		if answer != 233168 {
			t.Errorf("solution no.%d function give wrong answer : %d", i+1, answer)
			return
		}
	}
}

func BenchmarkProblemOneA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemOneA(1000)
	}
}

func BenchmarkProblemOneB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemOneB(1000)
	}
}
