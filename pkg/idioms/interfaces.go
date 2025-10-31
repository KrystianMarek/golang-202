// Package idioms demonstrates Go-specific patterns and best practices.
package idioms

import (
	"fmt"
	"io"
	"strings"
)

// Interfaces demonstrate duck typing and interface composition.
//
// Why? Go's implicit interface satisfaction enables loose coupling
// and makes code more testable and maintainable.

// Reader is a minimal interface (interface segregation).
type Reader interface {
	Read() (string, error)
}

// Writer is a minimal interface.
type Writer interface {
	Write(data string) error
}

// ReadWriter composes Reader and Writer.
type ReadWriter interface {
	Reader
	Writer
}

// StringReader implements Reader for strings.
type StringReader struct {
	data string
	pos  int
}

// NewStringReader creates a string reader.
func NewStringReader(data string) *StringReader {
	return &StringReader{data: data, pos: 0}
}

// Read reads from the string.
func (s *StringReader) Read() (string, error) {
	if s.pos >= len(s.data) {
		return "", io.EOF
	}
	chunk := s.data[s.pos:]
	s.pos = len(s.data)
	return chunk, nil
}

// BufferWriter implements Writer with a buffer.
type BufferWriter struct {
	buffer strings.Builder
}

// NewBufferWriter creates a buffer writer.
func NewBufferWriter() *BufferWriter {
	return &BufferWriter{}
}

// Write writes to the buffer.
func (b *BufferWriter) Write(data string) error {
	b.buffer.WriteString(data)
	return nil
}

// String returns the buffered content.
func (b *BufferWriter) String() string {
	return b.buffer.String()
}

// Copy copies from Reader to Writer.
func Copy(r Reader, w Writer) error {
	data, err := r.Read()
	if err != nil && err != io.EOF {
		return err
	}
	return w.Write(data)
}

// Closer represents resources that need cleanup.
type Closer interface {
	Close() error
}

// File represents a file with multiple interfaces.
type File struct {
	name    string
	content string
	pos     int
	closed  bool
}

// NewFile creates a new file.
func NewFile(name, content string) *File {
	return &File{name: name, content: content, pos: 0, closed: false}
}

// Read implements Reader.
func (f *File) Read() (string, error) {
	if f.closed {
		return "", fmt.Errorf("file closed")
	}
	if f.pos >= len(f.content) {
		return "", io.EOF
	}
	chunk := f.content[f.pos:]
	f.pos = len(f.content)
	return chunk, nil
}

// Write implements Writer.
func (f *File) Write(data string) error {
	if f.closed {
		return fmt.Errorf("file closed")
	}
	f.content += data
	return nil
}

// Close implements Closer.
func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("already closed")
	}
	f.closed = true
	fmt.Printf("File %s closed\n", f.name)
	return nil
}

// Accept any type (empty interface pattern - pre-generics).
type Container struct {
	items []interface{}
}

// NewContainer creates a container.
func NewContainer() *Container {
	return &Container{items: make([]interface{}, 0)}
}

// Add adds an item.
func (c *Container) Add(item interface{}) {
	c.items = append(c.items, item)
}

// Get returns an item with type assertion.
func (c *Container) Get(index int) (interface{}, bool) {
	if index < 0 || index >= len(c.items) {
		return nil, false
	}
	return c.items[index], true
}

// GetAs returns an item with type assertion.
func (c *Container) GetAs(index int, target interface{}) bool {
	item, ok := c.Get(index)
	if !ok {
		return false
	}
	// Type assertion would happen at call site
	_ = target
	_ = item
	return true
}

// Stringer demonstrates implementing standard interfaces.
type Person struct {
	Name string
	Age  int
}

// String implements fmt.Stringer.
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// Validator is a common validation interface.
type Validator interface {
	Validate() error
}

// Email implements Validator.
type Email struct {
	Address string
}

// Validate validates the email.
func (e Email) Validate() error {
	if !strings.Contains(e.Address, "@") {
		return fmt.Errorf("invalid email: %s", e.Address)
	}
	return nil
}

// ValidateAll validates multiple validators.
func ValidateAll(validators ...Validator) error {
	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Processor demonstrates interface for dependency injection.
type Processor interface {
	Process(data string) string
}

// UpperCaseProcessor converts to uppercase.
type UpperCaseProcessor struct{}

// Process converts to uppercase.
func (u UpperCaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}

// TrimProcessor trims whitespace.
type TrimProcessor struct{}

// Process trims whitespace.
func (t TrimProcessor) Process(data string) string {
	return strings.TrimSpace(data)
}

// ProcessorChain chains multiple processors.
type ProcessorChain struct {
	processors []Processor
}

// NewProcessorChain creates a processor chain.
func NewProcessorChain(processors ...Processor) *ProcessorChain {
	return &ProcessorChain{processors: processors}
}

// Process processes data through all processors.
func (pc *ProcessorChain) Process(data string) string {
	result := data
	for _, p := range pc.processors {
		result = p.Process(result)
	}
	return result
}

// ExampleInterfaces demonstrates Go interface patterns.
func ExampleInterfaces() {
	fmt.Println("=== Go Interfaces ===")

	// Basic interfaces
	reader := NewStringReader("Hello, World!")
	writer := NewBufferWriter()

	err := Copy(reader, writer)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
	}
	fmt.Printf("Copied: %s\n\n", writer.String())

	// Multiple interfaces
	file := NewFile("test.txt", "initial content")

	data, _ := file.Read()
	fmt.Printf("Read from file: %s\n", data)

	file.Write(" appended")
	file.Close()


	// Empty interface (pre-generics pattern)
	container := NewContainer()
	container.Add(42)
	container.Add("hello")
	container.Add(Person{Name: "Alice", Age: 30})

	if item, ok := container.Get(2); ok {
		if person, ok := item.(Person); ok {
			fmt.Printf("Person: %s\n", person)
		}
	}


	// Validator interface
	email := Email{Address: "user@example.com"}
	if err := email.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Email is valid")
	}


	// Processor chain (dependency injection)
	chain := NewProcessorChain(
		TrimProcessor{},
		UpperCaseProcessor{},
	)

	result := chain.Process("  hello world  ")
	fmt.Printf("Processed: '%s'\n", result)
}

