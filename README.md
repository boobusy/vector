# Getting Started
### Pull in the dependency
```zsh
go get github.com/boobusy/vector
```

### Add the import to your project
```go
import (
    "github.com/boobusy/vector"
)
```


###Simple
```go
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
```
###goroutine safe
```go
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
```


# Examples
* [Simple](https://github.com/boobusy/vector/blob/master/examples/simple.go)
* [Vector](https://github.com/boobusy/vector/blob/master/examples/vector.go)

## author
问题或者联系作者：十二楼五城

![boobusy](http://boobusy.com/wechat.jpg)