package kata02

type Chop func(int, []int) int

func Chop1(needle int, haystack []int) int {
	haystackLength := len(haystack)
	switch haystackLength {
	case 0:
		return -1
	case 1:
		if haystack[0] == needle {
			return 0
		}
		return -1
	default:
		half1 := haystack[:haystackLength/2]
		half2 := haystack[haystackLength/2:]
		inHalf1 := Chop1(needle, half1)
		if inHalf1 >= 0 {
			return inHalf1
		}
		inHalf2 := Chop1(needle, half2)
		if inHalf2 >= 0 {
			return haystackLength/2 + inHalf2
		}
		return -1
	}
}

func chop2(index int, needle int, haystack []int, resultChannel chan int) {
	if haystackLength := len(haystack); haystackLength > 1 {
		chop2(index, needle, haystack[:haystackLength/2], resultChannel)
		chop2(index+(haystackLength/2), needle, haystack[haystackLength/2:], resultChannel)
	} else if haystackLength == 1 {
		if haystack[0] == needle {
			resultChannel <- index
		} else {
			resultChannel <- -1
		}
	}
}

func Chop2(needle int, haystack []int) int {
	resultChannel := make(chan int, len(haystack))
	go chop2(0, needle, haystack, resultChannel)
	for i := 0; i < len(haystack); i++ {
		if result := <-resultChannel; result >= 0 {
			return result
		}
	}
	return -1
}
