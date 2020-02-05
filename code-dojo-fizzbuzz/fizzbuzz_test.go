package fizzbuzz

import (
	"testing"
)

func TestStage1(t *testing.T) {

	tables := []struct {
		input  int
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{15, "FizzBuzz"},
		{16, "16"},
	}

	for _, table := range tables {
		result := stage1(table.input)
		if result != table.result {
			t.Errorf("FizzBuzz of %v incorrect, got: %v, want: %v.", table.input, result, table.result)
		}
	}
}

func TestStage2(t *testing.T) {

	tables := []struct {
		input  int
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{13, "Fizz"},
		{15, "FizzBuzz"},
		{16, "16"},
		{35, "FizzBuzz"},
		{53, "FizzBuzz"},
		{54, "FizzBuzz"},
		{56, "Buzz"},
		{61, "61"},
	}

	for _, table := range tables {
		result := stage2(table.input)
		if result != table.result {
			t.Errorf("FizzBuzz of %v incorrect, got: %v, want: %v.", table.input, result, table.result)
		}
	}
}
