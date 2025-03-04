package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

}

func twoSum(nums []int, target int) []int {
	var numbersMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num2 := target - nums[i]
		for key, value := range numbersMap {
			if num2 == value {
				return []int{key, i}
			}
		}
		numbersMap[i] = nums[i]
	}

	return nil
}
