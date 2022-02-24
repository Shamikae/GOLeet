package main

import (
	"fmt"
	"GOLeet/mypackage"
)

// var nums = [3]int{2,2,1}



func main() {

	var nums = []int{2,2,1,1,3}
	fmt.Println("Hello, Module")
	goLeet.PrintMessage()
	var xyz = singleNumber(nums)
	fmt.Println(xyz)
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}