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
		//11: {6, isIntPalindrome, 999000000999},
		//12: {6, isStrPalindrome, 999000000999},
		//13: {7, isIntPalindrome, 99956644665999},
		//14: {7, isStrPalindrome, 99956644665999},
		//15: {8, isIntPalindrome, 9999000000009999},
		//16: {8, isStrPalindrome, 9999000000009999},
		//17: {9, isIntPalindrome, 999900665566009999},
		//18: {9, isStrPalindrome, 999900665566009999},
		//19: {10, isIntPalindrome, 99999000000000099999},
		//20: {10, isStrPalindrome, 99999000000000099999},
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
//BenchmarkProblemThreeA1
//BenchmarkProblemThreeA1-8   	    1276	    839759 ns/op
//BenchmarkProblemThreeA2
//BenchmarkProblemThreeA2-8   	     138	   8296592 ns/op
//BenchmarkProblemThreeB1
//BenchmarkProblemThreeB1-8   	   20264	     58269 ns/op
//BenchmarkProblemThreeB2
//BenchmarkProblemThreeB2-8   	    2478	    575498 ns/op
//BenchmarkProblemThreeC1
//BenchmarkProblemThreeC1-8   	   22626	     52154 ns/op
//BenchmarkProblemThreeC2
//BenchmarkProblemThreeC2-8   	   10000	    126212 ns/op
//PASS
