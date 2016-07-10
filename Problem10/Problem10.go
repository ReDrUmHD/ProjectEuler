package main

import (
	"fmt"
	"math"
)

func main() {
	var list []int
	for i := 2; i < 2000000; i++ {
		if isPrime(int64(i)) == true {
			list = append(list, i)
		}
	}
	fmt.Println(addSlice(list))
}

func isPrime(num int64) bool {
	if num <= 1 {
		return false
	}
	for i := quickSqrt(num); i >= 1; i-- {
		if i == 1 {
			return true
		}
		if num%i == 0 {
			return false
		}
	}
	return true
}

func quickSqrt(number int64) int64 {
	return int64(math.Sqrt(float64(number)))
}

func addSlice(list []int) int {
	stor := 0
	for _, v := range list {
		stor += v
	}
	return stor
}
