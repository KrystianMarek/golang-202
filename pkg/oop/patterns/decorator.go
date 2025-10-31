package patterns

import "fmt"

// Decorator pattern demonstrates adding behavior to objects dynamically.
//
// Why? Decorators provide flexible alternatives to subclassing, allowing
// behavior to be added at runtime through composition.

// Coffee is the base interface.
type Coffee interface {
	Cost() float64
	Description() string
}

// SimpleCoffee is the base implementation.
type SimpleCoffee struct{}

// Cost returns the base cost.
func (s *SimpleCoffee) Cost() float64 {
	return 2.00
}

// Description returns the description.
func (s *SimpleCoffee) Description() string {
	return "Simple coffee"
}

// MilkDecorator adds milk to coffee.
type MilkDecorator struct {
	coffee Coffee
}

// Cost adds milk cost.
func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.50
}

// Description adds milk description.
func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", milk"
}

// SugarDecorator adds sugar to coffee.
type SugarDecorator struct {
	coffee Coffee
}

// Cost adds sugar cost.
func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.25
}

// Description adds sugar description.
func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", sugar"
}

// WhipDecorator adds whipped cream.
type WhipDecorator struct {
	coffee Coffee
}

// Cost adds whip cost.
func (w *WhipDecorator) Cost() float64 {
	return w.coffee.Cost() + 0.75
}

// Description adds whip description.
func (w *WhipDecorator) Description() string {
	return w.coffee.Description() + ", whipped cream"
}

// DataSource is an interface for reading/writing data.
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// FileDataSource writes to a file.
type FileDataSource struct {
	filename string
	data     string
}

// NewFileDataSource creates a file data source.
func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

// WriteData writes data.
func (f *FileDataSource) WriteData(data string) {
	f.data = data
	fmt.Printf("Writing to file %s: %s\n", f.filename, data)
}

// ReadData reads data.
func (f *FileDataSource) ReadData() string {
	fmt.Printf("Reading from file %s\n", f.filename)
	return f.data
}

// EncryptionDecorator adds encryption.
type EncryptionDecorator struct {
	wrapped DataSource
}

// NewEncryptionDecorator creates an encryption decorator.
func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{wrapped: source}
}

// WriteData encrypts and writes.
func (e *EncryptionDecorator) WriteData(data string) {
	encrypted := fmt.Sprintf("ENCRYPTED(%s)", data)
	e.wrapped.WriteData(encrypted)
}

// ReadData reads and decrypts.
func (e *EncryptionDecorator) ReadData() string {
	data := e.wrapped.ReadData()
	// Simulate decryption
	if len(data) > 10 {
		return data[10 : len(data)-1]
	}
	return data
}

// CompressionDecorator adds compression.
type CompressionDecorator struct {
	wrapped DataSource
}

// NewCompressionDecorator creates a compression decorator.
func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{wrapped: source}
}

// WriteData compresses and writes.
func (c *CompressionDecorator) WriteData(data string) {
	compressed := fmt.Sprintf("COMPRESSED(%s)", data)
	c.wrapped.WriteData(compressed)
}

// ReadData reads and decompresses.
func (c *CompressionDecorator) ReadData() string {
	data := c.wrapped.ReadData()
	// Simulate decompression
	if len(data) > 11 {
		return data[11 : len(data)-1]
	}
	return data
}

// Notifier sends notifications.
type Notifier interface {
	Send(message string)
}

// BaseNotifier is the base implementation.
type BaseNotifier struct{}

// Send sends a basic notification.
func (b *BaseNotifier) Send(message string) {
	fmt.Printf("[BASE] %s\n", message)
}

// SMSDecorator adds SMS notifications.
type SMSDecorator struct {
	wrapped Notifier
}

// Send sends via wrapped and SMS.
func (s *SMSDecorator) Send(message string) {
	s.wrapped.Send(message)
	fmt.Printf("[SMS] %s\n", message)
}

// SlackDecorator adds Slack notifications.
type SlackDecorator struct {
	wrapped Notifier
}

// Send sends via wrapped and Slack.
func (s *SlackDecorator) Send(message string) {
	s.wrapped.Send(message)
	fmt.Printf("[SLACK] %s\n", message)
}

// ExampleDecorator demonstrates the Decorator pattern.
func ExampleDecorator() {
	fmt.Println("=== Decorator Pattern ===")

	// Coffee decorators
	coffee := &SimpleCoffee{}
	fmt.Printf("%s: $%.2f\n", coffee.Description(), coffee.Cost())

	coffeeWithMilk := &MilkDecorator{coffee: coffee}
	fmt.Printf("%s: $%.2f\n",
		coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	fancyCoffee := &WhipDecorator{
		coffee: &SugarDecorator{
			coffee: &MilkDecorator{
				coffee: coffee,
			},
		},
	}
	fmt.Printf("%s: $%.2f\n\n",
		fancyCoffee.Description(), fancyCoffee.Cost())

	// Data source decorators
	fileSource := NewFileDataSource("data.txt")
	encryptedSource := NewEncryptionDecorator(fileSource)
	compressedEncrypted := NewCompressionDecorator(encryptedSource)

	compressedEncrypted.WriteData("sensitive data")
	readData := compressedEncrypted.ReadData()
	fmt.Printf("Read: %s\n\n", readData)

	// Notification decorators
	notifier := Notifier(&BaseNotifier{})
	notifier = &SMSDecorator{wrapped: notifier}
	notifier = &SlackDecorator{wrapped: notifier}

	notifier.Send("Server alert: High CPU usage!")
}
