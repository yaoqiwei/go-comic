package cache

import (
	"sync"
	"time"
)

type LockVal struct {
	lock  *sync.RWMutex
	value interface{}
}

var locks = make(map[string]*LockVal)
var cacheLock sync.RWMutex

func getLock(key string) *LockVal {
	cacheLock.RLock()
	loc, ok := locks[key]
	cacheLock.RUnlock()
	if !ok {
		cacheLock.Lock()
		loc = &LockVal{}
		loc.lock = &sync.RWMutex{}
		loc.value = nil
		locks[key] = loc
		cacheLock.Unlock()
	}
	return loc
}

func GetCache(key string) interface{} {
	lock := getLock(key)
	lock.lock.RLock()
	defer lock.lock.RUnlock()
	return lock.value
}

func SetCache(key string, c interface{}, second ...time.Duration) {
	lock := getLock(key)
	lock.lock.Lock()
	defer lock.lock.Unlock()
	lock.value = c

	go func() {
		if len(second) > 0 {
			time.Sleep(second[0])
		} else {
			time.Sleep(5 * time.Second)
		}
		lock.lock.Lock()
		defer lock.lock.Unlock()
		lock.value = nil
	}()
}
