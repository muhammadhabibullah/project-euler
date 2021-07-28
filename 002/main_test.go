package main

import (
	"testing"
)

var (
	below          = 4000000
	xyInputOutputs = map[int]int{
		10:                10,
		100:               44,
		4000000:           4613732,
		40000000000000000: 49597426547377748,
	}
)

func TestSolutions(t *testing.T) {
	solutions := map[string]SolutionOf002{
		"Brute Force":             BruteForce,
		"Even Number Brute Force": EvenNumberBruteForce,
	}

	for limit, correct := range xyInputOutputs {
		for name, solution := range solutions {
			got := solution(limit)
			if got != correct {
				t.Errorf("%s solution give wrong answer : %d", name, got)
			}
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(below)
	}
}

func BenchmarkEvenNumberBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenNumberBruteForce(below)
	}
}
