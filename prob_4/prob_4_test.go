package prob4

import "testing"

type palindromeValidation func(int) bool

func TestProblemFourA(t *testing.T) {
	solutions := map[string]func(int, func(int) bool) int{
		"super brute force":  ProblemFourA,
		"faster brute force": ProblemFourB,
	}
	tests := map[int]struct {
		input  int
		isPal  palindromeValidation
		output int
	}{
		1: {1, isIntPalindrome, 9},
		2: {1, isStrPalindrome, 9},
		3: {2, isIntPalindrome, 9009},
		4: {2, isStrPalindrome, 9009},
		5: {3, isIntPalindrome, 906609},
		6: {3, isStrPalindrome, 906609},
		7: {4, isIntPalindrome, 99000099},
		8: {4, isStrPalindrome, 99000099},
	}

	for solutionName, solutionFunc := range solutions {
		for _, ts := range tests {
			got := solutionFunc(ts.input, ts.isPal)
			expected := ts.output
			if got != expected {
				t.Errorf("%s solution for %d-digit number give wrong answer: %d, expected: %d",
					solutionName, ts.input, got, expected)
			}
		}
	}
}

func BenchmarkProblemThreeA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourA(3, isIntPalindrome)
	}
}

func BenchmarkProblemThreeA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourA(3, isStrPalindrome)
	}
}

func BenchmarkProblemThreeB1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourB(3, isIntPalindrome)
	}
}

func BenchmarkProblemThreeB2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourB(3, isStrPalindrome)
	}
}
