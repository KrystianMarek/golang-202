// Package runner provides utilities for running examples.
package runner

import (
	"fmt"
	"time"
)

// Example represents a runnable example.
type Example struct {
	Name        string
	Description string
	Run         func()
}

// Runner manages and executes examples.
type Runner struct {
	examples []Example
}

// NewRunner creates a new runner.
func NewRunner() *Runner {
	return &Runner{
		examples: make([]Example, 0),
	}
}

// Register registers an example.
func (r *Runner) Register(example Example) {
	r.examples = append(r.examples, example)
}

// Run runs a specific example by name.
func (r *Runner) Run(name string) bool {
	for _, ex := range r.examples {
		if ex.Name == name {
			r.runExample(ex)
			return true
		}
	}
	return false
}

// RunAll runs all registered examples.
func (r *Runner) RunAll() {
	for i, ex := range r.examples {
		r.runExample(ex)
		if i < len(r.examples)-1 {

		}
	}
}

func (r *Runner) runExample(ex Example) {
	fmt.Printf("Running: %s\n", ex.Name)
	if ex.Description != "" {
		fmt.Printf("Description: %s\n", ex.Description)
	}


	start := time.Now()
	ex.Run()
	duration := time.Since(start)

	fmt.Printf("\nCompleted in: %v\n", duration)
}

// List lists all registered examples.
func (r *Runner) List() {
	fmt.Println("Available examples:")
	for i, ex := range r.examples {
		fmt.Printf("%d. %s", i+1, ex.Name)
		if ex.Description != "" {
			fmt.Printf(" - %s", ex.Description)
		}

	}
}

