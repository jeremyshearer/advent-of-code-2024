package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day8b/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	var antennaMap [][]string
	locationsByFrequency := map[string][]Point{}
	antinodes := map[int]map[int]bool{}
	for y, row := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		cells := strings.Split(row, "")
		antennaMap = append(antennaMap, cells)
		for x, cell := range cells {
			if cell == "." {
				continue
			}
			p := Point{X: x, Y: y}
			locationsByFrequency[cell] = append(locationsByFrequency[cell], p)
		}
	}
	maxY := len(antennaMap) - 1
	maxX := len(antennaMap[0]) - 1
	for _, locs := range locationsByFrequency {
		for _, locationA := range locs {
			for _, locationB := range locs {
				if locationA == locationB {
					continue
				}

				yDelta := locationA.Y - locationB.Y
				xDelta := locationA.X - locationB.X
				var upper Point
				var lower Point

				if yDelta < 0 {
					upper = locationA
					lower = locationB
				} else {
					upper = locationB
					lower = locationA
				}
				uap := Point{X: upper.X + xDelta, Y: upper.Y + yDelta}
				for uap.Y >= 0 && uap.Y <= maxY && uap.X >= 0 && uap.X <= maxX {
					if antinodes[uap.Y] == nil {
						antinodes[uap.Y] = map[int]bool{}
					}
					antinodes[uap.Y][uap.X] = true
					uap.X += xDelta
					uap.Y += yDelta
				}

				lap := Point{X: lower.X - xDelta, Y: lower.Y - yDelta}
				for lap.Y >= 0 && lap.Y <= maxY && lap.X >= 0 && lap.X <= maxX {
					if antinodes[lap.Y] == nil {
						antinodes[lap.Y] = map[int]bool{}
					}
					antinodes[lap.Y][lap.X] = true
					lap.X -= xDelta
					lap.Y -= yDelta
				}
			}
		}
	}

	var antinodeCount int

	for y, row := range antennaMap {
		for x, cell := range row {
			if cell == "." && antinodes[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
	fmt.Printf("%+v\n", antinodes)
	for _, row := range antinodes {
		for _, cell := range row {
			if cell {
				antinodeCount++
			}
		}
	}

	fmt.Printf("count: %d\n", antinodeCount)
}

type Point struct {
	X int
	Y int
}
