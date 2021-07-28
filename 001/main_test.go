package main

import (
	"testing"
)

var (
	x, y, below    = 3, 5, 1000
	xyInputOutputs = map[int]int{
		10:    23,
		100:   2318,
		1000:  233168,
		10000: 23331668,
	}
)

func TestSolutions(t *testing.T) {
	solutions := map[string]SolutionOf001{
		"Brute Force": BruteForce,
		"Arithmetic":  Arithmetic,
	}

	for limit, correct := range xyInputOutputs {
		for name, solution := range solutions {
			got := solution(x, y, limit)
			if got != correct {
				t.Errorf("%s solution give wrong answer : %d", name, got)
			}
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(x, y, below)
	}
}

func BenchmarkArithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Arithmetic(x, y, below)
	}
}
