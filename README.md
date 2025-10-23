# smallset

`smallset` is a generic sorted set implementation for Go, backed by a slice. It is designed to be highly memory-efficient and offers fast range access, making it an excellent choice for **small sets (under 1000 elements)**.

## Features

  * **Go Generics:** Works with any `cmp.Ordered` type (e.g., `int`, `string`, `float64` ...).

  * **Sorted Data:** Elements are always maintained in ascending order within the internal slice.

  * **Fast Accessors:** Provides O(1) performance for `Min()`, `Max()`, `MinK()`, and `MaxK()`.

  * **Memory Efficiency:** Avoids the memory overhead of map keys and pointers associated with traditional hash maps, leading to better cache locality for small data sets.

  * **Efficient Set Operations:** Binary operations like `Intersect`, `Union`, and `Difference` are performed using the efficient two-pointer merge algorithm, which is often much faster than map-based iteration for sorted data.

## Performance Trade-offs

| Operation | `smallset` | `map[T]struct{}` |
| --------- | ------------------------- | ---------------------------------- |
| **Lookup (`Contains`)** | $O(\log N)$ (Binary Search) | $O(1)$ average (Hash Lookup) |
| **Range Access (MinK/MaxK)** | $O(1)$ | $O(N)$ (Requires map iteration or conversion) |
| **Mutation (Add/Remove)** | $O(N)$ (Slice shift dominates) | $O(1)$ average |
| **Set Operations (Intersect, Union, Difference)** | $O(N)$ (Two-Pointer Merge) | $O(N)$ (Iteration + $O(1)$ lookups) |

The most common advice is to use hash maps for their guaranteed $O(1)$ complexity. However, for small $N$, the map's large constant factor is so dominant that it outweighs the slice's slightly worse asymptotic complexity. The reasons are:

- **Cache Locality**: A Go slice stores all its elements contiguously in memory. When performing set operations the CPU efficiently loads sequential elements into the cache with minimal latency. Map-based sets, however, store key-value pairs scattered across memory via hash buckets. This forces the CPU to constantly jump through pointers, resulting in more cache-misses.

- **Hashing Overhead**: Every single map operation requires calculating a hash for the key. This cryptographic overhead is a constant tax that the slice-based set bypasses entirely, replacing it with simpler, faster value comparison.

**The Sweet Spot:** `smallset` is best used for **sets where the size remains small (< 1000)** or in scenarios where additions/removals are infrequent, but lookups, iteration, and ordered access are common.

## Benchmarks

We tested `smallset` against `github.com/deckarep/golang-set`, the most used [map-based set implementation](https://github.com/deckarep/golang-set/tree/main).  
See the results [here](bench.md).

**TLDR:** `smallset` offers comparable performance for insertions and deletions up to 1000 elements, while outperforming the map based sets by **10-30x** for heavy set-like operations like `Intersect`, `Difference`, `SymetricDifference`...

## Installation

```
go get github.com/pippellia-btc/smallset
```


## Usage Examples

### Basic Operations

```golang
package main

import (
	"fmt"
	"github.com/pippellia-btc/smallset"
)

func main() {
	// 1. Create a new set with initial capacity
	set := smallset.New[int](5)

	// 2. Add elements (returns true if added, false if duplicate)
	set.Add(20) // true
	set.Add(10) // true
	set.Add(20) // false (duplicate)
	fmt.Println("Set size:", set.Size()) // 2

	// 3. Contains and Remove
	fmt.Println("Contains 10:", set.Contains(10)) // true
	set.Remove(10)
	fmt.Println("Contains 10:", set.Contains(10)) // false
}
```

### Min/Max and TopK

```golang
package main

import (
	"fmt"
	"github.com/pippellia-btc/smallset"
)

func main() {
	set := smallset.From(50, 10, 30, 20, 40) // Internally sorted: [10, 20, 30, 40, 50]

	// MinK (returns k smallest elements)
	mins := set.MinK(2)
	fmt.Println("2 Smallest:", mins) // [10 20]

	// MaxK (returns k largest elements)
	maxs := set.MaxK(3)
	fmt.Println("3 Largest:", maxs) // [30 40 50]

	// MinK with k > size (returns all elements)
	all := set.MinK(10)
	fmt.Println("All elements:", all) // [10 20 30 40 50]
}
```
