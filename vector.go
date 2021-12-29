package vector

import (
	"runtime"
)

type (
	vector struct {
		items []any
		size  int
		r     chan struct{}
		w     chan func()
	}

	any = interface{}
)


func New(size int) *vector {
	v := &vector{
		items: make([]any, 0, size),
		size:  size,
		r:     make(chan struct{}, 0),
		w:     make(chan func()),
	}

	go v.syncSliceData()
	return v
}

// Len returns the number of items.
func (v *vector) Len() int {
	return len(v.items)
}

// Cap returns the capacity of items.
func (v *vector) Cap() int {
	return cap(v.items)
}

func (v *vector) IsEmpty() bool {
	return v.Len() == 0
}

// Items returns the items.
func (v *vector) Items() []any {
	return v.items
}

// Push inserts a new item e with value v at the back of items.
func (v *vector) Push(item any) {
	v.w <- func() {
		v.items = append(v.items, item)
	}
}

// Remake items size
func (v *vector) Remake(items []any) {
	v.w <- func() {
		v.items = items
	}
}

// PopFront returns the first val of items.
func (v *vector) PopFront() (w any) {
	v.w <- func() {
		defer v.rChan()

		if len(v.items) == 0 {
			return
		}
		w = v.items[0]
		v.items[0] = nil
		v.items = v.items[1:]
	}
	<-v.r
	return
}

// PopBack returns the last val of items.
func (v *vector) PopBack() (w any) {
	v.w <- func() {
		defer v.rChan()

		l := len(v.items) - 1
		if l < 1 {
			return
		}
		w = v.items[l]
		v.items[l] = nil
		v.items = v.items[:l]
	}
	<-v.r
	return
}

func (v *vector) syncSliceData() {
	for {
		(<-v.w)()
		runtime.Gosched()
	}
}

func (v *vector) rChan() {
	v.r <- struct{}{}
}
