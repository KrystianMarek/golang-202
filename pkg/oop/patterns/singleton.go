// Package patterns implements Gang of Four design patterns using Go idioms.
package patterns

import (
	"fmt"
	"sync"
)

// Config represents a global configuration singleton.
// This demonstrates the Singleton pattern using sync.Once.
//
// Why? Singletons ensure only one instance exists globally.
// Go's sync.Once provides thread-safe initialization.
type Config struct {
	AppName  string
	Version  string
	Settings map[string]string
}

var (
	configInstance *Config
	configOnce     sync.Once
)

// GetConfig returns the singleton Config instance.
// Thread-safe initialization guaranteed by sync.Once.
func GetConfig() *Config {
	configOnce.Do(func() {
		configInstance = &Config{
			AppName: "MyApp",
			Version: "1.0.0",
			Settings: map[string]string{
				"debug": "false",
				"port":  "8080",
			},
		}
		fmt.Println("Config instance created")
	})
	return configInstance
}

// Database represents a database connection pool singleton.
type Database struct {
	ConnectionString string
	MaxConnections   int
	mu               sync.RWMutex
	connections      int
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

// GetDatabase returns the singleton Database instance.
func GetDatabase() *Database {
	dbOnce.Do(func() {
		dbInstance = &Database{
			ConnectionString: "postgres://localhost:5432/mydb",
			MaxConnections:   10,
			connections:      0,
		}
		fmt.Println("Database instance created")
	})
	return dbInstance
}

// Connect simulates acquiring a connection.
func (db *Database) Connect() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.connections >= db.MaxConnections {
		return fmt.Errorf("max connections reached")
	}

	db.connections++
	fmt.Printf("Connected (active: %d/%d)\n",
		db.connections, db.MaxConnections)
	return nil
}

// Disconnect simulates releasing a connection.
func (db *Database) Disconnect() {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.connections > 0 {
		db.connections--
		fmt.Printf("Disconnected (active: %d/%d)\n",
			db.connections, db.MaxConnections)
	}
}

// AppLogger singleton with lazy initialization.
type AppLogger struct {
	mu     sync.Mutex
	logs   []string
	prefix string
}

var (
	appLoggerInstance *AppLogger
	appLoggerOnce     sync.Once
)

// GetAppLogger returns the singleton AppLogger instance.
func GetAppLogger() *AppLogger {
	appLoggerOnce.Do(func() {
		appLoggerInstance = &AppLogger{
			logs:   make([]string, 0),
			prefix: "[APP]",
		}
		fmt.Println("AppLogger instance created")
	})
	return appLoggerInstance
}

// Log adds a log entry.
func (l *AppLogger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := fmt.Sprintf("%s %s", l.prefix, message)
	l.logs = append(l.logs, entry)
	fmt.Println(entry)
}

// GetLogs returns all log entries.
func (l *AppLogger) GetLogs() []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Return a copy to prevent external modification
	logs := make([]string, len(l.logs))
	copy(logs, l.logs)
	return logs
}

// ExampleSingleton demonstrates the Singleton pattern.
func ExampleSingleton() {
	fmt.Println("=== Singleton Pattern ===")

	// Config singleton
	config1 := GetConfig()
	config2 := GetConfig()

	fmt.Printf("Same instance: %v\n", config1 == config2)
	fmt.Printf("Config: %s v%s\n\n", config1.AppName, config1.Version)

	// Database singleton
	db1 := GetDatabase()
	db2 := GetDatabase()

	fmt.Printf("Same DB instance: %v\n", db1 == db2)
	_ = db1.Connect()
	_ = db2.Connect()
	db1.Disconnect()

	// AppLogger singleton
	logger1 := GetAppLogger()
	logger2 := GetAppLogger()

	fmt.Printf("Same AppLogger instance: %v\n", logger1 == logger2)
	logger1.Log("Application started")
	logger2.Log("Processing request")

	logs := logger1.GetLogs()
	fmt.Printf("\nTotal logs: %d\n", len(logs))
}
