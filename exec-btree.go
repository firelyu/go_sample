package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		// Walk left tree
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		// Walk right tree
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	var s1, s2 []int
	for v := range ch1 {
		s1 = append(s1, v)
	}
	for v := range ch2 {
		s2 = append(s2, v)
	}

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	// Get help from the following link.
	// http://stackoverflow.com/questions/42093495/use-range-channel-in-go-error-fatal#comment-71365649
	go func() {
		Walk(tree.New(3), ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
}
