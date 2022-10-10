package gobar

import "math/rand"

type skipListNode struct {
	val  int
	next []*skipListNode
}

type SkipList struct {
	head  *skipListNode
	level int
}

const maxLevel int = 32
const pFactor float64 = 0.25

func NewSkipList() SkipList {
	return SkipList{
		head: &skipListNode{
			val:  -1,
			next: make([]*skipListNode, maxLevel),
		},
		level: 0,
	}
}

func (s *SkipList) Add(num int) {
	update := make([]*skipListNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	lv := randomLevel()
	if lv > s.level {
		s.level = lv
	}
	newNode := &skipListNode{
		val:  num,
		next: make([]*skipListNode, lv),
	}
	for i := range update[:lv] {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (s *SkipList) Search(num int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
	}
	return curr.next[0] != nil && curr.next[0].val == num
}

func (s *SkipList) Delete(num int) bool {
	update := make([]*skipListNode, s.level)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	curr = curr.next[0]
	if curr == nil || curr.val != num {
		return false
	}
	for i := range update {
		if update[i].next[i] == curr {
			update[i].next[i] = curr.next[i]
		} else {
			break
		}
	}
	for i := s.level - 1; i >= 0; i-- {
		if s.head.next[i] == nil {
			s.level--
		} else {
			break
		}
	}
	return true
}

func randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}
