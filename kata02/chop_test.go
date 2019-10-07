package kata02

import "testing"

func TestChops(t *testing.T) {
	emptyHaystack := []int{}
	tinyHaystack := []int{1}
	smallHaystack := []int{1, 3, 5}
	medHaystack := []int{1, 3, 5, 7}
	largeHaystack := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 22, 24, 26, 28, 29, 30, 34, 35, 38, 39, 100}

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
		{7, largeHaystack, 3},
		{19, largeHaystack, 9},
		{35, largeHaystack, 18},
		{99, largeHaystack, -1},
		{101, largeHaystack, -1},
	}

	chops := []Chop{Chop1, Chop2}

	for i, chop := range chops {
		for _, table := range tables {
			result := chop(table.needle, table.haystack)
			if result != table.result {
				t.Errorf("Chop%d of (%v) in (%v) was incorrect, got: %v, want: %v.", i+1, table.needle, table.haystack, result, table.result)
			}
		}
	}

}
