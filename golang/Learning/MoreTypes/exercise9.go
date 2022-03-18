//Exercise: Equivalent Binary Trees
//1. Implement the Walk function.
//2. Test the Walk function.
//
//The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
//Create a new channel ch and kick off the walker:
//go Walk(tree.New(1), ch)
//Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
//3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
//4. Test the Same function.
//
//Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// used the preorder (left, root, right)
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go Walk(t1, channel1)
	go Walk(t2, channel2)
	for i := 0; i < 10; i++ {
		if <-channel1 != <-channel2 {
			return false
		}
	}

	return true
}

func main() {
	//tree := tree.New(1)
	//pipe := make(chan int)

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

	//go Walk(tree, pipe)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-pipe)
	//}

}
