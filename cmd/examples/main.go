package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KrystianMarek/golang-202/pkg/functional"
	"github.com/KrystianMarek/golang-202/pkg/go124"
	"github.com/KrystianMarek/golang-202/pkg/idioms"
	"github.com/KrystianMarek/golang-202/pkg/oop"
	"github.com/KrystianMarek/golang-202/pkg/oop/patterns"
)

func main() {
	if len(os.Args) > 1 {
		runExample(os.Args[1])
		return
	}

	fmt.Println("==========================================================")
	fmt.Println("    GoLang-202: Advanced Go Patterns & Features")
	fmt.Println("==========================================================")

	runAllExamples()
}

func runExample(name string) {
	examples := map[string]func(){
		"go124":       runGo124Examples,
		"oop":         runOOPExamples,
		"functional":  runFunctionalExamples,
		"idioms":      runIdiomsExamples,
		"patterns":    runPatternExamples,
	}

	if fn, ok := examples[name]; ok {
		fn()
	} else {
		fmt.Printf("Unknown example: %s\n", name)
		fmt.Println("\nAvailable examples:")
		fmt.Println("  go124      - Go 1.24 features")
		fmt.Println("  oop        - OOP patterns")
		fmt.Println("  functional - Functional programming")
		fmt.Println("  idioms     - Go idioms")
		fmt.Println("  patterns   - Design patterns")
	}
}

func runAllExamples() {
	runGo124Examples()
	separator()

	runOOPExamples()
	separator()

	runFunctionalExamples()
	separator()

	runIdiomsExamples()
	separator()

	runPatternExamples()
}

func runGo124Examples() {
	header("Go 1.24 Features")
	go124.ExampleIterators()

	go124.ExampleUnique()

	go124.ExampleCleanup()

	go124.ExampleGenericAliases()

	go124.ExampleGenerics()
}

func runOOPExamples() {
	header("Object-Oriented Programming")
	oop.ExampleComposition()
}

func runFunctionalExamples() {
	header("Functional Programming")
	functional.ExampleHigherOrder()

	functional.ExampleImmutability()

	functional.ExamplePipelines()
}

func runIdiomsExamples() {
	header("Go Idioms")
	idioms.ExampleInterfaces()

	idioms.ExampleErrors()

	idioms.ExampleConcurrency()

	idioms.ExampleChannels()

	idioms.ExampleZeroValues()
}

func runPatternExamples() {
	header("Design Patterns")
	patterns.ExampleSingleton()

	patterns.ExampleFactory()

	patterns.ExampleBuilder()

	patterns.ExampleObserver()

	patterns.ExampleGenericObserver()

	patterns.ExampleAdapter()

	patterns.ExampleDecorator()

	patterns.ExampleStrategy()
}

func header(title string) {

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  %s\n", title)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func separator() {
	fmt.Println("\n" + strings.Repeat("─", 60) + "")
}

