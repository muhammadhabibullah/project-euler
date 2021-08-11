package main

import (
	"testing"
)

type SolutionOf005 func(int, int) int

var (
	benchmarkInput = 20
	inputOutputs   = map[int]int{
		3:  6,
		10: 2520,
		20: 232792560,
	}
)

func TestSolutions(t *testing.T) {
	solutions := map[string]SolutionOf005{
		"Brute Force": BruteForce,
		"LCM and GCD": LCMnGCD,
	}

	for limit, correct := range inputOutputs {
		for name, solution := range solutions {
			got := solution(1, limit)
			if got != correct {
				t.Errorf("%s solution give wrong answer : %d", name, got)
			}
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(1, benchmarkInput)
	}
}

func BenchmarkLCMnGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCMnGCD(1, benchmarkInput)
	}
}
