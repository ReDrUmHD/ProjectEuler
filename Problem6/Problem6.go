package main

import "fmt"

func main() {
	var squares, second int
	for x := 1; x <= 100; x++ {
		squares += x * x
		second += x
	}
	second = second * second
	fmt.Println(second - squares)

}
