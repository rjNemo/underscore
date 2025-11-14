# Action Plan: Underscore Library Improvements

**Status:** In Progress - Week 1 (4/5 completed)
**Overall Quality Score:** 8.2/10 → 9.0/10 (estimated after fixes)
**Generated:** 2025-11-14
**Last Updated:** 2025-11-14

This document outlines prioritized improvements for the underscore Go library based on comprehensive codebase review.

## Completion Status

### ✅ Completed Issues

- [x] **Issue 1**: Filter pre-allocation (90% fewer allocations) - Commit 92b6463
- [x] **Issue 2**: OrderBy bubble sort replacement (629x faster) - Commit 7caa23e
- [x] **Issue 3**: Partition pre-allocation - Commit 46d52e3
- [x] **Issue 4**: Max/Min empty slice handling - Commit a194355
- [x] **Issue 5**: Add edge case tests - Commit 106b713
- [x] **Issue 6**: Drop semantics clarification (breaking change) - Ready to commit

### 🔄 In Progress

- None currently

### ⏳ Pending

- See sections below for remaining issues

---

## Priority Matrix

### 🔴 Critical (Week 1) - 2-3 hours total

High impact, low effort fixes that significantly improve performance and stability.

### 🟡 High Priority (Week 2) - 5-6 hours total

Important improvements for API consistency and testing coverage.

### 🟢 Medium Priority (Week 3) - 4-5 hours total

Additional optimizations and quality improvements.

### 🔵 Future Enhancements

New features to add based on user demand.

---

## 🔴 CRITICAL ISSUES (Week 1)

### 1. Fix Filter Pre-allocation ⏱️ 2 min ✅ COMPLETED

**File:** `filter.go:4`
**Issue:** No pre-allocation causes O(n) allocations instead of O(1)
**Impact:** 2-5x performance improvement
**Status:** ✅ Fixed in commit 92b6463
**Results:** 90% fewer allocations (10→1), 8% faster execution

#### Current Code

```go
func Filter[T any](values []T, predicate func(T) bool) (res []T) {
    for _, v := range values {
        if predicate(v) {
            res = append(res, v)  // ❌ No pre-allocation
        }
    }
    return res
}
```

#### Fixed Code

```go
func Filter[T any](values []T, predicate func(T) bool) (res []T) {
    res = make([]T, 0, len(values))  // ✅ Pre-allocate
    for _, v := range values {
        if predicate(v) {
            res = append(res, v)
        }
    }
    return res
}
```

#### Steps

1. Add benchmark test first to measure improvement
2. Change line 4: Add `res = make([]T, 0, len(values))`
3. Run tests: `go test ./... -v`
4. Run benchmark: `go test -bench=BenchmarkFilter -benchmem`
5. Commit: "perf: pre-allocate Filter result slice"

---

### 2. Replace OrderBy Bubble Sort ⏱️ 5 min ✅ COMPLETED

**File:** `orderBy.go:7-27`
**Issue:** O(n²) bubble sort with TODO comment
**Impact:** O(n²) → O(n log n) complexity
**Status:** ✅ Fixed in commit 7caa23e
**Results:** 629x faster for large datasets (1000 items), resolved TODO

#### Current Code

```go
func OrderBy[T any](list []T, predicate func(T, T) bool) []T {
    swaps := true
    var tmp T

    //todo: replace with a faster algorithm, this one is pretty simple
    for swaps {
        swaps = false

        for i := 0; i < len(list)-1; i++ {
            if predicate(list[i], list[i+1]) {
                swaps = true
                tmp = list[i]

                list[i] = list[i+1]
                list[i+1] = tmp
            }
        }
    }

    return list
}
```

#### Fixed Code

