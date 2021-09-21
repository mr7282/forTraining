package main

import (
	"fmt"
	// "sort"
)

// func Counter(s string) string {
// 	slice := []byte(s)
// 	return "ok"
// }

func main()  {
	slice := []byte("aaabbbccccc")

	type totalItem []byte
	var total []totalItem

	for iLetter, letter := range slice {
		for _, equal := range slice[iLetter + 1:] {
			var letterCount []byte
			if letter == equal {
				letterCount = append(letterCount, letter)
				slice = append(slice[:iLetter], slice[iLetter + 1:]...)

			}
			fmt.Printf("letterCount: %v\n", letterCount)
			total = append(total, letterCount)
		}
	}
	
	fmt.Printf("slice: %v\n", total)
}