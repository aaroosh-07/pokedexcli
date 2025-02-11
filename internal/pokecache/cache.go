package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	v map[string]cacheEntry
	mu sync.Mutex
}

func NewCache() (*Cache) {
	newCache := Cache{v: make(map[string]cacheEntry)}

	ticker := time.NewTicker(10 * time.Second)

	//fires a new go routine to clear map of old timestamps
	go func(){
		for range ticker.C {
			newCache.readLoop()
		}
	}()

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, isPresent := c.v[key]

	if isPresent {
		//check if need to update map or not
		return 
	}

	c.v[key] = cacheEntry{
					val: val,
					createdAt: time.Now(),
				}
	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ele, isPresent := c.v[key]

	if !isPresent {
		return []byte{}, false
	}

	return ele.val, true
}

func (c *Cache) readLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	curTime := time.Now()

	for key, ele := range c.v {
		if curTime.Sub(ele.createdAt) > 5 * time.Second {
			delete(c.v, key)
		}
	}
}