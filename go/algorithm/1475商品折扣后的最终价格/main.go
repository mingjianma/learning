package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	var total []int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				total = append(total, prices[i]-prices[j])
				break
			}
		}
		if i+1 > len(total) {
			total = append(total, prices[i])
		}
	}
	return total
}

func main() {
	var prices1 = []int{8, 4, 6, 2, 3}
	var prices2 = []int{1, 2, 3, 4, 5}
	var prices3 = []int{10, 1, 1, 6}
	fmt.Println(finalPrices(prices1)) // 4,2,4,2,3
	fmt.Println(finalPrices(prices2)) // 1,2,3,4,5
	fmt.Println(finalPrices(prices3)) // 9,0,1,6
}
