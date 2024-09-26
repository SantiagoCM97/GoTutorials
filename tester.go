package main

import (
	"fmt"
	"sort"
)

type TwoSum struct {
	cache []int
}

func Constructor() TwoSum {
	return TwoSum{
		cache: []int{},
	}
}

func (this *TwoSum) Add(number int) {
	this.cache = append(this.cache, number)
	sort.Ints(this.cache)
	return
}

//	func (this *TwoSum) Find(value int) bool {
//		for _, valA := range this.cache {
//			fmt.Println("Current a: %s", valA)
//			for _, valB := range this.cache {
//				fmt.Println("Current b: %s", valB)
//				if value == valA+valB {
//					return true
//				}
//			}
//		}
//		return false
//	}
func (this *TwoSum) Find(value int) bool {
	i := 0
	j := len(this.cache) - 1
	for i <= j {
		fmt.Println("Current i: %s", this.cache[i])
		fmt.Println("Current j: %s", this.cache[j])
		if this.cache[i]+this.cache[j] == value {
			return true
		} else if this.cache[i]+this.cache[j] > value {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	twoSum := Constructor()
	twoSum.Add(1)
	twoSum.Add(4)
	result := twoSum.Find(5)
	fmt.Println(result)
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
