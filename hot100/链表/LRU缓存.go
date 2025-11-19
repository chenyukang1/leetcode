package 链表

type LRUCache struct {
	table map[int]int
	keys  []int
	len   int
}

func Constructor(capacity int) LRUCache {
	keys := make([]int, capacity)
	for i := range keys {
		keys[i] = -1
	}
	return LRUCache{table: make(map[int]int), keys: keys, len: capacity}
}

func (c *LRUCache) Get(key int) int {
	if val, ok := c.table[key]; ok {
		c.sort(key, true)
		return val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	_, ok := c.table[key]
	c.sort(key, ok)
	c.table[key] = value
}

func (c *LRUCache) sort(first int, exists bool) {
	var end = c.len - 1
	if exists {
		for i := range c.keys {
			if c.keys[i] == first {
				end = i
				break
			}
		}
	}
	if end == c.len-1 && c.keys[c.len-1] != -1 {
		delete(c.table, c.keys[c.len-1])
	}
	for i := end; i >= 1; i-- {
		c.keys[i] = c.keys[i-1]
	}
	c.keys[0] = first
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
