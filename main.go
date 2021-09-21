package main

import (
	"fmt"
)

func main()  {
	slice := []byte("aaabbbccccc")

	type totalItem []byte
	var total []totalItem
	
	for iLetter, letter := range slice {
		var letterCount []byte
		for _, equal := range slice[iLetter + 1:] {
			if letter == equal  {
				letterCount = append(letterCount, equal)
				slice = append(slice[:iLetter], slice[iLetter + 1:]...) 
			}
			fmt.Printf("letterCount: %v\n", letterCount)
		}
		total = append(total, letterCount)
	}
	
	fmt.Printf("slice: %v\n", total)
}