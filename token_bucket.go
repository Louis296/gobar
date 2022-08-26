package gobar

import (
	"sync"
	"time"
)

type TokenBucket struct {
	size  int
	token int
	mutex *sync.Mutex
	cond  *sync.Cond
}

// CreateTokenBucket will return a TokenBucket with generate token
// every 'interval' millisecond and size of 'size'
func CreateTokenBucket(size, interval int, full bool) *TokenBucket {
	bucket := &TokenBucket{
		size:  size,
		token: 0,
		mutex: new(sync.Mutex),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	if full {
		bucket.token = size
	}
	go func() {
		for {
			time.Sleep(time.Duration(interval) * time.Millisecond)
			if bucket.token < bucket.size {
				bucket.token++
			}
			bucket.cond.Signal()
		}
	}()
	return bucket
}

func (b *TokenBucket) GetToken() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.token == 0 {
		b.cond.L.Lock()
		b.cond.Wait()
		b.cond.L.Unlock()
	}
	b.token--
}
