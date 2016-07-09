package main

import (
	"fmt"
	"math"
)

func main() {
	var num int64 = 600851475143
	for i := quickSqrt(num); i >= 1; i-- {
		if num%i == 0 && checkPrime(i) == true {
			fmt.Println(i)
			break
		}
	}

}

func checkPrime(number int64) bool {
	if number <= 1 {
		return false
	}
	for b := quickSqrt(number); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if number%b == 0 {
			return false
		}
	}
	return true
}

func quickSqrt(number int64) int64 {
	return int64(math.Sqrt(float64(number)))
}