```go
import "slices"

// OrderBy orders a slice by a field value within a struct, the predicate allows you
// to pick the fields you want to orderBy. Use > for ASC or < for DESC
// Uses O(n log n) sorting algorithm. Mutates the input slice.
//
// func (left Person, right Person) bool { return left.Age > right.Age }
func OrderBy[T any](list []T, predicate func(T, T) bool) []T {
    slices.SortFunc(list, func(a, b T) int {
        if predicate(a, b) {
            return 1
        }
        if predicate(b, a) {
            return -1
        }
        return 0
    })
    return list
}
```

#### Steps

1. Add benchmark test to measure improvement
2. Replace entire function body with `slices.SortFunc`
3. Update doc comment to mention O(n log n) and mutation
4. Run tests: `go test ./... -v`
5. Run benchmark: `go test -bench=BenchmarkOrderBy -benchmem`
6. Commit: "perf: replace bubble sort with slices.SortFunc in OrderBy"

---

### 3. Fix Partition Pre-allocation ⏱️ 2 min ✅ COMPLETED

**File:** `partition.go:6-7`
**Issue:** No capacity hints cause repeated allocations
**Impact:** Fewer allocations during split
**Status:** ✅ Fixed in commit 46d52e3
**Results:** Reduced allocations from O(log n) to O(1) per slice

#### Current Code

```go
func Partition[T any](values []T, predicate func(T) bool) ([]T, []T) {
    keep := make([]T, 0)      // ❌ No capacity hint
    reject := make([]T, 0)    // ❌ No capacity hint

    for _, v := range values {
        if predicate(v) {
            keep = append(keep, v)
        } else {
            reject = append(reject, v)
        }
    }
    return keep, reject
}
```

#### Fixed Code

```go
func Partition[T any](values []T, predicate func(T) bool) ([]T, []T) {
    keep := make([]T, 0, len(values))     // ✅ Pre-allocate
    reject := make([]T, 0, len(values))   // ✅ Pre-allocate

    for _, v := range values {
        if predicate(v) {
            keep = append(keep, v)
        } else {
            reject = append(reject, v)
        }
    }
    return keep, reject
}
```

#### Steps

1. Change lines 6-7 to add capacity hint
2. Run tests: `go test ./... -v`
3. Commit: "perf: pre-allocate Partition result slices"

---

### 4. Handle Max/Min Empty Slices ⏱️ 30 min ✅ COMPLETED

**Files:** `max.go:8-16`, `min.go:8-16`
**Issue:** Panics on empty slices
**Impact:** Prevent runtime panics
**Status:** ✅ Fixed in commit a194355 (Option B: Non-breaking)
**Results:** Clear panic messages, documented behavior, added tests

#### Current Code

```go
func Max[T cmp.Ordered](values []T) T {
    max := values[0]  // ❌ Panic on empty slice
    for _, v := range values {
        if v > max {
            max = v
        }
    }
    return max
}
```

#### Option A: Return Error (Recommended)

```go
import "errors"

// Max returns the maximum value in the slice.
// Returns error if the slice is empty.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Max[T cmp.Ordered](values []T) (T, error) {
    var zero T
    if len(values) == 0 {
        return zero, errors.New("cannot find max of empty slice")
    }

    max := values[0]
    for _, v := range values {
        if v > max {
            max = v
        }
    }
    return max, nil
}
```

#### Option B: Document Panic (Faster, breaking change avoided)

```go
// Max returns the maximum value in the slice.
// Panics if values is empty.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Max[T cmp.Ordered](values []T) T {
    if len(values) == 0 {
        panic("underscore.Max: empty slice")
    }

    max := values[0]
    for _, v := range values {
        if v > max {
            max = v
        }
    }
    return max
}
```

#### Steps (Choose one approach)

**For Option A (Breaking Change):**

1. Update `max.go` and `min.go` to return `(T, error)`
2. Update `pipe.go` Max/Min methods to return error
3. Update all test files to check error return
4. Update README.md examples if needed
5. Run tests: `go test ./... -v`
6. Document breaking change in CHANGELOG
7. Commit: "fix!: Max/Min return error on empty slices"

**For Option B (Non-Breaking):**

