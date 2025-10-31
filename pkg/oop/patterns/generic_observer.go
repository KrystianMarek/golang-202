package patterns

import (
	"fmt"
	"sync"
)

// Generic Observer pattern demonstrates type-safe event handling.
//
// Why? Generic observers provide compile-time type safety while
// maintaining the flexibility of the Observer pattern.

// GenericObserver is a type-safe observer interface.
type GenericObserver[T any] interface {
	OnEvent(event T)
	GetID() string
}

// GenericSubject manages generic observers.
type GenericSubject[T any] struct {
	mu        sync.RWMutex
	observers map[string]GenericObserver[T]
}

// NewGenericSubject creates a new generic subject.
func NewGenericSubject[T any]() *GenericSubject[T] {
	return &GenericSubject[T]{
		observers: make(map[string]GenericObserver[T]),
	}
}

// Attach adds an observer.
func (s *GenericSubject[T]) Attach(observer GenericObserver[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers[observer.GetID()] = observer
}

// Detach removes an observer.
func (s *GenericSubject[T]) Detach(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.observers, id)
}

// Notify sends an event to all observers.
func (s *GenericSubject[T]) Notify(event T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.OnEvent(event)
	}
}

// UserEvent represents a user-related event.
type UserEvent struct {
	Type     string
	UserID   int
	Username string
}

// UserEventLogger logs user events.
type UserEventLogger struct {
	ID   string
	logs []UserEvent
	mu   sync.Mutex
}

// OnEvent handles user events.
func (l *UserEventLogger) OnEvent(event UserEvent) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.logs = append(l.logs, event)
	fmt.Printf("[Logger] User event: %s for %s\n", event.Type, event.Username)
}

// GetID returns the observer ID.
func (l *UserEventLogger) GetID() string {
	return l.ID
}

// UserEventNotifier sends notifications.
type UserEventNotifier struct {
	ID string
}

// OnEvent sends notifications for user events.
func (n *UserEventNotifier) OnEvent(event UserEvent) {
	fmt.Printf("[Notifier] Sending notification: %s - User %s\n",
		event.Type, event.Username)
}

// GetID returns the observer ID.
func (n *UserEventNotifier) GetID() string {
	return n.ID
}

// OrderEvent represents an order-related event.
type OrderEvent struct {
	Type    string
	OrderID int
	Amount  float64
}

// OrderEventProcessor processes orders.
type OrderEventProcessor struct {
	ID           string
	TotalRevenue float64
	mu           sync.Mutex
}

// OnEvent processes order events.
func (p *OrderEventProcessor) OnEvent(event OrderEvent) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if event.Type == "order.completed" {
		p.TotalRevenue += event.Amount
		fmt.Printf("[Processor] Order #%d completed: $%.2f (Total: $%.2f)\n",
			event.OrderID, event.Amount, p.TotalRevenue)
	}
}

// GetID returns the observer ID.
func (p *OrderEventProcessor) GetID() string {
	return p.ID
}

// GenericChannelSubject uses channels for event distribution.
type GenericChannelSubject[T any] struct {
	mu          sync.RWMutex
	subscribers map[string]chan T
}

// NewGenericChannelSubject creates a channel-based subject.
func NewGenericChannelSubject[T any]() *GenericChannelSubject[T] {
	return &GenericChannelSubject[T]{
		subscribers: make(map[string]chan T),
	}
}

// Subscribe creates a new subscription.
func (s *GenericChannelSubject[T]) Subscribe(id string, bufferSize int) <-chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan T, bufferSize)
	s.subscribers[id] = ch
	return ch
}

// Unsubscribe removes a subscription.
func (s *GenericChannelSubject[T]) Unsubscribe(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if ch, ok := s.subscribers[id]; ok {
		close(ch)
		delete(s.subscribers, id)
	}
}

// Publish sends an event to all subscribers.
func (s *GenericChannelSubject[T]) Publish(event T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for id, ch := range s.subscribers {
		select {
		case ch <- event:
		default:
			fmt.Printf("Warning: Subscriber %s buffer full\n", id)
		}
	}
}

// Close closes all subscriptions.
func (s *GenericChannelSubject[T]) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ch := range s.subscribers {
		close(ch)
	}
	s.subscribers = make(map[string]chan T)
}

// ExampleGenericObserver demonstrates generic observer patterns.
func ExampleGenericObserver() {
	fmt.Println("=== Generic Observer Pattern ===")

	// Type-safe user event observers
	userSubject := NewGenericSubject[UserEvent]()

	logger := &UserEventLogger{
		ID:   "logger-1",
		logs: make([]UserEvent, 0),
	}

	notifier := &UserEventNotifier{ID: "notifier-1"}

	userSubject.Attach(logger)
	userSubject.Attach(notifier)

	userSubject.Notify(UserEvent{
		Type:     "user.login",
		UserID:   123,
		Username: "alice",
	})

	userSubject.Notify(UserEvent{
		Type:     "user.logout",
		UserID:   123,
		Username: "alice",
	})

	// Type-safe order event observers
	orderSubject := NewGenericSubject[OrderEvent]()

	processor := &OrderEventProcessor{ID: "processor-1"}
	orderSubject.Attach(processor)

	orderSubject.Notify(OrderEvent{
		Type:    "order.completed",
		OrderID: 1001,
		Amount:  99.99,
	})

	orderSubject.Notify(OrderEvent{
		Type:    "order.completed",
		OrderID: 1002,
		Amount:  149.99,
	})

	// Channel-based generic observer
	fmt.Println("\nChannel-based Generic Observer:")

	channelSubject := NewGenericChannelSubject[string]()
	defer channelSubject.Close()

	sub1 := channelSubject.Subscribe("sub-1", 10)
	sub2 := channelSubject.Subscribe("sub-2", 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for msg := range sub1 {
			fmt.Printf("[Sub-1] Received: %s\n", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range sub2 {
			fmt.Printf("[Sub-2] Received: %s\n", msg)
		}
	}()

	channelSubject.Publish("Message 1")
	channelSubject.Publish("Message 2")
	channelSubject.Publish("Message 3")

	channelSubject.Close()
	wg.Wait()
}

