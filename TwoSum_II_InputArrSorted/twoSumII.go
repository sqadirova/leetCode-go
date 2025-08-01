// Two Sum II - Input Array is Sorted
package main

import "fmt"

func main() {
	//numbers = [1,2,3,4], target = 3
	//Output: [1,2]
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// two pointer solution
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		curSum := numbers[l] + numbers[r]
		if curSum > target {
			r--
		} else if curSum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}

// brute force solution
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

// hashmap solution
func twoSum2(numbers []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		tmp := target - numbers[i]
		if val, exists := mp[tmp]; exists {
			return []int{val, i + 1}
		}
		mp[numbers[i]] = i + 1
	}
	return []int{}
}
