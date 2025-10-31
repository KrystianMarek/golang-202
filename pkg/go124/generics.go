package go124

import "fmt"

// Generics demonstrates Go's generic programming features.
//
// Why? Generics enable type-safe, reusable code without reflection or
// code duplication, making it easier to build robust data structures
// and algorithms that work with any type.

// Stack is a generic stack implementation.
type Stack[T any] []T

// NewStack creates a new stack.
func NewStack[T any]() *Stack[T] {
	s := make(Stack[T], 0)
	return &s
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop removes and returns the top element.
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx]
	return val, true
}

// Peek returns the top element without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	return (*s)[len(*s)-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the number of elements.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// Queue is a generic FIFO queue.
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an element to the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

// Dequeue removes and returns the first element.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Set is a generic set using map.
type Set[T comparable] map[T]struct{}

// NewSet creates a new set.
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add adds an element to the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Remove removes an element from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Contains checks if an element exists.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Size returns the number of elements.
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice converts the set to a slice.
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}

// Union returns the union of two sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s {
		result.Add(k)
	}
	for k := range other {
		result.Add(k)
	}
	return result
}

// Intersection returns the intersection of two sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// BinaryTree is a generic binary tree.
type BinaryTree[T any] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

// NewBinaryTree creates a new binary tree node.
func NewBinaryTree[T any](value T) *BinaryTree[T] {
	return &BinaryTree[T]{Value: value}
}

// Insert inserts a value (requires comparison).
func (t *BinaryTree[T]) Insert(value T, less func(T, T) bool) {
	if less(value, t.Value) {
		if t.Left == nil {
			t.Left = NewBinaryTree(value)
		} else {
			t.Left.Insert(value, less)
		}
	} else {
		if t.Right == nil {
			t.Right = NewBinaryTree(value)
		} else {
			t.Right.Insert(value, less)
		}
	}
}

// Number represents constraint-based number types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Min returns the minimum of two values.
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values.
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of all values.
func Sum[T Number](values ...T) T {
	var result T
	for _, v := range values {
		result += v
	}
	return result
}

// GenericResult represents a value or error (more advanced than Result).
type GenericResult[T any] struct {
	value T
	err   error
}

// OkResult creates a successful result.
func OkResult[T any](value T) GenericResult[T] {
	return GenericResult[T]{value: value}
}

// ErrResult creates an error result.
func ErrResult[T any](err error) GenericResult[T] {
	var zero T
	return GenericResult[T]{value: zero, err: err}
}

// IsOk returns true if the result is successful.
func (r GenericResult[T]) IsOk() bool {
	return r.err == nil
}

// Unwrap returns the value and error.
func (r GenericResult[T]) Unwrap() (T, error) {
	return r.value, r.err
}

// Map transforms the value if successful.
func (r GenericResult[T]) Map(fn func(T) T) GenericResult[T] {
	if r.IsOk() {
		return OkResult(fn(r.value))
	}
	return r
}

// Cache is a generic cache with type safety.
type Cache[K comparable, V any] struct {
	data map[K]V
}

// NewCache creates a new cache.
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

// Get retrieves a value.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, ok := c.data[key]
	return val, ok
}

// Set stores a value.
func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

// Delete removes a value.
func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}

// Keys returns all keys.
func (c *Cache[K, V]) Keys() []K {
	keys := make([]K, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

// ExampleGenerics demonstrates generic programming.
func ExampleGenerics() {
	fmt.Println("=== Generics ===")

	// Generic stack
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Integer Stack:")
	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// String stack
	stringStack := NewStack[string]()
	stringStack.Push("hello")
	stringStack.Push("world")

	val, _ := stringStack.Peek()
	fmt.Printf("Peek: %s, Size: %d\n", val, stringStack.Size())

	// Generic queue
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("\nQueue:")
	for !queue.IsEmpty() {
		val, _ := queue.Dequeue()
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Generic set
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	fmt.Printf("\nSet1: %v\n", set1.ToSlice())
	fmt.Printf("Set2: %v\n", set2.ToSlice())
	fmt.Printf("Union: %v\n", set1.Union(set2).ToSlice())
	fmt.Printf("Intersection: %v\n", set1.Intersection(set2).ToSlice())

	// Generic binary tree
	tree := NewBinaryTree(5)
	tree.Insert(3, func(a, b int) bool { return a < b })
	tree.Insert(7, func(a, b int) bool { return a < b })
	tree.Insert(1, func(a, b int) bool { return a < b })

	fmt.Printf("\nTree root: %d\n", tree.Value)
	fmt.Printf("Tree left: %d\n", tree.Left.Value)
	fmt.Printf("Tree right: %d\n", tree.Right.Value)

	// Generic math functions
	fmt.Printf("\nMin(5, 3): %d\n", Min(5, 3))
	fmt.Printf("Max(5, 3): %d\n", Max(5, 3))
	fmt.Printf("Sum(1, 2, 3, 4, 5): %d\n", Sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum(1.5, 2.5, 3.5): %.1f\n", Sum(1.5, 2.5, 3.5))

	// Generic result type
	successResult := OkResult(42)
	fmt.Printf("\nResult is OK: %v, value: %d\n",
		successResult.IsOk(), successResult.value)

	doubledResult := successResult.Map(func(v int) int { return v * 2 })
	fmt.Printf("Doubled: %d\n", doubledResult.value)

	// Generic cache
	cache := NewCache[string, int]()
	cache.Set("age", 25)
	cache.Set("score", 100)

	age, ok := cache.Get("age")
	fmt.Printf("\nCache get 'age': %d (found: %v)\n", age, ok)
	fmt.Printf("Cache keys: %v\n", cache.Keys())
}
