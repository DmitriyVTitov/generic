# Generic - Go Generics Utilities Library

[![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)](https://go.dev)
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
go install github.com/DmitriyVTitov/generic@latest
```

Requires Go 1.21 or later.

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

### Slices

- **`RemoveDuplicates[S ~[]E, E comparable](s S) S`** - Remove duplicates in place for comparable elements
- **`RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S`** - Remove duplicates using a custom equality function
- **`ContainsDuplicates[S ~[]E, E comparable](s S) bool`** - Check whether a slice contains duplicate elements
- **`ContainsDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) bool`** - Check duplicates using a custom equality function
- **`GetDuplicateIndexes[S ~[]E, E comparable](s S) map[int][]int`** - Get duplicate indexes for each element
- **`GetDuplicateIndexesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) map[int][]int`** - Get duplicate indexes using a custom equality function
- **`RemoveDuplicatesByComparableKey[T any, K comparable](slice []T, fn func(T) K) []T`** - Remove duplicates by derived comparable key

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
