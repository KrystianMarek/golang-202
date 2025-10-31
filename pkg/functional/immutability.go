package functional

import "fmt"

// Immutability demonstrates immutable data structures through copy-on-write.
//
// Why? Immutable data structures prevent unintended side effects and make
// concurrent code safer, though they trade memory for safety.

// Point represents an immutable 2D point.
type Point struct {
	x, y float64
}

// NewPoint creates a new point.
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// X returns the x coordinate.
func (p Point) X() float64 {
	return p.x
}

// Y returns the y coordinate.
func (p Point) Y() float64 {
	return p.y
}

// WithX returns a new point with updated x.
func (p Point) WithX(x float64) Point {
	return Point{x: x, y: p.y}
}

// WithY returns a new point with updated y.
func (p Point) WithY(y float64) Point {
	return Point{x: p.x, y: y}
}

// Move returns a new point moved by dx, dy.
func (p Point) Move(dx, dy float64) Point {
	return Point{x: p.x + dx, y: p.y + dy}
}

// ImmutableList represents an immutable list.
type ImmutableList[T any] struct {
	items []T
}

// NewImmutableList creates a new immutable list.
func NewImmutableList[T any](items ...T) ImmutableList[T] {
	copied := make([]T, len(items))
	copy(copied, items)
	return ImmutableList[T]{items: copied}
}

// Get returns the item at index.
func (l ImmutableList[T]) Get(index int) T {
	return l.items[index]
}

// Size returns the list size.
func (l ImmutableList[T]) Size() int {
	return len(l.items)
}

// Add returns a new list with the item added.
func (l ImmutableList[T]) Add(item T) ImmutableList[T] {
	newItems := make([]T, len(l.items)+1)
	copy(newItems, l.items)
	newItems[len(l.items)] = item
	return ImmutableList[T]{items: newItems}
}

// Remove returns a new list with the item at index removed.
func (l ImmutableList[T]) Remove(index int) ImmutableList[T] {
	if index < 0 || index >= len(l.items) {
		return l
	}
	newItems := make([]T, 0, len(l.items)-1)
	newItems = append(newItems, l.items[:index]...)
	newItems = append(newItems, l.items[index+1:]...)
	return ImmutableList[T]{items: newItems}
}

// Map returns a new list with mapper applied to each element.
func (l ImmutableList[T]) Map(mapper func(T) T) ImmutableList[T] {
	newItems := make([]T, len(l.items))
	for i, item := range l.items {
		newItems[i] = mapper(item)
	}
	return ImmutableList[T]{items: newItems}
}

// Filter returns a new list with only elements satisfying the predicate.
func (l ImmutableList[T]) Filter(predicate func(T) bool) ImmutableList[T] {
	newItems := make([]T, 0)
	for _, item := range l.items {
		if predicate(item) {
			newItems = append(newItems, item)
		}
	}
	return ImmutableList[T]{items: newItems}
}

// ToSlice returns a copy of the internal slice.
func (l ImmutableList[T]) ToSlice() []T {
	result := make([]T, len(l.items))
	copy(result, l.items)
	return result
}

// User represents an immutable user record.
type User struct {
	id       int
	username string
	email    string
	age      int
}

// NewUser creates a new user.
func NewUser(id int, username, email string, age int) User {
	return User{
		id:       id,
		username: username,
		email:    email,
		age:      age,
	}
}

// ID returns the user ID.
func (u User) ID() int {
	return u.id
}

// Username returns the username.
func (u User) Username() string {
	return u.username
}

// Email returns the email.
func (u User) Email() string {
	return u.email
}

// Age returns the age.
func (u User) Age() int {
	return u.age
}

// WithEmail returns a new user with updated email.
func (u User) WithEmail(email string) User {
	return User{
		id:       u.id,
		username: u.username,
		email:    email,
		age:      u.age,
	}
}

// WithAge returns a new user with updated age.
func (u User) WithAge(age int) User {
	return User{
		id:       u.id,
		username: u.username,
		email:    u.email,
		age:      age,
	}
}

// Config represents immutable configuration.
type Config struct {
	settings map[string]string
}

// NewConfig creates a new config.
func NewConfig(settings map[string]string) Config {
	copied := make(map[string]string, len(settings))
	for k, v := range settings {
		copied[k] = v
	}
	return Config{settings: copied}
}

// Get retrieves a setting.
func (c Config) Get(key string) string {
	return c.settings[key]
}

// WithSetting returns a new config with the setting added/updated.
func (c Config) WithSetting(key, value string) Config {
	newSettings := make(map[string]string, len(c.settings)+1)
	for k, v := range c.settings {
		newSettings[k] = v
	}
	newSettings[key] = value
	return Config{settings: newSettings}
}

// WithoutSetting returns a new config with the setting removed.
func (c Config) WithoutSetting(key string) Config {
	newSettings := make(map[string]string, len(c.settings))
	for k, v := range c.settings {
		if k != key {
			newSettings[k] = v
		}
	}
	return Config{settings: newSettings}
}

// GetAll returns a copy of all settings.
func (c Config) GetAll() map[string]string {
	result := make(map[string]string, len(c.settings))
	for k, v := range c.settings {
		result[k] = v
	}
	return result
}

// ExampleImmutability demonstrates immutable data structures.
func ExampleImmutability() {
	fmt.Println("=== Immutability ===")

	// Immutable points
	p1 := NewPoint(1.0, 2.0)
	p2 := p1.WithX(5.0)
	p3 := p1.Move(3.0, 4.0)

	fmt.Printf("p1: (%.1f, %.1f)\n", p1.X(), p1.Y())
	fmt.Printf("p2: (%.1f, %.1f)\n", p2.X(), p2.Y())
	fmt.Printf("p3: (%.1f, %.1f)\n\n", p3.X(), p3.Y())

	// Immutable list
	list1 := NewImmutableList(1, 2, 3)
	list2 := list1.Add(4).Add(5)
	list3 := list2.Remove(1)

	fmt.Printf("list1: %v\n", list1.ToSlice())
	fmt.Printf("list2: %v\n", list2.ToSlice())
	fmt.Printf("list3: %v\n\n", list3.ToSlice())

	// List transformations
	doubled := list1.Map(func(n int) int { return n * 2 })
	evens := list2.Filter(func(n int) bool { return n%2 == 0 })

	fmt.Printf("doubled: %v\n", doubled.ToSlice())
	fmt.Printf("evens: %v\n\n", evens.ToSlice())

	// Immutable user
	user1 := NewUser(1, "alice", "alice@example.com", 25)
	user2 := user1.WithEmail("alice.new@example.com")
	user3 := user2.WithAge(26)

	fmt.Printf("user1: %s (%s), age %d\n",
		user1.Username(), user1.Email(), user1.Age())
	fmt.Printf("user2: %s (%s), age %d\n",
		user2.Username(), user2.Email(), user2.Age())
	fmt.Printf("user3: %s (%s), age %d\n\n",
		user3.Username(), user3.Email(), user3.Age())

	// Immutable config
	config1 := NewConfig(map[string]string{
		"host": "localhost",
		"port": "8080",
	})
	config2 := config1.WithSetting("debug", "true")
	config3 := config2.WithoutSetting("port")

	fmt.Printf("config1: %v\n", config1.GetAll())
	fmt.Printf("config2: %v\n", config2.GetAll())
	fmt.Printf("config3: %v\n", config3.GetAll())
}

