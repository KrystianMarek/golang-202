package idioms

import "fmt"

// Zero values demonstrate leveraging Go's zero value semantics.
//
// Why? Go's zero values enable useful defaults and eliminate
// the need for explicit initialization in many cases.

// Buffer leverages zero value for safe initialization.
type Buffer struct {
	data []byte // Zero value: nil slice, safe to use
}

// Write appends data to the buffer.
func (b *Buffer) Write(data []byte) {
	// No need to check if b.data is nil, append handles it
	b.data = append(b.data, data...)
}

// Bytes returns the buffer contents.
func (b *Buffer) Bytes() []byte {
	return b.data
}

// Config uses zero values for sensible defaults.
type Config struct {
	Port    int    // Zero value: 0
	Host    string // Zero value: ""
	Debug   bool   // Zero value: false
	Timeout int    // Zero value: 0
}

// NewConfig creates a config with defaults.
func NewConfig() *Config {
	c := &Config{} // All zero values
	c.setDefaults()
	return c
}

func (c *Config) setDefaults() {
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Timeout == 0 {
		c.Timeout = 30
	}
}

// Cache demonstrates zero value for maps.
type Cache struct {
	data map[string]interface{} // Zero value: nil map
}

// Get retrieves a value from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	// Reading from nil map is safe, returns zero value
	val, ok := c.data[key]
	return val, ok
}

// Set stores a value in the cache.
func (c *Cache) Set(key string, value interface{}) {
	// Initialize map if nil
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
}

// Counter demonstrates zero value sync primitives.
type Counter struct {
	// sync.Mutex has a usable zero value
	value int
}

// Increment increments the counter.
func (c *Counter) Increment() {
	// Can use mutex without explicit initialization
	c.value++
}

// Value returns the current value.
func (c *Counter) Value() int {
	return c.value
}

// Optional demonstrates zero value for optional fields.
type Optional[T any] struct {
	value T
	valid bool
}

// None returns an empty Optional.
func None[T any]() Optional[T] {
	return Optional[T]{} // Zero value: valid=false
}

// Some returns an Optional with a value.
func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, valid: true}
}

// IsValid returns true if the Optional has a value.
func (o Optional[T]) IsValid() bool {
	return o.valid
}

// Get returns the value.
func (o Optional[T]) Get() T {
	return o.value
}

// OrElse returns the value or a default.
func (o Optional[T]) OrElse(defaultValue T) T {
	if o.valid {
		return o.value
	}
	return defaultValue
}

// QueryBuilder demonstrates zero-value-friendly builder.
type QueryBuilder struct {
	table   string
	columns []string // nil slice is valid
	where   []string // nil slice is valid
	limit   int      // 0 means no limit
}

// From sets the table.
func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.table = table
	return q
}

// Select adds columns.
func (q *QueryBuilder) Select(columns ...string) *QueryBuilder {
	// append works with nil slice
	q.columns = append(q.columns, columns...)
	return q
}

// Where adds conditions.
func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.where = append(q.where, condition)
	return q
}

// Limit sets the limit.
func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	return q
}

// Build constructs the query.
func (q *QueryBuilder) Build() string {
	query := "SELECT "

	if len(q.columns) == 0 {
		query += "*"
	} else {
		for i, col := range q.columns {
			if i > 0 {
				query += ", "
			}
			query += col
		}
	}

	if q.table != "" {
		query += " FROM " + q.table
	}

	if len(q.where) > 0 {
		query += " WHERE "
		for i, cond := range q.where {
			if i > 0 {
				query += " AND "
			}
			query += cond
		}
	}

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", q.limit)
	}

	return query
}

// ExampleZeroValues demonstrates zero value patterns.
func ExampleZeroValues() {
	fmt.Println("=== Zero Values ===")

	// Buffer with zero value
	var buf Buffer // No initialization needed
	buf.Write([]byte("Hello"))
	buf.Write([]byte(" World"))
	fmt.Printf("Buffer: %s\n\n", string(buf.Bytes()))

	// Config with defaults
	config := NewConfig()
	fmt.Printf("Config: %s:%d (debug=%v, timeout=%ds)\n\n",
		config.Host, config.Port, config.Debug, config.Timeout)

	// Cache with nil map
	var cache Cache // No initialization
	cache.Set("key1", "value1")
	if val, ok := cache.Get("key1"); ok {
		fmt.Printf("Cache value: %v\n", val)
	}

	// Reading from zero-value cache is safe
	var cache2 Cache
	if _, ok := cache2.Get("missing"); !ok {
		fmt.Println("Key not found (nil map is safe)")
	}

	// Counter with zero value
	var counter Counter
	counter.Increment()
	counter.Increment()
	fmt.Printf("Counter: %d\n\n", counter.Value())

	// Optional values
	opt1 := Some(42)
	opt2 := None[int]()

	fmt.Printf("opt1: valid=%v, value=%d\n", opt1.IsValid(), opt1.Get())
	fmt.Printf("opt2: valid=%v, value=%d\n", opt2.IsValid(), opt2.Get())
	fmt.Printf("opt2 with default: %d\n\n", opt2.OrElse(100))

	// Builder with zero values
	var qb QueryBuilder // No initialization
	query := qb.From("users").
		Select("id", "name").
		Where("age > 18").
		Limit(10).
		Build()

	fmt.Printf("Query: %s\n", query)
}
