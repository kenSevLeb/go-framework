package set

import (
	"fmt"
	"strings"
)

// threadUnsafeSet 非线程安全的集合
type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (set *threadUnsafeSet) Add(items ...interface{}) {
	for _, item := range items {
		(*set)[item] = struct{}{}
	}
}

func (set *threadUnsafeSet) Contain(item interface{}) bool {
	_, ok := (*set)[item]
	return ok
}

func (set *threadUnsafeSet) Remove(item interface{}) {
	delete((*set), item)
}

func (set *threadUnsafeSet) Size() int {
	return len((*set))
}

func (set *threadUnsafeSet) Clear() {
	*set = newThreadUnsafeSet()
}

func (set *threadUnsafeSet) Empty() bool {
	return set.Size() == 0
}

func (set *threadUnsafeSet) Duplicate() Set {
	duplicateSet := newThreadUnsafeSet()
	for item, _ := range *set {
		duplicateSet.Add(item)
	}
	return &duplicateSet
}

func (set *threadUnsafeSet) String() string {
	items := make([]string, 0, len(*set))

	for elem := range *set {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}

func (set *threadUnsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	for elem := range *set {
		keys = append(keys, elem)
	}

	return keys
}
