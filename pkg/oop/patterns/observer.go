package patterns

import (
	"fmt"
	"sync"
)

// Observer interface for the Observer pattern.
// This demonstrates the Observer pattern using channels and interfaces.
//
// Why? Go's channels provide a natural implementation for the Observer
// pattern, enabling event-driven architectures and pub/sub systems.
type Observer interface {
	Update(event Event)
	GetID() string
}

// Event represents an event in the system.
type Event struct {
	Type string
	Data interface{}
}

// Subject manages observers and notifies them of events.
type Subject struct {
	mu        sync.RWMutex
	observers map[string]Observer
}

// NewSubject creates a new Subject.
func NewSubject() *Subject {
	return &Subject{
		observers: make(map[string]Observer),
	}
}

// Attach adds an observer.
func (s *Subject) Attach(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers[observer.GetID()] = observer
	fmt.Printf("Observer %s attached\n", observer.GetID())
}

// Detach removes an observer.
func (s *Subject) Detach(observerID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.observers, observerID)
	fmt.Printf("Observer %s detached\n", observerID)
}

// Notify sends an event to all observers.
func (s *Subject) Notify(event Event) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Printf("Notifying %d observers of event: %s\n",
		len(s.observers), event.Type)

	for _, observer := range s.observers {
		observer.Update(event)
	}
}

// EmailObserver observes events and sends emails.
type EmailObserver struct {
	ID    string
	Email string
}

// Update handles the event.
func (e *EmailObserver) Update(event Event) {
	fmt.Printf("[%s] Sending email to %s: %s - %v\n",
		e.ID, e.Email, event.Type, event.Data)
}

// GetID returns the observer ID.
func (e *EmailObserver) GetID() string {
	return e.ID
}

// LogObserver logs events.
type LogObserver struct {
	ID   string
	logs []Event
	mu   sync.Mutex
}

// Update logs the event.
func (l *LogObserver) Update(event Event) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.logs = append(l.logs, event)
	fmt.Printf("[%s] Logged event: %s\n", l.ID, event.Type)
}

// GetID returns the observer ID.
func (l *LogObserver) GetID() string {
	return l.ID
}

// GetLogs returns all logged events.
func (l *LogObserver) GetLogs() []Event {
	l.mu.Lock()
	defer l.mu.Unlock()

	logs := make([]Event, len(l.logs))
	copy(logs, l.logs)
	return logs
}

// ChannelEventBus demonstrates channel-based pub/sub.
type ChannelEventBus struct {
	subscribers map[string][]chan Event
	mu          sync.RWMutex
}

// NewChannelEventBus creates a new event bus.
func NewChannelEventBus() *ChannelEventBus {
	return &ChannelEventBus{
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe creates a channel for a specific event type.
func (b *ChannelEventBus) Subscribe(eventType string) chan Event {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Event, 10)
	b.subscribers[eventType] = append(b.subscribers[eventType], ch)

	fmt.Printf("New subscriber for event type: %s\n", eventType)
	return ch
}

// Publish sends an event to all subscribers.
func (b *ChannelEventBus) Publish(event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	channels := b.subscribers[event.Type]
	fmt.Printf("Publishing event %s to %d subscribers\n",
		event.Type, len(channels))

	for _, ch := range channels {
		// Non-blocking send
		select {
		case ch <- event:
		default:
			fmt.Println("Channel full, skipping")
		}
	}
}

// Close closes all channels.
func (b *ChannelEventBus) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, channels := range b.subscribers {
		for _, ch := range channels {
			close(ch)
		}
	}
}

// ExampleObserver demonstrates the Observer pattern.
func ExampleObserver() {
	fmt.Println("=== Observer Pattern ===")

	// Traditional observer pattern
	subject := NewSubject()

	emailObs := &EmailObserver{
		ID:    "email-1",
		Email: "admin@example.com",
	}

	logObs := &LogObserver{
		ID:   "log-1",
		logs: make([]Event, 0),
	}

	subject.Attach(emailObs)
	subject.Attach(logObs)

	subject.Notify(Event{
		Type: "user.created",
		Data: map[string]string{"username": "alice"},
	})

	subject.Notify(Event{
		Type: "order.placed",
		Data: map[string]interface{}{"order_id": 123, "total": 99.99},
	})

	subject.Detach(emailObs.GetID())

	subject.Notify(Event{
		Type: "payment.received",
		Data: 50.00,
	})

	fmt.Printf("\nTotal logged events: %d\n\n", len(logObs.GetLogs()))

	// Channel-based observer pattern
	fmt.Println("Channel-based Event Bus:")

	eventBus := NewChannelEventBus()
	defer eventBus.Close()

	userEventsCh := eventBus.Subscribe("user.event")
	orderEventsCh := eventBus.Subscribe("order.event")

	// Start listeners
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for event := range userEventsCh {
			fmt.Printf("[User Listener] Received: %s - %v\n",
				event.Type, event.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for event := range orderEventsCh {
			fmt.Printf("[Order Listener] Received: %s - %v\n",
				event.Type, event.Data)
		}
	}()

	// Publish events
	eventBus.Publish(Event{Type: "user.event", Data: "User logged in"})
	eventBus.Publish(Event{Type: "order.event", Data: "Order created"})
	eventBus.Publish(Event{Type: "user.event", Data: "User updated profile"})

	eventBus.Close()
	wg.Wait()
}

