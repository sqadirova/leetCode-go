// Two Sum
package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

}

// hashmap one pass solution O(n) time complexity and O(n) space complexity
func twoSum(nums []int, target int) []int {
	numbersMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, found := numbersMap[diff]; found {
			return []int{j, i}
		}
		numbersMap[n] = i
	}
	return []int{}
}

// hashmap two pass solution O(n) time complexity and O(n) space complexity
//func twoSum(nums []int, target int) []int {
//	indices := make(map[int]int)
//
//	for i, n := range nums {
//		indices[n] = i
//	}
//
//	for i, n := range nums {
//		diff := target - n
//		if j, found := indices[diff]; found && j != i {
//			return []int{i, j}
//		}
//	}
//	return []int{}
//}

//O(n^2) time complexity and O(1) space complexity
//func twoSum(nums []int, target int) []int {
//	var numbersMap = make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		num2 := target - nums[i]
//		for key, value := range numbersMap {
//			if num2 == value {
//				return []int{key, i}
//			}
//		}
//		numbersMap[i] = nums[i]
//	}
//
//	return nil
//}
