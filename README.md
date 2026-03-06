# Generic - Go Generics Utilities Library

[![Go Version](https://img.shields.io/badge/go-1.18%2B-blue)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-green)](#license)

`generic` is a lightweight, production-ready Go library that provides reusable generic data structures, type utility functions, and concurrency patterns. Leverage Go 1.18+ generics to write type-safe, efficient code without casting or reflection.

## 🎯 Features

- **Generic Data Structures** - Type-safe FIFO queue with thread-safe operations
- **Generic Utility Functions** - Pointer handling, map operations, and type conversions without reflection
- **Generic Concurrency Patterns** - Channel multiplexing with the FanIn pattern
- **Zero Dependencies** - Pure Go standard library only
- **Production Ready** - Fully tested with comprehensive test coverage
- **Type Safe** - Leverage Go 1.18+ generics for compile-time type safety

## 📦 Installation

```bash
go get github.com/DmitriyVTitov/generic
```

Requires Go 1.18 or later.

## 📚 API Reference

### Datastructures

**Queue[T]** - Thread-safe FIFO queue for any type T.

- **`NewQueue[T]() *Queue[T]`** - Create a new queue
- **`Enqueue(el *T)`** - Add element to queue tail
- **`Dequeue() *T`** - Remove and return first element (nil if empty)
- **`Len() int`** - Return number of elements
- **`All() []*T`** - Get all elements and clear queue

### Functions

- **`Val[T](ptr *T) T`** - Dereference pointer or return zero value if nil
- **`Ptr[T](val T) *T`** - Create pointer to value
- **`MapItem[T, K, V](m map[K]V, key K) T`** - Safely retrieve and cast map values

### Patterns

- **`FanIn[T](channels ...<-chan T) <-chan T`** - Multiplex channels into single output

## 🚀 Quick Start

### Queue - Thread-Safe FIFO Data Structure

```go
package main

import (
	"fmt"
	"github.com/DmitriyVTitov/generic"
)

func main() {
	// Create a queue for integers
	q := generic.NewQueue[int]()

	// Enqueue elements
	val1, val2, val3 := 10, 20, 30
	q.Enqueue(&val1)
	q.Enqueue(&val2)
	q.Enqueue(&val3)

	// Dequeue elements
	first := q.Dequeue()
	fmt.Println(*first) // Output: 10

	// Check queue length
	fmt.Println(q.Len()) // Output: 2

	// Get all remaining elements
	remaining := q.All()
	fmt.Println(len(remaining)) // Output: 0 (queue is cleared)
}
```

### Utility Functions - Pointer and Value Handling

```go
package main

import (
	"github.com/DmitriyVTitov/generic"
)

func main() {
	// Val - safely dereference pointers
	ptr := (*int)(nil)
	value := generic.Val(ptr) // Returns 0 (zero value)

	// Ptr - create pointers to values
	num := 42
	pointerToNum := generic.Ptr(num)

	// Also works with function results that don't support &
	pointerToLiteral := generic.Ptr(100)
}
```

### Map Operations - Type-Safe Map Access

```go
package main

import (
	"github.com/DmitriyVTitov/generic"
)

func main() {
	m := map[string]interface{}{
		"count": 42,
		"name":  "example",
	}

	// Safely retrieve and cast values
	count := generic.MapItem[int](m, "count")     // Returns 42
	missing := generic.MapItem[int](m, "missing") // Returns 0

	// Works with any type
	name := generic.MapItem[string](m, "name") // Returns "example"
}
```

### FanIn Pattern - Multiplex Multiple Channels

```go
package main

import (
	"fmt"
	"github.com/DmitriyVTitov/generic"
)

func main() {
	// Create multiple source channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "message 1"
		close(ch1)
	}()

	go func() {
		ch2 <- "message 2"
		close(ch2)
	}()

	// Multiplex into single channel
	combined := generic.FanIn(ch1, ch2)

	for msg := range combined {
		fmt.Println(msg) // Receives from both channels
	}
}
```

## 🔒 Thread Safety

The `Queue` implementation is fully thread-safe using mutex locks for all operations. It's safe to use across multiple goroutines without additional synchronization.

## 💡 Why Use Generics?

Go 1.18+ generics enable:

- **Type Safety** - Catch type errors at compile time, not runtime
- **No Reflection** - Generic code is as fast as hand-written type-specific code
- **No Type Assertions** - Eliminate error-prone manual casting
- **Code Reuse** - Write once, use with any type

## 📄 License

MIT License - See LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## 🔗 Related Packages

- [Go Generics Documentation](https://go.dev/blog/intro-generics)

## Keywords

Go generics, generic queue, thread-safe data structures, Go 1.18, type-safe utilities, concurrent programming, Go patterns, channel multiplexing, FIFO queue, Go utilities library
