package main

import (
	"fmt"
	"math"
)

// const declares  a constant value
//ex
const s string = "constant value here"

func main() {
	fmt.Println(s)

	// a const statement can appear anywhere a 'var' statement can
	const n = 5000000

	// Constant expressions perform arithmetic with arbitrary precission
	const d = 3e20 / n

	fmt.Println(d)

	// A numeric constant without specific type has no type until otherwise given one.
	fmt.Println(int64(d))

	// a numerical constant can be given a type also by using it in a specific context
	// that requires that type, thus making the constant that type in the process.
	fmt.Println(math.Sin(n))
}