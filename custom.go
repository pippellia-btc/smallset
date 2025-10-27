package smallset

import (
	"fmt"
	"iter"
	"slices"
)

// Custom is a slice-based set sorted in ascending order, as determined by the
// cmp function provided in the contructor.
// If T is an ordered type, you should use [Ordered] for better performance.
//
// It's more performant that a map based approach for small collections (< 1000).
// The capacity of the set can dynamically grow, but the performance would start to deteriorate.
// Not safe for concurrent use.
type Custom[T any] struct {
	items []T
	cmp   compareFunc[T]
}

// The three-way comparison function:
//   - cmp(a, b) < 0 if a is less than b
//   - cmp(a, b) > 0 if a is greater than b
//   - cmp(a, b) == 0 if a is equivalent to b (duplicates)
//
// It's a custom type so it can have methods that makes code more readable.
type compareFunc[T any] func(a, b T) int

func (c compareFunc[T]) less(a, b T) bool  { return c(a, b) < 0 }
func (c compareFunc[T]) equal(a, b T) bool { return c(a, b) == 0 }

// NewCustom returns an initialized set with the provided compare function and capacity.
//
// The cmp function allows two elements, a and b, to be compared,
// following a similar convention to that of the slices package.
// - cmp(a, b) < 0 if a < b
// - cmp(a, b) > 0 if a > b
// - cmp(a, b) == 0 if a = b (duplicates)
//
// It panics if the cmp function is nil or capacity is <= 0.
func NewCustom[T any](cmp func(a, b T) int, capacity int) *Custom[T] {
	if capacity <= 0 {
		panic("smallset.NewCustom: capacity must be > 0")
	}
	if cmp == nil {
		panic("smallset.NewCustom: cmp cannot be nil")
	}

	return &Custom[T]{
		items: make([]T, 0, capacity),
		cmp:   compareFunc[T](cmp),
	}
}

// NewCustomFrom returns an initialized set that contains the provided elements,
// sorted by the provided compare function cmp.
//
// The cmp function allows two elements, a and b, to be compared,
// following a similar convention to that of the slices package.
// - cmp(a, b) < 0 if a < b
// - cmp(a, b) > 0 if a > b
// - cmp(a, b) == 0 if a = b (duplicates)
//
// It panics if cmp is nil.
func NewCustomFrom[T any](cmp func(a, b T) int, items ...T) *Custom[T] {
	if len(items) == 0 {
		return NewCustom(cmp, defaultCapacity)
	}
	if cmp == nil {
		panic("smallset.NewCustomFrom: cmp cannot be nil")
	}

	copy := slices.Clone(items)
	compare := compareFunc[T](cmp)
	slices.SortFunc(copy, compare)
	copy = slices.CompactFunc(copy, compare.equal)

	s := NewCustom(compare, max(len(copy), defaultCapacity))
	s.items = copy
	return s
}

