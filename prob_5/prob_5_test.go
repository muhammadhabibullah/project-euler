package prob5

import (
	"testing"
)

func TestProblemFive(t *testing.T) {
	type io struct {
		input1 int
		input2 int
		output int
	}
	inputOutput := []io{
		{1, 10, 2520},
		{1, 20, 232792560},
		{1, 30, 2329089562800},
		{1, 40, 5342931457063200},
	}

	testCases := map[string]struct {
		io       []io
		solution func(int, int) int
	}{
		"use lcm & gcd":     {inputOutput, ProblemFiveA},
		"without lcm & gcd": {inputOutput, ProblemFiveB},
		"use nth prime":     {inputOutput, ProblemFiveC},
	}

	for testName, ts := range testCases {
		for _, io := range ts.io {
			got := ts.solution(io.input1, io.input2)
			expected := io.output
			if got != expected {
				t.Errorf("%s solution function give wrong answer : %d; expected: %d", testName, got, expected)
			}
		}
	}
}

func BenchmarkProblemFiveA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFiveA(1, 20)
	}
}

func BenchmarkProblemFiveB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFiveB(1, 20)
	}
}

func BenchmarkProblemFiveC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFiveC(1, 20)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_5
//BenchmarkProblemFiveA
//BenchmarkProblemFiveA-8   	 2324925	       498 ns/op
//BenchmarkProblemFiveB
//BenchmarkProblemFiveB-8   	 2196723	       542 ns/op
//BenchmarkProblemFiveC
//BenchmarkProblemFiveC-8   	  758619	      1323 ns/op
//PASS
