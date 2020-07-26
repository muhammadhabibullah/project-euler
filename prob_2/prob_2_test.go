package prob2

import "testing"

func TestProblemTwo(t *testing.T) {
	solutions := []func(below int) int{
		ProblemTwoA,
		ProblemTwoA2,
	}

	for i, solution := range solutions {
		answer := solution(4000000)
		if answer != 4613732 {
			t.Errorf("solution no.%d function give wrong answer : %d", i+1, answer)
			return
		}
	}
}

func BenchmarkProblemTwoA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemTwoA(4000000)
	}
}

func BenchmarkProblemTwoA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemTwoA2(4000000)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_2
//BenchmarkProblemTwoA
//BenchmarkProblemTwoA-8    	65314743	        18.6 ns/op
//BenchmarkProblemTwoA2
//BenchmarkProblemTwoA2-8   	66219823	        18.2 ns/op
//PASS
