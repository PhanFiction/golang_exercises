package generic_exercise1

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	node2 := &List[int]{val: 2, next: nil}
	node1 := &List[int]{val: 1, next: node2}

	fmt.Println(node1.val, node2.val)
}
