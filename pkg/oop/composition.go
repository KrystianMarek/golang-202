// Package oop demonstrates object-oriented programming patterns in Go
// using composition, interfaces, and struct embedding.
package oop

import "fmt"

// Base represents a base type with common functionality.
// This demonstrates struct embedding for composition-based inheritance.
//
// Why? Go doesn't have classical inheritance, but embedding provides
// similar benefits: code reuse and polymorphism through interfaces.
type Base struct {
	ID   string
	Name string
}

// GetID returns the identifier.
func (b *Base) GetID() string {
	return b.ID
}

// GetName returns the name.
func (b *Base) GetName() string {
	return b.Name
}

// Describe provides a base description.
func (b *Base) Describe() string {
	return fmt.Sprintf("Base[%s]: %s", b.ID, b.Name)
}

// Extended embeds Base and adds additional functionality.
type Extended struct {
	Base // Embedding promotes Base's fields and methods
	Extra string
}

// Describe overrides the base method.
func (e *Extended) Describe() string {
	return fmt.Sprintf("Extended[%s]: %s (extra: %s)", e.ID, e.Name, e.Extra)
}

// Shape is an interface for geometric shapes.
// This demonstrates polymorphism through interfaces.
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Circle implements Shape.
type Circle struct {
	Radius float64
}

// Area calculates the circle's area.
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the circle's perimeter.
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// Name returns the shape name.
func (c Circle) Name() string {
	return "Circle"
}

// Rectangle implements Shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the rectangle's area.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the rectangle's perimeter.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Name returns the shape name.
func (r Rectangle) Name() string {
	return "Rectangle"
}

// PrintShapeInfo demonstrates interface polymorphism.
func PrintShapeInfo(s Shape) {
	fmt.Printf("%s: Area=%.2f, Perimeter=%.2f\n",
		s.Name(), s.Area(), s.Perimeter())
}

// Logger interface for different logging strategies.
type Logger interface {
	Log(message string)
}

// ConsoleLogger logs to console.
type ConsoleLogger struct{}

// Log prints to console.
func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[CONSOLE] %s\n", message)
}

// FileLogger simulates file logging.
type FileLogger struct {
	Filename string
}

// Log simulates writing to a file.
func (fl FileLogger) Log(message string) {
	fmt.Printf("[FILE:%s] %s\n", fl.Filename, message)
}

// Service demonstrates dependency injection via interfaces.
type Service struct {
	logger Logger
}

// NewService creates a service with injected logger.
func NewService(logger Logger) *Service {
	return &Service{logger: logger}
}

// DoWork performs work and logs it.
func (s *Service) DoWork(task string) {
	s.logger.Log(fmt.Sprintf("Starting task: %s", task))
	// Simulate work
	s.logger.Log(fmt.Sprintf("Completed task: %s", task))
}

// Engine component.
type Engine struct {
	Horsepower int
}

// Start starts the engine.
func (e *Engine) Start() string {
	return fmt.Sprintf("Engine started (%d HP)", e.Horsepower)
}

// Wheels component.
type Wheels struct {
	Count int
}

// Roll makes the wheels roll.
func (w *Wheels) Roll() string {
	return fmt.Sprintf("%d wheels rolling", w.Count)
}

// Car demonstrates composition of multiple components.
type Car struct {
	Engine *Engine
	Wheels *Wheels
	Model  string
}

// NewCar creates a car with components.
func NewCar(model string, hp int, wheelCount int) *Car {
	return &Car{
		Engine: &Engine{Horsepower: hp},
		Wheels: &Wheels{Count: wheelCount},
		Model:  model,
	}
}

// Drive uses composed components.
func (c *Car) Drive() string {
	return fmt.Sprintf("%s: %s, %s",
		c.Model, c.Engine.Start(), c.Wheels.Roll())
}

// ExampleComposition demonstrates composition patterns.
func ExampleComposition() {
	// Struct embedding
	base := Base{ID: "001", Name: "BaseObject"}
	fmt.Println(base.Describe())

	extended := Extended{
		Base:  Base{ID: "002", Name: "ExtendedObject"},
		Extra: "additional data",
	}
	fmt.Println(extended.Describe())
	fmt.Printf("Embedded field access: %s\n", extended.GetID())

	// Interface polymorphism
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	fmt.Println("\nShapes:")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}

	// Dependency injection
	fmt.Println("\nDependency Injection:")
	consoleService := NewService(ConsoleLogger{})
	consoleService.DoWork("process data")

	fileService := NewService(FileLogger{Filename: "app.log"})
	fileService.DoWork("save records")

	// Component composition
	fmt.Println("\nComponent Composition:")
	car := NewCar("Tesla Model 3", 283, 4)
	fmt.Println(car.Drive())
}

