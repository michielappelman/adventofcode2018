package generic

import "strconv"

func Sum(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}

func StringsToInts(list []string) []int {
	var rowInts []int
	for _, c := range list {
		toInt, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		rowInts = append(rowInts, toInt)
	}
	return rowInts
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IndexOfMax(list []int) int {
	var highest int
	for i, v := range list {
		if v > list[highest] {
			highest = i
		}
	}
	return highest
}

func Max(list []int) int {
	var highest int
	for _, v := range list {
		if v > highest {
			highest = v
		}
	}
	return highest
}
