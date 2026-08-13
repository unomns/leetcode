package main

import "fmt"

func main() {
	c := Constructor()

	fmt.Println(c.Ping(1))
	fmt.Println(c.Ping(100))
	fmt.Println(c.Ping(3001))
	fmt.Println(c.Ping(3002))
}

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	n := 0
	for n < len(this.requests) && this.requests[n] < t-3000 {
		n++
	}

	this.requests = this.requests[n:]

	return len(this.requests)
}
