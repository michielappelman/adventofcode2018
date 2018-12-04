package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

type Square struct {
	col, row int
}

func StarOne(input []string) string {
	fabric := make(map[Square][]int)
	pattern := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	for _, claim := range input {
		matches := pattern.FindAllStringSubmatch(claim, 1)
		ints, _ := generic.StringsToInts(matches[0][1:])
		for x := ints[1]; x < ints[1]+ints[3]; x++ {
			for y := ints[2]; y < ints[2]+ints[4]; y++ {
				fabric[Square{x, y}] = append(fabric[Square{x, y}], ints[0])
			}
		}
	}
	double := 0
	for _, s := range fabric {
		if len(s) > 1 {
			double++
		}
	}
	return strconv.Itoa(double)
}

func StarTwo(input []string) string {
	fabric := make(map[Square][]int)
	pattern := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	ids := make(map[int]bool)
	var maxX, maxY int
	for _, claim := range input {
		matches := pattern.FindAllStringSubmatch(claim, 1)
		ints, _ := generic.StringsToInts(matches[0][1:])
		ids[ints[0]] = true
		for x := ints[1]; x < ints[1]+ints[3]; x++ {
			for y := ints[2]; y < ints[2]+ints[4]; y++ {
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
				fabric[Square{x, y}] = append(fabric[Square{x, y}], ints[0])
			}
		}
	}
	for _, s := range fabric {
		if len(s) > 1 {
			for _, c := range s {
				delete(ids, c)
			}
		}
	}
	var left int
	for k, _ := range ids {
		left = k
	}

	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	for xy, s := range fabric {
		if len(s) > 1 {
			img.Set(xy.col, xy.row, color.RGBA{255, 0, 0, 255})
		} else if s[0] == left {
			img.Set(xy.col, xy.row, color.RGBA{0, 0, 255, 255})
		} else {
			img.Set(xy.col, xy.row, color.RGBA{0, 255, 0, 255})
		}
	}
	f, _ := os.OpenFile("fabric.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	return strconv.Itoa(left)
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
