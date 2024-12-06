package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day6a/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	var grid [][]string
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
	l = &lab{
		guardX:          guardX,
		guardY:          guardY,
		facingDirection: "N",
		grid:            grid,
		visited: map[int]map[int]bool{
			guardY: {
				guardX: true,
			},
		},
	}
	for {
		l.moveGuard()
		if l.guardExit {
			break
		}
	}
	fmt.Printf("%d\n%+v\n", l.visitedCount(), l.visited)
}

var directions = []string{"N", "E", "S", "W"}
var movements = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type lab struct {
	grid            [][]string
	visited         map[int]map[int]bool
	guardX          int
	guardY          int
	facingDirection string
	guardExit       bool
}

func (l *lab) turnRight() {
	curr := slices.Index(directions, l.facingDirection)
	newIndex := (curr + 1) % (len(directions))
	fmt.Println(newIndex)
	l.facingDirection = directions[newIndex]
}

func (l *lab) moveGuard() {
	movement := movements[slices.Index(directions, l.facingDirection)]
	newY := l.guardY + movement[0]
	newX := l.guardX + movement[1]
	if newY < 0 || newY >= len(l.grid) || newX < 0 || newY > len(l.grid) {
		l.guardExit = true
		return
	}
	if l.grid[newY][newX] == "#" {
		l.turnRight()
	} else {
		l.guardY = newY
		l.guardX = newX
		if l.visited[l.guardY] == nil {
			l.visited[l.guardY] = map[int]bool{}
		}
		l.visited[l.guardY][l.guardX] = true
	}
}

func (l *lab) visitedCount() int {
	var result int
	for _, v := range l.visited {
		for _, visited := range v {
			if visited {
				result++
			}
		}
	}
	return result
}
