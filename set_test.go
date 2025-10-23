package smallset

import (
	"fmt"
	"iter"
	"math/rand"
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestContains(t *testing.T) {
	initial := []int{5, 10, 15, 20}
	s := From(initial...)

	cases := []struct {
		element  int
		expected bool
	}{
		{element: 15, expected: true},
		{element: 1, expected: false},
		{element: 25, expected: false},
		{element: 13, expected: false},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := s.Contains(test.element)
			if result != test.expected {
				t.Errorf("Contains(%d) expected %t got %t", test.element, test.expected, result)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		toAdd    []int
		expected []bool
		items    []int
	}{
		{
			toAdd:    []int{10, 20, 30},
			expected: []bool{true, true, true},
			items:    []int{10, 20, 30},
		},
		{
			toAdd:    []int{7, 5, 5, 10, 8, 7},
			expected: []bool{true, true, false, true, true, false},
			items:    []int{5, 7, 8, 10},
		},
		{
			toAdd:    []int{5, 10, 7},
			expected: []bool{true, true, true},
			items:    []int{5, 7, 10},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := New[int](10)
			res := make([]bool, len(test.toAdd))
			for j, e := range test.toAdd {
				res[j] = s.Add(e)
			}

			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("Add results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !reflect.DeepEqual(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		initial  []int
		toRemove []int
		expected []bool
		items    []int
	}{
		{
			initial:  []int{10, 20, 30},
			toRemove: []int{20, 10, 50},
			expected: []bool{true, true, false},
			items:    []int{30},
		},
		{
			initial:  []int{5, 7, 8, 10},
			toRemove: []int{7, 5, 10, 8},
			expected: []bool{true, true, true, true},
			items:    []int{},
		},
		{
			initial:  []int{1, 2, 3},
			toRemove: []int{5, 6},
			expected: []bool{false, false},
			items:    []int{1, 2, 3},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := From(test.initial...)
			res := make([]bool, len(test.toRemove))
			for j, e := range test.toRemove {
				res[j] = s.Remove(e)
			}

			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("Remove results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !reflect.DeepEqual(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	s1 := From(1, 2, 3)
	s2 := From(3, 2, 1)
	s3 := From(1, 2, 3, 4)
	s4 := New[int](10)

	cases := []struct {
		setA     *Set[int]
		setB     *Set[int]
		expected bool
	}{
		{setA: s1, setB: s2, expected: true},
		{setA: s1, setB: s3, expected: false},
		{setA: s1, setB: s4, expected: false},
		{setA: s4, setB: New[int](30), expected: true},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			if res := test.setA.IsEqual(test.setB); res != test.expected {
				t.Errorf("IsEqual expected %t, got %t", test.expected, res)
			}
		})
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		set      *Set[int]
		expected int
	}{
		{set: From(10, 5, 20, 15), expected: 5},
		{set: From(1, 5, 20, 69), expected: 1},
		{set: From(7, 8, 4, 12, 221), expected: 4},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.Min()
			if result != test.expected {
				t.Errorf("Min() failed.\nExpected: %d\nActual:   %d", test.expected, result)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		set      *Set[int]
		expected int
	}{
		{set: From(10, 5, 20, 15), expected: 20},
		{set: From(1, 5, 20, 69), expected: 69},
		{set: From(7, 8, 4, 12, 221), expected: 221},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.Max()
			if result != test.expected {
				t.Errorf("Max() failed.\nExpected: %d\nActual:   %d", test.expected, result)
			}
		})
	}
}

func TestMinK(t *testing.T) {
	cases := []struct {
		set      *Set[int]
		k        int
		expected []int
	}{
		{set: From(10, 5, 20, 15), k: 2, expected: []int{5, 10}},
		{set: From(7, 8, 4, 12, 221), k: 150, expected: []int{4, 7, 8, 12, 221}},
		{set: From(1, 5, 20, 69), k: 0, expected: []int{}},
		{set: New[int](10), k: 5, expected: []int{}},
		{set: New[int](10), k: 0, expected: []int{}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.MinK(test.k)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MinK(%d) failed.\nExpected: %v\nActual: %v", test.k, test.expected, result)
			}
		})
	}
}

func TestMaxK(t *testing.T) {
	cases := []struct {
		set      *Set[int]
		k        int
		expected []int
	}{
		{set: From(10, 5, 20, 15), k: 2, expected: []int{15, 20}},
		{set: From(7, 8, 4, 12, 221), k: 150, expected: []int{4, 7, 8, 12, 221}},
		{set: From(1, 5, 20, 69), k: 0, expected: []int{}},
		{set: New[int](10), k: 5, expected: []int{}},
		{set: New[int](10), k: 0, expected: []int{}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.MaxK(test.k)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MaxK(%d) failed.\nExpected: %v\nActual: %v", test.k, test.expected, result)
			}
		})
	}
}

func collect[T any](seq iter.Seq2[int, T]) []T {
	var out []T
	for _, v := range seq {
		out = append(out, v)
	}
	return out
}

