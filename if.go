// if else in go can be simple like in c++, but there's better ways to do it

package main

import "fmt"

func main() {
	// basic example
	if 3%2 == 0 {
		fmt.Println("3 is even")
	} else {
		fmt.Println("3 is odd")
	}

	// logical opperators are the same as c++  && AND, || OR
	
	// statements can come before the conditional and all variables are available in the scope of the if statement.
	if i:=true; i == false {
		fmt.Println("i is false")
	} else {
		fmt.Println("i is true")
	}
}