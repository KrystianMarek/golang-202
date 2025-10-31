// Package functional demonstrates functional programming patterns in Go.
package functional

import "fmt"

// Higher-order functions demonstrate functions as first-class citizens.
//
// Why? Functions as values enable composition, callbacks, and declarative
// programming styles that can lead to more maintainable code.

// Predicate is a function type that tests a condition.
type Predicate[T any] func(T) bool

// Mapper is a function type that transforms values.
type Mapper[T, U any] func(T) U

// Reducer is a function type that combines values.
type Reducer[T, U any] func(U, T) U

// Filter returns a new slice containing only elements that satisfy the predicate.
func Filter[T any](slice []T, predicate Predicate[T]) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map transforms each element using the mapper function.
func Map[T, U any](slice []T, mapper Mapper[T, U]) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// Reduce combines all elements using the reducer function.
func Reduce[T, U any](slice []T, initial U, reducer Reducer[T, U]) U {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// ForEach applies a function to each element.
func ForEach[T any](slice []T, fn func(T)) {
	for _, item := range slice {
		fn(item)
	}
}

// Any returns true if any element satisfies the predicate.
func Any[T any](slice []T, predicate Predicate[T]) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate.
func All[T any](slice []T, predicate Predicate[T]) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Compose composes two functions: f(g(x)).
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}

// Curry converts a function of two arguments into a curried function.
func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

// Partial partially applies a function.
func Partial[A, B, C any](f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

// Pipe creates a pipeline of functions.
func Pipe[T any](initial T, functions ...func(T) T) T {
	result := initial
	for _, fn := range functions {
		result = fn(result)
	}
	return result
}

// Memoize caches function results.
func Memoize[K comparable, V any](fn func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := fn(k)
		cache[k] = v
		return v
	}
}

// Once ensures a function is called only once.
func Once[T any](fn func() T) func() T {
	var (
		result T
		called bool
	)
	return func() T {
		if !called {
			result = fn()
			called = true
		}
		return result
	}
}

// Debounce returns a debounced version of the function.
func Debounce[T any](fn func(T), threshold int) func(T) {
	count := 0
	return func(t T) {
		count++
		if count >= threshold {
			fn(t)
			count = 0
		}
	}
}

// ExampleHigherOrder demonstrates higher-order functions.
func ExampleHigherOrder() {
	fmt.Println("=== Higher-Order Functions ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	// Map
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)

	// Reduce
	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("Sum: %d\n", sum)

	// Any and All
	hasEven := Any(numbers, func(n int) bool { return n%2 == 0 })
	allPositive := All(numbers, func(n int) bool { return n > 0 })
	fmt.Printf("Has even: %v, All positive: %v\n\n", hasEven, allPositive)

	// Composition
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	addOneThenDouble := Compose(double, addOne)
	fmt.Printf("Compose (5+1)*2 = %d\n", addOneThenDouble(5))

	// Currying
	add := func(a, b int) int { return a + b }
	curriedAdd := Curry(add)
	add5 := curriedAdd(5)
	fmt.Printf("Curried add(5)(3) = %d\n", add5(3))

	// Partial application
	multiply := func(a, b int) int { return a * b }
	double2 := Partial(multiply, 2)
	fmt.Printf("Partial multiply(2, 7) = %d\n\n", double2(7))

	// Pipeline
	result := Pipe(5,
		func(x int) int { return x + 1 },
		func(x int) int { return x * 2 },
		func(x int) int { return x - 3 },
	)
	fmt.Printf("Pipeline (5+1)*2-3 = %d\n", result)

	// Memoization
	fibonacci := func(n int) int {
		if n <= 1 {
			return n
		}
		a, b := 0, 1
		for i := 2; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
	memoFib := Memoize(fibonacci)
	fmt.Printf("Memoized fib(10) = %d\n", memoFib(10))
	fmt.Printf("Memoized fib(10) = %d (cached)\n", memoFib(10))
}

