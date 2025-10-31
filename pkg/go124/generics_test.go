package go124

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}

	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	peek, ok := stack.Peek()
	if !ok || peek != 2 {
		t.Errorf("Expected peek 2, got %d", peek)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}
}

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()

	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	val, ok := queue.Dequeue()
	if !ok || val != "first" {
		t.Errorf("Expected 'first', got '%s'", val)
	}

	val, ok = queue.Dequeue()
	if !ok || val != "second" {
		t.Errorf("Expected 'second', got '%s'", val)
	}
}

func TestSet(t *testing.T) {
	set := NewSet(1, 2, 3)

	if !set.Contains(2) {
		t.Error("Set should contain 2")
	}

	set.Add(4)
	if set.Size() != 4 {
		t.Errorf("Expected size 4, got %d", set.Size())
	}

	set.Remove(2)
	if set.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}
}

func TestSetOperations(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)

	union := set1.Union(set2)
	if union.Size() != 5 {
		t.Errorf("Expected union size 5, got %d", union.Size())
	}

	intersection := set1.Intersection(set2)
	if intersection.Size() != 1 || !intersection.Contains(3) {
		t.Error("Expected intersection to contain only 3")
	}
}

func TestGenericMath(t *testing.T) {
	if Min(5, 3) != 3 {
		t.Error("Min failed")
	}

	if Max(5, 3) != 5 {
		t.Error("Max failed")
	}

	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Error("Sum failed")
	}
}

func TestCache(t *testing.T) {
	cache := NewCache[string, int]()

	cache.Set("age", 25)
	cache.Set("score", 100)

	val, ok := cache.Get("age")
	if !ok || val != 25 {
		t.Errorf("Expected age 25, got %d", val)
	}

	keys := cache.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	cache.Delete("age")
	_, ok = cache.Get("age")
	if ok {
		t.Error("Key should be deleted")
	}
}

func BenchmarkStackPush(b *testing.B) {
	stack := NewStack[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := NewStack[int]()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
		if stack.IsEmpty() {
			for j := 0; j < 1000; j++ {
				stack.Push(j)
			}
		}
	}
}

func BenchmarkSetAdd(b *testing.B) {
	set := NewSet[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkCacheSet(b *testing.B) {
	cache := NewCache[int, string]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Set(i, "value")
	}
}

func BenchmarkCacheGet(b *testing.B) {
	cache := NewCache[int, string]()
	for i := 0; i < 1000; i++ {
		cache.Set(i, "value")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