func TestRangeAsc(t *testing.T) {
	s := From(1, 3, 5, 7, 9)

	cases := []struct {
		min, max int
		expected []int
	}{
		{min: -1, max: 10, expected: []int{1, 3, 5, 7, 9}},
		{min: 3, max: 7, expected: []int{3, 5}},
		{min: 5, max: 6, expected: []int{5}},
		{min: 8, max: 8, expected: nil},
		{min: 0, max: 2, expected: []int{1}},
		{min: 10, max: 20, expected: nil},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := collect(s.RangeAsc(test.min, test.max))
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("RangeAsc(%d, %d) failed.\nExpected: %v\nActual: %v", test.min, test.max, result, test.expected)
			}
		})
	}
}

func TestRangeDesc(t *testing.T) {
	s := From(1, 3, 5, 7, 9)

	cases := []struct {
		max, min int
		expected []int
	}{
		{max: 10, min: -1, expected: []int{9, 7, 5, 3, 1}},
		{max: 7, min: 3, expected: []int{7, 5}},
		{max: 6, min: 4, expected: []int{5}},
		{min: 8, max: 8, expected: nil},
		{max: 2, min: 0, expected: []int{1}},
		{max: 20, min: 10, expected: nil},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := collect(s.RangeDesc(test.max, test.min))
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("RangeDesc(%d, %d) failed.\nExpected: %v\nActual: %v", test.max, test.min, test.expected, result)
			}
		})
	}
}

// --- Binary Set Operation Tests ---

