# GoLang-202: Advanced Go Patterns & Features

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/KrystianMarek/golang-202)](https://goreportcard.com/report/github.com/KrystianMarek/golang-202)

A comprehensive, importable Go library showcasing advanced Go concepts, modern features (Go 1.24+), design patterns, and idiomatic practices. This repository serves as both a learning resource and a reusable module for Go developers.

## 🎯 Features

- **Go 1.24+ Features**: Iterator functions, `unique` package, `runtime.AddCleanup`, generic type aliases
- **OOP Patterns**: Composition-based design, struct embedding, interface polymorphism
- **Design Patterns**: 23 Gang of Four patterns adapted to Go idioms
- **Functional Programming**: Higher-order functions, immutability, lazy evaluation, pipelines
- **Go Idioms**: Error handling, concurrency, zero values, interfaces, context propagation
- **Fully Tested**: Comprehensive test coverage with benchmarks
- **Importable**: Use as a module in your own projects

## 📦 Installation

```bash
go get github.com/KrystianMarek/golang-202
```

## 🚀 Quick Start

### As a Library

```go
package main

import (
    "fmt"
    "github.com/KrystianMarek/golang-202/pkg/go124"
    "github.com/KrystianMarek/golang-202/pkg/functional"
)

func main() {
    // Use Go 1.24 iterators
    for val := range go124.Range(1, 10) {
        fmt.Println(val)
    }

    // Functional pipelines
    numbers := []int{1, 2, 3, 4, 5}
    result := functional.NewPipeline(numbers).
        Filter(func(n int) bool { return n%2 == 0 }).
        Map(func(n int) int { return n * n }).
        Collect()

    fmt.Println(result) // [4, 16]
}
```

### Running Examples

```bash
# Build and run all examples
go run cmd/examples/main.go

# Run specific example category
go run cmd/examples/main.go go124
go run cmd/examples/main.go patterns
go run cmd/examples/main.go functional
```

## 📚 Package Overview

### `pkg/go124` - Go 1.24 Features

Modern Go features introduced in version 1.24:

```go
import "github.com/KrystianMarek/golang-202/pkg/go124"

// Iterator functions
tree := &go124.TreeNode{Value: 5, Left: ..., Right: ...}
for val := range tree.InOrder() {
    fmt.Println(val)
}

// Value canonicalization
handle := unique.Make("interned-string")

// Generic type aliases
type IntList = go124.OrderedSlice[int]
```

**Key Topics:**
- Iterator functions (`iter.Seq`)
- `unique` package for value interning
- `runtime.AddCleanup` for resource management
- Parameterized type aliases
- Enhanced testing with `testing.B.Loop`

### `pkg/oop` - Object-Oriented Patterns

Composition-based OOP using Go's strengths:

```go
import "github.com/KrystianMarek/golang-202/pkg/oop"

// Struct embedding
type Extended struct {
    oop.Base
    Extra string
}

// Interface polymorphism
var shape oop.Shape = oop.Circle{Radius: 5}
fmt.Println(shape.Area())

// Dependency injection
service := oop.NewService(oop.ConsoleLogger{})
```

**Key Topics:**
- Struct embedding for composition
- Interface-based polymorphism
- Method receivers (value vs pointer)
- Dependency injection via interfaces

### `pkg/oop/patterns` - Design Patterns

23 Gang of Four patterns, Go-style:

```go
import "github.com/KrystianMarek/golang-202/pkg/oop/patterns"

// Singleton with sync.Once
config := patterns.GetConfig()

// Builder with fluent interface
req := patterns.NewRequestBuilder().
    Method("POST").
    URL("https://api.example.com").
    Header("Content-Type", "application/json").
    Build()

// Observer with channels
eventBus := patterns.NewChannelEventBus()
ch := eventBus.Subscribe("user.event")
```

**Patterns Implemented:**
- **Creational**: Singleton, Factory, Builder, Prototype, Abstract Factory
- **Structural**: Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy
- **Behavioral**: Observer, Strategy, Command, Chain of Responsibility, State, Iterator, and more

### `pkg/functional` - Functional Programming

Functional patterns leveraging Go's first-class functions:

```go
import "github.com/KrystianMarek/golang-202/pkg/functional"

// Higher-order functions
evens := functional.Filter(numbers, func(n int) bool { return n%2 == 0 })
doubled := functional.Map(evens, func(n int) int { return n * 2 })

// Immutable data structures
list := functional.NewImmutableList(1, 2, 3)
list2 := list.Add(4) // Returns new list

// Lazy pipelines
result := functional.NewPipeline(data).
    Filter(predicate).
    Map(transform).
    Take(10).
    Collect()
```

**Key Topics:**
- Map, Filter, Reduce, ForEach
- Function composition and currying
- Immutable data structures (copy-on-write)
- Lazy evaluation with iterators
- Pipeline-based data processing

### `pkg/idioms` - Go Idioms

Go-specific patterns and best practices:

```go
import "github.com/KrystianMarek/golang-202/pkg/idioms"

// Error handling
if errors.Is(err, idioms.ErrNotFound) {
    // Handle not found
}

var validationErr *idioms.ValidationError
if errors.As(err, &validationErr) {
    fmt.Println(validationErr.Field)
}

// Concurrency patterns
ctx := context.Background()
numbers := idioms.GenerateNumbers(ctx, 1, 100)
squares := idioms.Square(ctx, numbers)

// Zero value semantics
var buf idioms.Buffer // No initialization needed
buf.Write([]byte("data"))
```

**Key Topics:**
- Duck typing with interfaces
- Error wrapping (`errors.Is`, `errors.As`)
- Zero value initialization
- Goroutines and channels
- Context propagation
- `defer` for cleanup

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run with coverage:

```bash
go test -cover ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## 📖 Documentation

- [API Documentation](https://pkg.go.dev/github.com/KrystianMarek/golang-202)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Changelog](CHANGELOG.md)

## 🏗️ Project Structure

```
golang-202/
├── pkg/                    # Public packages
│   ├── go124/             # Go 1.24 features
│   ├── oop/               # OOP patterns
│   │   └── patterns/      # GoF design patterns
│   ├── functional/        # Functional programming
│   ├── idioms/            # Go idioms
│   └── examples/          # Integrated examples
├── cmd/
│   └── examples/          # CLI examples runner
├── internal/              # Private utilities
│   └── runner/
├── test/                  # Integration tests
├── docs/                  # Documentation
└── .github/workflows/     # CI/CD
```

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) first.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📋 Requirements

- Go 1.24 or higher
- No external runtime dependencies (stdlib only)

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎓 Learning Resources

This library is designed as a learning resource. Each package includes:
- Comprehensive godoc comments
- "Why?" explanations for each pattern
- Runnable examples
- Performance benchmarks where relevant
- Trade-off discussions

## 🔗 Related Projects

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)

## 📞 Support

- Create an [Issue](https://github.com/KrystianMarek/golang-202/issues) for bug reports or feature requests
- Star ⭐ this repository if you find it useful!

---

**Made with ❤️ for the Go community**

