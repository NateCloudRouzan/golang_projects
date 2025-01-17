package main

import (
	"fmt"
	"sort"
	"time"
)

/*
I have 3 different ideas for implementation.

-- sort strings and then iterate over them
acb bca -> abc abc (then check if equal)
Complexity would be 2(O(N)LogN)  where N is the length of the string

--- Put strings into hashmap then compare (more memory intensive)
aabbcc -> a[2] b[2] c[2]
bbaacc -> a[2] b[2] c[2]
(then loop over the hashmaps to check if they are matching)
Complexity would be 2N + M (N is length of string, M is unique charachters [around 30])
* can quickly check the length of the keys

---- remove from string B whatever char you are on for string a
Kind of like a bubble sort
ABBBC     BACBB
^         ^
ABBBC     BACBB
^          ^

BBBC      BCBB
^         ^
Complexity of this solution is N^2 at worst case


I wrote out all 3 to and tested the time. if the strings are similar the fastest is actually the bubble sort
*/

func main() {
	_question := `Given two strings s and t, return true if t is an anagram of s, and false otherwise. 
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.`

	fmt.Printf("Question: %v\n\n\n", _question)
	_randomString := "ahsfjklashdfgujkhsjklaghuirhguiaerhguohbafdnjkegbhojurnaeouienuhntayuibhnyuiaeuiytaiuhbnjuiklnerauihrguaeogj"
	_randomStringReverse := reverse(_randomString)

	_reallyLongRandomString := _randomString
	for i := 0; i < 10000; i++ {
		_reallyLongRandomString += _randomString
	}

	_reallyLongReversString := reverse(_reallyLongRandomString)

	_testQueue := []struct {
		name           string
		inputA         string
		inputB         string
		expectedOutput bool
	}{
		{
			name:           "Empty input",
			inputA:         "",
			inputB:         "",
			expectedOutput: true,
		},
		{
			name:           "doesnt match",
			inputA:         _randomString,
			inputB:         "iasjgiohjr8oeuahg8o9uaerthyuiogheruiohuivrehaugheruhi",
			expectedOutput: false,
		},
		{
			name:           "simple anagram",
			inputA:         "abcd",
			inputB:         "dcba",
			expectedOutput: true,
		},
		{
			name:           "simple anagram",
			inputA:         "tacotacotaco",
			inputB:         "octaoctaocta",
			expectedOutput: true,
		},

		{
			name:           "same string",
			inputA:         _randomString,
			inputB:         _randomString,
			expectedOutput: true,
		},
		{
			name:           "random string forward and backward",
			inputA:         _randomString,
			inputB:         _randomStringReverse,
			expectedOutput: true,
		},

		{
			name:           "same string but really long",
			inputA:         _reallyLongRandomString,
			inputB:         _reallyLongRandomString,
			expectedOutput: true,
		},

		{
			name:           "really long reverses of one another",
			inputA:         _reallyLongRandomString,
			inputB:         _reallyLongReversString,
			expectedOutput: true,
		},
	}

	for _, in := range _testQueue {
		fmt.Printf("Running Test Case: %v\n", in.name)

		// sorted approach
		sortStart := time.Now()
		out := isAnagramSort(in.inputA, in.inputB)
		dur := time.Now().Sub(sortStart).Milliseconds()
		if out != in.expectedOutput {
			fmt.Printf("\tSorted Approach Fail : (Unexpected return) %v\n", in.name)
		} else {
			fmt.Printf("\tSorted Approach PASS : %v Time elapsed %vms\n", in.name, dur)
		}

		// map approach
		mapStart := time.Now()
		out = isAnagramMap(in.inputA, in.inputB)
		dur = time.Now().Sub(mapStart).Milliseconds()
		if out != in.expectedOutput {
			fmt.Printf("\tMap Approach   Fail : (Unexpected return) %v\n", in.name)
		} else {
			fmt.Printf("\tMap Approach    PASS : %v Time elapsed %vms\n", in.name, dur)
		}

		// bubble approach
		bubbleStart := time.Now()
		out = isAnagramBubble(in.inputA, in.inputB)
		bubbleEnd := time.Now()
		dur = bubbleEnd.Sub(bubbleStart).Milliseconds()
		if out != in.expectedOutput {
			fmt.Printf("\tBubble Approach Fail : (Unexpected return) %v\n", in.name)
		} else {
			fmt.Printf("\tBubble Approach PASS : %v Time elapsed %vms\n", in.name, dur)
		}

	}

}

func isAnagramMap(a string, b string) bool {
	// if length isnt the same then return false
	if len(a) != len(b) {
		return false
	}

	anagramMap := make(map[rune]int, 0)

	for _, c := range a {
		anagramMap[c]++
	}

	for _, c := range b {
		anagramMap[c]--
	}

	// now look for any nonzero values
	for _, val := range anagramMap {
		if val != 0 {
			return false
		}
	}
	return true
}

func isAnagramSort(a string, b string) bool {
	// if length isnt same return false
	if len(a) != len(b) {
		return false
	}

	s1 := []rune(a)
	sort.Slice(s1, func(i int, j int) bool { return s1[i] < s1[j] })

	s2 := []rune(b)
	sort.Slice(s2, func(i int, j int) bool { return s2[i] < s2[j] })

	return string(s1) == string(s2)
}

func isAnagramBubble(a string, b string) bool {
	// if length isnt the same then return false
	if len(a) != len(b) {
		return false
	}
	newb := []rune(b)
	//fmt.Println(newb)
	for _, aChar := range a {
		found := false
		//fmt.Printf("searching for %v len of b = %v a : %v b : %v\n", string(aChar), len(newb), a, string(newb))
		for i, bChar := range newb {
			if aChar == bChar {
				found = true
				if len(b) < 2 { // if you reach the last char then you have suceeded
					return true
				}
				if i == 0 {
					newb = newb[1:]
				} else if i == (len(newb) - 1) {
					//fmt.Println("in the ending case")
					_updated := newb[0 : len(newb)-1]
					newb = _updated
				} else {
					newb = append(newb[0:i], newb[i+1:]...)
				}
				break
			}

		}
		if !found {
			//fmt.Printf("final b per bubble : %v\n", b)
			return false
		}
	}

	//fmt.Printf("final b per bubble : %v\n", b)
	return len(newb) == 0
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
