# Contributing to GoLang-202

First off, thank you for considering contributing to GoLang-202! It's people like you that make this a great learning resource for the Go community.

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please be respectful and constructive in all interactions.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the behavior
- **Expected behavior** vs actual behavior
- **Go version** (`go version`)
- **Code sample** demonstrating the issue

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

- **Use case** - why is this enhancement needed?
- **Describe the solution** you'd like
- **Describe alternatives** you've considered
- **Example code** showing the proposed API

### Pull Requests

1. **Fork** the repository and create your branch from `main`
2. **Add tests** for any new functionality
3. **Update documentation** including godoc comments
4. **Ensure tests pass**: `go test ./...`
5. **Run linter**: `golangci-lint run`
6. **Follow Go conventions** and the project's style

## Development Setup

### Prerequisites

- Go 1.24+
- golangci-lint (optional but recommended)

### Getting Started

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/golang-202.git
cd golang-202

# Install dependencies
go mod download

# Run tests
go test ./...

# Run linter
golangci-lint run
```

## Style Guidelines

### Go Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` for formatting (enforced by CI)
- Keep functions focused and small (< 50 lines when possible)
- Prefer composition over embedding when appropriate
- Write descriptive variable names (avoid single letters except for loops/closures)

### Documentation Style

- **Every exported type/function** must have a godoc comment
- Start comments with the name of the thing being described
- Include **"Why?"** explanations for patterns and design decisions
- Provide **usage examples** in godoc
- Keep examples **compilable** and **self-contained**

Example:

```go
// Cache stores key-value pairs with automatic expiration.
// This demonstrates the lazy initialization pattern using sync.Once.
//
// Why? Lazy initialization defers expensive setup until first use,
// improving startup time and memory usage for optional features.
//
// Example:
//
//	cache := NewCache()
//	cache.Set("key", "value", 5*time.Minute)
//	val, found := cache.Get("key")
type Cache struct {
    // ...
}
```

### Testing Guidelines

- Write **table-driven tests** for multiple test cases
- Include **benchmarks** for performance-critical code
- Aim for **>80% coverage** for new code
- Use **descriptive test names**: `TestCacheConcurrentAccess`
- Test **edge cases** and **error paths**

Example:

```go
func TestFilter(t *testing.T) {
    tests := []struct {
        name      string
        input     []int
        predicate func(int) bool
        want      []int
    }{
        {
            name:      "even numbers",
            input:     []int{1, 2, 3, 4, 5},
            predicate: func(n int) bool { return n%2 == 0 },
            want:      []int{2, 4},
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Filter(tt.input, tt.predicate)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Reference issues and pull requests liberally
- Keep first line under 72 characters
- Add detailed description for complex changes

Good examples:
```
Add iterator support for binary trees

Implement InOrder and PreOrder traversal using Go 1.24's iter.Seq.
This allows lazy tree traversal without materializing all values.

Fixes #123
```

## Adding New Patterns or Features

When adding new patterns or features:

1. **Choose the right package**:
   - `pkg/go124/` - Go 1.24+ specific features
   - `pkg/oop/` - Composition and basic OOP
   - `pkg/oop/patterns/` - GoF design patterns
   - `pkg/functional/` - FP approaches
   - `pkg/idioms/` - Go-specific patterns

2. **Create a focused file**:
   - One pattern per file when possible
   - Keep files under 300 lines
   - Name files descriptively (e.g., `singleton.go`, `higher_order.go`)

3. **Include**:
   - Comprehensive godoc
   - Multiple examples showing different use cases
   - An `Example...()` function demonstrating usage
   - Tests covering main functionality
   - Benchmarks if performance is relevant

4. **Update documentation**:
   - Add to package `doc.go`
   - Update README.md if adding major features
   - Add to examples in `cmd/examples/main.go`

## Project Goals

Keep these goals in mind when contributing:

1. **Educational Focus**: Code should teach, not just work
2. **Idiomatic Go**: Prefer Go idioms over direct pattern translations
3. **Practical Examples**: Show real-world applications
4. **Zero External Dependencies**: Use only the standard library
5. **Backward Compatibility**: Maintain semver; don't break existing APIs

## Questions?

Don't hesitate to ask questions by opening an issue labeled "question". We're here to help!

## Recognition

Contributors will be recognized in the README and release notes. Thank you for helping make GoLang-202 better!

