package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))

}

func removeDuplicates(nums []int) int {
	nextIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[nextIndex] = nums[i]
			nextIndex++
		}
	}

	return nextIndex
}
