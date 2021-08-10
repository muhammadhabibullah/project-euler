package main

import (
	"testing"
)

var (
	benchmarkLimit = 1000000
	inputOutputs   = map[int]int{
		101110:  101101,
		800000:  793397,
		1000000: 906609,
	}
)

func TestSolutions(t *testing.T) {
	solutions := map[string]SolutionOf004{
		"Brute Force Iterative Palindrome": {BruteForce, iterativePalindromeValidator},
		"Brute Force Recursive Palindrome": {BruteForce, recursivePalindromeValidator},
	}

	for limit, correct := range inputOutputs {
		for name, solution := range solutions {
			got := solution.palindromeSearcher(limit, solution.palindromeValidator)
			if got != correct {
				t.Errorf("%s solution give wrong answer : %d", name, got)
			}
		}
	}
}

func BenchmarkBruteForceIterativePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(benchmarkLimit, iterativePalindromeValidator)
	}
}

func BenchmarkBruteForceRecursivePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(benchmarkLimit, recursivePalindromeValidator)
	}
}
