package kata02

import "testing"

func TestChop1(t *testing.T) {
	emptyHaystack := []int{}
	tinyHaystack := []int{1}
	smallHaystack := []int{1, 3, 5}
	medHaystack := []int{1, 3, 5, 7}

	tables := []struct {
		needle   int
		haystack []int
		result   int
	}{
		{3, emptyHaystack, -1},
		{3, tinyHaystack, -1},
		{1, tinyHaystack, 0},
		{1, smallHaystack, 0},
		{3, smallHaystack, 1},
		{5, smallHaystack, 2},
		{0, smallHaystack, -1},
		{4, smallHaystack, -1},
		{1, medHaystack, 0},
		{3, medHaystack, 1},
		{5, medHaystack, 2},
		{7, medHaystack, 3},
		{0, medHaystack, -1},
		{8, medHaystack, -1},
	}

	for _, table := range tables {
		result := Chop1(table.needle, table.haystack)
		if result != table.result {
			t.Errorf("Chop1 of (%v) in (%v) was incorrect, got: %v, want: %v.", table.needle, table.haystack, result, table.result)
		}
	}
}
