package smallset

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

type Person struct {
	ID   int
	Name string
	Age  int
	Data string
}

func PersonCmp(a, b Person) int {
	if a.ID < b.ID {
		return -1
	}
	if a.ID > b.ID {
		return 1
	}
	return 0
}

var (
	people1 = []Person{
		{ID: 2, Name: "Charlie", Age: 30},
		{ID: 3, Name: "Alice", Age: 25},
		{ID: 4, Name: "Eve", Age: 40},
		{ID: 2, Name: "Carly (Duplicate)", Age: 31},
		{ID: 1, Name: "Bob", Age: 50},
		{ID: 4, Name: "Eva (Duplicate)", Age: 41},
	}

	unique1 = []Person{
		{ID: 1, Name: "Bob", Age: 50},
		{ID: 2, Name: "Charlie", Age: 30},
		{ID: 3, Name: "Alice", Age: 25},
		{ID: 4, Name: "Eve", Age: 40},
	}

	people2 = []Person{
		{ID: 50, Name: "Alpha", Age: 5},
		{ID: 40, Name: "Beta", Age: 4},
		{ID: 40, Name: "Beta (Duplicate)", Age: 41},
		{ID: 30, Name: "Gamma", Age: 3},
		{ID: 30, Name: "Gamma (Duplicate)", Age: 31},
		{ID: 30, Name: "Gamma (Duplicate)", Age: 32},
		{ID: 20, Name: "Delta", Age: 2},
	}

	unique2 = []Person{
		{ID: 20, Name: "Delta", Age: 2},
		{ID: 30, Name: "Gamma", Age: 3},
		{ID: 40, Name: "Beta", Age: 4},
		{ID: 50, Name: "Alpha", Age: 5},
	}
)

func TestCustomContains(t *testing.T) {
	s := NewCustomFrom(PersonCmp, people1...)

	cases := []struct {
		element  Person
		expected bool
	}{
		{element: Person{ID: 1}, expected: true},
		{element: Person{ID: 69}, expected: false},
		{element: Person{ID: 11}, expected: false},
		{element: Person{}, expected: false},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := s.Contains(test.element)
			if result != test.expected {
				t.Errorf("Contains(%v) expected %v got %v", test.element, test.expected, result)
			}
		})
	}
}

