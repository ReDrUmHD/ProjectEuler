package main

import (
	"fmt"
)

func main() {
	var nums []int
	var total int
	//Loops 0-1000, storing integers that are multiples of 3 or 5 in slice "nums"
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			nums = append(nums, i)
		}
	}
	fmt.Println(addTogether(total, nums))
}

//Adds all integers stored in a slice
func addTogether(storage int, list []int) int {
	for _, a := range list {
		storage += a
	}
	return storage
}
