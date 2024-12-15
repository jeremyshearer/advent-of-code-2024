package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

const (
	wall  = "#"
	box   = "O"
	empty = "."
	robot = "@"
	left  = "<"
	right = ">"
	up    = "^"
	down  = "v"
)

var movements = map[string]image.Point{
	up:    {Y: -1, X: 0},
	right: {Y: 0, X: 1},
	down:  {Y: 1, X: 0},
	left:  {Y: 0, X: -1},
}

func main() {
	input, err := os.ReadFile("./cmd/day15a/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputParts := strings.Split(string(input), "\n\n")
	rows := strings.Split(inputParts[0], "\n")
	w := &Warehouse{
		grid:       map[int]map[int]string{},
		dimensions: image.Point{X: len(rows[0]), Y: len(rows)},
	}
	for y, row := range rows {
		for x, loc := range strings.Split(row, "") {
			if w.grid[y] == nil {
				w.grid[y] = map[int]string{}
			}
			if loc == robot {
				w.robot.X = x
				w.robot.Y = y
			}
			w.grid[y][x] = loc
		}
	}
	w.Print()

	for _, dir := range strings.Split(inputParts[1], "") {
		w.moveRobot(dir)
	}

	w.Print()
	fmt.Println(w.score())
}

type Warehouse struct {
	grid       map[int]map[int]string
	dimensions image.Point
	robot      image.Point
}

func (w *Warehouse) Print() {
	for y := range w.dimensions.Y {
		for x := range w.dimensions.X {
			v, ok := w.grid[y][x]
			if !ok {
				fmt.Print(empty)
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func (w *Warehouse) Space(l image.Point) string {
	v, ok := w.grid[l.Y][l.X]
	if !ok {
		return empty
	}
	return v
}

func (w *Warehouse) Shift(p, m image.Point) bool {
	v := w.Space(p)
	if v == empty {
		fmt.Println("tried to shift an empty space")
		return false
	}

	n := p.Add(m)
	nv := w.Space(n)
	if nv == wall {
		return false
	} else if nv == box {
		if w.Shift(n, m) {
			w.UpdatePosition(p, n)
			return true
		}
	} else if nv == empty {
		w.UpdatePosition(p, n)
		return true
	}

	return false
}

func (w *Warehouse) UpdatePosition(p, n image.Point) {
	v := w.grid[p.Y][p.X]
	w.grid[p.Y][p.X] = empty
	w.grid[n.Y][n.X] = v
	if v == robot {
		w.robot = n
	}
}

func (w *Warehouse) moveRobot(dir string) {
	move := movements[dir]
	w.Shift(w.robot, move)
}

func (w *Warehouse) score() int {
	var total int
	for y, cells := range w.grid {
		for x, cell := range cells {
			if cell == box {
				total += y*100 + x
			}
		}
	}
	return total
}
