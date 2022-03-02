package datastructer

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

//skip list 跳表实现增删查
// redis 源码 https://developpaper.com/redis-source-code-analysis-1-skiplist/

const UP_LEVELS_ABILITY = 500
const UP_LEVELS_TOTAL = 1000

type skipListNode struct {
	score int64
	val   interface{}
	next  *skipListNode
	pre   *skipListNode
	up    *skipListNode
	down  *skipListNode
}

type skipList struct {
	head   *skipListNode
	tail   *skipListNode
	size   int
	levels int
}

func NewSkipList() *skipList {
	sl := new(skipList)
	sl.head = new(skipListNode)
	sl.tail = new(skipListNode)
	sl.head.score = math.MinInt64
	sl.tail.score = math.MaxInt64

	sl.head.next = sl.tail
	sl.tail.pre = sl.head

	sl.size = 0
	sl.levels = 1

	return sl
}

func (sl *skipList) Size() int {
	return sl.size
}

func (sl *skipList) Levels() int {
	return sl.levels
}

func (sl *skipList) Get(score int64) interface{} {
	node := sl.findNode(score)
	if node.score == score {
		return node.val
	} else {
		return nil
	}
}

func (sl *skipList) Insert(score int64, val interface{}) {
	f := sl.findNode(score)
	if f.score == score {
		f.val = val
		return
	}
	curNode := new(skipListNode)
	curNode.score = score
	curNode.val = val

	sl.insertAfter(f, curNode)

	rander := rand.New(rand.NewSource(time.Now().UnixNano()))

	curlevels := 1
	for rander.Intn(UP_LEVELS_TOTAL) < UP_LEVELS_ABILITY {
		curlevels++
		if curlevels > sl.levels {
			sl.newlevels()
		}

		for f.up == nil {
			f = f.pre
		}
		f = f.up
		tmpNode := &skipListNode{score: score}

		curNode.up = tmpNode
		tmpNode.down = curNode
		sl.insertAfter(f, tmpNode)

		curNode = tmpNode
	}

	sl.size++
}

func (sl *skipList) Remove(score int64) interface{} {
	f := sl.findNode(score)
	if f.score != score {
		return nil
	}
	v := f.val

	for f != nil {
		f.pre.next = f.next
		f.next.pre = f.pre
		f = f.up
	}
	return v
}

func (sl *skipList) newlevels() {
	nhead := &skipListNode{score: math.MinInt64}
	ntail := &skipListNode{score: math.MaxInt64}
	nhead.next = ntail
	ntail.pre = nhead

	sl.head.up = nhead
	nhead.down = sl.head
	sl.tail.up = ntail
	ntail.down = sl.tail

	sl.head = nhead
	sl.tail = ntail
	sl.levels++
}

func (sl *skipList) insertAfter(pNode *skipListNode, curNode *skipListNode) {
	curNode.next = pNode.next
	curNode.pre = pNode
	pNode.next.pre = curNode
	pNode.next = curNode
}

func (sl *skipList) findNode(score int64) *skipListNode {
	p := sl.head

	for p != nil {
		if p.score == score {
			if p.down == nil {
				return p
			}
			p = p.down
		} else if p.score < score {
			if p.next.score > score {
				if p.down == nil {
					return p
				}
				p = p.down
			} else {
				p = p.next
			}
		}
	}
	return p
}

func (sl *skipList) Print() {

	mapScore := make(map[int64]int)

	p := sl.head
	for p.down != nil {
		p = p.down
	}
	index := 0
	for p != nil {
		mapScore[p.score] = index
		p = p.next
		index++
	}
	p = sl.head
	for i := 0; i < sl.levels; i++ {
		q := p
		preIndex := 0
		for q != nil {
			s := q.score
			if s == math.MinInt64 {
				fmt.Printf("%s", "BEGIN")
				q = q.next
				continue
			}
			index := mapScore[s]
			c := (index - preIndex - 1) * 6
			for m := 0; m < c; m++ {
				fmt.Print("-")
			}
			if s == math.MaxInt64 {
				fmt.Printf("-->%s\n", "END")
			} else {
				fmt.Printf("-->%3d", s)
				preIndex = index
			}
			q = q.next
		}
		p = p.down
	}
}

func TestSkipList(t *testing.T) {
	sk := NewSkipList()

	sk.Insert(100, "lala")
	sk.Insert(11, "sx")
	sk.Insert(22, "11")
	sk.Insert(3, "dd")
	sk.Insert(80, "bb")
	sk.Insert(77, "bb")
	sk.Insert(6, "bb")
	sk.Insert(88, "bb")
	sk.Insert(33, "bb")
	sk.Insert(44, "bb")
	sk.Print()
}


//——————————————————————————————————————————————————————
// 其他实现方式：https://github.com/ryszard/goskiplist
//——————————————————————————————————————————————————————

type node struct {
	forward    []*node
	backward   *node
	key, value interface{}
}

type SkipList struct {
	lessThan func(l, r interface{}) bool
	header   *node
	footer   *node
	length   int
	// MaxLevel determines how many items the SkipList can store
	// efficiently (2^MaxLevel).
	//
	// It is safe to increase MaxLevel to accomodate more
	// elements. If you decrease MaxLevel and the skip list
	// already contains nodes on higer levels, the effective
	// MaxLevel will be the greater of the new MaxLevel and the
	// level of the highest node.
	//
	// A SkipList with MaxLevel equal to 0 is equivalent to a
	// standard linked list and will not have any of the nice
	// properties of skip lists (probably not what you want).
	MaxLevel int
}

type Iterator interface {
	// Next returns true if the iterator contains subsequent elements
	// and advances its state to the next element if that is possible.
	Next() (ok bool)
	// Previous returns true if the iterator contains previous elements
	// and rewinds its state to the previous element if that is possible.
	Previous() (ok bool)
	// Key returns the current key.
	Key() interface{}
	// Value returns the current value.
	Value() interface{}
	// Seek reduces iterative seek costs for searching forward into the Skip List
	// by remarking the range of keys over which it has scanned before.  If the
	// requested key occurs prior to the point, the Skip List will start searching
	// as a safeguard.  It returns true if the key is within the known range of
	// the list.
	Seek(key interface{}) (ok bool)
	// Close this iterator to reap resources associated with it.  While not
	// strictly required, it will provide extra hints for the garbage collector.
	Close()
}