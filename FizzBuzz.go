// FizzBuzz Game
// Rules: Count up to a predetermined number. If
//		  the number is divisible by 3, say Fizz.
//	   	  If it is divisible by 5, say Buzz. If
// 		  it is divisible by both, then say
// 		  FizzBuzz.

package main

import (
	"fmt"
	"strconv"
)

// Array map of what we are testing
var m = make(map[string]int)

func main() {
	// Fizz and Buzz to the testing array map
	m["Fizz"] = 3
	m["Buzz"] = 5

	// Loop through our counter
	for i := 1; i <= 100; i++ {
		// Print out the return value from our function
		fmt.Println(buildString(i))
	}
}

func buildString(i int) string {

	// Start with a NULL output
	output := ""

	// Loop through the array maps to check the modulo
	for k, v := range m {

		// If the test passed them append the map key to the output.
		if i%v == 0 {
			output += k
		}
	}

	// If still blank then just pass the value back as a string
	if output == "" {
		output = strconv.Itoa(i)
	}

	// Return the output string
	return output
}
