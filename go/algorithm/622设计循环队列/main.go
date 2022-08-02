package main

import (
	"fmt"
)

type element struct {
	value int
	next  *element
	pre   *element
}

type MyCircularQueue struct {
	size        int
	currentSize int
	first       *element
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currentSize >= this.size {
		return false
	}
	if this.first == nil {
		this.first = &element{
			value: value,
		}
		this.first.next = this.first
		this.first.pre = this.first
	} else {
		new := &element{
			value: value,
			next:  this.first,
			pre:   this.first.pre,
		}
		// 原来队尾的next指向新增元素
		this.first.pre.next = new
		// 原来队首的pre指针指向新增元素
		this.first.pre = new
	}
	this.currentSize++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.currentSize <= 0 {
		return false
	}
	if this.currentSize == 1 {
		this.first = nil
	} else {
		// 原来队尾的next指向队首的下一个元素
		this.first.pre.next = this.first.next
		// 原来队首下一个元素的pre指针指向队尾元素
		this.first.next.pre = this.first.pre
		// 更新队首元素
		this.first = this.first.next
	}

	this.currentSize--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.currentSize == 0 {
		return -1
	}
	return this.first.value
}

func (this *MyCircularQueue) Rear() int {
	if this.currentSize == 0 {
		return -1
	}
	return this.first.pre.value
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.currentSize == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyCircularQueue) IsFull() bool {
	if this.currentSize >= this.size {
		return true
	} else {
		return false
	}
}

func main() {
	// ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","deQueue","deQueue","isEmpty","isEmpty","Rear","Rear","deQueue"]
	// [[8],[3],[9],[5],[0],[],[],[],[],[],[],[]]
	var circularQueue = Constructor(8) // 设置长度为 3
	result := circularQueue.EnQueue(3) // 返回 true
	fmt.Println(result)
	result = circularQueue.EnQueue(9) // 返回 true
	fmt.Println(result)
	result = circularQueue.EnQueue(5) // 返回 true
	fmt.Println(result)
	result = circularQueue.EnQueue(0) // 返回 false，队列已满
	fmt.Println(result)
	result = circularQueue.DeQueue() // 返回 true
	fmt.Println(result)
	result = circularQueue.DeQueue() // 返回 true
	fmt.Println(result)
	result = circularQueue.IsEmpty() // 返回 true
	fmt.Println(result)
	result = circularQueue.IsEmpty() // 返回 true
	fmt.Println(result)
	resultInt := circularQueue.Rear() // 返回 3
	fmt.Println(resultInt)
	resultInt = circularQueue.Rear() // 返回 3
	fmt.Println(resultInt)
	result = circularQueue.DeQueue() // 返回 true
	fmt.Println(result)
}
