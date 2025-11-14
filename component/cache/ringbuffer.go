// 本地环形队列
//
package cache

import "errors"

type T interface{}

var ErrIsEmpty = errors.New("error is empty")

// RingQueue 本地环形队列
type RingQueue struct {
	buf []T // 元素队列

	initializeSize int // 初始容量
	size           int // 当前容量
	r              int // 读指针计数
	w              int // 写指针计数
}

// NewRingQueue 实例化一个
func NewRingQueue(initializeSize int) *RingQueue {
	return &RingQueue{
		buf:            make([]T, initializeSize),
		initializeSize: initializeSize,
		size:           initializeSize,
		r:              0,
		w:              0,
	}
}

// IsEmpty 判断是否为空
func (rq *RingQueue) IsEmpty() bool {
	return rq.r == rq.w
}

// Read 从队列中读取数据
func (rq *RingQueue) Read() (T, error) {
	if rq.IsEmpty() {
		return nil, ErrIsEmpty
	}

	// 读取指针处获取，并向后移动一位
	v := rq.buf[rq.r]
	rq.r++

	// 绕一圈后重置为0
	if rq.r == rq.size {
		rq.r = 0
	}
	return v, nil
}

// Write 写入队列
func (rq *RingQueue) Write(v T) {
	// 通过写指针存储数据，并向后移动一位
	rq.buf[rq.w] = v
	rq.w++

	// 绕一圈后，重置为0
	if rq.w == rq.size {
		rq.w = 0
	}

	// 写入后，写指针遇到读指针，即队列已写满
	if rq.w == rq.r {
		rq.grow()
	}

}

// Len 获取队列长度
func (rq *RingQueue) Len() int {
	return rq.w - rq.r
}

// grow 扩容
func (rq *RingQueue) grow() {
	var size int
	// 参考golang的slice扩容反感
	if rq.size < 1024 {
		size = rq.size * 2 // 小于1024的按2倍扩容
	} else {
		size = rq.size + rq.size/4 // 超过1024的时候就以1.25倍扩容
	}
	buf := make([]T, size)
	// 拷贝读指针 前 待读取内容
	copy(buf[0:], rq.buf[rq.r:])

	// 拷贝读指针 前 待读取内容
	copy(buf[rq.size-rq.r:], rq.buf[0:rq.r])

	// 拷贝后读指针为头部
	rq.r = 0
	// 写指针为尾部（原队列大小）
	rq.w = rq.size
	rq.size = size
	rq.buf = buf
}
