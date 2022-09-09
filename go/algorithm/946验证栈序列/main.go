package main

import (
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	// 遍历数组 pushed，将 pushed 的每个元素依次入栈；
	// 每次将 pushed 的元素入栈之后，如果栈不为空且栈顶元素与 popped 的当前元素相同，则将栈顶元素出栈，同时遍历数组
	// popped，直到栈为空或栈顶元素与 popped 的当前元素不同。
	st := []int{}
	j := 0
	for _, x := range pushed {
		st = append(st, x)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}
