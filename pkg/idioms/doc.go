// Package idioms demonstrates Go-specific patterns and best practices.
//
// This package covers idiomatic Go patterns that differentiate Go
// from other languages:
//   - Duck typing through implicit interface satisfaction
//   - Explicit error handling with errors.Is and errors.As
//   - Zero value semantics for usable defaults
//   - Goroutines and channels for concurrency
//   - Context propagation for cancellation and timeouts
//   - Defer for resource cleanup
//
// Key Go idioms:
//   - Accept interfaces, return structs
//   - Error handling at each call site
//   - Leverage zero values for initialization
//   - Use defer for cleanup (LIFO ordering)
//   - Context for cancellation propagation
//   - Channels for goroutine communication
//
// Example usage:
//
//	import "github.com/KrystianMarek/golang-202/pkg/idioms"
//
//	func main() {
//		// Interface-based dependency injection
//		var processor idioms.Processor = idioms.UpperCaseProcessor{}
//		result := processor.Process("hello")
//
//		// Error handling with errors.Is
//		if errors.Is(err, idioms.ErrNotFound) {
//			// Handle not found
//		}
//
//		// Concurrency with channels
//		ctx := context.Background()
//		numbers := idioms.GenerateNumbers(ctx, 1, 10)
//		squares := idioms.Square(ctx, numbers)
//	}
package idioms

