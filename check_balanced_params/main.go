package main

import (
	"fmt"
	"time"	
)
	

func main() {
	_question := `Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.`


fmt.Printf("Question: %v\n\n\n", _question)

_testQueue := []struct {
	name string
	input string
	expectedOutput bool


} {
	{
		name: "Empty input",
		input : "",
		expectedOutput: false,
	}, 
	{
		name: "invalid characters",
		input : "aa",
		expectedOutput: false,
	}, 
	{
		name: "single open",
		input : "(",
		expectedOutput: false,
	}, 
	{
		name: "odd number of parens []]",
		input : "[]]",
		expectedOutput: false,
	}, 
	{
		name: "single pair ()",
		input : "()",
		expectedOutput: true,
	}, 
	{
		name: "single pair {}",
		input : "{}",
		expectedOutput: true,
	}, 
	{
		name: "single pair []",
		input : "[]",
		expectedOutput: true,
	}, 
	{
		name: "embedded {{[]}}",
		input : "{{[]}}",
		expectedOutput: true,
	}, 
	{
		name: "doubly embedded {{[][]}}",
		input : "{{[][]}}",
		expectedOutput: true,
	}, 
	{
		name: "badly embedded {[{][]}}",
		input : "{[{][]}}",
		expectedOutput: false,
	}, 
	{
		name: "never closes [[[[[[[[[[[[[[[[[",
		input : "[[[[[[[[[[[[[[[[[",
		expectedOutput: false,
	}, 
	{
		name: "starts with close )()",
		input : ")()",
		expectedOutput: false,
	}, 
}

for _, in := range _testQueue {
	start := time.Now()
	out := isBalanced(in.input)
	t := time.Now()
	if out != in.expectedOutput {
		fmt.Printf("Fail : (Unexpected return) %v\n",in.name)
	} else {
		fmt.Printf("PASS : %v Time elapsed %v\n", in.name, t.Sub(start) )
	}
}



}

func isBalanced(s string) bool {
	// if no chars it cannot be balanced
	if len(s) < 2  || len(s) % 2 != 0{
		return false
	}
	_openQueue := ""

	openToCloseMap := map[byte]byte{
		'{': '}',
		'(': ')',
		'[' : ']',
	}

	for i := range s {
		var mostRecentOpen byte 
		if len(_openQueue) > 0 {
			mostRecentOpen = _openQueue[len(_openQueue) - 1 ]// set to whatever is most recent in the queue
		}

		// if its an opener add to queue
		_, inMap := openToCloseMap[s[i]]
		if inMap { 			// this means its an opener
			_openQueue += string(s[i])
			mostRecentOpen = s[i]
		} else if openToCloseMap[mostRecentOpen] != s[i] { // if it isnt returning what the most recent opener is then mal formatted
			fmt.Printf("\tgot a rogue close param that has no open.(expected=%v actual=%v)\n", string(openToCloseMap[mostRecentOpen]) , string(s[i]))
			return false
		}else { // remove from the queue
			//fmt.Println("\tclosing param and removing value")
			_openQueue = _openQueue[0:len(_openQueue)-1]
		} 
		//fmt.Printf("\tcur char : %v outstanding open : %v, bank %v \n", string(s[i]), string(mostRecentOpen),_openQueue)
	}  

	return len(_openQueue) == 0 // if string is empty 
}

