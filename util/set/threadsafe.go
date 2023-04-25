package set

import "sync"

type threadSafeSet struct {
	s threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() *threadSafeSet {
	return &threadSafeSet{s: newThreadUnsafeSet()}
}

func (set *threadSafeSet) Add(items ...interface{}) {
	set.Lock()
	set.s.Add(items...)
	set.Unlock()
}

func (set *threadSafeSet) Contain(item interface{}) bool {
	set.RLock()
	defer set.RUnlock()
	return set.s.Contain(item)
}

func (set *threadSafeSet) Remove(item interface{}) {
	set.Lock()
	set.s.Remove(item)
	set.Unlock()
}

func (set *threadSafeSet) Size() int {
	set.RLock()
	defer set.RUnlock()
	return set.s.Size()
}

func (set *threadSafeSet) Clear() {
	set.Lock()
	set.s.Clear()
	set.Unlock()
}

func (set *threadSafeSet) Empty() bool {
	return set.Size() == 0
}

func (set *threadSafeSet) Duplicate() Set {
	set.RLock()
	defer set.RUnlock()
	s := set.s.Duplicate()
	return &threadSafeSet{s: *(s.(*threadUnsafeSet))}
}

func (set *threadSafeSet) ToSlice() []interface{} {
	set.RLock()
	defer set.RUnlock()
	return set.s.ToSlice()
}
