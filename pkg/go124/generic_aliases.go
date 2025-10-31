package go124

import "fmt"

// OrderedSlice is a parameterized type alias for slices of ordered types.
// Go 1.24 allows generic type aliases, enabling concise type definitions.
//
// Why? Generic type aliases reduce boilerplate and improve readability
// when working with complex generic types.
type OrderedSlice[T comparable] = []T

// KeyValue represents a generic key-value pair.
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

// KVMap is a generic type alias for maps.
type KVMap[K comparable, V any] = map[K]V

// Result represents a computation result or error.
type Result[T any] struct {
	Value T
	Error error
}

// ResultSlice is a type alias for slices of results.
type ResultSlice[T any] = []Result[T]

// Optional represents an optional value.
type Optional[T any] struct {
	value   T
	present bool
}

// Some creates an Optional with a value.
func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v, present: true}
}

// None creates an empty Optional.
func None[T any]() Optional[T] {
	return Optional[T]{present: false}
}

// IsPresent returns true if the Optional has a value.
func (o Optional[T]) IsPresent() bool {
	return o.present
}

// Get returns the value or zero value if not present.
func (o Optional[T]) Get() T {
	return o.value
}

// OptionalSlice is a type alias for slices of optionals.
type OptionalSlice[T any] = []Optional[T]

// Pair represents a generic pair of values.
type Pair[A, B any] struct {
	First  A
	Second B
}

// PairList is a type alias for lists of pairs.
type PairList[A, B any] = []Pair[A, B]

// NewPair creates a new pair.
func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

// Swap returns a new pair with swapped values.
func (p Pair[A, B]) Swap() Pair[B, A] {
	return Pair[B, A]{First: p.Second, Second: p.First}
}

// Transform applies functions to both elements.
func Transform[A, B, C, D any](
	p Pair[A, B],
	f func(A) C,
	g func(B) D,
) Pair[C, D] {
	return Pair[C, D]{
		First:  f(p.First),
		Second: g(p.Second),
	}
}

// ExampleGenericAliases demonstrates parameterized type aliases.
func ExampleGenericAliases() {
	// Using generic slice alias
	numbers := OrderedSlice[int]{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %v\n", numbers)

	// Using generic map alias
	scores := KVMap[string, int]{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	fmt.Printf("Scores: %v\n", scores)

	// Using Optional
	maybeValue := Some(42)
	noValue := None[int]()

	fmt.Printf("Maybe value present: %v, value: %d\n",
		maybeValue.IsPresent(), maybeValue.Get())
	fmt.Printf("No value present: %v\n", noValue.IsPresent())

	// Using Pair
	pair := NewPair("age", 25)
	fmt.Printf("Original pair: (%s, %d)\n", pair.First, pair.Second)

	swapped := pair.Swap()
	fmt.Printf("Swapped pair: (%d, %s)\n", swapped.First, swapped.Second)

	// Transform pair
	transformed := Transform(
		pair,
		func(s string) int { return len(s) },
		func(n int) string { return fmt.Sprintf("%d years", n) },
	)
	fmt.Printf("Transformed: (%d, %s)\n",
		transformed.First, transformed.Second)

	// Result slice
	results := ResultSlice[string]{
		{Value: "success", Error: nil},
		{Value: "", Error: fmt.Errorf("failed")},
	}

	for i, r := range results {
		if r.Error != nil {
			fmt.Printf("Result %d: error - %v\n", i, r.Error)
		} else {
			fmt.Printf("Result %d: %s\n", i, r.Value)
		}
	}
}

