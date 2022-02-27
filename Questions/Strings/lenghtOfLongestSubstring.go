package main

import (
	"fmt"
	"math"
)

func lenghtOfLongestString(str string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(str))

	for left < len(str) {
		if idx, ok := indexes[str[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[str[left]] = left
		left++
		res = int(math.Max(float64(res), float64(left-right)))
	}
	return res
}

func main() {
	fmt.Println(lenghtOfLongestString("aabswda"))
	//lenghtOfLongestString("aabswda")
}
