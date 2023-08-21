// https://leetcode.com/problems/lru-cache/

package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
	key  int
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
	} else {
		if len(this.cache) >= this.capacity {
			// Evict the least recently used item
			delete(this.cache, this.tail.key)
			if this.tail.prev != nil {
				this.tail.prev.next = nil
			}
			this.tail = this.tail.prev
		}
		newNode := &Node{
			key: key,
			val: value,
		}
		this.cache[key] = newNode
		this.moveToHead(newNode)
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		if node != this.head {
			if node.prev != nil {
				node.prev.next = node.next
			}
			if node.next != nil {
				node.next.prev = node.prev
			}
			node.prev = nil
			node.next = this.head
			this.head.prev = node
			this.head = node
			if node == this.tail {
				for this.tail != nil && this.tail.next != nil {
					this.tail = this.tail.next
				}
			}

		}
	}
}

func main() {
	lru_cache := Constructor(2)

	lru_cache.Put(1, 1)
	lru_cache.Put(2, 2)

	fmt.Println(lru_cache.Get(2))

	lru_cache.Put(3, 3)

	fmt.Println(lru_cache.Get(2)) // Output: -1 (not found)
	fmt.Println(lru_cache.Get(3)) // Output: 3
	fmt.Println("LRU Cache", lru_cache)
}
