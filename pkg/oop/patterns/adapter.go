package patterns

import "fmt"

// Adapter pattern demonstrates how to make incompatible interfaces work together.
//
// Why? Adapters allow legacy code or third-party libraries with different
// interfaces to work with your system without modifying their source code.

// MediaPlayer is the target interface our client code expects.
type MediaPlayer interface {
	Play(filename string) error
}

// AudioPlayer plays audio files using the MediaPlayer interface.
type AudioPlayer struct{}

// Play plays an audio file.
func (a *AudioPlayer) Play(filename string) error {
	fmt.Printf("Playing audio file: %s\n", filename)
	return nil
}

// LegacyVideoPlayer is an old interface we need to adapt.
type LegacyVideoPlayer struct{}

// PlayVideo has a different method signature.
func (l *LegacyVideoPlayer) PlayVideo(videoFile, format string) {
	fmt.Printf("Playing %s video: %s\n", format, videoFile)
}

// VideoPlayerAdapter adapts LegacyVideoPlayer to MediaPlayer interface.
type VideoPlayerAdapter struct {
	legacyPlayer *LegacyVideoPlayer
	format       string
}

// NewVideoPlayerAdapter creates an adapter.
func NewVideoPlayerAdapter(format string) *VideoPlayerAdapter {
	return &VideoPlayerAdapter{
		legacyPlayer: &LegacyVideoPlayer{},
		format:       format,
	}
}

// Play implements MediaPlayer interface.
func (v *VideoPlayerAdapter) Play(filename string) error {
	v.legacyPlayer.PlayVideo(filename, v.format)
	return nil
}

// ThirdPartyPayment is an external payment service with different interface.
type ThirdPartyPayment struct{}

// ProcessTransaction is the third-party method.
func (t *ThirdPartyPayment) ProcessTransaction(amount float64, currency, account string) bool {
	fmt.Printf("Third-party processing: %.2f %s to account %s\n", amount, currency, account)
	return true
}

// PaymentProcessor is our expected interface.
type PaymentProcessor interface {
	Pay(amount float64, recipient string) error
}

// PaymentAdapter adapts ThirdPartyPayment to PaymentProcessor.
type PaymentAdapter struct {
	thirdParty *ThirdPartyPayment
	currency   string
}

// NewPaymentAdapter creates a payment adapter.
func NewPaymentAdapter(currency string) *PaymentAdapter {
	return &PaymentAdapter{
		thirdParty: &ThirdPartyPayment{},
		currency:   currency,
	}
}

// Pay implements PaymentProcessor interface.
func (p *PaymentAdapter) Pay(amount float64, recipient string) error {
	success := p.thirdParty.ProcessTransaction(amount, p.currency, recipient)
	if !success {
		return fmt.Errorf("payment failed")
	}
	return nil
}

// OldLogger is a legacy logging system.
type OldLogger struct{}

// WriteLog has a different signature.
func (o *OldLogger) WriteLog(level int, msg string) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
	if level < len(levels) {
		fmt.Printf("[OLD][%s] %s\n", levels[level], msg)
	}
}

// Logger is our new logging interface.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

// LoggerAdapter adapts OldLogger to new Logger interface.
type LoggerAdapter struct {
	oldLogger *OldLogger
}

// NewLoggerAdapter creates a logger adapter.
func NewLoggerAdapter() *LoggerAdapter {
	return &LoggerAdapter{
		oldLogger: &OldLogger{},
	}
}

// Debug logs debug message.
func (l *LoggerAdapter) Debug(msg string) {
	l.oldLogger.WriteLog(0, msg)
}

// Info logs info message.
func (l *LoggerAdapter) Info(msg string) {
	l.oldLogger.WriteLog(1, msg)
}

// Error logs error message.
func (l *LoggerAdapter) Error(msg string) {
	l.oldLogger.WriteLog(3, msg)
}

// TemperatureSensor is an old sensor using Fahrenheit.
type TemperatureSensor struct{}

// ReadFahrenheit returns temperature in Fahrenheit.
func (t *TemperatureSensor) ReadFahrenheit() float64 {
	return 68.0 // Simulated reading
}

// CelsiusReader is the interface we want.
type CelsiusReader interface {
	ReadCelsius() float64
}

// TempSensorAdapter adapts Fahrenheit to Celsius.
type TempSensorAdapter struct {
	sensor *TemperatureSensor
}

// NewTempSensorAdapter creates a temperature adapter.
func NewTempSensorAdapter() *TempSensorAdapter {
	return &TempSensorAdapter{
		sensor: &TemperatureSensor{},
	}
}

// ReadCelsius converts and returns Celsius.
func (t *TempSensorAdapter) ReadCelsius() float64 {
	fahrenheit := t.sensor.ReadFahrenheit()
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

// ExampleAdapter demonstrates the Adapter pattern.
func ExampleAdapter() {
	fmt.Println("=== Adapter Pattern ===")

	// Media player adapters
	players := []MediaPlayer{
		&AudioPlayer{},
		NewVideoPlayerAdapter("MP4"),
		NewVideoPlayerAdapter("AVI"),
	}

	files := []string{"song.mp3", "movie.mp4", "video.avi"}
	for i, player := range players {
		_ = player.Play(files[i])
	}

	// Payment adapter
	paymentProcessor := NewPaymentAdapter("USD")
	err := paymentProcessor.Pay(99.99, "merchant-123")
	if err != nil {
		fmt.Printf("Payment error: %v\n", err)
	}

	// Logger adapter
	logger := NewLoggerAdapter()
	logger.Debug("Application started")
	logger.Info("Processing request")
	logger.Error("An error occurred")

	// Temperature sensor adapter
	tempReader := NewTempSensorAdapter()
	celsius := tempReader.ReadCelsius()
	fmt.Printf("Temperature: %.1f°C\n", celsius)
}
