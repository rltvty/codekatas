package kata04

import (
	"reflect"
	"testing"
)

func TestWeather(t *testing.T) {
	if weather() != "9" {
		t.Error("found incorrect day with the largest temperature span")
	}
}

func TestTeam(t *testing.T) {
	if team() != "Aston_Villa" {
		t.Error("found incorrect team with the least goals difference")
	}
}

func TestGetColumns(t *testing.T) {
	row := " 9  86    32*   59       6  61.5       0.00         240  7.6 220  12  6.0  78 46 1018.6"
	query := []int{1, 2}
	want := []string{"86", "32"}
	result, err := getColumns(row, query, "*")
	if err != nil {
		t.Error("getColumns should not have errored")
	}
	if !reflect.DeepEqual(want, result) {
		t.Errorf("getColumns of (%v) in (%v) was incorrect, got: %v, want: %v.", query, row, result, want)
	}
	result, err = getColumns("", query, "")
	if err == nil {
		t.Error("getColumns should have errored when more columns are requested than exist")
	}
	if result != nil {
		t.Errorf("getColumns result should be nil when more columns are requested than exist")
	}
}
