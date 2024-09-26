package main

import (
	"fmt"
	"sync"
)

func leftMostDigit(n int) int {
	if n < 0 {
		n = -n // Handle negative numbers
	}
	for n >= 10 {
		n /= 10 // Divide by 10 until n is less than 10
	}
	return n
}

//	func testGoRoutine(min, high, id int) {
//		fmt.Printf("Started Go routine with id: %d \n", id)
//		results := calculatePrimes(min, high)
//		for _, val := range results {
//			fmt.Printf("Routine id: %d | prime: %d \n", id, val)
//		}
//	}
func testGoRoutine(min, max, id int, receiver chan<- int) {
	fmt.Printf("Started Go routine with id: %d \n", id)
	result := calculatePrimes(min, max)
	for _, val := range result {
		receiver <- val
	}
}

func calculatePrimes(min, high int) []int {
	fmt.Printf("Calculating primes with min: %d \n", min)
	results := make([]int, 0)

	for num := min; num < high; num++ {
		if num < 2 {
			continue
		}
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			results = append(results, num)
		}
	}
	return results
}

func main() {
	var waiter sync.WaitGroup
	var userInput int
	receiver := make(chan int, 200)
	fmt.Printf("Set the max Prime number. \nMax: ")
	fmt.Scan(&userInput)

	intervals := leftMostDigit(userInput)
	intervalSize := 100
	currMin := 0
	currMax := intervalSize

	for i := 0; i <= intervals; i++ {
		waiter.Add(1)
		fmt.Printf("Ran in main function -> CurrMin: %d | currMax: %d | id: %d \n", currMin, currMax, i)

		go func(min, max, id int) {
			defer waiter.Done()
			fmt.Printf("Ran through the GoRoutine -> CurrMin: %d | currMax: %d | id: %d \n", min, max, id)
			testGoRoutine(min, max, id, receiver)
		}(currMin, currMax, i)

		currMin = currMax + 1
		currMax = currMax + intervalSize
		if currMax > userInput {
			currMax = userInput
		}
	}

	// Close the channel once all goroutines are done
	go func() {
		waiter.Wait()
		close(receiver)
	}()

	for val := range receiver {
		fmt.Printf("prime: %d \n", val)
	}

}
