package main

import "fmt"

func main() {
	slice1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(buyStock(slice1))
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(buyStock(slice2))
	slice3 := []int{7, 6, 4, 3, 1}
	fmt.Println(buyStock(slice3))
}

func buyStock(slice []int) int {
	if len(slice) < 0 {
		return 0
	}
	var lastStock, result int
	for i := 0; i < len(slice); i++ {
		if lastStock == 0 {
			lastStock = slice[i]
			continue
		}
		if slice[i] > lastStock {
			result += slice[i] - lastStock
		}
		lastStock = slice[i]
	}

	return result
}
