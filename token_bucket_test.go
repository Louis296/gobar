package gobar

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestOneGoroutine(t *testing.T) {
	bucket := CreateTokenBucket(10, 10000, true)
	for {
		bucket.GetToken()
		fmt.Println("get token at time ", time.Now())
	}
}

func TestMultiGoroutine(t *testing.T) {
	bucket := CreateTokenBucket(10, 1000, true)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		j := i
		go func() {
			for {
				bucket.GetToken()
				fmt.Println("go routine [ ", j, " ]get token at time ", time.Now())
			}
		}()
	}
	wg.Wait()
}
