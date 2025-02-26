package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums)) // Output: [4, 0, 1, 1, 3]
}

func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums) // copy the nums to sortedNums
	sort.Ints(sortedNums)  // sort the copy

	numMap := make(map[int]int)

	for i, num := range sortedNums {
		if _, exist := numMap[num]; !exist {
			numMap[num] = i
		}
	}

	// nums = [8, 1, 2, 2, 3]
	// sortedNums = [1, 2, 2, 3, 8]
	// numMap= {1: 0, 2: 1, 3: 3, 8: 4}   //num : i

	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = numMap[num]
	}

	return result
}
