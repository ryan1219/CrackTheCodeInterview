package main

import (
	"fmt"
	"strconv"
)

const (
	INITIAL_SIZE = 16
	LOAD_fACTOR  = 0.75
)

// question: https://leetcode.com/problems/design-hashmap/
type MyHashMap struct {
	size    int
	buckets []*node
}

type node struct {
	key   int
	value int
	next  *node
}

func (this node) String() string {
	return "{" + strconv.Itoa(this.key) + ":" + strconv.Itoa(this.value) + "}"
}

func (this MyHashMap) String() string {
	s := "["
	for i, v := range this.buckets {
		if i > 0 {
			s += ", "
		}
		for v != nil {
			s += fmt.Sprintf("%v->", v)
			v = v.next
		}
	}
	return s + "]"
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.buckets)
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]*node, INITIAL_SIZE)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	// double the size of bucket when load factor ratio > 0.75
	if float64(this.size)/float64(len(this.buckets)) > LOAD_fACTOR {
		newBuckets := make([]*node, 2*len(this.buckets))
		// rehash all values from old buckets to new buckets where the order is maintained for old hash == new hash
		for _, val := range this.buckets {
			if val == nil {
				continue
			}
			for val != nil {
				hashCode := val.key % len(newBuckets)
				head := newBuckets[hashCode]
				if head == nil {
					newBuckets[hashCode] = &node{key: val.key, value: val.value}
				} else {
					for head.next != nil {
						head = head.next
					}
					head.next = &node{key: val.key, value: val.value}
				}
				val = val.next
			}
		}
		this.buckets = newBuckets
	}

	hashCode := this.hash(key)
	head := this.buckets[hashCode]
	if head == nil {
		this.buckets[hashCode] = &node{key: key, value: value}
	} else {
		for head.key != key && head.next != nil {
			head = head.next
		}
		if head.key == key {
			head.value = value
		} else {
			head.next = &node{key: key, value: value}
		}
	}

	this.size++
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashCode := this.hash(key)
	if this.buckets[hashCode] == nil {
		return -1
	}

	head := this.buckets[hashCode]
	for head != nil {
		if head.key == key {
			return head.value
		}
		head = head.next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashCode := this.hash(key)

	if this.buckets[hashCode] == nil {
		return
	}

	cur := this.buckets[hashCode]
	next := cur.next
	if cur.key == key {
		this.buckets[hashCode] = cur.next
	} else {
		for next != nil && next.key != key {
			cur = next
			next = cur.next
		}

		if next == nil {
			return
		}

		cur.next = next.next
		next.next = nil
	}

	this.size--
}
