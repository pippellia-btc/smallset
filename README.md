# miss: Mini Sorted Set

`miss` (Mini Sorted Set) is a generic set implementation for Go, backed by a sorted slice. It is designed to be highly memory-efficient and offers fast range access, making it an excellent choice for **small sets (under 1000 elements)**.

## Features

  * **Go Generics:** Works with any type that satisfies the `cmp.Ordered` constraint (e.g., `int`, `string`, `float64` ...).

  * **Sorted Data:** Elements are always maintained in ascending order within the internal slice.

  * **Fast Accessors:** Provides O(1) performance for `Min()`, `Max()`, `MinK()`, and `MaxK()`.

  * **Memory Efficiency:** Avoids the memory overhead of map keys and pointers associated with traditional hash maps, leading to better cache locality for small data sets.

  * **Efficient Set Operations:** Binary operations like `Intersect`, `Union`, and `Difference` are performed using the efficient two-pointer merge algorithm, achieving O(N+M) complexity, which is often much faster than map-based iteration for sorted data.

## Performance Trade-offs

| **Operation** | **Complexity** | **Notes** |
| **Lookup (`Contains`)** | O(log N) | Achieved via `slices.BinarySearch`. |
| **Range Access (MinK/MaxK)** | O(1) | Direct slice access and cloning. |
| **Mutation (Add/Remove)** | O(log N + N) | Requires an O(log N) search followed by an O(N) slice insertion/deletion. |
| **Set Operations (Intersection, Union)** | O(N + M) | Highly optimized two-pointer merge. |

**The Sweet Spot:** Due to the O(N) complexity of mutations, `miss` is best used for **utility sets where the size remains small (e.g., N < 1000)** or in scenarios where additions/removals are infrequent, but lookups, iteration, and ordered access are common.

## Benchmarks

We tested `miss` against the most used (map-based set implementation)[https://github.com/deckarep/golang-set/tree/main] `github.com/deckarep/golang-set`. See the results [here](bench.md)

**TLDR:** `miss` offers comparable performance for insertions and deletions up to 1000 elements, while outperforming the map based sets by **10-30x** for heavy set-like operations like `Intersect`, `Difference`, `SymetricDifference`...

## Installation

```
go get github.com/pippellia-btc/miss
```


## Usage Examples

### Basic Operations

```golang
package main

import (
	"fmt"
	"github.com/pippellia-btc/miss"
)

func main() {
	// 1. Create a new set with initial capacity
	s := miss.New[int](5)

	// 2. Add elements (returns true if added, false if duplicate)
	s.Add(20) // true
	s.Add(10) // true
	s.Add(20) // false (duplicate)
	fmt.Println("Set size:", s.Size()) // 2

	// 3. Contains and Remove
	fmt.Println("Contains 10:", s.Contains(10)) // true
	s.Remove(10)
	fmt.Println("Contains 10:", s.Contains(10)) // false
}
```

### Min/Max and TopK

```golang
package main

import (
	"fmt"
	"your.repo/path/miss"
)

func main() {
	s := miss.From(50, 10, 30, 20, 40) // Internally sorted: [10, 20, 30, 40, 50]

	// MinK (returns k smallest elements)
	mins := s.MinK(2)
	fmt.Println("2 Smallest:", mins) // [10 20]

	// MaxK (returns k largest elements)
	maxs := s.MaxK(3)
	fmt.Println("3 Largest:", maxs) // [30 40 50]

	// MinK with k > size (returns all elements)
	all := s.MinK(10)
	fmt.Println("All elements:", all) // [10 20 30 40 50]
}
```