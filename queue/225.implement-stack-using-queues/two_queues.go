package main

type MyStack2 struct {
	q1 []int
	q2 []int
}

func NewStack2() *MyStack2 {
	return &MyStack2{}
}

func (this *MyStack2) Push(x int) {
	this.q2 = append(this.q2, x)
	for len(this.q1) > 0 {
		f := this.q1[0]
		this.q1 = this.q1[1:]
		this.q2 = append(this.q2, f)
	}

	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack2) Pop() int {
	p := this.q1[0]
	this.q1 = this.q1[1:]

	return p
}

func (this *MyStack2) Top() int {
	return this.q1[0]
}

func (this *MyStack2) Empty() bool {
	return len(this.q1) == 0
}