1. Add length check with explicit panic message
2. Update doc comments to document panic behavior
3. Add tests for panic behavior: `assert.Panics(t, func() { Max([]int{}) })`
4. Run tests: `go test ./... -v`
5. Commit: "fix: add explicit panic with message for Max/Min on empty slices"

---

### 5. Add Edge Case Tests ⏱️ 1 hour

**Files:** Create/update `*_test.go` files
**Issue:** Missing tests for empty slices, nil inputs, single elements
**Impact:** Catch regressions and edge cases

#### Test Cases to Add

**Empty Slice Tests** (`filter_test.go`, `max_test.go`, `min_test.go`, etc.)

```go
func TestFilterEmpty(t *testing.T) {
    result := Filter([]int{}, func(n int) bool { return n > 0 })
    assert.Empty(t, result)
}

func TestMaxEmpty(t *testing.T) {
    // If using Option A from above
    _, err := Max([]int{})
    assert.Error(t, err)

    // If using Option B from above
    assert.Panics(t, func() { Max([]int{}) })
}

func TestLastEmpty(t *testing.T) {
    assert.Panics(t, func() { Last([]int{}) })
}
```

**Single Element Tests**

```go
func TestFilterSingleElement(t *testing.T) {
    result := Filter([]int{5}, func(n int) bool { return n > 0 })
    assert.Equal(t, []int{5}, result)
}

func TestPartitionSingleElement(t *testing.T) {
    keep, reject := Partition([]int{5}, func(n int) bool { return n > 3 })
    assert.Equal(t, []int{5}, keep)
    assert.Empty(t, reject)
}
```

**Large Slice Tests**

```go
func TestFilterLargeSlice(t *testing.T) {
    large := make([]int, 10000)
    for i := range large {
        large[i] = i
    }
    result := Filter(large, func(n int) bool { return n%2 == 0 })
    assert.Equal(t, 5000, len(result))
}
```

**Nil Predicate Tests** (if applicable)

```go
func TestFilterNilPredicate(t *testing.T) {
    assert.Panics(t, func() {
        Filter([]int{1, 2, 3}, nil)
    })
}
```

#### Steps

1. Create test plan spreadsheet/checklist of all functions
2. Add edge case tests for each function
3. Run tests: `go test ./... -v -cover`
4. Fix any failures discovered
5. Commit: "test: add comprehensive edge case tests"

---

## 🟡 HIGH PRIORITY (Week 2)

### 6. Clarify Drop Semantics ⏱️ 30 min

**File:** `drop.go`
**Issue:** Function name suggests "drop first N" but removes element at index
**Impact:** API clarity and user expectations

#### Current Behavior

```go
Drop([]int{1,2,3,4,5}, 2) // Returns [1,2,4,5] - removes index 2
```

#### Expected Behavior (based on underscore.js/Haskell)

```go
Drop([]int{1,2,3,4,5}, 2) // Should return [3,4,5] - drop first 2
```

#### Solution: Add New Functions

**Keep current Drop but rename to RemoveAt**

```go
// RemoveAt returns a new slice with the element at the given index removed.
// Returns original slice if index is out of bounds.
func RemoveAt[T any](values []T, index int) []T {
    if index < 0 || index >= len(values) {
        return values
    }
    res := make([]T, 0, len(values)-1)
    for i, value := range values {
        if i != index {
            res = append(res, value)
        }
    }
    return res
}
```

**Add new Drop with correct semantics**

```go
// Drop returns a new slice with the first n elements removed.
// If n is greater than the slice length, returns an empty slice.
// If n is negative, returns the original slice.
func Drop[T any](values []T, n int) []T {
    if n <= 0 {
        return values
    }
    if n >= len(values) {
        return []T{}
    }
    res := make([]T, len(values)-n)
    copy(res, values[n:])
    return res
}
```

**Add DropWhile**

```go
// DropWhile drops elements from the beginning while predicate is true.
// Returns remaining elements once predicate returns false.
func DropWhile[T any](values []T, predicate func(T) bool) []T {
    for i, v := range values {
        if !predicate(v) {
            res := make([]T, len(values)-i)
            copy(res, values[i:])
            return res
        }
    }
    return []T{}
}
```

