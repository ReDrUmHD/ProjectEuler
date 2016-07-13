package main

import (
	"fmt"
)

func main() {
	var total int

	//Stores multiples of three up to 1000 in a slice called "nums"
	var nums []int
	for i := 0; i < 1000; i += 3 {
		nums = append(nums, i)
	}

	//Stores multiples of five up to 1000 that are not already in the slice "nums"
	for v := 0; v < 1000; v += 5 {
		if isIn(v, nums) == false {
			nums = append(nums, v)
		}
	}
	fmt.Println(addTogether(nums, total))
}

//Checks to see if the value is already in slice
func isIn(value int, list []int) bool {
	for _, e := range list {
		if e == value {
			return true
		}
	}
	return false
}

//Loops through the slice "nums" and all values together, storing it to "total"
func addTogether(list []int, storage int) int {
	for _, f := range list {
		storage += f
	}
	return storage
}
