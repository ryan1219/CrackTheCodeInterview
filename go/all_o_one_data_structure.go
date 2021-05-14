package main

// question: https://leetcode.com/problems/all-oone-data-structure/
/*
to achieve O(1) for all operations: https://leetcode.com/problems/all-oone-data-structure/discuss/91383/An-accepted-JAVA-solution-detailed-explanation.(HashMap-%2B-double-linked-list)
*/
type Node struct {
	prev  *Node
	next  *Node
	value int
	keys  map[string]int
}

func NewNode(value int, key string) Node {
	newNode := Node{keys: make(map[string]int)}
	newNode.value = value
	newNode.keys[key] = 1

	return newNode
}

func (this *Node) AddKey(key string) {
	this.keys[key] = 1
}

func (this *Node) RemoveKey(key string) {
	delete(this.keys, key)
}

type AllOne struct {
	head   *Node
	tail   *Node
	keyMap map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{keyMap: make(map[string]*Node)}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	_, contained := this.keyMap[key]
	if this.tail == nil {
		newNode := NewNode(1, key)
		this.head = &newNode
		this.tail = &newNode
		this.keyMap[key] = &newNode
	} else if !contained {
		if this.tail.value == 1 {
			this.tail.keys[key] = 1
			this.keyMap[key] = this.tail
		} else {
			newNode := NewNode(1, key)
			newNode.prev = this.tail
			this.tail.next = &newNode
			this.tail = &newNode
			this.keyMap[key] = &newNode
		}
	} else {
		// node := this.keyMap[key]

	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {

}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
