package patterns

import (
	"fmt"
	"strings"
)

// Strategy pattern demonstrates selecting algorithms at runtime.
//
// Why? Strategy pattern allows changing behavior at runtime by
// encapsulating algorithms in interchangeable objects.

// PaymentStrategy defines the interface for payment algorithms.
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardStrategy implements credit card payment.
type CreditCardStrategy struct {
	CardNumber string
	CVV        string
}

// Pay processes credit card payment.
func (c *CreditCardStrategy) Pay(amount float64) string {
	masked := "****-****-****-" + c.CardNumber[len(c.CardNumber)-4:]
	return fmt.Sprintf("Paid $%.2f using credit card %s", amount, masked)
}

// PayPalStrategy implements PayPal payment.
type PayPalStrategy struct {
	Email string
}

// Pay processes PayPal payment.
func (p *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, p.Email)
}

// CryptoStrategy implements cryptocurrency payment.
type CryptoStrategy struct {
	WalletAddress string
}

// Pay processes crypto payment.
func (c *CryptoStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using crypto wallet %s",
		amount, c.WalletAddress[:10]+"...")
}

// ShoppingCart uses a payment strategy.
type ShoppingCart struct {
	items           []string
	total           float64
	paymentStrategy PaymentStrategy
}

// NewShoppingCart creates a shopping cart.
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: make([]string, 0),
		total: 0,
	}
}

// AddItem adds an item to the cart.
func (s *ShoppingCart) AddItem(item string, price float64) {
	s.items = append(s.items, item)
	s.total += price
}

// SetPaymentStrategy sets the payment strategy.
func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	s.paymentStrategy = strategy
}

// Checkout processes the payment.
func (s *ShoppingCart) Checkout() string {
	if s.paymentStrategy == nil {
		return "No payment method selected"
	}
	result := s.paymentStrategy.Pay(s.total)
	fmt.Printf("Items: %v\n", s.items)
	return result
}

// CompressionStrategy defines compression algorithms.
type CompressionStrategy interface {
	Compress(data string) string
}

// ZipCompression implements ZIP compression.
type ZipCompression struct{}

// Compress simulates ZIP compression.
func (z *ZipCompression) Compress(data string) string {
	return fmt.Sprintf("[ZIP:%s]", strings.ToUpper(data))
}

// RarCompression implements RAR compression.
type RarCompression struct{}

// Compress simulates RAR compression.
func (r *RarCompression) Compress(data string) string {
	return fmt.Sprintf("[RAR:%s]", strings.ToLower(data))
}

// FileCompressor uses compression strategies.
type FileCompressor struct {
	strategy CompressionStrategy
}

// NewFileCompressor creates a compressor.
func NewFileCompressor(strategy CompressionStrategy) *FileCompressor {
	return &FileCompressor{strategy: strategy}
}

// SetStrategy changes the compression strategy.
func (f *FileCompressor) SetStrategy(strategy CompressionStrategy) {
	f.strategy = strategy
}

// Compress compresses data using the current strategy.
func (f *FileCompressor) Compress(filename, data string) string {
	compressed := f.strategy.Compress(data)
	return fmt.Sprintf("File %s: %s", filename, compressed)
}

// SortStrategy defines sorting algorithms.
type SortStrategy interface {
	Sort(data []int) []int
	Name() string
}

// BubbleSort implements bubble sort.
type BubbleSort struct{}

// Sort performs bubble sort.
func (b *BubbleSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// Name returns the algorithm name.
func (b *BubbleSort) Name() string {
	return "Bubble Sort"
}

// QuickSort implements quick sort.
type QuickSort struct{}

// Sort performs quick sort.
func (q *QuickSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	q.quickSort(result, 0, len(result)-1)
	return result
}

func (q *QuickSort) quickSort(arr []int, low, high int) {
	if low < high {
		pi := q.partition(arr, low, high)
		q.quickSort(arr, low, pi-1)
		q.quickSort(arr, pi+1, high)
	}
}

func (q *QuickSort) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Name returns the algorithm name.
func (q *QuickSort) Name() string {
	return "Quick Sort"
}

// Sorter uses sorting strategies.
type Sorter struct {
	strategy SortStrategy
}

// NewSorter creates a sorter.
func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

// Sort sorts data using the current strategy.
func (s *Sorter) Sort(data []int) {
	fmt.Printf("Using %s\n", s.strategy.Name())
	sorted := s.strategy.Sort(data)
	fmt.Printf("Input: %v\n", data)
	fmt.Printf("Sorted: %v\n", sorted)
}

// ExampleStrategy demonstrates the Strategy pattern.
func ExampleStrategy() {
	fmt.Println("=== Strategy Pattern ===")

	// Payment strategies
	cart := NewShoppingCart()
	cart.AddItem("Laptop", 999.99)
	cart.AddItem("Mouse", 29.99)

	cart.SetPaymentStrategy(&CreditCardStrategy{
		CardNumber: "1234567890123456",
		CVV:        "123",
	})
	fmt.Println(cart.Checkout())

	cart2 := NewShoppingCart()
	cart2.AddItem("Book", 19.99)
	cart2.SetPaymentStrategy(&PayPalStrategy{Email: "user@example.com"})
	fmt.Println(cart2.Checkout())

	// Compression strategies
	compressor := NewFileCompressor(&ZipCompression{})
	fmt.Println(compressor.Compress("data.txt", "Hello World"))

	compressor.SetStrategy(&RarCompression{})
	fmt.Println(compressor.Compress("archive.txt", "Hello World"))

	// Sorting strategies
	data := []int{64, 34, 25, 12, 22, 11, 90}

	sorter := NewSorter(&BubbleSort{})
	sorter.Sort(data)

	sorter = NewSorter(&QuickSort{})
	sorter.Sort(data)
}
