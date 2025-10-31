package patterns

import "fmt"

// Notification is an interface for different notification types.
// This demonstrates the Factory Method pattern.
//
// Why? Factory functions return interfaces, allowing runtime
// selection of concrete types while hiding implementation details.
type Notification interface {
	Send(message string) error
	GetType() string
}

// EmailNotification sends notifications via email.
type EmailNotification struct {
	To string
}

// Send sends an email notification.
func (e *EmailNotification) Send(message string) error {
	fmt.Printf("[EMAIL to %s] %s\n", e.To, message)
	return nil
}

// GetType returns the notification type.
func (e *EmailNotification) GetType() string {
	return "email"
}

// SMSNotification sends notifications via SMS.
type SMSNotification struct {
	PhoneNumber string
}

// Send sends an SMS notification.
func (s *SMSNotification) Send(message string) error {
	fmt.Printf("[SMS to %s] %s\n", s.PhoneNumber, message)
	return nil
}

// GetType returns the notification type.
func (s *SMSNotification) GetType() string {
	return "sms"
}

// PushNotification sends push notifications.
type PushNotification struct {
	DeviceID string
}

// Send sends a push notification.
func (p *PushNotification) Send(message string) error {
	fmt.Printf("[PUSH to %s] %s\n", p.DeviceID, message)
	return nil
}

// GetType returns the notification type.
func (p *PushNotification) GetType() string {
	return "push"
}

// NewNotification is a factory function that creates notifications.
func NewNotification(notifType, target string) Notification {
	switch notifType {
	case "email":
		return &EmailNotification{To: target}
	case "sms":
		return &SMSNotification{PhoneNumber: target}
	case "push":
		return &PushNotification{DeviceID: target}
	default:
		return &EmailNotification{To: target}
	}
}

// Document interface for different document types.
type Document interface {
	Open() string
	Save(content string) error
	GetFormat() string
}

// PDFDocument represents a PDF document.
type PDFDocument struct {
	Filename string
}

// Open opens the PDF.
func (p *PDFDocument) Open() string {
	return fmt.Sprintf("Opening PDF: %s", p.Filename)
}

// Save saves the PDF.
func (p *PDFDocument) Save(content string) error {
	fmt.Printf("Saving PDF %s: %s\n", p.Filename, content)
	return nil
}

// GetFormat returns the format.
func (p *PDFDocument) GetFormat() string {
	return "PDF"
}

// WordDocument represents a Word document.
type WordDocument struct {
	Filename string
}

// Open opens the Word document.
func (w *WordDocument) Open() string {
	return fmt.Sprintf("Opening Word: %s", w.Filename)
}

// Save saves the Word document.
func (w *WordDocument) Save(content string) error {
	fmt.Printf("Saving Word %s: %s\n", w.Filename, content)
	return nil
}

// GetFormat returns the format.
func (w *WordDocument) GetFormat() string {
	return "DOCX"
}

// DocumentFactory creates documents.
type DocumentFactory struct{}

// CreateDocument is a factory method.
func (f *DocumentFactory) CreateDocument(format, filename string) Document {
	switch format {
	case "pdf":
		return &PDFDocument{Filename: filename}
	case "docx":
		return &WordDocument{Filename: filename}
	default:
		return &PDFDocument{Filename: filename}
	}
}

// Transport interface for different shipping methods.
type Transport interface {
	Deliver(destination string) string
}

// Truck represents truck transport.
type Truck struct{}

// Deliver delivers by truck.
func (t *Truck) Deliver(destination string) string {
	return fmt.Sprintf("Delivering to %s by truck", destination)
}

// Ship represents ship transport.
type Ship struct{}

// Deliver delivers by ship.
func (s *Ship) Deliver(destination string) string {
	return fmt.Sprintf("Delivering to %s by ship", destination)
}

// Logistics is an abstract factory creator.
type Logistics interface {
	CreateTransport() Transport
	Plan(destination string) string
}

// RoadLogistics creates truck transport.
type RoadLogistics struct{}

// CreateTransport creates a truck.
func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// Plan plans road delivery.
func (r *RoadLogistics) Plan(destination string) string {
	transport := r.CreateTransport()
	return transport.Deliver(destination)
}

// SeaLogistics creates ship transport.
type SeaLogistics struct{}

// CreateTransport creates a ship.
func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

// Plan plans sea delivery.
func (s *SeaLogistics) Plan(destination string) string {
	transport := s.CreateTransport()
	return transport.Deliver(destination)
}

// ExampleFactory demonstrates Factory patterns.
func ExampleFactory() {
	fmt.Println("=== Factory Pattern ===")

	// Simple factory function
	email := NewNotification("email", "user@example.com")
	sms := NewNotification("sms", "+1234567890")
	push := NewNotification("push", "device-123")

	notifications := []Notification{email, sms, push}
	for _, n := range notifications {
		_ = n.Send(fmt.Sprintf("Hello from %s!", n.GetType()))
	}

	// Factory method pattern
	docFactory := &DocumentFactory{}
	pdf := docFactory.CreateDocument("pdf", "report.pdf")
	word := docFactory.CreateDocument("docx", "letter.docx")

	fmt.Println(pdf.Open())
	_ = pdf.Save("PDF content")

	fmt.Println(word.Open())
	_ = word.Save("Word content")

	// Abstract factory pattern
	roadLogistics := &RoadLogistics{}
	seaLogistics := &SeaLogistics{}

	fmt.Println(roadLogistics.Plan("New York"))
	fmt.Println(seaLogistics.Plan("London"))
}
