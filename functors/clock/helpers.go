package clock

var Unit = func(i int) int {
	return i
}

var AmPmMapper = func(i int) int {
	return (i + 12) % 24
}

func AmHoursFn() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}

