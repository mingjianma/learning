package main

import (
	"fmt"
	"time"
)

//main()启动genJob获取存放任务的通道jobCh，然后创建retCh，它的缓存空间是200，并使用workerPool启动一个有5个协程的协程池。
//1s之后，关闭retCh，然后开始从retCh中读取协程池处理结果，并打印。
func main() {
	jobCh := genJob(10000)
	retCh := make(chan string, 10000)
	workerPool(5, jobCh, retCh)

	time.Sleep(time.Second)
	close(retCh)
	for ret := range retCh {
		fmt.Println(ret)
	}
}

//genJob启动一个协程，并生产n个任务，写入到jobCh。
func genJob(n int) <-chan int {
	jobCh := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()

	return jobCh
}

//workerPool()会创建1个简单的协程池，协程的数量可以通入参数n执行，并且还指定了jobCh和retCh两个参数。
func workerPool(n int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < n; i++ {
		go worker(i, jobCh, retCh)
	}
}

//worker()是协程池中的协程，入参分布是它的ID、job通道和结果通道。
//使用for-range从jobCh读取任务，直到jobCh关闭，然后一个最简单的任务：生成1个字符串，证明自己处理了某个任务，并把字符串作为结果写入retCh。
func worker(id int, jobCh <-chan int, retCh chan<- string) {
	cnt := 0
	for job := range jobCh {
		cnt++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", id, job, cnt)
		retCh <- ret
	}
}
