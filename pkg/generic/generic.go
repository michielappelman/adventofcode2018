package generic

import "strconv"

// Sum takes a slice of ints and returns their total sum.
func Sum(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}

// StringsToInts takes a slice of strings containing ints and returns the slice of ints or an
// error.
func StringsToInts(list []string) ([]int, error) {
	var rowInts []int
	for _, c := range list {
		toInt, err := strconv.Atoi(c)
		if err != nil {
			return rowInts, err
		}
		rowInts = append(rowInts, toInt)
	}
	return rowInts, nil
}

// Abs returns the absolute value of an integer.
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// IndexOfMax returns the index of the highest integer in a slice of integers.
func IndexOfMax(list []int) int {
	var highest int
	for i, v := range list {
		if v > list[highest] {
			highest = i
		}
	}
	return highest
}

// Max returns the highest integer in a slice of integers.
func Max(list []int) int {
	var highest int
	for _, v := range list {
		if v > highest {
			highest = v
		}
	}
	return highest
}

// Min returns the lowest integer in a slice of integers.
func Min(list []int) int {
	var lowest int
	for _, v := range list {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}
