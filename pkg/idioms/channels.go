package idioms

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Channels demonstrates Go 1.24 enhanced channel patterns.
//
// Why? Go 1.24 improves channel safety with guaranteed for-range termination,
// optimized select with context, and better diagnostics. Channels are now
// safer, faster, and more composable than ever.

// SafePipeline demonstrates guaranteed termination with for-range.
// Go 1.24 guarantees the loop exits when the channel is closed.
func SafePipeline(ctx context.Context, input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output) // Guarantees for-range will terminate

		for item := range input { // Go 1.24: guaranteed termination
			select {
			case <-ctx.Done():
				return
			case output <- item * 2:
			}
		}
	}()

	return output
}

// Generator creates a channel that produces values.
func Generator(ctx context.Context, values ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}

// FanOutFanIn demonstrates concurrent processing with channels.
func FanOutFanIn(ctx context.Context, input <-chan int, workers int) <-chan int {
	// Fan-out: distribute work to multiple workers
	workerChannels := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		workerChannels[i] = worker(ctx, input, i)
	}

	// Fan-in: merge results from all workers
	return merge(ctx, workerChannels...)
}

func worker(ctx context.Context, input <-chan int, id int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for val := range input {
			select {
			case <-ctx.Done():
				return
			case output <- val * val:
				fmt.Printf("Worker %d processed: %d\n", id, val)
			}
		}
	}()

	return output
}

func merge(ctx context.Context, channels ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()

			for val := range c { // Safe termination
				select {
				case <-ctx.Done():
					return
				case output <- val:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Tee splits a channel into two output channels.
func Tee(ctx context.Context, input <-chan int) (out1, out2 <-chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out1, out2 = ch1, ch2

	go func() {
		defer close(ch1)
		defer close(ch2)

		for val := range input {
			select {
			case <-ctx.Done():
				return
			case ch1 <- val:
			}

			select {
			case <-ctx.Done():
				return
			case ch2 <- val:
			}
		}
	}()

	return
}

// Broadcaster sends values to multiple subscribers.
type Broadcaster[T any] struct {
	mu          sync.RWMutex
	subscribers []chan T
	input       chan T
	ctx         context.Context
	cancel      context.CancelFunc
}

// NewBroadcaster creates a new broadcaster.
func NewBroadcaster[T any]() *Broadcaster[T] {
	ctx, cancel := context.WithCancel(context.Background())
	b := &Broadcaster[T]{
		subscribers: make([]chan T, 0),
		input:       make(chan T, 10),
		ctx:         ctx,
		cancel:      cancel,
	}

	go b.run()
	return b
}

func (b *Broadcaster[T]) run() {
	for {
		select {
		case <-b.ctx.Done():
			b.closeAll()
			return
		case val := <-b.input:
			b.broadcast(val)
		}
	}
}

func (b *Broadcaster[T]) broadcast(val T) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers {
		select {
		case ch <- val:
		default:
			// Skip slow consumers
		}
	}
}

// Subscribe creates a new subscription channel.
func (b *Broadcaster[T]) Subscribe() <-chan T {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan T, 10)
	b.subscribers = append(b.subscribers, ch)
	return ch
}

// Send broadcasts a value to all subscribers.
func (b *Broadcaster[T]) Send(val T) {
	select {
	case b.input <- val:
	case <-b.ctx.Done():
	}
}

// Close shuts down the broadcaster.
func (b *Broadcaster[T]) Close() {
	b.cancel()
}

func (b *Broadcaster[T]) closeAll() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subscribers {
		close(ch)
	}
}

// OrDone wraps a channel with context cancellation.
func OrDone[T any](ctx context.Context, ch <-chan T) <-chan T {
	output := make(chan T)

	go func() {
		defer close(output)

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case output <- val:
				}
			}
		}
	}()

	return output
}

// Bridge flattens a channel of channels.
func Bridge[T any](ctx context.Context, chanStream <-chan <-chan T) <-chan T {
	output := make(chan T)

	go func() {
		defer close(output)

		for {
			select {
			case <-ctx.Done():
				return
			case ch, ok := <-chanStream:
				if !ok {
					return
				}

				for val := range OrDone(ctx, ch) {
					select {
					case <-ctx.Done():
						return
					case output <- val:
					}
				}
			}
		}
	}()

	return output
}

// ExampleChannels demonstrates Go 1.24 channel patterns.
func ExampleChannels() {
	fmt.Println("=== Go 1.24 Channels ===")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Safe pipeline with guaranteed termination
	fmt.Println("\nSafe Pipeline:")
	input := Generator(ctx, 1, 2, 3, 4, 5)
	doubled := SafePipeline(ctx, input)

	for val := range doubled {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Fan-out/Fan-in pattern
	fmt.Println("\nFan-Out/Fan-In:")
	input2 := Generator(ctx, 1, 2, 3, 4, 5)
	results := FanOutFanIn(ctx, input2, 3)

	collected := make([]int, 0)
	for val := range results {
		collected = append(collected, val)
	}
	fmt.Printf("Results: %v\n", collected)

	// Broadcaster pattern
	fmt.Println("\nBroadcaster:")
	broadcaster := NewBroadcaster[string]()
	defer broadcaster.Close()

	sub1 := broadcaster.Subscribe()
	sub2 := broadcaster.Subscribe()

	go func() {
		for msg := range sub1 {
			fmt.Printf("Sub1: %s\n", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Printf("Sub2: %s\n", msg)
		}
	}()

	broadcaster.Send("Hello")
	broadcaster.Send("World")
	time.Sleep(100 * time.Millisecond)

	// OrDone pattern
	fmt.Println("\nOrDone with Context:")
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer timeoutCancel()

	slowChan := make(chan int)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slowChan <- 42
	}()

	safeChan := OrDone(timeoutCtx, slowChan)
	select {
	case val := <-safeChan:
		fmt.Printf("Received: %d\n", val)
	case <-timeoutCtx.Done():
		fmt.Println("Context cancelled - safe exit")
	}
}
