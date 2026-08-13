package main

type MyStack1 struct {
	q []int
}

func NewStack1() *MyStack1 {
	return &MyStack1{}
}

func (this *MyStack1) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ {
		h := this.q[0]
		this.q = this.q[1:]
		this.q = append(this.q, h)
	}
}

func (this *MyStack1) Pop() int {
	h := this.q[0]
	this.q = this.q[1:]

	return h
}

func (this *MyStack1) Top() int {
	return this.q[0]
}

func (this *MyStack1) Empty() bool {
	return len(this.q) == 0
}
