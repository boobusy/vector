//go:build ignore
// +build ignore


package main

import (
	"fmt"
	"github.com/boobusy/vector"
	"sync"
	"time"
)

type Task struct {
	Id int
}

func main() {

	queue := vector.New(10000)

	// auto clean
	queue.UsePurge(1<<20, 100 * time.Millisecond)

	w := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			for k := 0; k < 1000; k++ {
				task := &Task{
					Id: k,
				}
				queue.Push(task)
			}
		}()
	}

	w.Wait()
	fmt.Println(queue.IsEmpty(), queue.Len(), queue.Cap())

	var val vector.Val
	for !queue.IsEmpty() {
		val = queue.PopBack()
		fmt.Println(val.(*Task).Id) // or fmt.Println(vector.PopFront())
	}


}
