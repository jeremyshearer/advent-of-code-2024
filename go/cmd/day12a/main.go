package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

var directions = []image.Point{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func main() {
	input, err := os.ReadFile("./cmd/day12a/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	plots := []plot{}
	var garden [][]string
	for _, row := range strings.Split(string(input), "\n") {
		plants := strings.Split(row, "")
		garden = append(garden, plants)
	}
	visited := map[int]map[int]bool{}
	for y, row := range garden {
		for x, plant := range row {
			p := buildPlot(garden, visited, image.Point{X: x, Y: y}, plant)
			if len(p) > 0 {
				plots = append(plots, p)
			}
		}
	}
	var total int
	for _, p := range plots {
		total += p.cost()
	}
	fmt.Println(total)
}

func buildPlot(garden [][]string, visited map[int]map[int]bool, start image.Point, plant string) plot {
	p := &plot{}
	check(garden, p, visited, start, plant)
	return *p
}

func check(garden [][]string, p *plot, visited map[int]map[int]bool, curr image.Point, plant string) {
	if curr.Y < 0 || curr.Y > len(garden)-1 || curr.X < 0 || curr.X > len(garden[curr.Y])-1 {
		return
	}

	if visited[curr.Y] == nil {
		visited[curr.Y] = map[int]bool{}
	}

	match := plant == garden[curr.Y][curr.X]

	if !match || visited[curr.Y][curr.X] {
		return
	}

	visited[curr.Y][curr.X] = true

	*p = append(*p, curr)
	for _, dir := range directions {
		check(garden, p, visited, curr.Add(dir), plant)
	}
}

type plot []image.Point

func (p plot) perimeter() int {
	if len(p) == 1 {
		return 4
	}

	var total int
	for _, plant := range p {
		plantBorder := 4
		for _, op := range p {
			delta := plant.Sub(op)
			if delta.X == 0 && delta.Y == 0 {
				continue
			}

			if (delta.X == 1 && delta.Y == 0) || (delta.X == -1 && delta.Y == 0) || (delta.X == 0 && delta.Y == 1) || (delta.X == 0 && delta.Y == -1) {
				plantBorder--
			}
		}
		total += plantBorder
	}

	return total
}

func (p plot) cost() int {
	return p.perimeter() * len(p)
}
