package set

type Set interface {
	// 添加元素
	Add(items ...interface{})
	// 包含
	Contain(item interface{}) bool
	// 删除
	Remove(item interface{})
	// 集合大小
	Size() int
	// 清空
	Clear()
	// 判断是否为空
	Empty() bool
	// 创建副本
	Duplicate() Set
	// 数组
	ToSlice() []interface{}
}

func ThreadSafe(items ...interface{}) Set {
	s := newThreadSafeSet()
	s.Add(items...)
	return s
}

func ThreadUnsafe(items ...interface{}) Set {
	s := newThreadUnsafeSet()
	s.Add(items...)
	return &s
}
