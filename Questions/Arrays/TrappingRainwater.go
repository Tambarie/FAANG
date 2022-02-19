package main

import (
	"fmt"
)

//func TrappingRainwater(rainWater []int) int {
//	var totalRainWater  float64 = 0
//
//	for p := 0; p < len(rainWater); p++ {
//		pLeft := p
//		pRight := p
//		var maxLeft float64 = 0
//		var maxRight float64  = 0
//		for pLeft >= 0 {
//			maxLeft = math.Max(maxLeft,float64(rainWater[pLeft]))
//			pLeft--
//		}
//		for pRight < len(rainWater){
//			maxRight = math.Max(maxRight,float64(rainWater[pRight]))
//			pRight++
//		}
//		currentWater := math.Min(maxLeft,maxRight) - float64(rainWater[p])
//		if currentWater >=0{
//			totalRainWater += currentWater
//		}
//	}
//	return int(totalRainWater)
//}

//1. Identify the pointer with the lesser value
//2. Is the pointer greater than or equa to the max value on that side
// yes -> update max on the side
//no -> get water for pointer value and add to total
//3. Move pointer inwards
//4. Repeat for other pointer

func OptimalSolutionTrappingRainwater(rainWater []int) int {
	var totalRainWater int = 0
	leftPointer := 0
	rightPointer := len(rainWater) - 1
	leftMax := 0
	rightMax := 0

	for leftPointer < rightPointer {
		if rainWater[leftPointer] <= rainWater[rightPointer] {
			if rainWater[leftPointer] >= leftMax {
				leftMax = rainWater[leftPointer]
			} else {
				totalRainWater += leftMax - rainWater[leftPointer]
			}
			leftPointer++
		} else {
			if rainWater[rightPointer] >= rightMax {
				rightMax = rainWater[rightPointer]
			} else {
				totalRainWater += rightMax - rainWater[rightPointer]
			}
			rightPointer--
		}

	}

	return totalRainWater

}

func main() {
	rainwater := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
	// fmt.Println(TrappingRainwater(rainwater))
	fmt.Println(OptimalSolutionTrappingRainwater(rainwater))
}
