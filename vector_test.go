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

	v := New(1000)

	for i := 0; i < 100; i++ {
		v.Push(rand.Int())
	}

	if v.Len() != 100 {
		t.Fatal("len != 100. len:", v.Len())
	}

	if v.Cap() != 1000 {
		t.Fatal("cap != 1000")
	}

	t.Log(v.IsEmpty(), v.Len(), v.Cap())

	for !v.IsEmpty() {
		v.PopBack()
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
	v := New(1000)

	// auto clean
	v.UsePurge(1000, 10*time.Millisecond)

	for i := 0; i < 10000; i++ {
		v.Push(i)
	}

	if v.Len() != 10000 {
		t.Fatal("len error. len:", v.Len())
	}

	t.Log(v.Len(), v.Cap())

	for !v.IsEmpty() {
		v.PopFront()
	}

	time.Sleep(30 * time.Millisecond)

	if v.Cap() != 1000 {
		t.Fatal("Purge error. cap:", v.Cap())
	}

	t.Log("ok")
}
