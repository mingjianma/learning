package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 32)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

func main() {
	fmt.Println(ChangeNumber(float64(1256)/float64(3600), 2))
	fmt.Printf("%.2f", float64(1256)/float64(3600))
}
