// Package oop demonstrates object-oriented programming patterns in Go
// using composition, interfaces, and struct embedding.
//
// Go doesn't have traditional class-based inheritance, but provides
// powerful alternatives through:
//   - Struct embedding for composition
//   - Interfaces for polymorphism
//   - Methods for behavior
//   - Dependency injection via interfaces
//
// This package covers:
//   - Composition over inheritance
//   - Interface-based polymorphism
//   - Component-based design
//   - Dependency injection
//   - Gang of Four design patterns (see patterns subpackage)
//
// Example usage:
//
//	import (
//		"github.com/KrystianMarek/golang-202/pkg/oop"
//		"github.com/KrystianMarek/golang-202/pkg/oop/patterns"
//	)
//
//	func main() {
//		oop.ExampleComposition()
//		patterns.ExampleSingleton()
//	}
package oop
