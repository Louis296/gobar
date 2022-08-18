package gobar

import (
	"fmt"
	"testing"
)

func TestPrefixTree(t *testing.T) {
	tree := PrefixTree{}
	tree.Insert("hello", 1)
	tree.Insert("22", 2)
	tree.Insert("hello world", 3)
	tree.Insert("234", 4)
	if i, ok := tree.Get("hello"); ok {
		if i != 1 {
			t.Fatalf("tree can't do Get")
		} else {
			fmt.Println(i)
		}
	} else {
		t.Fatalf("tree can't do Get")
	}
	tree.Delete("hello world")
	if _, ok := tree.Get("hello world"); ok {
		t.Fatalf("tree can't do Delete")
	}
}
