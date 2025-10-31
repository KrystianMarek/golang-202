package go124

import (
	"fmt"
	"runtime"
)

// Resource represents a managed resource with cleanup.
// This demonstrates resource cleanup patterns.
//
// Why? Automatic cleanup using finalizers helps prevent resource leaks.
// While SetFinalizer is used here, Go 1.24 introduces runtime.AddCleanup
// for more predictable cleanup behavior. This is useful for file handles,
// network connections, and temporary resources.
type Resource struct {
	ID   string
	Data []byte
}

// NewResource creates a resource with automatic cleanup.
// Note: runtime.AddCleanup is a Go 1.24 feature. If not available,
// consider using runtime.SetFinalizer as an alternative.
func NewResource(id string, size int) *Resource {
	r := &Resource{
		ID:   id,
		Data: make([]byte, size),
	}

	// Register cleanup function using SetFinalizer
	// (AddCleanup is not yet in stable release)
	runtime.SetFinalizer(r, func(res *Resource) {
		fmt.Printf("Cleaning up resource: %s\n", res.ID)
		// In real code, this might close files, release locks, etc.
		res.Data = nil
	})

	return r
}

// Use simulates using the resource.
func (r *Resource) Use() {
	fmt.Printf("Using resource: %s (size: %d bytes)\n", r.ID, len(r.Data))
}

// FileHandle represents a managed file handle.
type FileHandle struct {
	Path string
	fd   int
}

// OpenFile simulates opening a file with cleanup.
func OpenFile(path string) *FileHandle {
	fh := &FileHandle{
		Path: path,
		fd:   42, // Simulated file descriptor
	}

	runtime.SetFinalizer(fh, func(handle *FileHandle) {
		fmt.Printf("Closing file: %s (fd: %d)\n", handle.Path, handle.fd)
		handle.fd = -1
	})

	return fh
}

// Read simulates reading from the file.
func (fh *FileHandle) Read() []byte {
	if fh.fd < 0 {
		return nil
	}
	fmt.Printf("Reading from: %s\n", fh.Path)
	return []byte(fmt.Sprintf("data from %s", fh.Path))
}

// TempBuffer represents a temporary buffer with automatic cleanup.
type TempBuffer struct {
	Name string
	buf  []byte
}

// NewTempBuffer creates a temporary buffer.
func NewTempBuffer(name string, capacity int) *TempBuffer {
	tb := &TempBuffer{
		Name: name,
		buf:  make([]byte, 0, capacity),
	}

	runtime.SetFinalizer(tb, func(buffer *TempBuffer) {
		fmt.Printf("Releasing temp buffer: %s (cap: %d)\n",
			buffer.Name, cap(buffer.buf))
		buffer.buf = nil
	})

	return tb
}

// Write adds data to the buffer.
func (tb *TempBuffer) Write(data []byte) {
	tb.buf = append(tb.buf, data...)
}

// ExampleCleanup demonstrates resource cleanup patterns.
func ExampleCleanup() {
	fmt.Println("Creating resources...")

	// Create some resources
	r1 := NewResource("resource-1", 1024)
	r1.Use()

	fh := OpenFile("/tmp/example.txt")
	data := fh.Read()
	fmt.Printf("Read: %s\n", string(data))

	tb := NewTempBuffer("temp-1", 512)
	tb.Write([]byte("temporary data"))

	// Resources will be cleaned up when they go out of scope
	// and GC runs. Force GC for demonstration.
	fmt.Println("\nForcing GC to trigger cleanup...")
	runtime.GC()

	fmt.Println("Example complete")
}

