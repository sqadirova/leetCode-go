// Top K Frequent Elements
package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	//freq = count:number
	for _, num := range nums {
		freq[num]++
	}
	// fmt.println(freq)

	buckets := make([][]int, len(nums)+1)
	for num, freqCount := range freq {
		buckets[freqCount] = append(buckets[freqCount], num)
	}

	result := []int{}

	for i := len(nums); i >= 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func main() {
	// Example 1:
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d, Output: %v (Expected: [1,2] - order may vary)\n", nums1, k1, topKFrequent(nums1, k1))

	// Example 2:
	nums2 := []int{1}
	k2 := 1
	fmt.Printf("Input: nums = %v, k = %d, Output: %v (Expected: [1])\n", nums2, k2, topKFrequent(nums2, k2))

	// Additional Test Cases:
	nums3 := []int{1, 2}
	k3 := 2
	fmt.Printf("Input: nums = %v, k = %d, Output: %v (Expected: [1,2] or [2,1])\n", nums3, k3, topKFrequent(nums3, k3))

	nums4 := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	k4 := 2
	fmt.Printf("Input: nums = %v, k = %d, Output: %v (Expected: [4,3] or [3,4])\n", nums4, k4, topKFrequent(nums4, k4))

	nums5 := []int{5, 5, 5, 5, 4, 4, 4, 3, 3, 2, 1}
	k5 := 3
	fmt.Printf("Input: nums = %v, k = %d, Output: %v (Expected: [5,4,3] in any order)\n", nums5, k5, topKFrequent(nums5, k5))
}
