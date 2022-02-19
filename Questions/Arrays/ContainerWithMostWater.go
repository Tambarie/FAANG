package main

import (
	"fmt"
	"math"
)

// ConatainerWithMostWater Normal solution
//func ConatainerWithMostWater(numbers []int) int {
//	var maxArea int = 0
//	for i := 0; i < len(numbers); i++ {
//		for j := i+1; j < len(numbers); j++ {
//			min := int(math.Min(float64(numbers[i]), float64(numbers[j])))
//			width := j - i
//			area := min * width
//			if area > maxArea{
//				maxArea = area
//			}
//		}
//	}
//	return maxArea
//}

// OptimalContainerWithMostWater Optimal Solution
func OptimalContainerWithMostWater(numbers []int) float64 {
	pointerOne := 0
	pointerTwo := len(numbers) - 1
	var maxArea float64 = 0

	for pointerOne < pointerTwo {
		width := float64(pointerTwo - pointerOne)
		height := math.Min(float64(numbers[pointerOne]), float64(numbers[pointerTwo]))
		area := width * height
		maxArea = math.Max(maxArea, area)
		if numbers[pointerOne] <= numbers[pointerTwo] {
			pointerOne++
		} else {
			pointerTwo--
		}
	}
	return maxArea
}

func main() {
	numbers := []int{1, 4, 3, 5}
	//fmt.Println(ConatainerWithMostWater(numbers))
	fmt.Println(OptimalContainerWithMostWater(numbers))
}
