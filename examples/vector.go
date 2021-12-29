//go:build ignore
// +build ignore


package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/boobusy/vector"
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

	for !queue.IsEmpty() {
		fmt.Println(queue.PopBack().(*Task).Id) // or fmt.Println(vector.PopFront())
	}

}
