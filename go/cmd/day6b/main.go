package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day6b/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var grid maze
	rows := strings.Split(string(input), "\n")
	var l *lab
	var guardX int
	var guardY int
	for y, row := range rows {
		cells := strings.Split(row, "")
		for x, cell := range cells {
			if cell == "^" {
				guardX = x
				guardY = y
			}
		}
		grid = append(grid, cells)
	}

	var loopCount int
	for y, row := range grid {
		for x := range row {
			loopGrid := make(maze, 0, len(grid))
			for _, rowToCopy := range grid {
				loopRow := make([]string, len(rowToCopy))
				copy(loopRow, rowToCopy)
				loopGrid = append(loopGrid, loopRow)
			}
			if !(y == guardY && x == guardX) {
				loopGrid[y][x] = "O"
			}
			l = &lab{
				guardX:          guardX,
				guardY:          guardY,
				facingDirection: "N",
				grid:            loopGrid,
				seenObstacles: map[int]map[int]int{
					guardY: {
						guardX: 1,
					},
				},
			}
			l.run()
			if l.guardLoop {
				loopCount++
			}
		}
	}
	fmt.Println(loopCount)
}

var directions = []string{"N", "E", "S", "W"}
var movements = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type maze [][]string

func (l *lab) Print() {
	for y, row := range l.grid {
		for x, cell := range row {
			if y == l.guardY && x == l.guardX {
				fmt.Print("G")
			} else if l.seenObstacles[y][x] > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
	fmt.Println("==========")
}

type lab struct {
	grid            [][]string
	seenObstacles   map[int]map[int]int
	guardX          int
	guardY          int
	facingDirection string
	guardExit       bool
	guardLoop       bool
}

func (l *lab) run() {
	for {
		if l.guardExit {
			break
		}

		if l.guardLoop {
			break
		}

		l.moveGuard()
	}
}

func (l *lab) turnRight() {
	curr := slices.Index(directions, l.facingDirection)
	newIndex := (curr + 1) % (len(directions))
	l.facingDirection = directions[newIndex]
}

func (l *lab) moveGuard() {
	movement := movements[slices.Index(directions, l.facingDirection)]
	newY := l.guardY + movement[0]
	newX := l.guardX + movement[1]
	if newY < 0 || newY >= len(l.grid) || newX < 0 || newX >= len(l.grid[newY]) {
		l.guardExit = true
		return
	}

	if l.grid[newY][newX] == "#" || l.grid[newY][newX] == "O" {
		if l.seenObstacles[newY][newX] > 2 {
			l.guardLoop = true
			return
		}

		if l.seenObstacles[newY] == nil {
			l.seenObstacles[newY] = map[int]int{}
		}
		l.seenObstacles[newY][newX]++

		l.turnRight()
	} else {
		l.guardY = newY
		l.guardX = newX
	}
}
