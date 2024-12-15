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
	input, err := os.ReadFile("./cmd/day12b/sample.txt")
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
		fmt.Printf("%s\n", garden[p[0].Y][p[0].X])
		total += p.sideCount(garden) * len(p)
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

func (p plot) sideCount(garden [][]string) int {
	var edges []image.Point
	for _, pos := range p {
		for _, dir := range directions {
			chkPos := pos.Add(dir)

			if chkPos.Y < 0 || chkPos.Y > len(garden)-1 || chkPos.X < 0 || chkPos.X > len(garden[chkPos.Y])-1 || garden[pos.Y][pos.X] != garden[chkPos.Y][chkPos.X] {
				edges = append(edges, pos)
				break
			}
		}
	}
	if len(edges) == 1 {
		return 4
	}
	horizontal := map[int][]int{}
	vertical := map[int][]int{}
	fmt.Println(edges)
	for _, e := range edges {
		r := e.Add(image.Point{X: 1})
		l := e.Add(image.Point{X: -1})
		if r.X > len(garden[e.Y])-1 || garden[r.Y][r.X] == garden[e.Y][e.X] || l.X < 0 || garden[l.Y][l.X] == garden[e.Y][e.X] {
			if horizontal[e.Y] == nil {
				horizontal[e.Y] = []int{e.X}
			} else {
				horizontal[e.Y] = append(horizontal[e.Y], e.X)
			}
		}
		u := e.Add(image.Point{Y: -1})
		d := e.Add(image.Point{Y: 1})
		if d.Y > len(garden)-1 || garden[d.Y][d.X] == garden[e.Y][e.X] || u.Y < 0 || garden[e.Y][e.X] == garden[u.Y][u.X] {
			if vertical[e.X] == nil {
				vertical[e.X] = []int{e.Y}
			} else {
				vertical[e.X] = append(vertical[e.X], e.Y)
			}
		}
	}
	fmt.Printf("\nv: %+v\n h: %+v\n", vertical, horizontal)

	return len(edges)
}
