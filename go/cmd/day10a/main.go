package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

var movements = []image.Point{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func main() {
	input, err := os.ReadFile("./cmd/day10a/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]int

	var trailHeads []image.Point
	for y, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var row []int
		for x, cell := range strings.Split(line, "") {
			if cell == "." {
				row = append(row, 100)
			} else {
				elevation, _ := strconv.Atoi(cell)
				row = append(row, elevation)
				if elevation == 0 {
					trailHeads = append(trailHeads, image.Point{X: x, Y: y})
				}
			}
		}
		grid = append(grid, row)
	}

	var total int
	for _, th := range trailHeads {
		total += solve(th, grid)
	}
	fmt.Printf("trails: %d\n", total)
}

func solve(start image.Point, grid [][]int) int {
	visited := make(map[int]map[int]bool)
	var trail []image.Point
	walk(start, -1, grid, &trail, visited)
	var score int
	for _, p := range trail {
		if grid[p.Y][p.X] == 9 {
			score++
		}
	}
	return score
}

func walk(curr image.Point, previousElevation int, grid [][]int, path *[]image.Point, visited map[int]map[int]bool) bool {
	if curr.Y < 0 || curr.Y > len(grid)-1 || curr.X < 0 || curr.X > len(grid[curr.Y])-1 {
		return false
	}

	currElevation := grid[curr.Y][curr.X]

	if visited[curr.Y] != nil && visited[curr.Y][curr.X] {
		return false
	}

	if currElevation-previousElevation != 1 {
		return false
	}

	if visited[curr.Y] == nil {
		visited[curr.Y] = map[int]bool{curr.X: true}
	} else {
		visited[curr.Y][curr.X] = true
	}
	*path = append(*path, curr)

	for _, m := range movements {
		walk(curr.Add(m), currElevation, grid, path, visited)
	}

	return false
}