func TestCustomAdd(t *testing.T) {
	cases := []struct {
		toAdd    []Person
		expected []bool
		items    []Person
	}{
		{
			toAdd:    people1,
			expected: []bool{true, true, true, false, true, false},
			items:    unique1,
		},
		{
			toAdd:    people2,
			expected: []bool{true, true, false, true, false, false, true},
			items:    unique2,
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := NewCustom(PersonCmp, 10)
			res := make([]bool, len(test.toAdd))
			for j, e := range test.toAdd {
				res[j] = s.Add(e)
			}

			if !slices.Equal(res, test.expected) {
				t.Errorf("Add results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !slices.Equal(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestCustomRemove(t *testing.T) {
	cases := []struct {
		initial  []Person
		toRemove []Person
		expected []bool
		items    []Person
	}{
		{
			initial:  people1,
			toRemove: []Person{{ID: 1}, {ID: 2}, {ID: 69}},
			expected: []bool{true, true, false},
			items:    []Person{{ID: 3, Name: "Alice", Age: 25}, {ID: 4, Name: "Eve", Age: 40}},
		},
		{
			initial:  people2,
			toRemove: []Person{{ID: 20}, {ID: 50}},
			expected: []bool{true, true},
			items:    []Person{{ID: 30, Name: "Gamma", Age: 3}, {ID: 40, Name: "Beta", Age: 4}},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := NewCustomFrom(PersonCmp, test.initial...)
			res := make([]bool, len(test.toRemove))
			for j, e := range test.toRemove {
				res[j] = s.Remove(e)
			}

			if !slices.Equal(res, test.expected) {
				t.Errorf("Remove results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !slices.Equal(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}
func TestCustomRemoveBefore(t *testing.T) {
	cases := []struct {
		initial  []Person
		max      Person
		expected int
		items    []Person
	}{
		{
			initial:  people1,
			max:      Person{ID: 3},
			expected: 2,
			items:    []Person{{ID: 3, Name: "Alice", Age: 25}, {ID: 4, Name: "Eve", Age: 40}},
		},
		{
			initial:  people2,
			max:      Person{ID: 20},
			expected: 0,
			items:    unique2,
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := NewCustomFrom(PersonCmp, test.initial...)
			res := s.RemoveBefore(test.max)

			if res != test.expected {
				t.Errorf("Remove results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !slices.Equal(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestCustomRemoveFrom(t *testing.T) {
	cases := []struct {
		initial  []Person
		min      Person
		expected int
		items    []Person
	}{
		{
			initial:  people1,
			min:      Person{ID: 3},
			expected: 2,
			items:    []Person{{ID: 1, Name: "Bob", Age: 50}, {ID: 2, Name: "Charlie", Age: 30}},
		},
		{
			initial:  people2,
			min:      Person{ID: 20},
			expected: 4,
			items:    []Person{},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := NewCustomFrom(PersonCmp, test.initial...)
			res := s.RemoveFrom(test.min)

			if res != test.expected {
				t.Errorf("Remove results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !slices.Equal(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestCustomRemoveBetween(t *testing.T) {
	cases := []struct {
		initial  []Person
		min, max Person
		expected int
		items    []Person
	}{
		{
			initial: people1,
			min:     Person{ID: 2}, max: Person{ID: 4},
			expected: 2,
			items:    []Person{{ID: 1, Name: "Bob", Age: 50}, {ID: 4, Name: "Eve", Age: 40}},
		},
		{
			initial: people2,
			min:     Person{ID: 20}, max: Person{ID: 50},
			expected: 3,
			items:    []Person{{ID: 50, Name: "Alpha", Age: 5}},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			s := NewCustomFrom(PersonCmp, test.initial...)
			res := s.RemoveBetween(test.min, test.max)

			if res != test.expected {
				t.Errorf("Remove results mismatch.\nExpected: %v\nActual: %v", test.expected, res)
			}

			if !slices.Equal(s.items, test.items) {
				t.Errorf("Items mismatch.\nExpected: %v\nActual: %v", test.items, s.items)
			}
		})
	}
}

func TestCustomIsEqual(t *testing.T) {
	s1 := NewCustomFrom(cmp.Compare[int], 1, 2, 3)
	s2 := NewCustomFrom(cmp.Compare[int], 3, 2, 1)
	s3 := NewCustomFrom(cmp.Compare[int], 1, 2, 3, 4)
	s4 := NewCustom(cmp.Compare[int], 10)

	cases := []struct {
		setA     *Custom[int]
		setB     *Custom[int]
		expected bool
	}{
		{setA: s1, setB: s2, expected: true},
		{setA: s1, setB: s3, expected: false},
		{setA: s1, setB: s4, expected: false},
		{setA: s4, setB: NewCustom(cmp.Compare[int], 30), expected: true},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			if res := test.setA.IsEqual(test.setB); res != test.expected {
				t.Errorf("IsEqual expected %t, got %t", test.expected, res)
			}
		})
	}
}

func TestCustomMin(t *testing.T) {
	cases := []struct {
		set      *Custom[int]
		expected int
	}{
		{set: NewCustomFrom(cmp.Compare[int], 10, 5, 20, 15), expected: 5},
		{set: NewCustomFrom(cmp.Compare[int], 1, 5, 20, 69), expected: 1},
		{set: NewCustomFrom(cmp.Compare[int], 7, 8, 4, 12, 221), expected: 4},
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

func TestCustomMax(t *testing.T) {
	cases := []struct {
		set      *Custom[int]
		expected int
	}{
		{set: NewCustomFrom(cmp.Compare[int], 10, 5, 20, 15), expected: 20},
		{set: NewCustomFrom(cmp.Compare[int], 1, 5, 20, 69), expected: 69},
		{set: NewCustomFrom(cmp.Compare[int], 7, 8, 4, 12, 221), expected: 221},
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

func TestCustomMinK(t *testing.T) {
	cases := []struct {
		set      *Custom[int]
		k        int
		expected []int
	}{
		{set: NewCustomFrom(cmp.Compare[int], 10, 5, 20, 15), k: 2, expected: []int{5, 10}},
		{set: NewCustomFrom(cmp.Compare[int], 7, 8, 4, 12, 221), k: 150, expected: []int{4, 7, 8, 12, 221}},
		{set: NewCustomFrom(cmp.Compare[int], 1, 5, 20, 69), k: 0, expected: []int{}},
		{set: NewCustom(cmp.Compare[int], 10), k: 5, expected: []int{}},
		{set: NewCustom(cmp.Compare[int], 10), k: 0, expected: []int{}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.MinK(test.k)
			if !slices.Equal(result, test.expected) {
				t.Errorf("MinK(%d) failed.\nExpected: %v\nActual: %v", test.k, test.expected, result)
			}
		})
	}
}

func TestCustomMaxK(t *testing.T) {
	cases := []struct {
		set      *Custom[int]
		k        int
		expected []int
	}{
		{set: NewCustomFrom(cmp.Compare[int], 10, 5, 20, 15), k: 2, expected: []int{15, 20}},
		{set: NewCustomFrom(cmp.Compare[int], 7, 8, 4, 12, 221), k: 150, expected: []int{4, 7, 8, 12, 221}},
		{set: NewCustomFrom(cmp.Compare[int], 1, 5, 20, 69), k: 0, expected: []int{}},
		{set: NewCustom(cmp.Compare[int], 10), k: 5, expected: []int{}},
		{set: NewCustom(cmp.Compare[int], 10), k: 0, expected: []int{}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			result := test.set.MaxK(test.k)
			if !slices.Equal(result, test.expected) {
				t.Errorf("MaxK(%d) failed.\nExpected: %v\nActual: %v", test.k, test.expected, result)
			}
		})
	}
}

func TestCustomBetweenAsc(t *testing.T) {
	s := NewCustomFrom(cmp.Compare[int], 1, 3, 5, 7, 9)

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
			result := collect(s.BetweenAsc(test.min, test.max))
			if !slices.Equal(result, test.expected) {
				t.Errorf("BetweenAsc(%d, %d) failed.\nExpected: %v\nActual: %v", test.min, test.max, result, test.expected)
			}
		})
	}
}

func TestCustomBetweenDesc(t *testing.T) {
	s := NewCustomFrom(cmp.Compare[int], 1, 3, 5, 7, 9)

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
			result := collect(s.BetweenDesc(test.max, test.min))
			if !slices.Equal(result, test.expected) {
				t.Errorf("BetweenDesc(%d, %d) failed.\nExpected: %v\nActual: %v", test.max, test.min, test.expected, result)
			}
		})
	}
}

// --- Binary Set Operation TestCustoms ---

func TestCustomIntersect(t *testing.T) {
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
			s1 := NewCustomFrom(cmp.Compare[int], test.s1...)
			s2 := NewCustomFrom(cmp.Compare[int], test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			inter := s1.Intersect(s2)
			if !slices.Equal(inter.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, inter.items)
			}

			if !s1.IsEqual(o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !s2.IsEqual(o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestCustomDifference(t *testing.T) {
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
			s1 := NewCustomFrom(cmp.Compare[int], test.s1...)
			s2 := NewCustomFrom(cmp.Compare[int], test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			diff := s1.Difference(s2)
			if !slices.Equal(diff.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, diff.items)
			}

			if !s1.IsEqual(o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !s2.IsEqual(o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestCustomUnion(t *testing.T) {
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
			s1 := NewCustomFrom(cmp.Compare[int], test.s1...)
			s2 := NewCustomFrom(cmp.Compare[int], test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			union := s1.Union(s2)
			if !slices.Equal(union.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, union.items)
			}

			if !s1.IsEqual(o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !s2.IsEqual(o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestCustomSymmetricDifference(t *testing.T) {
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
			s1 := NewCustomFrom(cmp.Compare[int], test.s1...)
			s2 := NewCustomFrom(cmp.Compare[int], test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			sdiff := s1.SymmetricDifference(s2)
			if !slices.Equal(sdiff.items, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, sdiff.items)
			}

			if !s1.IsEqual(o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !s2.IsEqual(o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}

func TestCustomPartition(t *testing.T) {
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
			s1 := NewCustomFrom(cmp.Compare[int], test.s1...)
			s2 := NewCustomFrom(cmp.Compare[int], test.s2...)

			o1 := s1.Clone()
			o2 := s2.Clone()

			d12, inter, d21 := s1.Partition(s2)

			if !slices.Equal(d12.items, test.expectedD12) {
				t.Errorf("D12 Expected %v, got %v", test.expectedD12, d12.items)
			}
			if !slices.Equal(inter.items, test.expectedInter) {
				t.Errorf("Inter Expected %v, got %v", test.expectedInter, inter.items)
			}
			if !slices.Equal(d21.items, test.expectedD21) {
				t.Errorf("D21 Expected %v, got %v", test.expectedD21, d21.items)
			}

			if !s1.IsEqual(o1) {
				t.Errorf("s1 mutated. before %v, after %v", o1, s1)
			}
			if !s2.IsEqual(o2) {
				t.Errorf("s2 mutated. before %v, after %v", o2, s2)
			}
		})
	}
}