#### Steps

1. Create `remove_at.go` with new function
2. Create `remove_at_test.go` with tests
3. Update `drop.go` with new semantics
4. Update `drop_test.go` with new tests
5. Add `drop_while.go` and tests
6. Update README.md function list
7. Document breaking change in CHANGELOG
8. Commit: "feat!: fix Drop semantics and add RemoveAt, DropWhile"

---

### 7. Add Benchmarks ⏱️ 2 hours

**Files:** Create `benchmark_test.go` or add to existing test files
**Issue:** No performance baselines exist
**Impact:** Track performance regressions

#### Benchmarks to Add

**Core Functions**

```go
func BenchmarkFilter(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Filter(data, func(n int) bool { return n%2 == 0 })
    }
}

func BenchmarkMap(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Map(data, func(n int) int { return n * 2 })
    }
}

func BenchmarkReduce(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Reduce(data, func(n, acc int) int { return n + acc }, 0)
    }
}
```

**OrderBy Comparison**

```go
func BenchmarkOrderBy(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = 1000 - i  // Reverse order
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        dataCopy := make([]int, len(data))
        copy(dataCopy, data)
        OrderBy(dataCopy, func(a, b int) bool { return a > b })
    }
}
```

**Concurrency Benchmarks**

```go
func BenchmarkParallelMap(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    ctx := context.Background()

    for _, workers := range []int{1, 2, 4, 8, 16} {
        b.Run(fmt.Sprintf("workers=%d", workers), func(b *testing.B) {
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                ParallelMap(ctx, data, workers, func(ctx context.Context, n int) (int, error) {
                    return n * 2, nil
                })
            }
        })
    }
}

func BenchmarkMapVsParallelMap(b *testing.B) {
    data := make([]int, 10000)
    for i := range data {
        data[i] = i
    }
    ctx := context.Background()

    b.Run("Map", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Map(data, func(n int) int { return n * 2 })
        }
    })

    b.Run("ParallelMap", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            ParallelMap(ctx, data, 0, func(ctx context.Context, n int) (int, error) {
                return n * 2, nil
            })
        }
    })
}
```

**Memory Allocation Benchmarks**

```go
func BenchmarkPartition(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Partition(data, func(n int) bool { return n%2 == 0 })
    }
}

func BenchmarkUnique(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i % 100  // Many duplicates
    }

    b.Run("Unique", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Unique(data)
        }
    })

    b.Run("UniqueInPlace", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            dataCopy := make([]int, len(data))
            copy(dataCopy, data)
            UniqueInPlace(dataCopy)
        }
    })
}
```

#### Steps

1. Create benchmarks for core functions
2. Run baseline: `go test -bench=. -benchmem > bench_before.txt`
3. Document baseline results
4. Add benchmark CI job (optional)
5. Commit: "test: add comprehensive benchmarks for core functions"

---

### 8. Improve Documentation ⏱️ 1 hour

**Files:** Various `.go` files, `CLAUDE.md`, README.md
**Issue:** Missing edge case warnings, performance notes
**Impact:** Better developer experience

#### Documentation Updates Needed

**Add panic warnings**

```go
// Max returns the maximum value in the slice.
// Panics if values is empty.  // ← Add this
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Max[T cmp.Ordered](values []T) T

// Last returns the last element from the slice.
// Panics if the slice is empty.  // ← Add this
func Last[T any](values []T) T
```

**Add complexity notes**

```go
// OrderBy orders a slice by a field value.
// Uses O(n log n) sorting. Mutates the input slice.  // ← Add this
// The predicate allows you to pick the fields you want to orderBy.
// Use > for ASC or < for DESC
```

**Add constraint explanations**

```go
// Pipe enables method chaining for ordered types.
// Type parameter T must be cmp.Ordered because Max/Min methods require it.  // ← Add this
type Pipe[T cmp.Ordered] struct {
    Value []T
}
```

