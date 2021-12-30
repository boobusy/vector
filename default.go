package vector

import (
	"time"
)

var (
	defautlVector = New(1 << 10)
)

// Len returns the number of items.
func Len() int {
	return defautlVector.Len()
}

// Cap returns the capacity of items.
func Cap() int {
	return defautlVector.Cap()
}

func IsEmpty() bool {
	return defautlVector.Len() == 0
}

// Items returns the items.
func Items() []any {
	return defautlVector.Items()
}

// Push inserts a new item e with value v at the back of items.
func Push(item any) {
	defautlVector.Push(item)
}

// Remake items size
func Remake(size int) {
	defautlVector.Remake(size)
}

// PopFront returns the first val of items.
func PopFront() any {
	return defautlVector.PopFront()
}

// PopBack returns the last val of items.
func PopBack() any {
	return defautlVector.PopBack()
}

func UsePurge(maxSize int, interval time.Duration) *Type {
	return defautlVector.UsePurge(maxSize, interval)
}