func TestIntersect(t *testing.T) {
	cases := []struct {
		s1       []int
		s2       []int
		expected []int
	}{
		{s1: []int{1, 3, 5, 7}, s2: []int{2, 3, 6, 7}, expected: []int{3, 7}},
		{s1: []int{1, 2, 3}, s2: []int{4, 5, 6}, expected: []int{}},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{s1: []int{}, s2: []int{1, 2, 3}, expected: []int{}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s1 := From(test.s1...)
			s2 := From(test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			inter := s1.Intersect(s2)
			if !reflect.DeepEqual(inter.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, inter.items)
			}

			if !reflect.DeepEqual(s1, o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !reflect.DeepEqual(s2, o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		s1       []int
		s2       []int
		expected []int // s1 - s2
	}{
		{s1: []int{}, s2: []int{2, 3, 6, 7}, expected: []int{}},
		{s1: []int{2, 3, 6, 7}, s2: []int{}, expected: []int{2, 3, 6, 7}},
		{s1: []int{1, 3, 5, 7}, s2: []int{2, 3, 6, 7}, expected: []int{1, 5}},
		{s1: []int{1, 2, 3}, s2: []int{4, 5, 6}, expected: []int{1, 2, 3}},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: []int{}},
		{s1: []int{1, 2, 3}, s2: []int{}, expected: []int{1, 2, 3}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s1 := From(test.s1...)
			s2 := From(test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			diff := s1.Difference(s2)
			if !reflect.DeepEqual(diff.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, diff.items)
			}

			if !reflect.DeepEqual(s1, o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !reflect.DeepEqual(s2, o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		s1       []int
		s2       []int
		expected []int
	}{
		{s1: []int{}, s2: []int{2, 3, 6, 7}, expected: []int{2, 3, 6, 7}},
		{s1: []int{1, 3, 5, 7}, s2: []int{2, 3, 6, 7}, expected: []int{1, 2, 3, 5, 6, 7}},
		{s1: []int{1, 2, 3}, s2: []int{4, 5, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{s1: []int{}, s2: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s1 := From(test.s1...)
			s2 := From(test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			union := s1.Union(s2)
			if !reflect.DeepEqual(union.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, union.items)
			}

			if !reflect.DeepEqual(s1, o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !reflect.DeepEqual(s2, o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	cases := []struct {
		s1       []int
		s2       []int
		expected []int
	}{
		{s1: []int{}, s2: []int{2, 3, 6, 7}, expected: []int{2, 3, 6, 7}},
		{s1: []int{1, 3, 5, 7}, s2: []int{2, 3, 6, 7}, expected: []int{1, 2, 5, 6}},
		{s1: []int{1, 2, 3}, s2: []int{4, 5, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: []int{}},
		{s1: []int{10, 20}, s2: []int{}, expected: []int{10, 20}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s1 := From(test.s1...)
			s2 := From(test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			sdiff := s1.SymmetricDifference(s2)
			if !reflect.DeepEqual(sdiff.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, sdiff.items)
			}

			if !reflect.DeepEqual(s1, o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !reflect.DeepEqual(s2, o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	cases := []struct {
		s1            []int
		s2            []int
		expectedD12   []int
		expectedInter []int
		expectedD21   []int
	}{
		{
			s1:            []int{},
			s2:            []int{2, 3, 6, 7},
			expectedD12:   []int{},
			expectedInter: []int{},
			expectedD21:   []int{2, 3, 6, 7},
		},
		{
			s1:            []int{1, 3, 5, 7},
			s2:            []int{2, 3, 6, 7},
			expectedD12:   []int{1, 5},
			expectedInter: []int{3, 7},
			expectedD21:   []int{2, 6},
		},
		{
			s1:            []int{1, 2, 3},
			s2:            []int{4, 5, 6},
			expectedD12:   []int{1, 2, 3},
			expectedInter: []int{},
			expectedD21:   []int{4, 5, 6},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s1 := From(test.s1...)
			s2 := From(test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			d12, inter, d21 := s1.Partition(s2)

			if !reflect.DeepEqual(d12.items, test.expectedD12) {
				t.Errorf("D12 Expected %v, got %v", test.expectedD12, d12.items)
			}
			if !reflect.DeepEqual(inter.items, test.expectedInter) {
				t.Errorf("Inter Expected %v, got %v", test.expectedInter, inter.items)
			}
			if !reflect.DeepEqual(d21.items, test.expectedD21) {
				t.Errorf("D21 Expected %v, got %v", test.expectedD21, d21.items)
			}

			if !reflect.DeepEqual(s1, o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !reflect.DeepEqual(s2, o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

type bench struct {
	size int
	vals []int
}

var (
	benchSizes = []int{10, 100, 1000}
	benchs     = make([]bench, len(benchSizes))
)

func init() {
	for i, size := range benchSizes {
		vals := make([]int, size)
		for i := range size {
			vals[i] = rand.Int()
		}

		benchs[i] = bench{
			size: size,
			vals: vals,
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {

			b.Run("slice_set", func(b *testing.B) {
				set := New[int](bench.size)
				b.ResetTimer()
				for i := range b.N {
					set.Add(bench.vals[i%bench.size])
				}
			})

			b.Run("map_set", func(b *testing.B) {
				set := mapset.NewThreadUnsafeSetWithSize[int](bench.size)
				b.ResetTimer()
				for i := range b.N {
					set.Add(bench.vals[i%bench.size])
				}
			})
		})
	}
}

func BenchmarkContains(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {

			b.Run("slice_set", func(b *testing.B) {
				set := New[int](bench.size)
				for _, v := range bench.vals {
					set.Add(v)
				}

				b.ResetTimer()
				for i := range b.N {
					set.Contains(i)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				set := mapset.NewThreadUnsafeSetWithSize[int](bench.size)
				for _, v := range bench.vals {
					set.Add(v)
				}

				b.ResetTimer()
				for i := range b.N {
					set.Contains(i)
				}
			})
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {

			b.Run("slice_set", func(b *testing.B) {
				set := New[int](bench.size)
				for _, v := range bench.vals {
					set.Add(v)
				}

				b.ResetTimer()
				for i := range b.N {
					set.Remove(i)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				set := mapset.NewThreadUnsafeSetWithSize[int](bench.size)
				for _, v := range bench.vals {
					set.Add(v)
				}

				b.ResetTimer()
				for i := range b.N {
					set.Remove(i)
				}
			})
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			// Create two sets with ~50% overlap for a challenging scenario
			vals1 := bench.vals[:bench.size]
			vals2 := make([]int, bench.size)
			copy(vals2, vals1)
			for i := bench.size / 2; i < bench.size; i++ {
				vals2[i] += 100000 // Ensure half are different
			}

			set1 := From(vals1...)
			set2 := From(vals2...)

			map1 := mapset.NewThreadUnsafeSet[int](vals1...)
			map2 := mapset.NewThreadUnsafeSet[int](vals2...)

			b.Run("slice_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					set1.Intersect(set2)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					map1.Intersect(map2)
				}
			})
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			// Create two sets with no overlap for worst-case union
			vals1 := bench.vals[:bench.size/2]
			vals2 := bench.vals[bench.size/2:]

			set1 := From(vals1...)
			set2 := From(vals2...)

			map1 := mapset.NewThreadUnsafeSet(vals1...)
			map2 := mapset.NewThreadUnsafeSet(vals2...)

			b.Run("slice_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					set1.Union(set2)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					map1.Union(map2)
				}
			})
		})
	}
}

func BenchmarkDifference(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			// sets with ~50% overlap for a challenging scenario
			vals1 := bench.vals[:bench.size]
			vals2 := make([]int, bench.size)
			copy(vals2, vals1)
			for i := bench.size / 2; i < bench.size; i++ {
				vals2[i] += 100000
			}

			set1 := From(vals1...)
			set2 := From(vals2...)

			map1 := mapset.NewThreadUnsafeSet(vals1...)
			map2 := mapset.NewThreadUnsafeSet(vals2...)

			b.Run("slice_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					set1.Difference(set2)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					map1.Difference(map2)
				}
			})
		})
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			// Create two sets with ~50% overlap
			vals1 := bench.vals[:bench.size]
			vals2 := make([]int, bench.size)
			copy(vals2, vals1)
			for i := bench.size / 2; i < bench.size; i++ {
				vals2[i] += 100000
			}

			set1 := From(vals1...)
			set2 := From(vals2...)

			map1 := mapset.NewThreadUnsafeSet(vals1...)
			map2 := mapset.NewThreadUnsafeSet(vals2...)

			b.Run("slice_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					set1.SymmetricDifference(set2)
				}
			})

			b.Run("map_set", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					map1.SymmetricDifference(map2)
				}
			})
		})
	}
}