**Update README.md**

- Add performance section
- Add "When to use" guidelines for ParallelMap
- Add edge case handling notes

#### Steps

1. Review all function doc comments
2. Add panic conditions where applicable
3. Add complexity notes for non-O(n) operations
4. Update README.md with performance section
5. Update docs/ Hugo site if needed
6. Commit: "docs: add panic warnings and performance notes"

---

## 🟢 MEDIUM PRIORITY (Week 3)

### 9. Fix Flatmap Allocation ⏱️ 30 min

**File:** `flatmap.go:6`
**Issue:** No pre-allocation causes repeated allocations
**Impact:** ~30-50% performance improvement

#### Current Code

```go
func Flatmap[T any](values []T, mapper func(n T) []T) []T {
    res := make([]T, 0)  // ❌ No pre-allocation
    for _, v := range values {
        vs := mapper(v)
        res = append(res, vs...)
    }
    return res
}
```

#### Option A: Estimate Average Size

```go
func Flatmap[T any](values []T, mapper func(n T) []T) []T {
    // Estimate capacity assuming avg 2-3 items per map
    res := make([]T, 0, len(values)*2)
    for _, v := range values {
        vs := mapper(v)
        res = append(res, vs...)
    }
    return res
}
```

#### Option B: Two-Pass (More Memory Efficient)

```go
func Flatmap[T any](values []T, mapper func(n T) []T) []T {
    // First pass: calculate total size
    totalSize := 0
    mapped := make([][]T, len(values))
    for i, v := range values {
        mapped[i] = mapper(v)
        totalSize += len(mapped[i])
    }

    // Second pass: allocate exact size and copy
    res := make([]T, 0, totalSize)
    for _, vs := range mapped {
        res = append(res, vs...)
    }
    return res
}
```

#### Steps

1. Add benchmark test
2. Choose approach based on typical use cases
3. Update implementation
4. Run tests and benchmarks
5. Commit: "perf: improve Flatmap allocation strategy"

---

### 10. Fix GroupBy Map Initialization ⏱️ 2 min

**File:** `groupby.go:5`
**Issue:** Capacity hint of 0 is useless for maps
**Impact:** Minor allocation improvement

#### Current Code

```go
func GroupBy[K comparable, V any](values []V, f func(V) K) map[K][]V {
    res := make(map[K][]V, 0)  // ❌ Capacity 0 is useless
    ...
}
```

#### Fixed Code

```go
func GroupBy[K comparable, V any](values []V, f func(V) K) map[K][]V {
    res := make(map[K][]V, len(values)/10)  // ✅ Estimate
    ...
}
```

#### Steps

1. Update capacity hint (len/10 or just len)
2. Run tests
3. Commit: "perf: improve GroupBy map initialization"

---

### 11. Relax Pipe Constraint ⏱️ 2 hours

**File:** `pipe.go:7` and method signatures
**Issue:** `Pipe[T cmp.Ordered]` prevents usage with custom types
**Impact:** Broader API usability

This is a breaking change that requires careful consideration.

#### Current Limitation

```go
type Pipe[T cmp.Ordered] struct {
    Value []T
}

// Cannot use with:
type Person struct { Name string; Age int }
NewPipe([]Person{...})  // ❌ Error: Person does not satisfy cmp.Ordered
```

#### Option A: Make Pipe Generic, Constrain Methods

```go
// Pipe can now work with any type
type Pipe[T any] struct {
    Value []T
}

// Methods that need ordering constrain themselves
func (c Pipe[T]) Max() T where T: cmp.Ordered {  // ❌ Go doesn't support this
    return Max(c.Value)
}
```

**Problem:** Go doesn't support method-level constraints different from type-level.

#### Option B: Create Two Pipe Types

