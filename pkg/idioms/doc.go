// Package idioms demonstrates Go-specific patterns and best practices.
//
// This package covers idiomatic Go patterns that differentiate Go
// from other languages:
//   - Duck typing through implicit interface satisfaction
//   - Explicit error handling with errors.Is and errors.As
//   - Zero value semantics for usable defaults
//   - Goroutines and channels for concurrency
//   - Go 1.24 enhanced channel patterns (safe for-range, context integration)
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
//   - Go 1.24: Guaranteed channel termination with for-range
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
//		// Concurrency with channels (Go 1.24)
//		ctx := context.Background()
//		numbers := idioms.GenerateNumbers(ctx, 1, 10)
//		squares := idioms.Square(ctx, numbers)
//
//		// Safe pipeline with guaranteed termination
//		output := idioms.SafePipeline(ctx, numbers)
//		for val := range output { // Guaranteed to terminate
//			fmt.Println(val)
//		}
//	}
package idioms
