package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix)) // Output: [1,2,3,4,8,12,11,10,9,5,6,7]
}

func spiralOrder(matrix [][]int) []int {
	var result []int

	for len(matrix) > 0 {
		// Step 1: Append the first row
		result = append(result, matrix[0]...)
		matrix = matrix[1:]

		// Step 2: Append last element of each row (right column)
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			for i := 0; i < len(matrix); i++ {
				result = append(result, matrix[i][len(matrix[i])-1])
				matrix[i] = matrix[i][:len(matrix[i])-1] // Remove last element
			}
		}

		// Step 3: Append the last row in reverse order
		if len(matrix) > 0 {
			lastRow := matrix[len(matrix)-1]
			for i := len(lastRow) - 1; i >= 0; i-- {
				result = append(result, lastRow[i])
			}
			matrix = matrix[:len(matrix)-1] // Remove last row
		}

		// Step 4: Append first element of each row in reverse order
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			for i := len(matrix) - 1; i >= 0; i-- {
				result = append(result, matrix[i][0])
				matrix[i] = matrix[i][1:] // Remove first element
			}
		}
	}

	return result
}