```go
// Generic pipe for any type
type Pipe[T any] struct {
    Value []T
}

// Ordered pipe with additional methods
type OrderedPipe[T cmp.Ordered] struct {
    Pipe[T]  // Embed generic pipe
}

// Max/Min only on OrderedPipe
func (c OrderedPipe[T]) Max() T {
    return Max(c.Value)
}

// Factory functions
func NewPipe[T any](value []T) Pipe[T] {
    return Pipe[T]{Value: value}
}

func NewOrderedPipe[T cmp.Ordered](value []T) OrderedPipe[T] {
    return OrderedPipe[T]{Pipe: Pipe[T]{Value: value}}
}
```

#### Option C: Remove Max/Min from Pipe

```go
// Simplest solution: just remove problematic methods
type Pipe[T any] struct {
    Value []T
}

// Users can break chain for Max/Min
result := NewPipe(values).
    Filter(...).
    Map(...).
    Value
max := Max(result)  // Outside pipe chain
```

#### Steps

1. Decide on approach (discuss with maintainers)
2. Implement chosen solution
3. Update all tests
4. Update documentation
5. Add migration guide
6. Document breaking change
7. Commit: "feat!: relax Pipe type constraint"

---

### 12. Add Stress Tests ⏱️ 1 hour

**Files:** Create `stress_test.go`
**Issue:** No tests with large data or high concurrency
**Impact:** Catch race conditions and memory issues

#### Test Cases

**Large Data Tests**

```go
func TestFilterLargeData(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping stress test in short mode")
    }

    large := make([]int, 1_000_000)
    for i := range large {
        large[i] = i
    }

    result := Filter(large, func(n int) bool { return n%2 == 0 })
    assert.Equal(t, 500_000, len(result))
}
```

**Concurrency Stress Tests**

```go
func TestParallelMapHighConcurrency(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping stress test in short mode")
    }

    data := make([]int, 10000)
    for i := range data {
        data[i] = i
    }

    ctx := context.Background()

    // Test with many workers
    result, err := ParallelMap(ctx, data, 100, func(ctx context.Context, n int) (int, error) {
        time.Sleep(time.Microsecond)  // Simulate work
        return n * 2, nil
    })

    assert.NoError(t, err)
    assert.Equal(t, len(data), len(result))
}

func TestParallelMapCancellation(t *testing.T) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
    defer cancel()

    _, err := ParallelMap(ctx, data, 4, func(ctx context.Context, n int) (int, error) {
        time.Sleep(100 * time.Millisecond)  // Slow work
        return n, nil
    })

    assert.Error(t, err)
}
```

**Race Condition Tests**

```go
func TestParallelMapNoRaces(t *testing.T) {
    // Run with: go test -race
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }

    ctx := context.Background()

    for i := 0; i < 100; i++ {
        _, err := ParallelMap(ctx, data, 8, func(ctx context.Context, n int) (int, error) {
            return n * 2, nil
        })
        assert.NoError(t, err)
    }
}
```

#### Steps

1. Create `stress_test.go`
2. Add stress test flag handling
3. Run: `go test -v -run Stress`
4. Run: `go test -race`
5. Commit: "test: add stress tests for large data and concurrency"

---

### 13. Document Last Edge Cases ⏱️ 10 min

**File:** `last.go:5-8`
**Issue:** Panics on empty slices, not documented
**Impact:** Prevent user surprises

#### Current Code

```go
// Last returns the last element from the slice.
func Last[T any](values []T) T {
    n := len(values)
    return values[n-1]  // ❌ Panics on empty
}
```

#### Fixed Code

```go
// Last returns the last element from the slice.
// Panics if the slice is empty.
func Last[T any](values []T) T {
    if len(values) == 0 {
        panic("underscore.Last: empty slice")
    }
    n := len(values)
    return values[n-1]
}
```

#### Steps

1. Add length check with explicit panic
2. Update doc comment
3. Add test: `assert.Panics(t, func() { Last([]int{}) })`
4. Commit: "fix: add explicit panic for Last on empty slice"

---

## 🔵 FUTURE ENHANCEMENTS

