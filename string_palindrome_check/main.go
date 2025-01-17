package main

import (
	"fmt"
	"time"
)

func main() {
	_question := `Implement a function to check if a given string is a palindrome. (same forward and backward)`

	fmt.Printf("Question: %v\n\n\n", _question)

	_testQueue := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: true,
		},
		{
			name:           "single char input",
			input:          "b",
			expectedOutput: true,
		},
		{
			name:           "two of same characters aa",
			input:          "aa",
			expectedOutput: true,
		},
		{
			name:           "two diff charachters ab",
			input:          "ab",
			expectedOutput: false,
		},
		{
			name:           "odd palindrome aba",
			input:          "aba",
			expectedOutput: true,
		},
		{
			name:           "longer palindrome odd racecar",
			input:          "racecar",
			expectedOutput: true,
		},
		{
			name:           "longer palindrome odd abbaabba",
			input:          "abbaabba",
			expectedOutput: true,
		},
		{
			name:           "long and isnt a palindrome",
			input:          "racingcar",
			expectedOutput: false,
		},
	}

	for _, in := range _testQueue {
		start := time.Now()
		out := isPalindrome(in.input)
		t := time.Now()
		if out != in.expectedOutput {
			fmt.Printf("Fail : (Unexpected return) %v\n", in.name)
		} else {
			fmt.Printf("PASS : %v Time elapsed %v\n", in.name, t.Sub(start))
		}
	}

}

func isPalindrome(s string) bool {
	if len(s) < 2 { // empty or single letter should return true
		return true
	}

	for i := 0; i < (len(s) / 2); i++ {
		//fmt.Printf("\tindex %v front : %v back : %v\n", i, string(s[i]), string(s[len(s) - i - 1]))

		// check if palindrome
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
