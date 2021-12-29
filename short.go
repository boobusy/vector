package vector

var (
	defautlVector = New(1<<7 - 1)
)

// Len returns the number of items.
func Len() int {
	return len(defautlVector.items)
}

// Cap returns the capacity of items.
func Cap() int {
	return cap(defautlVector.items)
}

func IsEmpty() bool {
	return defautlVector.Len() == 0
}

// Items returns the items.
func Items() []any {
	return defautlVector.items
}

// Push inserts a new item e with value v at the back of items.
func Push(item any) {
	defautlVector.Push(item)
}

// Remake items size
func Remake(items []any) {
	defautlVector.Remake(items)
}

// PopFront returns the first val of items.
func PopFront() any {
	return defautlVector.PopFront()
}

// PopBack returns the last val of items.
func PopBack() any {
	return defautlVector.PopBack()
}
