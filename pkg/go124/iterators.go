// Package go124 demonstrates features introduced in Go 1.24 (released February 2025).
package go124

import (
	"fmt"
	"iter"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// InOrder returns an iterator that traverses the tree in-order.
// This demonstrates Go 1.24's iterator functions in for-range loops.
//
// Why? Iterator functions allow custom iteration patterns without
// materializing collections, enabling memory-efficient tree traversals.
func (t *TreeNode) InOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		var traverse func(*TreeNode) bool
		traverse = func(node *TreeNode) bool {
			if node == nil {
				return true
			}
			if !traverse(node.Left) {
				return false
			}
			if !yield(node.Value) {
				return false
			}
			return traverse(node.Right)
		}
		traverse(t)
	}
}

// PreOrder returns an iterator for pre-order traversal.
func (t *TreeNode) PreOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		var traverse func(*TreeNode) bool
		traverse = func(node *TreeNode) bool {
			if node == nil {
				return true
			}
			if !yield(node.Value) {
				return false
			}
			if !traverse(node.Left) {
				return false
			}
			return traverse(node.Right)
		}
		traverse(t)
	}
}

// Range returns an iterator that generates integers from start to end (exclusive).
// Demonstrates creating custom numeric sequences with iterators.
func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Filter returns an iterator that yields only values satisfying the predicate.
// This shows functional composition with iterators.
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map transforms values from the source iterator using the given function.
func Map[T, U any](seq iter.Seq[T], fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// ExampleIterators demonstrates iterator usage with Go 1.24's for-range support.
func ExampleIterators() {
	tree := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: 6,
			Left:  &TreeNode{Value: 5},
			Right: &TreeNode{Value: 7},
		},
	}

	fmt.Println("In-order traversal:")
	for val := range tree.InOrder() {
		fmt.Printf("%d ", val)
	}


	fmt.Println("Even numbers from 0 to 10:")
	evens := Filter(Range(0, 10), func(n int) bool { return n%2 == 0 })
	for val := range evens {
		fmt.Printf("%d ", val)
	}


	fmt.Println("Squares of 1 to 5:")
	squares := Map(Range(1, 6), func(n int) int { return n * n })
	for val := range squares {
		fmt.Printf("%d ", val)
	}

}

