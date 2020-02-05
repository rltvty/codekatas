package fizzbuzz

import (
	"strconv"
)

func stage1(input int) string {
	x := ""
	if input%3 == 0 {
		x += "Fizz"
	}
	if input%5 == 0 {
		x += "Buzz"
	}
	if x == "" {
		x = strconv.Itoa(input)
	}
	return x
}

func stage2(input int) string {
	fizz := false
	buzz := false
	if input%3 == 0 {
		fizz = true
	}
	if input%5 == 0 {
		buzz = true
	}
	s := strconv.Itoa(input)
	for _, r := range s {
		if r == '3' {
			fizz = true
		}
		if r == '5' {
			buzz = true
		}
	}
	switch {
	case fizz && buzz:
		return "FizzBuzz"
	case fizz:
		return "Fizz"
	case buzz:
		return "Buzz"
	default:
		return s
	}
}
