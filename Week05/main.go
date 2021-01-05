package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Bucket struct {
	success int
	failed int
	lastTime time.Time
}

type WindowsBucket struct {
	maxSize int
	Buckets [10]Bucket
	head    int
	tail    int
}

func main()  {
	WindowsBuckets := WindowsBucket{
		maxSize: 10,
		Buckets: [10]Bucket{},
		head: 0,
		tail: 0,
	}
	ticker := time.NewTicker(time.Second * 1)
	go func(WindowsBuckets WindowsBucket ) {
		for  {
			select {
			case <-ticker.C:
				success := rand.Intn(1000)
				failed := rand.Intn(1000)
				lastTime := time.Now()
				_ = WindowsBuckets.push(Bucket{success,failed,lastTime})
				fmt.Println(WindowsBuckets.size())
			}
		}
	}(WindowsBuckets)
	select {}
}


func (this *WindowsBucket) push(val Bucket) (err error) {
	if this.isFull() {
		return errors.New("队列满了")
	}
	this.Buckets[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

func (this *WindowsBucket) pop() (val Bucket, err error) {
	if this.isEmpty() {
		return Bucket{}, errors.New("队列是空的")
	}
	//取出,head指向队首并且含队首元素
	this.head = (this.head + 1) % this.maxSize
	return
}
func (this *WindowsBucket) size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func (this *WindowsBucket) isFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

func (this *WindowsBucket) isEmpty() bool {
	return this.tail == this.head
}


