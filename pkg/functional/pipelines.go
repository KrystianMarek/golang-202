package functional

import (
	"fmt"
	"iter"
	"strings"
)

// Pipelines demonstrate iterator-based data processing.
//
// Why? Lazy evaluation through iterators enables memory-efficient
// processing of large datasets without materializing intermediate results.

// Generator creates an iterator from a slice.
func Generator[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

// Take limits the number of items from an iterator.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for item := range seq {
			if count >= n {
				return
			}
			if !yield(item) {
				return
			}
			count++
		}
	}
}

// Skip skips the first n items from an iterator.
func Skip[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for item := range seq {
			if count >= n {
				if !yield(item) {
					return
				}
			}
			count++
		}
	}
}

// Chain concatenates multiple iterators.
func Chain[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for item := range seq {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Zip combines two iterators into pairs.
func Zip[A, B any](seqA iter.Seq[A], seqB iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		nextB, stopB := iter.Pull(seqB)
		defer stopB()

		for a := range seqA {
			b, ok := nextB()
			if !ok {
				return
			}
			if !yield(a, b) {
				return
			}
		}
	}
}

// Enumerate adds indices to an iterator.
func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for item := range seq {
			if !yield(index, item) {
				return
			}
			index++
		}
	}
}

// Collect materializes an iterator into a slice.
func Collect[T any](seq iter.Seq[T]) []T {
	result := make([]T, 0)
	for item := range seq {
		result = append(result, item)
	}
	return result
}

// Pipeline represents a composable data pipeline.
type Pipeline[T any] struct {
	source iter.Seq[T]
}

// NewPipeline creates a new pipeline from a slice.
func NewPipeline[T any](items []T) *Pipeline[T] {
	return &Pipeline[T]{source: Generator(items)}
}

// FromSeq creates a pipeline from an iterator.
func FromSeq[T any](seq iter.Seq[T]) *Pipeline[T] {
	return &Pipeline[T]{source: seq}
}

// Filter applies a filter to the pipeline.
func (p *Pipeline[T]) Filter(predicate func(T) bool) *Pipeline[T] {
	return &Pipeline[T]{
		source: func(yield func(T) bool) {
			for item := range p.source {
				if predicate(item) {
					if !yield(item) {
						return
					}
				}
			}
		},
	}
}

// Map applies a transformation to the pipeline.
func (p *Pipeline[T]) Map(mapper func(T) T) *Pipeline[T] {
	return &Pipeline[T]{
		source: func(yield func(T) bool) {
			for item := range p.source {
				if !yield(mapper(item)) {
					return
				}
			}
		},
	}
}

// Take limits the pipeline to n items.
func (p *Pipeline[T]) Take(n int) *Pipeline[T] {
	return &Pipeline[T]{source: Take(p.source, n)}
}

// Skip skips the first n items.
func (p *Pipeline[T]) Skip(n int) *Pipeline[T] {
	return &Pipeline[T]{source: Skip(p.source, n)}
}

// Collect materializes the pipeline into a slice.
func (p *Pipeline[T]) Collect() []T {
	return Collect(p.source)
}

// ForEach applies a function to each item.
func (p *Pipeline[T]) ForEach(fn func(T)) {
	for item := range p.source {
		fn(item)
	}
}

// Reduce combines all items using a reducer.
func (p *Pipeline[T]) Reduce(initial T, reducer func(T, T) T) T {
	result := initial
	for item := range p.source {
		result = reducer(result, item)
	}
	return result
}

// Count returns the number of items.
func (p *Pipeline[T]) Count() int {
	count := 0
	for range p.source {
		count++
	}
	return count
}

// First returns the first item or zero value.
func (p *Pipeline[T]) First() T {
	for item := range p.source {
		return item
	}
	var zero T
	return zero
}

// ExamplePipelines demonstrates iterator-based pipelines.
func ExamplePipelines() {
	fmt.Println("=== Functional Pipelines ===")

	// Basic pipeline
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := NewPipeline(numbers).
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * n }).
		Take(3).
		Collect()

	fmt.Printf("Pipeline result: %v\n\n", result)

	// String processing pipeline
	words := []string{"hello", "world", "go", "functional", "programming"}

	upperWords := NewPipeline(words).
		Filter(func(s string) bool { return len(s) > 3 }).
		Map(func(s string) string { return strings.ToUpper(s) }).
		Collect()

	fmt.Printf("Uppercase words: %v\n\n", upperWords)

	// Chain iterators
	seq1 := Generator([]int{1, 2, 3})
	seq2 := Generator([]int{4, 5, 6})
	seq3 := Generator([]int{7, 8, 9})

	chained := Collect(Chain(seq1, seq2, seq3))
	fmt.Printf("Chained: %v\n\n", chained)

	// Zip iterators
	names := Generator([]string{"Alice", "Bob", "Carol"})
	ages := Generator([]int{25, 30, 35})

	fmt.Println("Zipped:")
	for name, age := range Zip(names, ages) {
		fmt.Printf("  %s: %d\n", name, age)
	}


	// Enumerate
	items := Generator([]string{"apple", "banana", "cherry"})
	fmt.Println("Enumerated:")
	for idx, item := range Enumerate(items) {
		fmt.Printf("  [%d] %s\n", idx, item)
	}


	// Complex pipeline with reduce
	sum := NewPipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * 2 }).
		Reduce(0, func(acc, n int) int { return acc + n })

	fmt.Printf("Sum of doubled evens: %d\n", sum)

	// Take and Skip
	page2 := NewPipeline(numbers).
		Skip(3).
		Take(3).
		Collect()

	fmt.Printf("Page 2 (skip 3, take 3): %v\n", page2)

	// Count
	count := NewPipeline(numbers).
		Filter(func(n int) bool { return n > 5 }).
		Count()

	fmt.Printf("Count of numbers > 5: %d\n", count)
}

