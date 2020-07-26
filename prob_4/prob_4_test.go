package prob4

import "testing"

func TestProblemFour(t *testing.T) {
	solutions := map[string]func(int, palindromeValidator) int{
		//"brute force from bottom": ProblemFourA,
		"brute force from upper": ProblemFourB,
		"with 11 divisible":      ProblemFourC,
	}
	palindromeValidators := []palindromeValidator{
		isIntPalindrome,
		isStrPalindrome,
	}
	tests := map[int]int{
		1: 9,
		2: 9009,
		3: 906609,
		4: 99000099,
		5: 9966006699,
		6: 999000000999,
		7: 99956644665999,
		//8: 9999000000009999,
	}

	for solutionName, solutionFunc := range solutions {
		for input, expected := range tests {
			for _, validator := range palindromeValidators {
				got := solutionFunc(input, validator)
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
