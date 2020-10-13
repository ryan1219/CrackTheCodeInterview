package main

type MinStack struct {
	list    []int
	minList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.minList) == 0 {
		this.minList = append(this.minList, x)
	} else if x < this.minList[len(this.minList)-1] {
		this.minList = append(this.minList, x)
	} else {
		this.minList = append(this.minList, this.minList[len(this.minList)-1])
	}

	this.list = append(this.list, x)
}

func (this *MinStack) Pop() {
	this.list = this.remove(this.list, len(this.list)-1)
	this.minList = this.remove(this.minList, len(this.minList)-1)
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	return this.minList[len(this.minList)-1]
}

func (this *MinStack) remove(src []int, index int) []int {
	return append(src[:index], src[index+1:]...)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
