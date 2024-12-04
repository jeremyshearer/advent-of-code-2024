package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

var directions = map[string][]int{
	"U":  {-1, 0},
	"UR": {-1, 1},
	"R":  {0, 1},
	"DR": {1, 1},
	"D":  {1, 0},
	"DL": {1, -1},
	"L":  {0, -1},
	"UL": {-1, -1},
}

type wordSearch struct {
	grid [][]string
}

func (w wordSearch) print() {
	for y := range len(w.grid) {
		for x := range len(w.grid[y]) {
			fmt.Println(w.grid[y][x])
		}
	}
}

func (w wordSearch) collectFour(dir string, yPos, xPos int) string {
	sb := strings.Builder{}
	currX := xPos
	currY := yPos
	deltas := directions[dir]
	for range 4 {
		if currY < 0 || currY > len(w.grid)-1 || currX < 0 || currX > len(w.grid[currY])-1 {
			return sb.String()
		}
		sb.WriteString(w.grid[currY][currX])
		currY += deltas[0]
		currX += deltas[1]
	}
	return sb.String()
}

func newWordSearch(input string) wordSearch {
	var grid [][]string
	for _, row := range strings.Split(input, "\n") {
		var cols []string
		letters := strings.Split(row, "")
		for _, letter := range letters {
			cols = append(cols, letter)
		}
		grid = append(grid, cols)
	}
	return wordSearch{
		grid: grid,
	}
}

func main() {
	input, err := os.ReadFile("./cmd/day4a/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	ws := newWordSearch(string(input))
	var wordCount int
	for y := range len(ws.grid) {
		for x := range len(ws.grid[y]) {
			for dir := range maps.Keys(directions) {
				word := ws.collectFour(dir, y, x)
				if word == "XMAS" {
					wordCount++
				}
			}
		}
	}
	fmt.Printf("%d\n", wordCount)
}
