// Package functional demonstrates functional programming patterns in Go.
//
// This package covers functional programming concepts adapted to Go:
//   - Higher-order functions (map, filter, reduce)
//   - Function composition and currying
//   - Immutable data structures with copy-on-write
//   - Lazy evaluation through iterators (Go 1.24+)
//   - Pipeline-based data processing
//
// Go supports functional programming through:
//   - First-class functions
//   - Closures for state encapsulation
//   - Generic types for type-safe operations
//   - Iterators for lazy evaluation (Go 1.24+)
//
// Trade-offs:
//   - Immutability increases memory usage but improves safety
//   - Lazy evaluation reduces memory but adds complexity
//   - Functional style can be more declarative but less performant
//
// Example usage:
//
//	import "github.com/KrystianMarek/golang-202/pkg/functional"
//
//	func main() {
//		numbers := []int{1, 2, 3, 4, 5}
//		evens := functional.Filter(numbers, func(n int) bool { return n%2 == 0 })
//		doubled := functional.Map(evens, func(n int) int { return n * 2 })
//
//		// Or use pipelines
//		result := functional.NewPipeline(numbers).
//			Filter(func(n int) bool { return n%2 == 0 }).
//			Map(func(n int) int { return n * 2 }).
//			Collect()
//	}
package functional
