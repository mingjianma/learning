package xtimewheel

import (
	"container/list"
	"container/ring"
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

//时间轮算法实现
var (
	ErrTimeWheelStop = errors.New("时间轮已停止")
)

type BucketItem struct {
	interval time.Duration
	value    interface{}
}

type TimeWheelCallback func(interface{}) bool

type TimeWheel struct {
	closeWait sync.WaitGroup

	interval       time.Duration
	ticker         *time.Ticker
	cb             TimeWheelCallback
	buckets        *ring.Ring
	currentPos     int
	requestChannel chan interface{}
	quitChannel    chan interface{}
	stopFlag       int32
}

func NewTimeWheel(interval time.Duration, bucketCount int, cb TimeWheelCallback) *TimeWheel {
	if cb == nil {
		return nil
	}

	tw := &TimeWheel{
		interval:       interval,
		buckets:        ring.New(bucketCount),
		cb:             cb,
		requestChannel: make(chan interface{}),
		quitChannel:    make(chan interface{}),
	}

	for i := 0; i < tw.buckets.Len(); i++ {
		tw.buckets.Value = list.New()
		tw.buckets = tw.buckets.Next()
	}
	return tw
}

func (tw *TimeWheel) Start() *TimeWheel {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.run()
	return tw
}

func (tw *TimeWheel) run() {
	for {
		select {
		case <-tw.quitChannel:
			tw.ticker.Stop()
			tw.cleanBucketsTasks()
			return
		case <-tw.ticker.C:
			tw.rangeBucket()
		case item := <-tw.requestChannel:
			if bitem, ok := item.(*BucketItem); ok && bitem != nil {
				ibucket := tw.buckets.Move(int(bitem.interval.Nanoseconds()) / int(tw.interval))
				if ilist, ok := ibucket.Value.(*list.List); ok && ilist != nil {
					ilist.PushBack(bitem)
				}
			}
		}
	}
}

func (tw *TimeWheel) cleanBucketsTasks() {
	defer tw.closeWait.Done()
	for i := 0; i < tw.buckets.Len(); i++ {
		tw.rangeBucket()
	}
}

func (tw *TimeWheel) rangeBucket() {
	bucket := tw.buckets.Value.(*list.List)
	var next *list.Element
	for curr := bucket.Front(); curr != nil; curr = next {
		if bitem, ok := curr.Value.(*BucketItem); ok && bitem != nil {
			next = curr.Next()
			func() {
				defer func() {
					r := recover()
					if r != nil {
						log.Printf("%v", r)
					}
				}()
				if tw.cb(bitem.value) {
					bucket.Remove(curr)
				}
			}()
		}
	}
	tw.buckets = tw.buckets.Next()
}

func (tw *TimeWheel) Stop() {
	if !atomic.CompareAndSwapInt32(&tw.stopFlag, 0, 1) {
		return
	}
	tw.closeWait.Add(1)
	close(tw.quitChannel)
	tw.closeWait.Wait()
}

func (tw *TimeWheel) Add(interval time.Duration, value interface{}) error {
	if atomic.LoadInt32(&tw.stopFlag) == 0 {
		tw.requestChannel <- &BucketItem{
			interval: interval,
			value:    value,
		}
		return nil
	} else {
		return ErrTimeWheelStop
	}
}
