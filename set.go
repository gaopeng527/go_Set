// set project set.go
package set

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

// 将集合other添加到集合one中
func AddSet(one Set, other Set) {
	if one == nil || other == nil || other.Len() == 0 {
		return
	}
	for _, v := range other.Elements() {
		one.Add(v)
	}
}

// 判断集合 one 是否是集合 other 的超集
func IsSuperset(one Set, other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen <= otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}

// 生成集合 one 和集合 other 的并集
func Union(one Set, other Set) Set {
	if one == nil && other == nil {
		return nil
	}
	unionedSet := NewSimpleSet()
	AddSet(unionedSet, one)
	AddSet(unionedSet, other)
	return unionedSet
}

// 生成集合 one 和集合 other 的交集
func Intersect(one Set, other Set) Set {
	if one == nil || other == nil {
		return nil
	}
	intersectedSet := NewSimpleSet()
	if one.Len() == 0 || other.Len() == 0 {
		return intersectedSet
	}
	if one.Len() < other.Len() {
		for _, v := range one.Elements() {
			if other.Contains(v) {
				intersectedSet.Add(v)
			}
		}
	} else {
		for _, v := range other.Elements() {
			if one.Contains(v) {
				intersectedSet.Add(v)
			}
		}
	}
	return intersectedSet
}

// 生成集合 one 对集合 other 的差集
func Difference(one Set, other Set) Set {
	if one == nil {
		return nil
	}
	differencedSet := NewSimpleSet()
	if other == nil || other.Len() == 0 {
		AddSet(differencedSet, one)
		return differencedSet
	}
	for _, v := range one.Elements() {
		if !other.Contains(v) {
			differencedSet.Add(v)
		}
	}
	return differencedSet
}

// 生成集合 one 和集合 other 的对称差集
func SymmetricDifference(one Set, other Set) Set {
	diffA := Difference(one, other)
	if other == nil || other.Len() == 0 {
		return diffA
	}
	diffB := Difference(other, one)
	return Union(diffA, diffB)
}

// 返回一个HashSet
func NewSimpleSet() Set {
	return NewHashSet()
}

// 判断给定value是否为集合
func IsSet(value interface{}) bool {
	if _, ok := value.(Set); ok {
		return true
	}
	return false
}
