package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

const (
	minX = -2000
	maxX = 2000
	minY = -2000
	maxY = 2000
)

type Point struct {
	x, y int
}

func distance(a Point, b Point) int {
	return generic.Abs(a.x-b.x) + generic.Abs(a.y-b.y)
}

func StarOne(input []string) string {
	var locations []Point
	for _, l := range input {
		coords := strings.Split(l, ", ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		locations = append(locations, Point{x, y})
	}
	grid := make(map[Point]int)

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			p := Point{x, y}
			lowestVal := maxX
			lowestLoc := 0
			for l, loc := range locations {
				d := distance(loc, p)
				if d < lowestVal {
					lowestVal = d
					lowestLoc = l
				}
			}
			grid[p] = lowestLoc
		}
	}
	areas := make(map[int]int)
	for _, closest := range grid {
		areas[closest]++
	}

	biggestFinite := 0
	for _, area := range areas {
		if area > 10000 {
			continue
		}
		if area > biggestFinite {
			biggestFinite = area
		}
	}

	return strconv.Itoa(biggestFinite)
}

func StarTwo(input []string) string {
	return input[0]
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("1:", StarOne(input))
	fmt.Println("2:", StarTwo(input))
}
