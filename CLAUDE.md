# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`underscore` is a Go library providing functional programming helpers inspired by underscore.js, built on Go 1.18+ generics. The library is organized as a flat structure with individual files for each function, plus a `maps` subpackage for map-specific utilities.

## Development Commands

### Testing

```sh
# Run all tests (local)
go test ./...

# Run all tests with coverage (local)
go test ./... -coverpkg=./... -coverprofile cov.out -covermode=count
go tool cover -func cov.out
rm cov.out

# Run tests in Docker (preferred for CI/validation)
make test

# Run a single test
go test -run TestFunctionName

# Run tests for a specific file
go test -run TestMap
```

### Building

```sh
# Build Docker image
make build

# Install dependencies
go mod download
```

### Linting & Security

```sh
# Scan Docker image for vulnerabilities
make scan

# Scan config files
make scan-config
```

### Documentation

```sh
# Serve docs locally at http://localhost:1313
make docs

# Build static docs
make build-docs
```

## Architecture

### Code Organization

The library uses a **flat structure** where each function is implemented in its own file:

- `<function>.go` - implementation
- `<function>_test.go` - tests

Example: `filter.go` + `filter_test.go`, `map.go` + `map_test.go`

### Core Patterns

**Generic Functions**: Most functions use Go generics with constraints from `cmp.Ordered` or custom type parameters. Functions operate on slices and return new slices (immutable style).

**Pipe Chain**: The `Pipe[T]` struct enables method chaining for ordered types. Methods that return slices continue the chain, while methods that return values (like `All`, `Any`, `Reduce`) break the chain and return the final value.

```go
// pipe.go defines Pipe[T cmp.Ordered]
// Chain-continuing: Filter, Map
// Chain-breaking: All, Any, Reduce, Min, Max, Partition, Find, Each
```

**Concurrency Helpers**: `ParallelMap` and `ParallelFilter` use worker pools with:

- Context-based cancellation
- Order preservation (results match input order)
- First-error-wins semantics
- Default workers = GOMAXPROCS if workers <= 0

Implementation detail: Uses `sync.Once` to capture first error and cancel context immediately.

**Subpackages**:

- `maps/` - Map-specific utilities (`Keys`, `Values`, `Map`)
  - Uses type alias `M[K, V] = map[K]V` for cleaner signatures
  - `Map` function allows transforming map entries

### Testing Conventions

- Use `testify/assert` for assertions
- Test file names match source files with `_test.go` suffix
- Table-driven tests are common (see `map_test.go`, `filter_test.go`)
- Internal tests (using `package underscore` rather than `package underscore_test`) are used sparingly for testing unexported functions

## Key Constraints

- **Minimum Go version**: 1.24.2 (see go.mod)
- **Generic constraints**: Most collection functions require `cmp.Ordered` types; some use `comparable` or no constraints
- **Order preservation**: `ParallelMap` and `ParallelFilter` guarantee output order matches input order
- **No mutation**: Functions return new slices; `UniqueInPlace` is the exception (in-place deduplication)

## Known Limitations

### Recently Fixed (2025-11-14)

1. ✅ **Filter allocation** - Now pre-allocates with `make([]T, 0, len(values))` (90% fewer allocations)
2. ✅ **OrderBy algorithm** - Replaced bubble sort with `slices.SortFunc` (629x faster for large datasets)
3. ✅ **Partition allocation** - Now pre-allocates both result slices
4. ✅ **Max/Min empty slices** - Now panics with clear message: "underscore.Max: empty slice"
5. ✅ **Drop semantics** - Fixed to drop first N elements (breaking change). Old behavior available as `RemoveAt`

### API Design Issues

1. **Pipe constraint** (`pipe.go:7`) - `Pipe[T cmp.Ordered]` prevents usage with custom types
2. **Last panics** (`last.go:5-8`) - No empty slice handling

### Missing Features

Popular FP utilities not yet implemented: `TakeWhile`, `DropWhile`, `Scan`, `First/FirstN`, `Init`, `Intersperse`, `Sliding`, `FoldRight`, `Tap`, `Transpose`, `Unzip`, `ParallelReduce`, `Replicate`

## Performance Characteristics

### Good Performance Patterns
- `Filter` pre-allocates: `make([]T, 0, len(values))` ✅ (Fixed 2025-11-14)
- `Map` pre-allocates: `make([]P, 0, len(values))`
- `Partition` pre-allocates: `make([]T, 0, len(values))` for both slices ✅ (Fixed 2025-11-14)
- `Chunk` pre-calculates capacity: `(l+n-1)/n`
- `ParallelFilter` pre-counts before allocation
- `OrderBy` uses `slices.SortFunc`: O(n log n) performance ✅ (Fixed 2025-11-14)

### Remaining Performance Issues
- `Flatmap`: Accumulation overhead from repeated appends
- `GroupBy`: Map initialized with capacity 0 (useless hint)

### When to Use ParallelMap vs Map

Use `ParallelMap` when:
- Processing 100+ elements with expensive operations (>1ms per element)
- Operations are CPU-bound (not I/O-bound with shared resources)
- Order preservation is required
- Context cancellation is needed

Use regular `Map` when:
- Small slices (<100 elements)
- Fast operations (<100µs per element)
- Avoiding goroutine overhead matters
- Simple transformations without error handling

**Worker count guidelines:**
- Default (workers=0): Uses `runtime.GOMAXPROCS(0)` - good starting point
- CPU-bound: Use GOMAXPROCS or GOMAXPROCS*2
- I/O-bound: Can use higher values (10-100) if not sharing resources

## Contributing Notes

When adding new functions:

1. Create both `<function>.go` and `<function>_test.go`
2. Add examples in comments using Go doc format
3. Pre-allocate slices with `make([]T, 0, len(input))` when output size is similar to input
4. Document panic conditions (empty slices, nil inputs, invalid indices)
5. Add edge case tests (empty, single element, nil)
6. If the function applies to Pipe chains, add a method to `pipe.go`
7. Update README.md function list if adding new collection functions
8. Follow SemVer for version numbers
9. Ensure all tests pass: `make test`

When fixing bugs:
- Add regression tests before fixing
- Run benchmarks if performance-related: `go test -bench=. -benchmem`
- Check for similar issues in other functions
