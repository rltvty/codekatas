package kata05

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitHash(t *testing.T) {
	tables := []struct {
		hash           string
		bitsPerSection int
		result         []uint16
	}{
		{"", 8, []uint16{}},
		{"00", 8, []uint16{0}},
		{"01", 8, []uint16{1}},
		{"0f", 8, []uint16{15}},
		{"10", 8, []uint16{16}},
		{"f0", 8, []uint16{240}},
		{"ff", 8, []uint16{255}},
		{"ff", 3, []uint16{7, 7, 3}},
	}

	for _, table := range tables {
		fmt.Printf("STARTING TEST OF %v spilt into %v bit chunks.\n", table.hash, table.bitsPerSection)
		result := splitHash(table.hash, table.bitsPerSection)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("splitHash of %v was incorrect, got: %v, want: %v.", table.hash, result, table.result)
		} else {
			fmt.Printf("PASS\n")
		}
	}
}
