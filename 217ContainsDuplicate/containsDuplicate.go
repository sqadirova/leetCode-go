package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // false
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true

}

func containsDuplicate(nums []int) bool {
	seenSet := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seenSet[num]; exists {
			return true
		}
		seenSet[num] = struct{}{}
	}

	return false
}
