package prob4

import "testing"

type palindromeValidator func(int) bool

func TestProblemFour(t *testing.T) {
	solutions := map[string]func(int, func(int) bool) int{
		//"brute force from bottom": ProblemFourA,
		"brute force from upper": ProblemFourB,
		"with 11 divisible":      ProblemFourC,
	}
	palindromeValidators := []palindromeValidator{
		isIntPalindrome,
		isStrPalindrome,
	}
	tests := map[int]struct {
		validators []palindromeValidator
		output     int
	}{
		1: {palindromeValidators, 9},
		2: {palindromeValidators, 9009},
		3: {palindromeValidators, 906609},
		4: {palindromeValidators, 99000099},
		5: {palindromeValidators, 9966006699},
		6: {palindromeValidators, 999000000999},
		7: {palindromeValidators, 99956644665999},
		//8: {palindromeValidators, 9999000000009999},
	}

	for solutionName, solutionFunc := range solutions {
		for input, ts := range tests {
			for _, validator := range ts.validators {
				got := solutionFunc(input, validator)
				expected := ts.output
				if got != expected {
					t.Errorf("%s solution for %d-digit number give wrong answer: %d, expected: %d",
						solutionName, input, got, expected)
				}
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
//BenchmarkProblemThreeA1-8   	    1748	    636240 ns/op
//BenchmarkProblemThreeA2
//BenchmarkProblemThreeA2-8   	     474	   2635308 ns/op
//BenchmarkProblemThreeB1
//BenchmarkProblemThreeB1-8   	   20824	     60098 ns/op
//BenchmarkProblemThreeB2
//BenchmarkProblemThreeB2-8   	    5618	    187848 ns/op
//BenchmarkProblemThreeC1
//BenchmarkProblemThreeC1-8   	   20672	     55237 ns/op
//BenchmarkProblemThreeC2
//BenchmarkProblemThreeC2-8   	   16497	     73894 ns/op
//PASS
