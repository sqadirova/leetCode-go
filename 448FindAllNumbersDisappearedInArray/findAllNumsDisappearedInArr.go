package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums)) // Output: [5,6]

	nums2 := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums2)) // Output: [2]

}

func findDisappearedNumbers(nums []int) []int {
	var result []int
	numMap := make(map[int]bool)

	//store all numbers in map
	for _, num := range nums {
		numMap[num] = true
	}

	//check which number is missing
	for i := 1; i < len(nums)+1; i++ {
		if !numMap[i] {
			result = append(result, i)
		}
	}

	return result
}
