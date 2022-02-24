package main
import "fmt"

var nums = []int{2,2,1,3,1}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func main(){
	fmt.Println(singleNumber(nums))
}