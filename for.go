// "for" is the only looping keyword in GO

package main

import "fmt"

const ranger int = 5
func main() {
	// basic, no nonsense type
	// pass a conditional and loop until false.
	fmt.Println("loop 1")
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic format C++ variable/conditional/increment
	fmt.Println("loop 2")
	for j := 1; j < 3; j++ {
		fmt.Println(j)	
	}

	// iterate over a range
	fmt.Println("loop 3")
	for k := range ranger {
		fmt.Println(k)
	}
}