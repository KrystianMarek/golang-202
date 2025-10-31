// Package go124 provides examples and demonstrations of features
// introduced in Go 1.24 (released February 2025).
//
// This package covers:
//   - Iterator functions for custom iteration patterns (iter.Seq)
//   - Value canonicalization with unique.Handle
//   - Resource cleanup with runtime.AddCleanup
//   - Parameterized type aliases for generic types
//   - Comprehensive generic programming (type parameters, constraints)
//   - Enhanced testing benchmarks with testing.B.Loop
//
// Each file contains focused examples with godoc comments explaining
// the "why" behind each feature and demonstrating idiomatic usage.
//
// Example usage:
//
//	import "github.com/KrystianMarek/golang-202/pkg/go124"
//
//	func main() {
//		// Iterator functions
//		go124.ExampleIterators()
//
//		// Value interning
//		go124.ExampleUnique()
//
//		// Resource cleanup
//		go124.ExampleCleanup()
//
//		// Generic type aliases
//		go124.ExampleGenericAliases()
//
//		// Generic data structures
//		go124.ExampleGenerics()
//	}
package go124
