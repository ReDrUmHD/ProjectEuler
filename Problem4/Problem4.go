package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var temp int
	nums := []int{}
	for x := 999; x >= 100; x-- {
		for y := 999; y >= 100; y-- {
			temp = x * y
			if strconv.Itoa(temp) == reverseString(strconv.Itoa(temp)) {
				nums = append(nums, temp)
			}
		}
	}
	meme := sort.IntSlice(nums)
	sort.Sort(meme)
	fmt.Println(meme[len(meme)-1])
}

func reverseString(a string) string {
	runes := []rune(a)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
