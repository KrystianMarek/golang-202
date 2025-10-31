// Package examples provides integrated examples combining multiple patterns.
package examples

import (
	"fmt"
	"sync"
)

// GameEngine demonstrates OOP + FP + patterns in a game context.

// Entity interface for game objects.
type Entity interface {
	Update(deltaTime float64)
	Render() string
	GetID() string
}

// Component represents a behavior that can be attached to entities.
type Component interface {
	Update(deltaTime float64)
	GetType() string
}

// BaseEntity provides common entity functionality.
type BaseEntity struct {
	id         string
	components []Component
}

// NewBaseEntity creates a base entity.
func NewBaseEntity(id string) *BaseEntity {
	return &BaseEntity{
		id:         id,
		components: make([]Component, 0),
	}
}

// GetID returns the entity ID.
func (e *BaseEntity) GetID() string {
	return e.id
}

// AddComponent adds a component.
func (e *BaseEntity) AddComponent(c Component) {
	e.components = append(e.components, c)
}

// Update updates all components.
func (e *BaseEntity) Update(deltaTime float64) {
	for _, c := range e.components {
		c.Update(deltaTime)
	}
}

// PositionComponent tracks position.
type PositionComponent struct {
	X, Y       float64
	VelX, VelY float64
}

// Update updates position based on velocity.
func (p *PositionComponent) Update(deltaTime float64) {
	p.X += p.VelX * deltaTime
	p.Y += p.VelY * deltaTime
}

// GetType returns the component type.
func (p *PositionComponent) GetType() string {
	return "Position"
}

// HealthComponent tracks health.
type HealthComponent struct {
	Current, Max int
}

// Update does nothing for health.
func (h *HealthComponent) Update(deltaTime float64) {
	// Health doesn't auto-update
}

// GetType returns the component type.
func (h *HealthComponent) GetType() string {
	return "Health"
}

// TakeDamage reduces health.
func (h *HealthComponent) TakeDamage(amount int) {
	h.Current -= amount
	if h.Current < 0 {
		h.Current = 0
	}
}

// IsAlive returns true if health > 0.
func (h *HealthComponent) IsAlive() bool {
	return h.Current > 0
}

// Player is a concrete entity.
type Player struct {
	*BaseEntity
	name     string
	position *PositionComponent
	health   *HealthComponent
}

// NewPlayer creates a player.
func NewPlayer(id, name string) *Player {
	base := NewBaseEntity(id)

	pos := &PositionComponent{X: 0, Y: 0, VelX: 0, VelY: 0}
	health := &HealthComponent{Current: 100, Max: 100}

	base.AddComponent(pos)
	base.AddComponent(health)

	return &Player{
		BaseEntity: base,
		name:       name,
		position:   pos,
		health:     health,
	}
}

// Render renders the player.
func (p *Player) Render() string {
	return fmt.Sprintf("Player[%s] at (%.1f, %.1f) HP:%d/%d",
		p.name, p.position.X, p.position.Y,
		p.health.Current, p.health.Max)
}

// Move sets velocity.
func (p *Player) Move(velX, velY float64) {
	p.position.VelX = velX
	p.position.VelY = velY
}

// GameEvent represents game events (Observer pattern).
type GameEvent struct {
	Type string
	Data interface{}
}

// EventListener listens to game events.
type EventListener interface {
	OnEvent(event GameEvent)
}

// EventManager manages game events (Singleton + Observer).
type EventManager struct {
	mu        sync.RWMutex
	listeners map[string][]EventListener
}

var (
	eventManagerInstance *EventManager
	eventManagerOnce     sync.Once
)

// GetEventManager returns the singleton event manager.
func GetEventManager() *EventManager {
	eventManagerOnce.Do(func() {
		eventManagerInstance = &EventManager{
			listeners: make(map[string][]EventListener),
		}
	})
	return eventManagerInstance
}

// Subscribe adds a listener for an event type.
func (em *EventManager) Subscribe(eventType string, listener EventListener) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.listeners[eventType] = append(em.listeners[eventType], listener)
}

// Publish publishes an event to all listeners.
func (em *EventManager) Publish(event GameEvent) {
	em.mu.RLock()
	listeners := em.listeners[event.Type]
	em.mu.RUnlock()

	for _, listener := range listeners {
		listener.OnEvent(event)
	}
}

// ScoreListener listens to score events.
type ScoreListener struct {
	score int
}

// OnEvent handles score events.
func (s *ScoreListener) OnEvent(event GameEvent) {
	if event.Type == "score" {
		if points, ok := event.Data.(int); ok {
			s.score += points
			fmt.Printf("Score updated: %d\n", s.score)
		}
	}
}

// Game orchestrates the game (Facade pattern).
type Game struct {
	entities     []Entity
	eventManager *EventManager
	running      bool
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		entities:     make([]Entity, 0),
		eventManager: GetEventManager(),
		running:      false,
	}
}

// AddEntity adds an entity to the game.
func (g *Game) AddEntity(e Entity) {
	g.entities = append(g.entities, e)
}

// Update updates all entities.
func (g *Game) Update(deltaTime float64) {
	for _, e := range g.entities {
		e.Update(deltaTime)
	}
}

// Render renders all entities.
func (g *Game) Render() {
	fmt.Println("\n=== Game State ===")
	for _, e := range g.entities {
		fmt.Println(e.Render())
	}
	fmt.Println("==================")
}

// ExampleGameEngine demonstrates the integrated game engine.
func ExampleGameEngine() {
	fmt.Println("=== Game Engine Example ===")

	game := NewGame()

	// Create players
	player1 := NewPlayer("p1", "Alice")
	player2 := NewPlayer("p2", "Bob")

	game.AddEntity(player1)
	game.AddEntity(player2)

	// Subscribe to events
	scoreListener := &ScoreListener{}
	game.eventManager.Subscribe("score", scoreListener)

	// Initial state
	game.Render()

	// Move players
	player1.Move(10, 0)
	player2.Move(-5, 5)

	// Update game state
	fmt.Println("Updating game (deltaTime: 0.1)...")
	game.Update(0.1)
	game.Render()

	// Take damage
	player1.health.TakeDamage(30)
	fmt.Printf("%s took damage!\n", player1.name)

	// Publish score event
	game.eventManager.Publish(GameEvent{
		Type: "score",
		Data: 100,
	})

	game.Render()
}
