package main

import "math"

// question: https://leetcode.com/problems/all-oone-data-structure/
/*
to achieve O(1) for all operations: https://leetcode.com/problems/all-oone-data-structure/discuss/91383/An-accepted-JAVA-solution-detailed-explanation.(HashMap-%2B-double-linked-list)
*/
type AllOne struct {
	max int
	min int
	// key - count
	keyMap map[string]int
	// count - keys
	countMap map[int]map[string]int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{keyMap: make(map[string]int), countMap: make(map[int]map[string]int)}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if _, contained := this.keyMap[key]; !contained {
		this.keyMap[key] = 1
	} else {
		this.keyMap[key]++
	}

	// manage countMap
	// add new count in countMap
	if _, contained := this.countMap[this.keyMap[key]]; !contained {
		this.countMap[this.keyMap[key]] = make(map[string]int)
	}
	this.countMap[this.keyMap[key]][key] = 1
	// clean previous count in countMap
	if _, contained := this.countMap[this.keyMap[key]-1][key]; contained {
		delete(this.countMap[this.keyMap[key]-1], key)
		if len(this.countMap[this.keyMap[key]-1]) == 0 {
			delete(this.countMap, this.keyMap[key]-1)
		}
	}

	if this.keyMap[key] > this.max {
		this.max = this.keyMap[key]
	}
	if len(this.countMap[this.min]) == 0 {
		this.min++
	}
	if len(this.countMap[1]) != 0 {
		this.min = 1
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if _, contained := this.keyMap[key]; !contained {
		return
	} else {
		this.keyMap[key]--
	}
	if this.keyMap[key] == 0 {
		delete(this.keyMap, key)
	}

	// manage count map
	// clean previous count in countMap
	delete(this.countMap[this.keyMap[key]+1], key)
	if len(this.countMap[this.keyMap[key]+1]) == 0 {
		delete(this.countMap, this.keyMap[key]+1)
	}
	// add new count in countMap
	if this.keyMap[key] != 0 {
		if _, contained := this.countMap[this.keyMap[key]]; !contained {
			this.countMap[this.keyMap[key]] = make(map[string]int)
		}
		this.countMap[this.keyMap[key]][key] = 1
	}

	if len(this.countMap[this.max]) == 0 {
		this.max--
	}
	if this.keyMap[key] < this.min && this.keyMap[key] != 0 {
		this.min = this.keyMap[key]
	} else {
		this.min = math.MaxInt64
		for k, _ := range this.countMap {
			if k < this.min {
				this.min = k
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	for k := range this.countMap[this.max] {
		return k
	}

	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	for k := range this.countMap[this.min] {
		return k
	}

	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
