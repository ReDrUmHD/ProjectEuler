package main

import "fmt"

func main() {
	nums := []int{1, 2}
	a, b := 1, 2
	var total int

	for i := 0; i <= 4000000; {
		i = a + b
		nums = append(nums, i)
		a = b
		b = i
	}
	fmt.Println(nums[9])
	for _, x := range nums {
		if x%2 == 0 {
			total += x
		}
	}
	fmt.Println(total)
}
