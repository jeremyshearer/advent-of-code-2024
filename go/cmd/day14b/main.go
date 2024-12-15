package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	h = 103
	w = 101
)

func main() {
	input, err := os.ReadFile("./cmd/day14a/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var robots []*Robot
	for _, line := range strings.Split(string(input), "\n") {
		fields := strings.Fields(line)
		posStr, _ := strings.CutPrefix(fields[0], "p=")
		xy := strings.Split(posStr, ",")
		px, _ := strconv.Atoi(xy[0])
		py, _ := strconv.Atoi(xy[1])
		vStr, _ := strings.CutPrefix(fields[1], "v=")
		vxy := strings.Split(vStr, ",")
		vx, _ := strconv.Atoi(vxy[0])
		vy, _ := strconv.Atoi(vxy[1])
		r := &Robot{P: image.Point{X: px, Y: py}, Velocity: image.Point{X: vx, Y: vy}}
		robots = append(robots, r)
	}
	for i := range 10000 {
		for _, r := range robots {
			r.Tick()
		}
		rm := NewRobotMap(robots)
		for _, v := range rm {
			if len(v) > 30 {
				rm.Print()
				fmt.Printf("============================%d================================\n", i)
			}
		}
	}
}

func NewRobotMap(robots []*Robot) RobotMap {
	robotPositions := map[int]map[int]int{}
	for _, r := range robots {
		if robotPositions[r.P.Y] == nil {
			robotPositions[r.P.Y] = map[int]int{}
		}
		robotPositions[r.P.Y][r.P.X]++
	}
	return robotPositions
}

type RobotMap map[int]map[int]int

func (rm RobotMap) Print() {
	for y := range h {
		for x := range w {
			//if y == h/2 || x == w/2 {
			//	fmt.Print(" ")
			//	continue
			//}
			robotCount := rm[y][x]
			if robotCount == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(robotCount)
			}
		}
		fmt.Println()
	}
}

type Robot struct {
	P        image.Point
	Velocity image.Point
}

func (r *Robot) Tick() {
	newPos := r.P.Add(r.Velocity)
	if newPos.Y < 0 {
		newPos.Y = h + newPos.Y
	} else {
		newPos.Y = newPos.Y % h
	}
	if newPos.X < 0 {
		newPos.X = w + newPos.X
	} else {
		newPos.X = newPos.X % w
	}
	r.P = newPos
}
