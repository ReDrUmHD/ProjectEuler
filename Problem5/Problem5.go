package main

import (
	"fmt"
)

func main() {
	for i := 40; i > 1; i++ {
		if isDivisible(i) == true {
			fmt.Println(i)
			break
		}
	}
}

func isDivisible(number int) bool {
	for x := 1; x <= 20; x++ {
		if number%x != 0 {
			return false
		}
	}
	return true
}
