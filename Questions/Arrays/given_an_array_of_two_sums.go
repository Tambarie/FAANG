package main

import "fmt"

func TwoSums(nums []int, target int) []int {
	m := make(map[int]int)

	for k, value := range nums {
		if idx, ok := m[target-value]; ok {
			return []int{idx, k}
		}
		m[value] = k
	}
	return nil
}

func main() {
	arr := []int{1, 24, 5, 3, 2}
	fmt.Println(TwoSums(arr, 5))
}
