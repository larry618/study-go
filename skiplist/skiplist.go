package skiplist

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

//const maxLevel = 32

type skipListNode struct {
	key     int
	value   string
	forward []*skipListNode // 各层节点的下一个
}

func (node *skipListNode) String() string {
	return fmt.Sprintf("key: %2d, value: %s", node.key, node.value)
}

type SkipList struct {
	head  *skipListNode
	total int // 有效节点数
	level int // 当前层数  level 越大, 节点数量越小 0 level 保存了所有的节点
}

func NewSkipList() *SkipList {
	return &SkipList{
		level: 0,
		head: &skipListNode{
			key:     int(math.MinInt64),
			forward: make([]*skipListNode, 1)}}

	//return &SkipList{}
}

func (sk *SkipList) Get(searchKey int) string {
	curt, _ := sk.seek(searchKey)

	curt = curt.forward[0]
	if curt != nil && curt.key == searchKey {
		return curt.value
	}

	return ""
}

func (sk *SkipList) Add(key int, value string) {

	level := sk.randomLevel()
	oldLevel := sk.level
	if level > oldLevel {
		sk.head.forward = append(sk.head.forward, make([]*skipListNode, level-oldLevel+1)...)
		sk.level = level

	}

	curt, update := sk.seek(key) // seek 不会返回 nil 有一个 head 节点

	curt = curt.forward[0]
	if curt != nil && curt.key == key {
		curt.value = value
		sk.level = oldLevel
		return
	}

	newNode := &skipListNode{key, value, make([]*skipListNode, level+1)}

	for i := level; i >= 0; i-- {
		prev := update[i]
		newNode.forward[i] = prev.forward[i]
		prev.forward[i] = newNode
	}

}

func (sk *SkipList) Delete(key int) string {

	curt, update := sk.seek(key)

	curt = curt.forward[0]
	if curt.key != key { // 不存在
		return ""
	}

	for i := len(curt.forward) - 1; i >= 0; i-- {
		prev := update[i]
		prev.forward[i] = curt.forward[i] // 删除 curt
	}

	sk.removeEmptyLevel()

	return curt.value
}

func (sk *SkipList) removeEmptyLevel() {
	for sk.level >= 0 && sk.head.forward[sk.level] == nil {
		sk.level--
	}
}

// 在某一层找到一个非小于searchKey的结点后，跳到下一层继续找，直到最底层为止。
// 找到小于等于 target key 的最后一个节点
func (sk *SkipList) seek(searchKey int) (*skipListNode, []*skipListNode) {

	update := make([]*skipListNode, sk.level+1)
	curt := sk.head
	for i := sk.level; i >= 0; i-- {
		update[i] = sk.head

		for curt.forward[i] != nil && curt.forward[i].key < searchKey {
			curt = curt.forward[i] // 同一层的下一个
		}
		// 走完一层了
		update[i] = curt
	}
	return curt, update
}

func (sk *SkipList) randomLevel() int {
	return rand.Intn(sk.level + 2) // [0, maxLevel)
}

func (sk *SkipList) String() string {
	if sk.head == nil {
		return ""
	}
	buf := make([]*bytes.Buffer, sk.level+1)
	for i := 0; i < len(buf); i++ {
		buf[i] = new(bytes.Buffer)
		buf[i].WriteString(fmt.Sprintf("level %2d: ", i))
	}
	node := sk.head.forward[0] // head 的下一个节点
	for node != nil {
		for i := 0; i <= sk.level; i++ {
			if i < len(node.forward) {
				//buf[i].WriteString(strconv.Itoa(node.key))
				//buf[i].WriteString(" ")
				buf[i].WriteString(fmt.Sprintf(" %2d ", node.key))
				//buf[i].WriteString(node.value)
			} else {
				buf[i].WriteString(" -- ")
			}
		}
		if len(node.forward) > 0 {
			node = node.forward[0]
		} else {
			node = nil
		}
	}
	re := new(bytes.Buffer)
	for i := len(buf) - 1; i >= 0; i-- {
		re.WriteString(buf[i].String())
		re.WriteString("\n")
	}
	return re.String()
}
