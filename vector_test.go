package vector

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestVectorOne(t *testing.T) {

	for i := 0; i < 100; i++ {
		v := New(100)

		size := rand.Intn(10000) + 1000
		for i := 0; i < size; i++ {
			v.Push(i)
		}

		if v.Len() != size {
			t.Fatalf("len err. len:%d, size:%d", v.Len(), size)
		}

	}

	t.Log("ok")
}

func TestVectorSlmple(t *testing.T) {

	Remake(1000)

	for i := 0; i < 100; i++ {
		Push(rand.Int())
	}

	if Len() != 100 {
		t.Fatal("len != 100. len:", Len())
	}

	if Cap() != 1000 {
		t.Fatal("cap != 1000", Cap())
	}

	t.Log(IsEmpty(), Len(), Cap())

	for !IsEmpty() {
		PopBack()
	}

	if PopBack() != nil {
		t.Fatal("PopBack err")
	}

	if PopFront() != nil {
		t.Fatal("PopFront err")
	}

	t.Log("ok")
}

func TestGoVectorSlmple(t *testing.T) {

	vector := New(1000)

	w := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			for k := 0; k < 1000; k++ {
				vector.Push(k)
			}
		}()
	}
	w.Wait()

	if vector.Len() != 100000 {
		t.Fatal("len != 100000. len:", vector.Len())
	}

	t.Log(vector.IsEmpty(), vector.Len(), vector.Cap())

	w = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			for !vector.IsEmpty() {
				vector.PopFront()
			}
		}()
	}
	w.Wait()

	t.Log("ok")
}

func TestGoVectorPurge(t *testing.T) {

	Remake(1000)

	// auto clean
	UsePurge(1000, 10*time.Millisecond)
	UsePurge(1000, 100*time.Millisecond)
	UsePurge(1000, 1000*time.Millisecond)

	for i := 0; i < 10000; i++ {
		Push(i)
	}

	if Len() != 10000 {
		t.Fatal("len error. len:", Len())
	}

	t.Log(Len(), Cap())

	for !IsEmpty() {
		PopFront()
	}
	t.Log(Len(), Cap())
	time.Sleep(30 * time.Millisecond)
	t.Log(Len(), Cap())
	if Cap() != 1000 {
		t.Fatal("Purge error. cap:", Cap())
	}

	t.Log("ok")
}

func TestGoVectorRemake(t *testing.T) {

	v := New(100)
	v.Remake(2000)
	if v.Cap() != 2000 {
		t.Fatal("make error. cap:", v.Cap())
	}

	t.Log("ok")
}

func TestVectorItems(t *testing.T) {

	Push(1)
	Push(2)
	Push(3)
	for i, v := range Items() {
		t.Log(i, v)
	}

	var val Val
	val = PopBack()
	t.Log(val.(int))

	t.Log("ok")
}
