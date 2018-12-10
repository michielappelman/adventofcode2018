package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

type Light struct {
	position Position
	velocity Velocity
}

type Position struct {
	x, y int
}

type Velocity struct {
	dx, dy int
}

func (l *Light) Step() {
	l.position.x += l.velocity.dx
	l.position.y += l.velocity.dy
}

func getArea(lights []*Light) int {
	var minX, minY, maxX, maxY int
	for _, l := range lights {
		if l.position.x < minX {
			minX = l.position.x
		}
		if l.position.y < minY {
			minY = l.position.y
		}
		if l.position.x > maxX {
			maxX = l.position.x
		}
		if l.position.y > maxY {
			maxY = l.position.y
		}
	}
	width := maxX - minX
	height := maxY - minY
	return width * height
}

func StarOne(input []string) string {
	var lights []*Light
	pattern := regexp.MustCompile(`^.+<\s*(.+),\s*(.+)>.+<\s*(.+),\s*(.+)>.*$`)

	for _, l := range input {
		matches := pattern.FindAllStringSubmatch(l, 1)
		ints, err := generic.StringsToInts(matches[0][1:])
		if err != nil {
			log.Fatalf("error in strings->ints: %v", err)
		}
		lights = append(lights, &Light{Position{ints[0], ints[1]}, Velocity{ints[2], ints[3]}})
	}

	steps := 15000

	var messageSky []Light
	var maxX, maxY int
	minArea := getArea(lights)

	for s := 0; s < steps; s++ {
		area := getArea(lights)
		if area < minArea {
			minArea = area
			messageSky = make([]Light, 0)
			maxX = 0
			maxY = 0
			for _, l := range lights {
				if l.position.x > maxX {
					maxX = l.position.x
				}
				if l.position.y > maxY {
					maxY = l.position.y
				}
				messageSky = append(messageSky, *l)
			}
		}
		for _, l := range lights {
			l.Step()
		}
	}

	magnify := 4
	img := image.NewRGBA(image.Rect(0, 0, maxX*magnify+magnify*2, maxY*magnify+magnify*2))

	for _, l := range messageSky {
		for x := l.position.x * magnify; x < l.position.x*magnify+magnify; x++ {
			for y := l.position.y * magnify; y < l.position.y*magnify+magnify; y++ {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}

	f, _ := os.OpenFile("sky.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	return "cannot test this (easily)"
}

func StarTwo(input []string) string {
	var lights []*Light
	pattern := regexp.MustCompile(`^.+<\s*(.+),\s*(.+)>.+<\s*(.+),\s*(.+)>.*$`)

	for _, l := range input {
		matches := pattern.FindAllStringSubmatch(l, 1)
		ints, _ := generic.StringsToInts(matches[0][1:])
		lights = append(lights, &Light{Position{ints[0], ints[1]}, Velocity{ints[2], ints[3]}})
	}

	steps := 15000
	minSecond := 0
	minArea := getArea(lights)

	for s := 0; s < steps; s++ {
		area := getArea(lights)
		if area < minArea {
			minArea = area
			minSecond = s
		}
		for _, l := range lights {
			l.Step()
		}
	}

	return strconv.Itoa(minSecond)
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
