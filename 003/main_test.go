package main

import (
	"testing"
)

var (
	below          = 600851475143
	xyInputOutputs = map[int]int{
		10:           5,
		17:           17,
		600851475143: 6857,
	}
)

func TestSolutions(t *testing.T) {
	solutions := map[string]SolutionOf003{
		"Brute Force": BruteForce,
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
