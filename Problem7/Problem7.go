package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int64
	for x := 2; x > 1; x++ {
		if isPrime(int64(x)) == true {
			counter++
		}
		if counter == 10001 {
			fmt.Println(x)
			break
		}
	}
}

func isPrime(b int64) bool {
	if b <= 1 {
		return false
	}
	for a := quickSqrt(b); a >= 1; a-- {
		if a == 1 {
			return true
		}
		if b%a == 0 {
			return false
		}
	}
	return true
}

func quickSqrt(number int64) int64 {
	return int64(math.Sqrt(float64(number)))
}
