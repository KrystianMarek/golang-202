# Getting Started with GoLang-202

## Quick Start Guide

### Prerequisites

- Go 1.24 or higher installed
- Git for version control
- (Optional) golangci-lint for code quality

### Installation

#### As a Library

```bash
# Add to your project
go get github.com/KrystianMarek/golang-202
```

#### Clone for Development

```bash
# Clone the repository
git clone git@github.com:KrystianMarek/golang-202.git
cd golang-202

# Download dependencies (if any added in future)
go mod download

# Verify installation
go build ./...
go test ./...
```

### Your First Program

Create a new file `example.go`:

```go
package main

import (
    "fmt"
    "github.com/KrystianMarek/golang-202/pkg/go124"
    "github.com/KrystianMarek/golang-202/pkg/functional"
)

func main() {
    // Example 1: Go 1.24 Iterators
    fmt.Println("Numbers 1-10:")
    for n := range go124.Range(1, 11) {
        fmt.Printf("%d ", n)
    }
    fmt.Println()

    // Example 2: Functional Pipeline
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result := functional.NewPipeline(numbers).
        Filter(func(n int) bool { return n%2 == 0 }).
        Map(func(n int) int { return n * n }).
        Collect()

    fmt.Printf("Even squares: %v\n", result)
}
```

Run it:

```bash
go run example.go
```

### Running Built-in Examples

The repository includes comprehensive examples:

```bash
# Run all examples
go run cmd/examples/main.go

# Run specific category
go run cmd/examples/main.go go124      # Go 1.24 features
go run cmd/examples/main.go oop        # OOP patterns
go run cmd/examples/main.go functional # Functional programming
go run cmd/examples/main.go idioms     # Go idioms
go run cmd/examples/main.go patterns   # Design patterns
```

### Exploring Packages

#### Go 1.24 Features

```go
import "github.com/KrystianMarek/golang-202/pkg/go124"

// Iterator for tree traversal
tree := &go124.TreeNode{Value: 5, ...}
for val := range tree.InOrder() {
    fmt.Println(val)
}

// String interning
handle := unique.Make("common-string")

// Generic type aliases
type MyList = go124.OrderedSlice[int]
```

#### Design Patterns

```go
import "github.com/KrystianMarek/golang-202/pkg/oop/patterns"

// Singleton
config := patterns.GetConfig()

// Builder
req := patterns.NewRequestBuilder().
    Method("POST").
    URL("https://api.example.com").
    Build()

// Observer
bus := patterns.NewChannelEventBus()
ch := bus.Subscribe("events")
```

#### Functional Programming

```go
import "github.com/KrystianMarek/golang-202/pkg/functional"

// Higher-order functions
evens := functional.Filter(numbers, func(n int) bool {
    return n%2 == 0
})

// Immutable data
list := functional.NewImmutableList(1, 2, 3)
list2 := list.Add(4) // Returns new list

// Pipelines
result := functional.NewPipeline(data).
    Filter(predicate).
    Map(transform).
    Collect()
```

#### Go Idioms

```go
import (
    "github.com/KrystianMarek/golang-202/pkg/idioms"
    "context"
    "errors"
)

// Error handling
if errors.Is(err, idioms.ErrNotFound) {
    // Handle not found
}

// Concurrency
ctx := context.Background()
nums := idioms.GenerateNumbers(ctx, 1, 100)
squares := idioms.Square(ctx, nums)

// Zero values
var buf idioms.Buffer // Ready to use
buf.Write([]byte("data"))
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# View coverage in browser
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. ./...

# Run specific package tests
go test ./pkg/go124/
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter (requires golangci-lint)
golangci-lint run
```

### Building

```bash
# Build all packages
go build ./...

# Build examples binary
go build -o bin/examples cmd/examples/main.go
./bin/examples

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o bin/examples-linux cmd/examples/main.go
GOOS=windows GOARCH=amd64 go build -o bin/examples.exe cmd/examples/main.go
```

### Next Steps

1. **Explore the Documentation**
   - Read package godocs: `go doc github.com/KrystianMarek/golang-202/pkg/go124`
   - Browse online: [pkg.go.dev](https://pkg.go.dev/github.com/KrystianMarek/golang-202)

2. **Study the Examples**
   - Review `cmd/examples/main.go` for usage patterns
   - Check `pkg/examples/game_engine.go` for integration examples

3. **Read the Patterns**
   - Each pattern file includes detailed explanations
   - Look for "Why?" sections explaining the rationale

4. **Try the Exercises**
   - Implement your own patterns using the library
   - Combine multiple patterns in a project
   - Benchmark different approaches

### Common Tasks

#### Add to Existing Project

```bash
cd your-project
go get github.com/KrystianMarek/golang-202
```

Then import packages as needed:

```go
import (
    "github.com/KrystianMarek/golang-202/pkg/functional"
    "github.com/KrystianMarek/golang-202/pkg/oop/patterns"
)
```

#### Update to Latest Version

```bash
go get -u github.com/KrystianMarek/golang-202
go mod tidy
```

#### Check Specific Version

```bash
# Use specific version
go get github.com/KrystianMarek/golang-202@v0.1.0

# Use latest commit
go get github.com/KrystianMarek/golang-202@latest
```

### Troubleshooting

**Issue: Go version too old**
```
Solution: Upgrade to Go 1.24+
$ go version
$ brew upgrade go  # macOS
```

**Issue: Import errors**
```
Solution: Run go mod tidy
$ go mod tidy
```

**Issue: Build failures**
```
Solution: Check Go version and run tests
$ go version
$ go build ./...
$ go test ./...
```

### Getting Help

- Check the [README](../README.md) for overview
- Read [CONTRIBUTING](../CONTRIBUTING.md) for guidelines
- Open an [Issue](https://github.com/KrystianMarek/golang-202/issues) for bugs
- Review existing examples in the `pkg/` directory

### Learning Path

**Beginner**: Start with Go idioms
1. `pkg/idioms/zero_values.go` - Understand Go's zero values
2. `pkg/idioms/errors.go` - Master error handling
3. `pkg/idioms/interfaces.go` - Learn interface patterns

**Intermediate**: Explore patterns
1. `pkg/oop/composition.go` - OOP in Go
2. `pkg/oop/patterns/singleton.go` - Basic patterns
3. `pkg/functional/higher_order.go` - Functional patterns

**Advanced**: Study modern features
1. `pkg/go124/iterators.go` - Iterator functions
2. `pkg/functional/pipelines.go` - Lazy evaluation
3. `pkg/examples/game_engine.go` - Integration

Happy coding! 🚀

