package main

// question: https://leetcode.com/problems/maximum-frequency-stack/
type FreqStack struct {
	// key: frequency, value: list of number with the same frequency
	stackMap map[int][]int
	// key: number, value: frequency
	freqMap map[int]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{stackMap: make(map[int][]int), freqMap: make(map[int]int), maxFreq: 0}
}

func (this *FreqStack) Push(x int) {
	freq, contained := this.freqMap[x]
	if contained {
		freq++
	} else {
		freq = 1
	}
	this.freqMap[x] = freq
	if freq > this.maxFreq {
		this.maxFreq = freq
	}

	nums, contained := this.stackMap[freq]
	if contained {
		this.stackMap[freq] = append(nums, x)
	} else {
		this.stackMap[freq] = []int{x}
	}
}

func (this *FreqStack) Pop() int {
	if len(this.stackMap) == 0 {
		return -1
	}
	nums, _ := this.stackMap[this.maxFreq]
	res := nums[len(nums)-1]
	if len(nums) == 1 {
		delete(this.stackMap, this.maxFreq)
		this.maxFreq--
	} else {
		this.stackMap[this.maxFreq] = nums[:len(nums)-1]
	}

	this.freqMap[res] = this.freqMap[res] - 1

	return res
}
