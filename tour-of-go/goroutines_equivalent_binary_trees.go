package main

import (
	"fmt"
	"math/rand"
)

// There can be many different binary trees with the same sequence of values stored in it. 
// A function to check whether two binary trees store the same sequence is quite complex in
// most languages. We'll use Go's concurrency and channels to write a simple solution.

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// NewRandomTree returns a new, random binary tree holding the values k, 2k, ..., nk.
func NewRandomTree(k int, n int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *Tree, ch chan int) {
	if t != nil {
		WalkRecursive(t.Left, ch)
		ch <- t.Value
		WalkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	fmt.Println(t1)
	fmt.Println(t2)

	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
        n1, ok1 := <- ch1
		n2, ok2 := <- ch2
		// Trees not same size or nodes not same value
        if ok1 != ok2 || n1 != n2 {
        	return false
		}
		// If we reach !ok1 (or !ok2) without return false above
		// then we have exhausted the trees without finding a difference
		// and the trees are the same
		if !ok1 {
        	break;
        }
	}
	return true
}

func main() {
	ch := make(chan int)
	tree := NewRandomTree(1, 10)
	fmt.Println("Tree to walk: ", tree)

	go Walk(tree, ch)

	for node := range ch {
		fmt.Println(node)
	}

	fmt.Println(Same(NewRandomTree(1, 10), NewRandomTree(1, 10)))
	fmt.Println(Same(NewRandomTree(1, 10), NewRandomTree(2, 10)))
	fmt.Println(Same(NewRandomTree(2, 10), NewRandomTree(1, 10)))
	fmt.Println(Same(NewRandomTree(1, 11), NewRandomTree(1, 10)))
	fmt.Println(Same(NewRandomTree(1, 10), NewRandomTree(1, 11)))

}

