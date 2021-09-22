package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	input := Counter("zzzzccc44444UUUuu")
	fmt.Printf("Output: %v\n", input)
}


func Counter(s string) string {
	// Представляем строку полученную на вход в виде среза рун
	slice := []rune(s)
	// total - сохранение промежуточных результатов для итогового вывода
	var total string
	
	// Сортировка среза рун
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	
	for _, letter := range slice {
		// countLetter - переменная для подсчета количества совпадений рун
		var countLetter []rune
		
		if letter != 0 {
			for iEqual, equal := range slice {
				if letter == equal {
					countLetter = append(countLetter, equal)
					slice[iEqual] = rune(0)
				}
			}	
			total = total + string(letter) + strconv.Itoa(len(countLetter))
		}
	}
	return total
}