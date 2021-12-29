//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/boobusy/vector"
)

func main() {
	for i := 0; i < 100; i++ {
		vector.Push(i)
	}

	fmt.Println(vector.IsEmpty(),vector.Len(),vector.Cap())
	fmt.Println(vector.Items())

	for !vector.IsEmpty() {
		fmt.Println(vector.PopBack()) // or fmt.Println(vector.PopFront())
	}

}
