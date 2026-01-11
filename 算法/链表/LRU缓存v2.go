package 链表

// https://leetcode.cn/problems/lru-cache/solutions/259678/lruhuan-cun-ji-zhi-by-leetcode-solution/?envType=study-plan-v2&envId=top-100-liked

// LRUCacheV2 用数组实现的版本，put/get都是O(n)时间复杂度，v2版本使用双向链表，put/get时间复杂度优化到O(1)
type LRUCacheV2 struct {
	table     map[int]*DoublyNode
	dummyHead *DoublyNode
	dummyTail *DoublyNode
	count     int
	limit     int
}

type DoublyNode struct {
	prev *DoublyNode
	next *DoublyNode
	key  int
	val  int
}

func ConstructorV2(capacity int) LRUCacheV2 {
	dummyHead := &DoublyNode{}
	dummyTail := &DoublyNode{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	return LRUCacheV2{table: make(map[int]*DoublyNode), dummyHead: dummyHead, dummyTail: dummyTail, count: 0, limit: capacity}
}

func (c *LRUCacheV2) Get(key int) int {
	if node, ok := c.table[key]; ok {
		c.pushToHead(node, true)
		return node.val
	}
	return -1
}

func (c *LRUCacheV2) Put(key int, value int) {
	if node, ok := c.table[key]; ok {
		node.val = value
		c.pushToHead(node, true)
	} else {
		node = &DoublyNode{key: key, val: value}
		c.pushToHead(node, false)
		c.table[key] = node
		c.count++
		if c.count > c.limit {
			tail := c.dummyTail.prev
			prev := tail.prev
			prev.next = c.dummyTail
			c.dummyTail.prev = prev
			delete(c.table, tail.key)
		}
	}
}

func (c *LRUCacheV2) pushToHead(node *DoublyNode, exists bool) {
	if exists {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.prev = c.dummyHead
	node.next = c.dummyHead.next
	c.dummyHead.next.prev = node
	c.dummyHead.next = node
}
