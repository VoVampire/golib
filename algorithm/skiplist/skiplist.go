package skiplist

import (
	"math/rand"
	"time"
)

const (
	SKIPLIST_MAXLEVEL = 64
	SKIPLIST_P        = 0.25
)

func init() { rand.Seed(time.Now().Unix()) }

func zslRandomLevel() int {
	level := 1
	for rand.Intn(100) < SKIPLIST_P*100 && level < SKIPLIST_MAXLEVEL {
		level++
	}
	return level
}

// SkipList
type zSkipList struct {
	Header, Tail *zSkipListNode
	Length       uint64
	Level        int
}
type zSkipListNode struct {
	Backward *zSkipListNode
	Level    []*zSkipListLevel
	Score    float64
	Ele      string
}
type zSkipListLevel struct {
	Forward *zSkipListNode
	Span    uint64
}

func ZslCreate() *zSkipList {
	sln := zSkipListNode{}
	sln.Level = make([]*zSkipListLevel, SKIPLIST_MAXLEVEL)
	for i := 0; i < SKIPLIST_MAXLEVEL; i++ {
		sln.Level[i] = &zSkipListLevel{}
	}

	sl := zSkipList{}
	sl.Header = &sln
	sl.Level = 1
	return &sl
}

func (zsl *zSkipList) ZslInsert(score float64, ele string) *zSkipListNode {
	var update [SKIPLIST_MAXLEVEL]*zSkipListNode // 插入节点时，需要更新被插入节点每层的前一个节点
	var rank [SKIPLIST_MAXLEVEL]uint64           // 记录当前层从header节点到update[i]节点所经历的步长

	// 查找要插入的位置
	x := zsl.Header
	for i := zsl.Level - 1; i >= 0; i-- {
		if i != zsl.Level-1 {
			rank[i] = rank[i+1]
		}
		for x.Level[i].Forward != nil &&
			(x.Level[i].Forward.Score < score ||
				(x.Level[i].Forward.Score == score &&
					x.Level[i].Forward.Ele < ele)) {
			rank[i] += x.Level[i].Span
			x = x.Level[i].Forward
		}
		update[i] = x
	}

	// 调整跳跃表高度
	level := zslRandomLevel()
	if level > zsl.Level {
		for i := zsl.Level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.Header
			update[i].Level[i].Span = zsl.Length
		}
		zsl.Level = level
	}

	// 插入目标节点
	x = &zSkipListNode{
		Score: score,
		Ele:   ele,
		Level: make([]*zSkipListLevel, level),
	}
	for i := 0; i < level; i++ {
		x.Level[i] = &zSkipListLevel{}
		x.Level[i].Forward = update[i].Level[i].Forward
		update[i].Level[i].Forward = x
		x.Level[i].Span = update[i].Level[i].Span - (rank[0] - rank[i])
		update[i].Level[i].Span = (rank[0] - rank[i]) + 1
	}
	for i := level; i < zsl.Level; i++ {
		update[i].Level[i].Span++
	}

	// 调整backward
	if update[0] != zsl.Header {
		x.Backward = update[0]
	}
	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = x
	} else {
		zsl.Tail = x
	}

	zsl.Length++
	return x
}

func (zsl *zSkipList) ZslQuery(score float64, ele string) *zSkipListNode {
	x := zsl.Header
	for i := zsl.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil &&
			(x.Level[i].Forward.Score < score ||
				(x.Level[i].Forward.Score == score &&
					x.Level[i].Forward.Ele < ele)) {
			x = x.Level[i].Forward
		}
	}
	if x = x.Level[0].Forward; x != nil && x.Score == score && x.Ele == ele {
		return x
	}
	return nil
}

func (zsl *zSkipList) ZslDelete(score float64, ele string) int {
	var update [SKIPLIST_MAXLEVEL]*zSkipListNode // 删除节点时，需要更新被删除节点每层的前一个节点

	x := zsl.Header
	for i := zsl.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil &&
			(x.Level[i].Forward.Score < score ||
				(x.Level[i].Forward.Score == score &&
					x.Level[i].Forward.Ele < ele)) {
			x = x.Level[i].Forward
		}
		update[i] = x
	}

	if x = x.Level[0].Forward; x != nil && x.Score == score && x.Ele == ele {
		for i := 0; i < zsl.Level; i++ {
			if update[i].Level[i].Forward == x {
				update[i].Level[i].Span += x.Level[i].Span - 1
				update[i].Level[i].Forward = x.Level[i].Forward
			} else {
				update[i].Level[i].Span -= 1
			}
		}
		if x.Level[0].Forward != nil{
			x.Level[0].Forward.Backward = x.Backward
		} else {
			zsl.Tail = x.Backward
		}
		for zsl.Level > 1 && zsl.Header.Level[zsl.Level-1].Forward == nil{
			zsl.Level--
		}

		zsl.Length--
		return 1
	}
	return 0 /* not found */
}
