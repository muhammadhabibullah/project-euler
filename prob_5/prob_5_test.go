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
	}

	testCases := map[string]struct {
		io       []io
		solution func(int, int) int
	}{
		"use lcm & gcd":     {inputOutput, ProblemFiveA},
		"without lcm & gcd": {inputOutput, ProblemFiveB},
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

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_5
//BenchmarkProblemFiveA-8   	  749502	      1583 ns/op
//BenchmarkProblemFiveB-8   	 1222333	       973 ns/op
//PASS
