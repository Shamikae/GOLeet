/*
Problem Statement
Given a sorted array of numbers, remove the duplicates in-place such that each element appear only once and return the new length.

Example 1
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.

Example 2
Given nums = [0,0,1,1,1,2,2,3,3,4],
Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
It doesn't matter what values are set beyond the returned length.
*/

////////////NOTES://///////////
/*

-len() function is used to find the length of an array
-A slice is a dynamically-sized array 
	The type []T is a slice with elements of type T.
	A slice is formed by specifying two indices, a low and high bound, separated by a colon:
			a[low : high]  ex: a[1:4] 
*/
package main
import "fmt"

func main() {
	// fmt.Println(removeDuplicates([]int{1}))
	// fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,3,3,4}))
	// fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8, 9}))
	fmt.Println(SuperRemoveDuplicates([]int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8, 9}))
}


func removeDuplicates(nums []int) []int {
	//Create variable to store array length
	n := len(nums)

	//If n <= 1 return array values from 0 to length 
	// if n <= 1 {
	// 	return nums[0:n]
	// }
	j := 1
	
	//iterate through array starting at second index [1]
	for i := 1; i < n; i++ {
		//check index value if not equal to previous index value[i-1]
		
		if nums[i] != nums[i-1] {
			//add to the array
			nums[j] = nums[i]
			//increment j to move to next value 
			j++
			
		} 
		
	}
//return nums array values from 0 to j [the highest index]  
	return nums[0:j]
}

func SuperRemoveDuplicates(nums []int) int {
    const k = 2
    if len(nums) <= k {
        return len(nums)
    }
    i := k
    for j := k; j < len(nums); j++ {
        if nums[j] != nums[i-k] {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}