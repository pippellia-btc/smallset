package smallset

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

var defaultCapacity int = 10

// Ordered is a slice-based set sorted in ascending order.
// It's more performant that a map based approach for small collections (< 1000) of ordered types.
// The capacity of the set can dynamically grow, but the performance would start to deteriorate.
// Not safe for concurrent use.
type Ordered[T cmp.Ordered] struct {
	items []T
}

// New returns an initialized set with the provided capacity.
// It panics if the capacity is <= 0.
func New[T cmp.Ordered](capacity int) *Ordered[T] {
	if capacity <= 0 {
		panic("smallset.New: capacity must be > 0")
	}

	return &Ordered[T]{
		items: make([]T, 0, capacity),
	}
}

// NewFrom returns an initialized set that contains the provided elements.
func NewFrom[T cmp.Ordered](items ...T) *Ordered[T] {
	if len(items) == 0 {
		return New[T](defaultCapacity)
	}

	copy := slices.Clone(items)
	slices.Sort(copy)
	copy = slices.Compact(copy)

	s := New[T](max(len(copy), defaultCapacity))
	s.items = copy
	return s
}

// Size returns the number of elements in the set.
func (s *Ordered[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns whether the set has no elements.
func (s *Ordered[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements.
func (s *Ordered[T]) Clear() {
	s.items = s.items[:0]
}

// Clone returns a clone of the set.
func (s *Ordered[T]) Clone() *Ordered[T] {
	return &Ordered[T]{
		items: slices.Clone(s.items),
	}
}

// Items returns a copy of the internal slice of the set.
func (s *Ordered[T]) Items() []T {
	return slices.Clone(s.items)
}

// Contains returns whether the element is in the set. Operation is O(log(N))
func (s *Ordered[T]) Contains(e T) bool {
	_, found := slices.BinarySearch(s.items, e)
	return found
}

// At returns the element at index i or panics if out of range.
func (s *Ordered[T]) At(i int) T {
	if i < 0 || i >= len(s.items) {
		panic("smallset.Ordered.At: index out of range")
	}
	return s.items[i]
}

// Find returns the index of an element, or the position where target would appear
// in the sort order. It also returns a bool saying whether the target is really found in the slice.
func (s *Ordered[T]) Find(e T) (int, bool) {
	return slices.BinarySearch(s.items, e)
}

// Add an element and returns whether is was added (true), or was already present (false).
func (s *Ordered[T]) Add(e T) bool {
	i, found := slices.BinarySearch(s.items, e)
	if found {
		return false
	}

	s.items = slices.Insert(s.items, i, e)
	return true
}

// Remove an element if present, and returns whether is was removed (true), or was never present (false).
func (s *Ordered[T]) Remove(e T) bool {
	i, found := slices.BinarySearch(s.items, e)
	if !found {
		return false
	}

	s.items = slices.Delete(s.items, i, i+1)
	return true
}

// RemoveBefore removes all elements e such that e < max. Returns num removed.
func (s *Ordered[T]) RemoveBefore(max T) int {
	end, _ := slices.BinarySearch(s.items, max)
	if end == 0 {
		return 0
	}

	s.items = slices.Delete(s.items, 0, end)
	return end
}

// RemoveFrom removed all elements e such that e >= min. Returns num removed.
func (s *Ordered[T]) RemoveFrom(min T) int {
	start, _ := slices.BinarySearch(s.items, min)
	if start == len(s.items) {
		return 0
	}

	removed := len(s.items) - start
	s.items = slices.Delete(s.items, start, len(s.items))
	return removed
}

// RemoveBetween removes all elements e such that min <= e < max. Returns num removed.
func (s *Ordered[T]) RemoveBetween(min, max T) int {
	if cmp.Less(max, min) {
		panic("smallset.Ordered.RemoveBetween: invalid range (max < min)")
	}

	start, _ := slices.BinarySearch(s.items, min)
	end, _ := slices.BinarySearch(s.items, max)
	if start == end {
		return 0
	}

	s.items = slices.Delete(s.items, start, end)
	return end - start
}

// Min returns the smallest element in the set.
// It panics if the set is empty.
func (s *Ordered[T]) Min() T {
	if s.IsEmpty() {
		panic("smallset.Ordered.Min: set is empty")
	}
	return s.items[0]
}

// Max returns the biggest element in the sets.
// It panics if the set is empty.
func (s *Ordered[T]) Max() T {
	if s.IsEmpty() {
		panic("smallset.Ordered.Max: set is empty")
	}
	return s.items[len(s.items)-1]
}

// MinK returns the k smallest elements in s, sorted in ascending order. O(k) complexity.
// It panics if k is negative. If k is bigger than the set size, it returns all the items.
func (s *Ordered[T]) MinK(k int) []T {
	if k < 0 {
		panic(fmt.Sprintf("smallset.Ordered.MinK: k must be positive: %d", k))
	}
	k = min(k, s.Size())
	return slices.Clone(s.items[:k])
}

// MaxK returns the k biggest elements in s, sorted in ascending order. O(k) complexity.
// It panics if k is negative. If k is bigger than the set size, it returns all the items.
func (s *Ordered[T]) MaxK(k int) []T {
	if k < 0 {
		panic(fmt.Sprintf("smallset.Ordered.MaxK: k must be positive: %d", k))
	}
	k = min(k, s.Size())
	return slices.Clone(s.items[len(s.items)-k:])
}

// Ascend returns an iterator over the set in ascending order.
func (s *Ordered[T]) Ascend() iter.Seq2[int, T] {
	return slices.All(s.items)
}

// Descend returns an iterator over the set in descending order.
func (s *Ordered[T]) Descend() iter.Seq2[int, T] {
	return slices.Backward(s.items)
}

// BetweenAsc iterates NewFrom min (inclusive) to max (exclusive) in ascending order.
// If min or max are not present in the set, iteration starts/ends at the position
// where they would appear in the sorted slice. Panics if max < min.
func (s *Ordered[T]) BetweenAsc(min, max T) iter.Seq2[int, T] {
	if cmp.Less(max, min) {
		panic("smallset.Ordered.BetweenAsc: invalid range (max < min)")
	}
	start, _ := slices.BinarySearch(s.items, min)

	return func(yield func(int, T) bool) {
		for i := start; i < len(s.items); i++ {
			v := s.items[i]
			if !cmp.Less(v, max) {
				return
			}
			if !yield(i, v) {
				return
			}
		}
	}
}

// BetweenDesc iterates NewFrom max (inclusive) down to min (exclusive) in descending order.
// If min or max are not present in the set, iteration starts/ends at the position
// where they would appear in the sorted slice. Panics if max < min.
func (s *Ordered[T]) BetweenDesc(max, min T) iter.Seq2[int, T] {
	if cmp.Less(max, min) {
		panic("smallset.Ordered.BetweenDesc: invalid range (max < min)")
	}

	end, found := slices.BinarySearch(s.items, max)
	if !found && end > 0 {
		end--
	}

	return func(yield func(int, T) bool) {
		for i := end; i >= 0; i-- {
			v := s.items[i]
			if !cmp.Less(min, v) {
				return
			}
			if !yield(i, v) {
				return
			}
		}
	}
}

// IsEqual returns whether the two sets have the same elements.
func (s *Ordered[T]) IsEqual(other *Ordered[T]) bool {
	return slices.Equal(s.items, other.items)
}

// Intersect returns the intersection of two sets, returning a New set
// containing only the common elements. O(N+M) complexity.
func (s *Ordered[T]) Intersect(other *Ordered[T]) *Ordered[T] {
	size := min(s.Size(), other.Size())
	if size == 0 {
		return New[T](defaultCapacity)
	}

	inter := New[T](size)

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s_i < o_i {
			// element in s not in other
			i++
		} else if o_i < s_i {
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
func (s *Ordered[T]) Difference(other *Ordered[T]) *Ordered[T] {
	if s.IsEmpty() {
		return New[T](defaultCapacity)
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	diff := New[T](s.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s_i < o_i {
			// element in s not in other
			diff.items = append(diff.items, s_i)
			i++
		} else if o_i < s_i {
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

// SymmetricDifference returns a New set with all elements which are
// in either this set or the other set but not in both. O(N+M) complexity.
func (s *Ordered[T]) SymmetricDifference(other *Ordered[T]) *Ordered[T] {
	if s.IsEmpty() {
		return other.Clone()
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	sdiff := New[T](s.Size() + other.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s_i < o_i {
			// element in s not in other
			sdiff.items = append(sdiff.items, s_i)
			i++
		} else if o_i < s_i {
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

// Union returns a New set with all elements in both sets. O(N+M) complexity.
func (s *Ordered[T]) Union(other *Ordered[T]) *Ordered[T] {
	if s.IsEmpty() {
		return other.Clone()
	}
	if other.IsEmpty() {
		return s.Clone()
	}

	union := New[T](s.Size() + other.Size())

	i := 0
	j := 0

	for i < s.Size() && j < other.Size() {
		s_i := s.items[i]
		o_i := other.items[j]

		if s_i < o_i {
			// element in s not in other
			union.items = append(union.items, s_i)
			i++
		} else if o_i < s_i {
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

// Partition returns three New sets:
// - d12: elements in s1 not in s2
// - inter: elements in both sets
// - d21: elements in s2 not in s1
// O(N+M) complexity.
func (s1 *Ordered[T]) Partition(s2 *Ordered[T]) (d12, inter, d21 *Ordered[T]) {
	if s1.IsEmpty() {
		return New[T](defaultCapacity), New[T](defaultCapacity), s2.Clone()
	}
	if s2.IsEmpty() {
		return s1.Clone(), New[T](defaultCapacity), New[T](defaultCapacity)
	}

	d12 = New[T](s1.Size())
	inter = New[T](min(s1.Size(), s2.Size()))
	d21 = New[T](s2.Size())

	i := 0
	j := 0

	for i < s1.Size() && j < s2.Size() {
		e1 := s1.items[i]
		e2 := s2.items[j]

		if e1 < e2 {
			// element in s1 not in s2
			d12.items = append(d12.items, e1)
			i++
		} else if e2 < e1 {
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
