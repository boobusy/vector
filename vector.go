package vector

import (
	"runtime"
	"sync"
	"time"
)

type (
	vector struct {
		items []any
		size  int
		r     chan struct{}
		w     chan func()
		purge *sync.Once
	}

	any = interface{}

	Type = vector
	Val  = any
)

func New(size int) *Type {

	v := &vector{
		items: make([]any, 0, size),
		size:  size,
		r:     make(chan struct{}, 0),
		w:     make(chan func()),
		purge: new(sync.Once),
	}

	go v.syncSliceData()
	return v
}

// Len returns the number of items.
func (v *vector) Len() int {
	var l int
	v.w <- func() {
		defer v.rChan()

		l = len(v.items)
	}
	<-v.r
	return l
}

// Cap returns the capacity of items.
func (v *vector) Cap() int {
	var l int
	v.w <- func() {
		defer v.rChan()
		l = cap(v.items)
	}
	<-v.r
	return l
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
func (v *vector) Remake(size int) {
	v.w <- func() {
		v.items = make([]any, 0, size)
		v.size = size
	}
}

// Resize vector cap
func (v *vector) Resize(size int) {
	v.w <- func() {
		vcap := cap(v.items)
		if vcap > size {
			v.items = v.items[:size]
		} else if vcap < size {
			v.items = append(make([]any, 0, size), v.items...)
		}
		v.size = size
	}
}

// clear items
func (v *vector) Clear() {
	v.w <- func() {
		v.items = v.items[:0]
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
		if l < 0 {
			return
		}
		w = v.items[l]
		v.items[l] = nil
		v.items = v.items[:l]
	}
	<-v.r
	return
}

// open auto purge
func (v *vector) UsePurge(maxSize int, interval time.Duration) *Type {
	v.purge.Do(func() {
		go func() {
			ch := time.NewTicker(interval)
			for {
				<-ch.C
				if v.Len() == 0 && v.Cap() > maxSize {
					v.Remake(v.size)
				}
			}
		}()
	})
	return v
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
