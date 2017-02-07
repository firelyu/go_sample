package main

import (
	"golang.org/x/tour/tree"
	"fmt"

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

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var s1, s2 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, <-ch1)
	}
	//fmt.Println(s1)
	for i := 0; i < 10; i++ {
		s2 = append(s2, <-ch2)
	}
	//fmt.Println(s2)

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

	go Walk(tree.New(5), ch)

	//for v := range ch {
	//	fmt.Println(v)
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	res := Same(tree.New(91), tree.New(1))
	fmt.Println(res)
}