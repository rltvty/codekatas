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
		{"00", 8, []uint16{0}},                      //0000-0000
		{"01", 8, []uint16{1}},                      //0000-0001
		{"0f", 8, []uint16{15}},                     //0000-1111
		{"10", 8, []uint16{16}},                     //0001-0000
		{"f0", 8, []uint16{240}},                    //1111-0000
		{"ff", 8, []uint16{255}},                    //1111-1111
		{"ff", 3, []uint16{7, 7, 3}},                //011 111 111
		{"f2838", 7, []uint16{56, 80, 60}},          //011-1100 101-0000 011-1000
		{"DA741EB8", 9, []uint16{184, 15, 157, 27}}, //0-0001-1011 0-1001-1101 0-0000-1111 0-1011-1000
		{"DA741EB8", 12, []uint16{3768, 1857, 218}}, //0000-1101-1010 0111-0100-0001 1110-1011-1000
		{"DA741EB8", 11, []uint16{1720, 1667, 873}}, //011-0110-1001 110-1000-0011 110-1011-1000
	}

	for _, table := range tables {
		result := splitHash(table.hash, table.bitsPerSection)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("splitHash of %v was incorrect, got: %v, want: %v.", table.hash, result, table.result)
		} else {
			fmt.Printf("PASS\n")
		}
	}
}

func TestGetBitmap(t *testing.T) {
	tables := []struct {
		hash   string
		result string
	}{
		{"", ""},
		{"00", "0000 0000"},
		{"01", "0000 0001"},
		{"0f", "0000 1111"},
		{"10", "0001 0000"},
		{"f0", "1111 0000"},
		{"ff", "1111 1111"},
		{"f2838", "1111 0010 1000 0011 1000"},
	}

	for _, table := range tables {
		result := ""
		for i, v := range getBitmap(table.hash) {
			if i > 0 && i%4 == 0 {
				result += " "
			}
			if v {
				result += "1"
			} else {
				result += "0"
			}

		}

		if result != table.result {
			t.Errorf("getBitmap of %v was incorrect, got: %v, want: %v.", table.hash, result, table.result)
		} else {
			fmt.Printf("PASS\n")
		}
	}
}

func TestReverseByes(t *testing.T) {
	tables := []struct {
		in     []byte
		result []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{2}, []byte{2}},
		{[]byte{1, 2}, []byte{2, 1}},
		{[]byte{1, 2, 3}, []byte{3, 2, 1}},
	}

	for _, table := range tables {
		result := reverseBytes(table.in)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("reverseBytes of %v was incorrect, got: %v, want: %v.", table.in, result, table.result)
		} else {
			fmt.Printf("PASS\n")
		}
	}
}
