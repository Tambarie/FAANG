package main

import (
	"fmt"
	"regexp"
)

func ValidPalindrome(s string) {
	re := regexp.MustCompile(`a(x*)b`)
	newtr := re.ReplaceAllString(s, "D")
	fmt.Println(newtr)
}

func main() {
	word := "A man, a plan, a canal: Panama"

	ValidPalindrome(word)
	//fmt.Print(strings.Trim("¡¡¡Hello,Gophers!!! @ ", " @"))
}