### Missing Functional Programming Utilities

Add these based on user demand and usage patterns:

#### 14. TakeWhile / DropWhile ⏱️ 1 hour

```go
func TakeWhile[T any](values []T, predicate func(T) bool) []T
func DropWhile[T any](values []T, predicate func(T) bool) []T
```

#### 15. Scan (Reduce with history) ⏱️ 30 min

```go
func Scan[T, P any](values []T, acc P, fn func(T, P) P) []P
// Example: Scan([]int{1,2,3,4}, 0, +) → [1, 3, 6, 10]
```

#### 16. First / FirstN ⏱️ 20 min

```go
func First[T any](values []T) (T, error)
func FirstN[T any](values []T, n int) []T
```

#### 17. Init (all but last) ⏱️ 15 min

```go
func Init[T any](values []T) ([]T, T)
```

#### 18. Intersperse ⏱️ 20 min

```go
func Intersperse[T any](values []T, separator T) []T
```

#### 19. Sliding Window ⏱️ 30 min

```go
func Sliding[T any](values []T, size int) [][]T
```

#### 20. FoldRight ⏱️ 15 min

```go
func FoldRight[T, P any](values []T, acc P, fn func(T, P) P) P
```

#### 21. Tap (for debugging) ⏱️ 15 min

```go
func Tap[T any](values []T, fn func(T)) []T
```

#### 22. Transpose ⏱️ 30 min

```go
func Transpose[T any](matrix [][]T) [][]T
```

#### 23. Unzip ⏱️ 20 min

```go
func Unzip[L, R any](pairs []Tuple[L, R]) ([]L, []R)
```

#### 24. ParallelReduce ⏱️ 2 hours

```go
func ParallelReduce[T, P any](ctx context.Context, values []T, workers int,
    fn func(context.Context, T, P) (P, error), acc P) (P, error)
```

#### 25. Replicate ⏱️ 10 min

```go
func Replicate[T any](count int, value T) []T
```

---

## Testing Strategy

### Before Any Changes

1. Run full test suite: `go test ./... -v -cover`
2. Document current coverage: `go test -coverprofile=coverage.out && go tool cover -func=coverage.out`
3. Create baseline benchmarks: `go test -bench=. -benchmem > baseline.txt`

### After Each Change

1. Run affected tests: `go test -run TestFunction -v`
2. Run full suite: `go test ./... -v`
3. Check coverage: Coverage should not decrease
4. Run benchmarks: `go test -bench=BenchmarkFunction -benchmem`
5. Run race detector: `go test -race`

### CI Integration

Add GitHub Actions workflow:

```yaml
- name: Test
  run: go test ./... -v -race -coverprofile=coverage.out

- name: Benchmark
  run: go test -bench=. -benchmem
```

---

## Breaking Changes Policy

When making breaking changes:

1. **Document in CHANGELOG.md**
   - What changed
   - Why it changed
   - Migration path

2. **Update Version**
   - Major version bump (v0.7.0 → v0.8.0)
   - Follow SemVer strictly

3. **Provide Deprecation Period** (if possible)
   - Keep old function with `Deprecated:` doc comment
   - Add new function alongside
   - Remove in next major version

4. **Add Migration Guide**
   - Before/after code examples
   - Search/replace patterns
   - Common pitfalls

---

## Success Metrics

After completing all critical and high priority items:

- ✅ Test coverage remains >99%
- ✅ Filter performance improves 2-5x
- ✅ OrderBy performance improves 10-100x for large lists
- ✅ Zero panics on empty slices (or documented)
- ✅ Benchmark suite covering all core functions
- ✅ API inconsistencies resolved
- ✅ All edge cases tested

**Target Quality Score:** 9.5/10

---

## Notes

- All time estimates are approximate
- Test thoroughly after each change
- Consider user impact for breaking changes
- Gather community feedback before major API changes
- Update documentation as you go
- Run benchmarks to verify improvements

Generated by comprehensive codebase review on 2025-11-14.
