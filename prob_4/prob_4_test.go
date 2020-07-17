package prob4

import "testing"

type palindromeValidation func(int) bool

func TestProblemFour(t *testing.T) {
	solutions := map[string]func(int, func(int) bool) int{
		"brute force from bottom": ProblemFourA,
		"brute force from upper":  ProblemFourB,
		"with 11 divisible":       ProblemFourC,
	}
	tests := map[int]struct {
		input  int
		isPal  palindromeValidation
		output int
	}{
		1:  {1, isIntPalindrome, 9},
		2:  {1, isStrPalindrome, 9},
		3:  {2, isIntPalindrome, 9009},
		4:  {2, isStrPalindrome, 9009},
		5:  {3, isIntPalindrome, 906609},
		6:  {3, isStrPalindrome, 906609},
		7:  {4, isIntPalindrome, 99000099},
		8:  {4, isStrPalindrome, 99000099},
		9:  {5, isIntPalindrome, 9966006699},
		10: {5, isStrPalindrome, 9966006699},
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

func BenchmarkProblemThreeC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourC(3, isIntPalindrome)
	}
}

func BenchmarkProblemThreeC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProblemFourC(3, isStrPalindrome)
	}
}

//goos: linux
//goarch: amd64
//pkg: project-euler/prob_4
//BenchmarkProblemThreeA1-8   	      64	  18399625 ns/op
//BenchmarkProblemThreeA2-8   	     100	  10155324 ns/op
//BenchmarkProblemThreeB1-8   	    1054	   1161568 ns/op
//BenchmarkProblemThreeB2-8   	    1662	    632030 ns/op
//BenchmarkProblemThreeC1-8   	    3910	    264803 ns/op
//BenchmarkProblemThreeC2-8   	    6909	    178976 ns/op
//PASS
