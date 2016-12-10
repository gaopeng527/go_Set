// hash_set
package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

// 创建和初始化HashSet的方法
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// 向HashSet中添加元素的方法
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

// 删除HashSet中指定的元素
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

// 清除HashSet中的所有元素
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

// 判断HashSet是否包含指定元素
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// 获取HashSet中元素值数量
func (set *HashSet) Len() int {
	return len(set.m)
}

// 判断两个Set类型值是否相同
func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// 生成HashSet的一个快照
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// 获取HashSet自身字符串表示形式
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}
