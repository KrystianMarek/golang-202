package go124

import (
	"fmt"
	"unique"
)

// StringCache demonstrates using unique.Handle for value canonicalization.
// This is useful for deduplicating immutable values in memory-intensive applications.
//
// Why? The unique package allows interning values, ensuring only one copy
// exists in memory. This is particularly useful for string-heavy workloads
// like log processing, configuration management, or compiler symbol tables.
type StringCache struct {
	handles []unique.Handle[string]
}

// NewStringCache creates a new cache for string interning.
func NewStringCache() *StringCache {
	return &StringCache{
		handles: make([]unique.Handle[string], 0),
	}
}

// Intern adds a string to the cache, returning its unique handle.
// Multiple calls with equal strings return the same underlying memory.
func (sc *StringCache) Intern(s string) unique.Handle[string] {
	handle := unique.Make(s)
	sc.handles = append(sc.handles, handle)
	return handle
}

// Get retrieves the original value from a handle.
func (sc *StringCache) Get(handle unique.Handle[string]) string {
	return handle.Value()
}

// LogEntry represents a structured log entry with interned strings
// to reduce memory footprint.
type LogEntry struct {
	Level   unique.Handle[string]
	Message unique.Handle[string]
	Source  unique.Handle[string]
}

// LogAggregator collects log entries with string interning.
type LogAggregator struct {
	entries []LogEntry
	cache   *StringCache
}

// NewLogAggregator creates a new log aggregator.
func NewLogAggregator() *LogAggregator {
	return &LogAggregator{
		entries: make([]LogEntry, 0),
		cache:   NewStringCache(),
	}
}

// AddLog adds a log entry with automatic string interning.
func (la *LogAggregator) AddLog(level, message, source string) {
	entry := LogEntry{
		Level:   unique.Make(level),
		Message: unique.Make(message),
		Source:  unique.Make(source),
	}
	la.entries = append(la.entries, entry)
}

// GetLogs returns all log entries as strings.
func (la *LogAggregator) GetLogs() []string {
	logs := make([]string, len(la.entries))
	for i, entry := range la.entries {
		logs[i] = fmt.Sprintf("[%s] %s (from: %s)",
			entry.Level.Value(),
			entry.Message.Value(),
			entry.Source.Value())
	}
	return logs
}

// ExampleUnique demonstrates unique.Handle for value canonicalization.
func ExampleUnique() {
	// String interning example
	cache := NewStringCache()

	h1 := cache.Intern("hello")
	h2 := cache.Intern("hello")
	h3 := cache.Intern("world")

	fmt.Printf("h1 == h2: %v (same string)\n", h1 == h2)
	fmt.Printf("h1 == h3: %v (different strings)\n", h1 == h3)
	fmt.Printf("Value: %s\n", cache.Get(h1))

	// Log aggregation with interning
	aggregator := NewLogAggregator()
	aggregator.AddLog("ERROR", "Connection failed", "db-service")
	aggregator.AddLog("ERROR", "Connection failed", "api-service")
	aggregator.AddLog("INFO", "Request processed", "api-service")

	fmt.Println("\nLog entries:")
	for _, log := range aggregator.GetLogs() {
		fmt.Println(log)
	}
}
