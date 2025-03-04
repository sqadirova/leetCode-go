package main

import (
	"fmt"
	"math"
)

func main() {
	points1 := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	points2 := [][]int{{3, 2}, {-2, 2}}

	fmt.Println(minTimeToVisitAllPoints(points1)) // Output: 7
	fmt.Println(minTimeToVisitAllPoints(points2)) // Output: 5
}

func minTimeToVisitAllPoints(points [][]int) int {
	totalTime := 0

	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1][0], points[i-1][1]
		x2, y2 := points[i][0], points[i][1]

		//compute the time required to move from (x1, y1) to (x2, y2)
		totalTime += int(math.Max(math.Abs(float64(x2-x1)), math.Abs(float64(y2-y1))))
	}

	return totalTime
}
