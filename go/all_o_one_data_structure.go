package main

// question: https://leetcode.com/problems/all-oone-data-structure/
type AllOne struct {
	// key-value
	m1 map[string]int
	// value-keys
	m2  map[int]map[string]int
	max int
	min int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{m1: make(map[string]int), m2: map[int]map[string]int{}, max: 0, min: 0}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if _, contained := this.m1[key]; contained {
		this.m1[key]++
	} else {
		this.m1[key] = 1
	}
	if _, contained := this.m2[this.m1[key]]; !contained {
		this.m2[this.m1[key]] = make(map[string]int)
	}
	this.m2[this.m1[key]][key] = 1
	if this.m1[key]-1 > 0 {
		delete(this.m2[this.m1[key]-1], key)
	}
	if this.m1[key] > this.max {
		this.max++
	}
	if len(this.m2[this.min]) == 0 {
		this.min++
	}
	if len(this.m2[1]) != 0 {
		this.min = 1
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if _, contained := this.m1[key]; contained {
		this.m1[key]--
		if this.m1[key] == 0 {
			delete(this.m1, key)
		} else {
			this.m2[this.m1[key]][key] = 1
		}
		delete(this.m2[this.m1[key]+1], key)
		if len(this.m2[this.max]) == 0 {
			this.max--
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.max == 0 {
		return ""
	}
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.min == 0 {
		return ""
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
