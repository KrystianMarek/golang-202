package idioms

import (
	"errors"
	"fmt"
)

// Error handling demonstrates Go's error patterns.
//
// Why? Go's explicit error handling makes error paths visible
// and forces developers to handle errors at each call site.

// Sentinel errors for comparison with errors.Is.
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

// ValidationError is a custom error type.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface.
func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", v.Field, v.Message)
}

// DatabaseError wraps database-related errors.
type DatabaseError struct {
	Query string
	Err   error
}

// Error implements the error interface.
func (d *DatabaseError) Error() string {
	return fmt.Sprintf("database error on query '%s': %v", d.Query, d.Err)
}

// Unwrap enables errors.Is and errors.As.
func (d *DatabaseError) Unwrap() error {
	return d.Err
}

// UserService demonstrates error handling patterns.
type UserService struct {
	users map[string]string
}

// NewUserService creates a user service.
func NewUserService() *UserService {
	return &UserService{
		users: map[string]string{
			"alice": "alice@example.com",
		},
	}
}

// GetUser retrieves a user or returns an error.
func (s *UserService) GetUser(username string) (string, error) {
	email, ok := s.users[username]
	if !ok {
		return "", fmt.Errorf("user %s: %w", username, ErrNotFound)
	}
	return email, nil
}

// CreateUser creates a user with validation.
func (s *UserService) CreateUser(username, email string) error {
	if username == "" {
		return &ValidationError{Field: "username", Message: "cannot be empty"}
	}
	if email == "" {
		return &ValidationError{Field: "email", Message: "cannot be empty"}
	}

	if _, exists := s.users[username]; exists {
		return fmt.Errorf("username %s already exists", username)
	}

	s.users[username] = email
	return nil
}

// MultiError holds multiple errors.
type MultiError struct {
	Errors []error
}

// Error implements the error interface.
func (m *MultiError) Error() string {
	return fmt.Sprintf("multiple errors: %d error(s) occurred", len(m.Errors))
}

// Add adds an error to the collection.
func (m *MultiError) Add(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

// HasErrors returns true if there are any errors.
func (m *MultiError) HasErrors() bool {
	return len(m.Errors) > 0
}

// Unwrap returns the wrapped errors (for Go 1.20+ errors.Join compatibility).
func (m *MultiError) Unwrap() []error {
	return m.Errors
}

// ProcessBatch demonstrates collecting multiple errors.
func ProcessBatch(items []string) error {
	var multiErr MultiError

	for _, item := range items {
		if err := processItem(item); err != nil {
			multiErr.Add(fmt.Errorf("item %s: %w", item, err))
		}
	}

	if multiErr.HasErrors() {
		return &multiErr
	}
	return nil
}

func processItem(item string) error {
	if item == "" {
		return ErrInvalidInput
	}
	if item == "forbidden" {
		return ErrUnauthorized
	}
	return nil
}

// Result represents a result or error (alternative to multiple returns).
type Result[T any] struct {
	Value T
	Err   error
}

// Ok creates a successful result.
func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value}
}

// Err creates an error result.
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{Value: zero, Err: err}
}

// IsOk returns true if the result is successful.
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

// Unwrap returns the value and error.
func (r Result[T]) Unwrap() (T, error) {
	return r.Value, r.Err
}

// Divide returns a result instead of value and error.
func Divide(a, b float64) Result[float64] {
	if b == 0 {
		return Err[float64](errors.New("division by zero"))
	}
	return Ok(a / b)
}

// Must panics if there's an error (use sparingly, e.g., in init).
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// WrapError demonstrates error wrapping for context.
func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s failed: %w", operation, err)
}

// ExampleErrors demonstrates error handling patterns.
func ExampleErrors() {
	fmt.Println("=== Error Handling ===")

	service := NewUserService()

	// Basic error handling
	email, err := service.GetUser("alice")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found: %s\n", email)
	}

	// Sentinel error with errors.Is
	_, err = service.GetUser("bob")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found (detected with errors.Is)")
	}


	// Custom error with errors.As
	err = service.CreateUser("", "test@example.com")
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("Validation error: field=%s, msg=%s\n",
			validationErr.Field, validationErr.Message)
	}


	// Error wrapping
	dbErr := &DatabaseError{
		Query: "SELECT * FROM users",
		Err:   ErrNotFound,
	}
	fmt.Printf("Database error: %v\n", dbErr)
	if errors.Is(dbErr, ErrNotFound) {
		fmt.Println("Underlying error is ErrNotFound")
	}


	// Multiple errors
	items := []string{"valid", "", "forbidden", "ok"}
	if err := ProcessBatch(items); err != nil {
		fmt.Printf("Batch processing error: %v\n", err)

		var multiErr *MultiError
		if errors.As(err, &multiErr) {
			fmt.Printf("Individual errors (%d):\n", len(multiErr.Errors))
			for i, e := range multiErr.Errors {
				fmt.Printf("  %d: %v\n", i+1, e)
			}
		}
	}


	// Result type
	result := Divide(10, 2)
	if result.IsOk() {
		fmt.Printf("Division result: %.2f\n", result.Value)
	}

	result2 := Divide(10, 0)
	if !result2.IsOk() {
		fmt.Printf("Division error: %v\n", result2.Err)
	}


	// Error wrapping for context
	originalErr := errors.New("connection timeout")
	wrappedErr := WrapError("database query", originalErr)
	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// errors.Join (Go 1.20+)
	joinedErr := errors.Join(
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	)
	fmt.Printf("Joined errors: %v\n", joinedErr)
}

