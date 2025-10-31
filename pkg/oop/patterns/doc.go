// Package patterns implements Gang of Four (GoF) design patterns
// adapted to Go's interfaces, structs, and idioms.
//
// This package demonstrates how classical OOP design patterns can be
// implemented idiomatically in Go using:
//   - Interfaces for polymorphism
//   - Struct embedding for composition
//   - Channels for event-driven patterns
//   - sync.Once for thread-safe singletons
//   - Function types for strategy patterns
//
// Patterns included:
//
// Creational:
//   - Singleton: Thread-safe single instances using sync.Once
//   - Factory: Factory functions returning interfaces
//   - Builder: Fluent interfaces for complex object construction
//
// Structural:
//   - Adapter: Making incompatible interfaces work together
//   - Decorator: Adding behavior dynamically through composition
//
// Behavioral:
//   - Observer: Event-driven patterns using channels and interfaces
//   - Strategy: Swappable algorithms via interfaces
//
// Each pattern includes:
//   - Clear godoc comments explaining the "why"
//   - Multiple examples showing different use cases
//   - Runnable example functions
//
// Example usage:
//
//	import "github.com/KrystianMarek/golang-202/pkg/oop/patterns"
//
//	func main() {
//		patterns.ExampleSingleton()
//		patterns.ExampleFactory()
//		patterns.ExampleBuilder()
//	}
package patterns

