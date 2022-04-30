package design

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

// Implement the LRUCache class:

// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
// If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

type LinkedNode struct {
	key   int
	value int
	pre   *LinkedNode
	next  *LinkedNode
}

type LRUCache struct {
	hashMap   map[int]*LinkedNode
	head      *LinkedNode
	tail      *LinkedNode
	available int
}

func Constructor(capacity int) LRUCache {
	hashMap := make(map[int]*LinkedNode, capacity)
	head := &LinkedNode{}
	tail := &LinkedNode{}
	head.next = tail
	tail.pre = head
	cache := LRUCache{
		hashMap:   hashMap,
		head:      head,
		tail:      tail,
		available: capacity,
	}
	return cache
}

// h -> 1 -> 2 -> 3 -> t
func (this *LRUCache) Get(key int) int {
	cache, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	// 拿到返回值
	value := cache.value
	// key对应的节点脱离
	cache.pre.next = cache.next
	cache.next.pre = cache.pre
	// 放到队头
	cache.next = this.head.next
	cache.pre = this.head
	cache.next.pre = cache
	cache.pre.next = cache
	return value
}

func (this *LRUCache) Put(key int, value int) {
	cache, ok := this.hashMap[key]
	if ok {
		cache.value = value
		// k对应的
		pre := cache.pre
		next := cache.next
		// 取出cache
		pre.next = next
		next.pre = pre
		// 放到队头
		sec := this.head.next
		cache.next = sec
		cache.pre = this.head
		sec.pre = cache
		this.head.next = cache
	} else if this.available > 0 {
		cache := &LinkedNode{key: key, value: value}
		// 放在队头
		cache.pre = this.head
		cache.next = this.head.next
		cache.pre.next = cache
		cache.next.pre = cache
		this.available--
		this.hashMap[key] = cache
	} else {
		// 删除尾节点
		last := this.tail.pre
		this.tail.pre = last.pre
		last.pre.next = this.tail
		last.next, last.pre = nil, nil
		delete(this.hashMap, last.key)

		// 插入新节点
		cache := &LinkedNode{key: key, value: value}
		cache.pre = this.head
		cache.next = this.head.next
		cache.pre.next = cache
		cache.next.pre = cache
		this.hashMap[key] = cache
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
