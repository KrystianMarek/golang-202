package idioms

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Concurrency demonstrates goroutines, channels, and patterns.
//
// Why? Go's built-in concurrency primitives make concurrent
// programming accessible and idiomatic.

// GenerateNumbers creates a channel that produces values.
func GenerateNumbers(ctx context.Context, start, end int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

// Square demonstrates channel-based pipelines.
func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()

	return out
}

// FanOut distributes work across multiple workers.
func FanOut(ctx context.Context, in <-chan int, workers int) []<-chan int {
	channels := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		channels[i] = Square(ctx, in)
	}

	return channels
}

// FanIn merges multiple channels into one.
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case <-ctx.Done():
					return
				case out <- n:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// WorkerPool demonstrates the worker pool pattern.
type WorkerPool struct {
	workers int
	jobs    chan func()
	wg      sync.WaitGroup
}

// NewWorkerPool creates a worker pool.
func NewWorkerPool(workers int) *WorkerPool {
	pool := &WorkerPool{
		workers: workers,
		jobs:    make(chan func(), 100),
	}
	pool.start()
	return pool
}

func (p *WorkerPool) start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for job := range p.jobs {
		job()
	}
}

// Submit submits a job to the pool.
func (p *WorkerPool) Submit(job func()) {
	p.jobs <- job
}

// Close closes the pool and waits for completion.
func (p *WorkerPool) Close() {
	close(p.jobs)
	p.wg.Wait()
}

// SelectExample demonstrates the select statement.
func SelectExample(ctx context.Context) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-ctx.Done():
			fmt.Println("Context cancelled")
			return
		}
	}
}

// DoWithTimeout demonstrates timeout pattern with context.
func DoWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan bool)

	go func() {
		// Simulate work
		time.Sleep(200 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RateLimiter implements a token bucket rate limiter.
type RateLimiter struct {
	tokens chan struct{}
	rate   time.Duration
	stop   chan struct{}
}

// NewRateLimiter creates a rate limiter.
func NewRateLimiter(capacity int, rate time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, capacity),
		rate:   rate,
		stop:   make(chan struct{}),
	}

	// Fill initial tokens
	for i := 0; i < capacity; i++ {
		rl.tokens <- struct{}{}
	}

	// Refill tokens
	go func() {
		ticker := time.NewTicker(rate)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case rl.tokens <- struct{}{}:
				default:
					// Bucket full
				}
			case <-rl.stop:
				return
			}
		}
	}()

	return rl
}

// Allow checks if an action is allowed.
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

// Wait waits for a token.
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

// Close stops the rate limiter.
func (rl *RateLimiter) Close() {
	close(rl.stop)
}

// SafeCounter demonstrates synchronized access.
type SafeCounter struct {
	mu    sync.RWMutex
	count int
}

// Increment increments the counter.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the counter value.
func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

// ExampleConcurrency demonstrates concurrency patterns.
func ExampleConcurrency() {
	fmt.Println("=== Concurrency ===")

	ctx := context.Background()

	// Basic channel pipeline
	fmt.Println("Pipeline:")
	numbers := GenerateNumbers(ctx, 1, 5)
	squares := Square(ctx, numbers)

	for n := range squares {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// Worker pool
	fmt.Println("Worker Pool:")
	pool := NewWorkerPool(3)

	for i := 1; i <= 5; i++ {
		id := i
		pool.Submit(func() {
			fmt.Printf("Job %d executing\n", id)
			time.Sleep(100 * time.Millisecond)
		})
	}

	pool.Close()

	// Select statement
	fmt.Println("Select:")
	SelectExample(ctx)

	// Timeout
	fmt.Println("Timeout:")
	if err := DoWithTimeout(100 * time.Millisecond); err != nil {
		fmt.Printf("Operation timed out: %v\n", err)
	} else {
		fmt.Println("Operation completed")
	}

	// Rate limiter
	fmt.Println("Rate Limiter:")
	limiter := NewRateLimiter(3, 100*time.Millisecond)
	defer limiter.Close()

	for i := 1; i <= 5; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: allowed\n", i)
		} else {
			fmt.Printf("Request %d: rate limited\n", i)
		}
	}

	// Safe counter
	fmt.Println("Safe Counter:")
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())
}
