package main

import "fmt"

func CompareBackSpac(s, t string) bool {
	newS := validate(s)
	newT := validate(t)

	if len(newT) != len(newS) {
		return false
	} else {
		for i := 0; i < len(newS); i++ {
			if newT[i] != newS[i] {
				return false
			}
		}
	}
	return true
}

func validate(s string) []string {
	newS := []string{}
	for _, v := range s {
		if string(v) != "#" {
			newS = append(newS, string(v))
		} else {
			newS = newS[:len(newS)-1]
		}
	}
	return newS
}

func main() {
	s := "ab#C"
	t := "ab#c"
	fmt.Println(CompareBackSpac(s, t))
}