// Size returns the number of elements in the set.
func (s *Custom[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns whether the set has no elements.
func (s *Custom[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements.
func (s *Custom[T]) Clear() {
	s.items = s.items[:0]
}

// Clone returns a clone of the set, that shares the cmp comparator function.
func (s *Custom[T]) Clone() *Custom[T] {
	return &Custom[T]{
		items: slices.Clone(s.items),
		cmp:   s.cmp,
	}
}

// Items returns a copy of the internal slice of the set.
func (s *Custom[T]) Items() []T {
	return slices.Clone(s.items)
}

// Contains returns whether the element is in the set. Operation is O(log(N))
func (s *Custom[T]) Contains(e T) bool {
	_, found := slices.BinarySearchFunc(s.items, e, s.cmp)
	return found
}

// At returns the element at index i or panics if out of range.
func (s *Custom[T]) At(i int) T {
	if i < 0 || i >= len(s.items) {
		panic("smallset.Custom.At: index out of range")
	}
	return s.items[i]
}

// Find returns the index of an element, or the position where target would appear
// in the sort order. It also returns a bool saying whether the target is really found in the slice.
func (s *Custom[T]) Find(e T) (int, bool) {
	return slices.BinarySearchFunc(s.items, e, s.cmp)
}

// Add an element and returns whether is was added (true), or was already present (false).
func (s *Custom[T]) Add(e T) bool {
	i, found := slices.BinarySearchFunc(s.items, e, s.cmp)
	if found {
		return false
	}

	s.items = slices.Insert(s.items, i, e)
	return true
}

// Remove an element if present, and returns whether is was removed (true), or was never present (false).
func (s *Custom[T]) Remove(e T) bool {
	i, found := slices.BinarySearchFunc(s.items, e, s.cmp)
	if !found {
		return false
	}

	s.items = slices.Delete(s.items, i, i+1)
	return true
}

// RemoveBefore removes all elements e such that e < max. Returns num removed.
func (s *Custom[T]) RemoveBefore(max T) int {
	end, _ := slices.BinarySearchFunc(s.items, max, s.cmp)
	if end == 0 {
		return 0
	}

	s.items = slices.Delete(s.items, 0, end)
	return end
}

// RemoveFrom removed all elements e such that e >= min. Returns num removed.
func (s *Custom[T]) RemoveFrom(min T) int {
	start, _ := slices.BinarySearchFunc(s.items, min, s.cmp)
	if start == len(s.items) {
		return 0
	}

	removed := len(s.items) - start
	s.items = slices.Delete(s.items, start, len(s.items))
	return removed
}

// RemoveBetween removes all elements e such that min <= e < max. Returns num removed.
func (s *Custom[T]) RemoveBetween(min, max T) int {
	if s.cmp.less(max, min) {
		panic("smallset.Custom.RemoveBetween: invalid range (max < min)")
	}

	start, _ := slices.BinarySearchFunc(s.items, min, s.cmp)
	end, _ := slices.BinarySearchFunc(s.items, max, s.cmp)
	if start == end {
		return 0
	}

	s.items = slices.Delete(s.items, start, end)
	return end - start
}

// Min returns the smallest element in the set.
// It panics if the set is empty.
func (s *Custom[T]) Min() T {
	if s.IsEmpty() {
		panic("smallset.Custom.Min: set is empty")
	}
	return s.items[0]
}

// Max returns the biggest element in the sets.
// It panics if the set is empty.
func (s *Custom[T]) Max() T {
	if s.IsEmpty() {
		panic("smallset.Custom.Max: set is empty")
	}
	return s.items[len(s.items)-1]
}

// MinK returns the k smallest elements in s, sorted in ascending order. O(k) complexity.
// It panics if k is negative. If k is bigger than the set size, it returns all the items.
func (s *Custom[T]) MinK(k int) []T {
	if k < 0 {
		panic(fmt.Sprintf("smallset.Custom.MinK: k must be positive: %d", k))
	}
	k = min(k, s.Size())
	return slices.Clone(s.items[:k])
}

// MaxK returns the k biggest elements in s, sorted in ascending order. O(k) complexity.
// It panics if k is negative. If k is bigger than the set size, it returns all the items.
func (s *Custom[T]) MaxK(k int) []T {
	if k < 0 {
		panic(fmt.Sprintf("smallset.Custom.MaxK: k must be positive: %d", k))
	}
	k = min(k, s.Size())
	return slices.Clone(s.items[len(s.items)-k:])
}

// Ascend returns an iterator over the set in ascending order.
func (s *Custom[T]) Ascend() iter.Seq2[int, T] {
	return slices.All(s.items)
}

// Descend returns an iterator over the set in descending order.
func (s *Custom[T]) Descend() iter.Seq2[int, T] {
	return slices.Backward(s.items)
}

// BetweenAsc iterates NewCustomFrom min (inclusive) to max (exclusive) in ascending order.
// If min or max are not present in the set, iteration starts/ends at the position
// where they would appear in the sorted slice. Panics if max < min.
func (s *Custom[T]) BetweenAsc(min, max T) iter.Seq2[int, T] {
	if s.cmp.less(max, min) {
		panic("smallset.Custom.BetweenAsc: invalid range (max < min)")
	}
	start, _ := slices.BinarySearchFunc(s.items, min, s.cmp)

	return func(yield func(int, T) bool) {
		for i := start; i < len(s.items); i++ {
			v := s.items[i]
			if !s.cmp.less(v, max) {
				return
			}
			if !yield(i, v) {
				return
			}
		}
	}
}

// BetweenDesc iterates NewCustomFrom max (inclusive) down to min (exclusive) in descending order.
// If min or max are not present in the set, iteration starts/ends at the position
// where they would appear in the sorted slice. Panics if max < min.
func (s *Custom[T]) BetweenDesc(max, min T) iter.Seq2[int, T] {
	if s.cmp.less(max, min) {
		panic("smallset.Custom.BetweenDesc: invalid range (max < min)")
	}

	end, found := slices.BinarySearchFunc(s.items, max, s.cmp)
	if !found && end > 0 {
		end--
	}

	return func(yield func(int, T) bool) {
		for i := end; i >= 0; i-- {
			v := s.items[i]
			if !s.cmp.less(min, v) {
				return
			}
			if !yield(i, v) {
				return
			}
		}
	}
}

// IsEqual returns whether the two sets have the same elements.
func (s *Custom[T]) IsEqual(other *Custom[T]) bool {
	return slices.EqualFunc(s.items, other.items, s.cmp.equal)
}

// Intersect returns the intersection of two sets, returning a NewCustom set
// containing only the common elements. O(N+M) complexity.
func (s *Custom[T]) Intersect(other *Custom[T]) *Custom[T] {
	size := min(s.Size(), other.Size())
	if size == 0 {
		return NewCustom[T](s.cmp, defaultCapacity)
	}

	inter := NewCustom[T](s.cmp, size)

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s.cmp.less(s_i, o_i) {
			// element in s not in other
			i++
		} else if s.cmp.less(o_i, s_i) {
			// element in other not in s
			j++
		} else {
			// element in both
			inter.items = append(inter.items, s_i)
			i++
			j++
		}
	}

	return inter
}

// Difference returns the difference between this set and other. The returned set will contain
// all elements of this set that are not elements of other. O(N+M) complexity.
func (s *Custom[T]) Difference(other *Custom[T]) *Custom[T] {
	if s.IsEmpty() {
		return NewCustom[T](s.cmp, defaultCapacity)
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	diff := NewCustom[T](s.cmp, s.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s.cmp.less(s_i, o_i) {
			// element in s not in other
			diff.items = append(diff.items, s_i)
			i++
		} else if s.cmp.less(o_i, s_i) {
			// element in other not in s
			j++
		} else {
			// element in both
			i++
			j++
		}
	}

	diff.items = append(diff.items, s.items[i:]...)
	return diff
}

// SymmetricDifference returns a NewCustom set with all elements which are
// in either this set or the other set but not in both. O(N+M) complexity.
func (s *Custom[T]) SymmetricDifference(other *Custom[T]) *Custom[T] {
	if s.IsEmpty() {
		return other.Clone()
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	sdiff := NewCustom[T](s.cmp, s.Size()+other.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s.cmp.less(s_i, o_i) {
			// element in s not in other
			sdiff.items = append(sdiff.items, s_i)
			i++
		} else if s.cmp.less(o_i, s_i) {
			// element in other not in s
			sdiff.items = append(sdiff.items, o_i)
			j++
		} else {
			// element in both
			i++
			j++
		}
	}

	sdiff.items = append(sdiff.items, s.items[i:]...)
	sdiff.items = append(sdiff.items, other.items[j:]...)
	return sdiff
}

// Union returns a NewCustom set with all elements in both sets. O(N+M) complexity.
func (s *Custom[T]) Union(other *Custom[T]) *Custom[T] {
	if s.IsEmpty() {
		return other.Clone()
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	union := NewCustom[T](s.cmp, s.Size()+other.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s.cmp.less(s_i, o_i) {
			// element in s not in other
			union.items = append(union.items, s_i)
			i++
		} else if s.cmp.less(o_i, s_i) {
			// element in other not in s
			union.items = append(union.items, o_i)
			j++
		} else {
			// element in both
			union.items = append(union.items, s_i)
			i++
			j++
		}
	}

	union.items = append(union.items, s.items[i:]...)
	union.items = append(union.items, other.items[j:]...)
	return union
}

// Partition returns three NewCustom sets:
// - d12: elements in s1 not in s2
// - inter: elements in both sets
// - d21: elements in s2 not in s1
// O(N+M) complexity.
func (s1 *Custom[T]) Partition(s2 *Custom[T]) (d12, inter, d21 *Custom[T]) {
	if s1.IsEmpty() {
		return NewCustom[T](s1.cmp, defaultCapacity), NewCustom[T](s1.cmp, defaultCapacity), s2.Clone()
	}
	if s2.IsEmpty() {
		return s1.Clone(), NewCustom[T](s1.cmp, defaultCapacity), NewCustom[T](s1.cmp, defaultCapacity)
	}

	d12 = NewCustom[T](s1.cmp, s1.Size())
	inter = NewCustom[T](s1.cmp, min(s1.Size(), s2.Size()))
	d21 = NewCustom[T](s1.cmp, s2.Size())

	i := 0
	j := 0

	for i < s1.Size() && j < s2.Size() {
		e1 := s1.items[i]
		e2 := s2.items[j]

		if s1.cmp.less(e1, e2) {
			// element in s1 not in s2
			d12.items = append(d12.items, e1)
			i++
		} else if s1.cmp.less(e2, e1) {
			// element in s2 not in s1
			d21.items = append(d21.items, e2)
			j++
		} else {
			// element in both
			inter.items = append(inter.items, e1)
			i++
			j++
		}
	}

	d12.items = append(d12.items, s1.items[i:]...)
	d21.items = append(d21.items, s2.items[j:]...)
	return d12, inter, d21
}
