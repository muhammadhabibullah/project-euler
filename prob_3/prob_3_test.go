package prob3

import "testing"

func TestProblemThree(t *testing.T) {
	solutions := []func(below int) int{
		ProblemThreeA,
		ProblemThreeB,
	}

	for i, solution := range solutions {
		answer := solution(600851475143)
		if answer != 6857 {
			t.Errorf("solution no.%d function give wrong answer : %d", i+1, answer)
			return
		}
	}
}

func BenchmarkProblemThreeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemThreeA(600851475143)
	}
}

func BenchmarkProblemThreeB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemThreeB(600851475143)
	}
}
