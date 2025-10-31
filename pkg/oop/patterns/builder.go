package patterns

import "fmt"

// Builder pattern demonstrates fluent interface for constructing complex objects.
//
// Why? Builders provide a clean way to construct objects with many optional
// parameters, avoiding telescoping constructors and improving readability.

// HTTPRequest represents an HTTP request built with the builder pattern.
type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
	Timeout int
}

// RequestBuilder builds HTTP requests fluently.
type RequestBuilder struct {
	request HTTPRequest
}

// NewRequestBuilder creates a new builder.
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: HTTPRequest{
			Headers: make(map[string]string),
			Timeout: 30,
		},
	}
}

// Method sets the HTTP method.
func (b *RequestBuilder) Method(method string) *RequestBuilder {
	b.request.Method = method
	return b
}

// URL sets the URL.
func (b *RequestBuilder) URL(url string) *RequestBuilder {
	b.request.URL = url
	return b
}

// Header adds a header.
func (b *RequestBuilder) Header(key, value string) *RequestBuilder {
	b.request.Headers[key] = value
	return b
}

// Body sets the request body.
func (b *RequestBuilder) Body(body string) *RequestBuilder {
	b.request.Body = body
	return b
}

// Timeout sets the timeout.
func (b *RequestBuilder) Timeout(seconds int) *RequestBuilder {
	b.request.Timeout = seconds
	return b
}

// Build returns the constructed request.
func (b *RequestBuilder) Build() HTTPRequest {
	return b.request
}

// EmailMessage represents an email.
type EmailMessage struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments []string
	Priority    int
}

// EmailBuilder builds emails fluently.
type EmailBuilder struct {
	email EmailMessage
}

// NewEmailBuilder creates a new email builder.
func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{
		email: EmailMessage{
			To:          make([]string, 0),
			CC:          make([]string, 0),
			BCC:         make([]string, 0),
			Attachments: make([]string, 0),
			Priority:    1,
		},
	}
}

// From sets the sender.
func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.From = from
	return b
}

// To adds a recipient.
func (b *EmailBuilder) To(to ...string) *EmailBuilder {
	b.email.To = append(b.email.To, to...)
	return b
}

// CC adds CC recipients.
func (b *EmailBuilder) CC(cc ...string) *EmailBuilder {
	b.email.CC = append(b.email.CC, cc...)
	return b
}

// Subject sets the subject.
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.Subject = subject
	return b
}

// Body sets the body.
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.Body = body
	return b
}

// Attachment adds an attachment.
func (b *EmailBuilder) Attachment(path string) *EmailBuilder {
	b.email.Attachments = append(b.email.Attachments, path)
	return b
}

// Priority sets the priority.
func (b *EmailBuilder) Priority(priority int) *EmailBuilder {
	b.email.Priority = priority
	return b
}

// Build returns the constructed email.
func (b *EmailBuilder) Build() EmailMessage {
	return b.email
}

// Send simulates sending the email.
func (e *EmailMessage) Send() {
	fmt.Printf("Sending email from %s to %v\n", e.From, e.To)
	fmt.Printf("Subject: %s\n", e.Subject)
	fmt.Printf("Body: %s\n", e.Body)
	if len(e.Attachments) > 0 {
		fmt.Printf("Attachments: %v\n", e.Attachments)
	}
}

// QueryBuilder builds SQL queries (simplified).
type QueryBuilder struct {
	table      string
	columns    []string
	conditions []string
	orderBy    string
	limit      int
}

// NewQueryBuilder creates a query builder.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		columns:    make([]string, 0),
		conditions: make([]string, 0),
	}
}

// Select sets the columns to select.
func (b *QueryBuilder) Select(columns ...string) *QueryBuilder {
	b.columns = append(b.columns, columns...)
	return b
}

// From sets the table.
func (b *QueryBuilder) From(table string) *QueryBuilder {
	b.table = table
	return b
}

// Where adds a condition.
func (b *QueryBuilder) Where(condition string) *QueryBuilder {
	b.conditions = append(b.conditions, condition)
	return b
}

// OrderBy sets the ordering.
func (b *QueryBuilder) OrderBy(column string) *QueryBuilder {
	b.orderBy = column
	return b
}

// Limit sets the limit.
func (b *QueryBuilder) Limit(limit int) *QueryBuilder {
	b.limit = limit
	return b
}

// Build constructs the SQL string.
func (b *QueryBuilder) Build() string {
	query := "SELECT "

	if len(b.columns) == 0 {
		query += "*"
	} else {
		for i, col := range b.columns {
			if i > 0 {
				query += ", "
			}
			query += col
		}
	}

	query += fmt.Sprintf(" FROM %s", b.table)

	if len(b.conditions) > 0 {
		query += " WHERE "
		for i, cond := range b.conditions {
			if i > 0 {
				query += " AND "
			}
			query += cond
		}
	}

	if b.orderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", b.orderBy)
	}

	if b.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", b.limit)
	}

	return query
}

// ExampleBuilder demonstrates the Builder pattern.
func ExampleBuilder() {
	fmt.Println("=== Builder Pattern ===")

	// HTTP Request builder
	req := NewRequestBuilder().
		Method("POST").
		URL("https://api.example.com/users").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"name": "Alice"}`).
		Timeout(60).
		Build()

	fmt.Printf("HTTP Request: %s %s\n", req.Method, req.URL)
	fmt.Printf("Headers: %v\n", req.Headers)
	fmt.Printf("Timeout: %ds\n\n", req.Timeout)

	// Email builder
	email := NewEmailBuilder().
		From("sender@example.com").
		To("recipient1@example.com", "recipient2@example.com").
		CC("manager@example.com").
		Subject("Monthly Report").
		Body("Please find the monthly report attached.").
		Attachment("/reports/monthly.pdf").
		Priority(2).
		Build()

	email.Send()

	// Query builder
	query := NewQueryBuilder().
		Select("id", "name", "email").
		From("users").
		Where("age > 18").
		Where("status = 'active'").
		OrderBy("name").
		Limit(10).
		Build()

	fmt.Printf("SQL Query: %s\n", query)
}
