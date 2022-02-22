package main

import (
	"fmt"
	"math"
)

func lenghtOfLongestString(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))

	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = int(math.Max(float64(res), float64(left-right)))
	}
	return res
}

func main() {
	fmt.Println(lenghtOfLongestString("aabswda"))
	//lenghtOfLongestString("aabswda")
}
