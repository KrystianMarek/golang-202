package go124

import (
	"testing"
	"unique"
)

func TestTreeIterator(t *testing.T) {
	tree := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: 6,
			Left:  &TreeNode{Value: 5},
			Right: &TreeNode{Value: 7},
		},
	}

	// Test in-order traversal
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	var result []int
	for val := range tree.InOrder() {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d values, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestRange(t *testing.T) {
	var result []int
	for val := range Range(0, 5) {
		result = append(result, val)
	}

	expected := []int{0, 1, 2, 3, 4}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d values, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	evens := Filter(Range(0, 10), func(n int) bool { return n%2 == 0 })
	var result []int
	for val := range evens {
		result = append(result, val)
	}

	expected := []int{0, 2, 4, 6, 8}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d values, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestUniqueHandles(t *testing.T) {
	h1 := unique.Make("test")
	h2 := unique.Make("test")
	h3 := unique.Make("other")

	if h1 != h2 {
		t.Error("Expected same handles for equal strings")
	}

	if h1 == h3 {
		t.Error("Expected different handles for different strings")
	}

	if h1.Value() != "test" {
		t.Errorf("Expected 'test', got '%s'", h1.Value())
	}
}

func TestLogAggregator(t *testing.T) {
	agg := NewLogAggregator()
	agg.AddLog("ERROR", "Test error", "test-service")
	agg.AddLog("INFO", "Test info", "test-service")

	logs := agg.GetLogs()
	if len(logs) != 2 {
		t.Fatalf("Expected 2 logs, got %d", len(logs))
	}
}

func TestOptional(t *testing.T) {
	some := Some(42)
	none := None[int]()

	if !some.IsPresent() {
		t.Error("Expected Some to be present")
	}

	if some.Get() != 42 {
		t.Errorf("Expected 42, got %d", some.Get())
	}

	if none.IsPresent() {
		t.Error("Expected None to not be present")
	}
}

func TestPair(t *testing.T) {
	pair := NewPair("key", 123)

	if pair.First != "key" {
		t.Errorf("Expected 'key', got '%s'", pair.First)
	}

	if pair.Second != 123 {
		t.Errorf("Expected 123, got %d", pair.Second)
	}

	swapped := pair.Swap()
	if swapped.First != 123 || swapped.Second != "key" {
		t.Error("Swap failed")
	}
}

func BenchmarkIteratorVsSlice(b *testing.B) {
	tree := &TreeNode{
		Value: 50,
		Left:  buildTree(1, 49),
		Right: buildTree(51, 100),
	}

	b.Run("Iterator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			for val := range tree.InOrder() {
				sum += val
			}
		}
	})

	b.Run("Materialized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			values := materializeTree(tree)
			sum := 0
			for _, val := range values {
				sum += val
			}
		}
	})
}

func buildTree(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &TreeNode{
		Value: mid,
		Left:  buildTree(start, mid-1),
		Right: buildTree(mid+1, end),
	}
}

func materializeTree(t *TreeNode) []int {
	var result []int
	for val := range t.InOrder() {
		result = append(result, val)
	}
	return result
}
