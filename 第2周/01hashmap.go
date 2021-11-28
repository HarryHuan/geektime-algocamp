package main

import "math"

func main() {

}

type ListNode struct {
	key string
	val int
	nxt *ListNode
}

var hashmap [99991]*ListNode

func initialize() {
	for i := range hashmap {
		hashmap[i] = &ListNode{"", 0, nil}
	}
}

func hash(key string) int {
	l := len(key)
	d := float64(0)
	for i := 0; i < l; i++ {
		f := key[i] - 'a' + 1
		d += float64(f) * math.Pow(float64(27), float64(l-i-1))
	}
	in := int(d)
	k := in % 99991
	return k
}

func put(key string, val int) {
	k := hash(key)
	newNode := ListNode{key, val, nil}
	hashmap[k].nxt = &newNode
}

func find(key string) int {
	k := hash(key)
	node := hashmap[k]
	node = node.nxt
	for node != nil {
		if node.key == key {
			return node.val
		}
		node = node.nxt
	}
	return 0
}

func remove(key string) {
	k := hash(key)
	node := hashmap[k]
	for node.nxt != nil {
		if node.nxt.key == key {
			node.nxt = node.nxt.nxt
		}
		node = node.nxt
	}
}
