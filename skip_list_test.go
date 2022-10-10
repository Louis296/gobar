package gobar

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList()
	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(88)
	fmt.Println(skipList.Search(2))
	fmt.Println(skipList.Delete(2))
	fmt.Println(skipList.Search(2))
}
