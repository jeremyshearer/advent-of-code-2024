package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day2a/input.`txt")
	if err != nil {
		log.Fatal(err)
	}

	var safeLevels int
	for _, reportLine := range strings.Split(string(input), "\n") {
		report, err := NewReport(strings.Split(reportLine, " "))
		if err != nil {
			log.Fatal(err)
		}
		if report.IsSafe() {
			fmt.Printf("safe: %+v\n", report)
			safeLevels++
		} else {
			fmt.Printf("unsafe: %+v\n", report)
			if report.isSafeWithTolerance() {
				fmt.Printf("safe: %+v\n", report)
				safeLevels++
			} else {
				//fmt.Printf("unsafe: %+v\n", report)
			}
		}
	}
	fmt.Println(safeLevels)
}

func NewReport(input []string) (Report, error) {
	var levels []int
	for _, char := range input {
		level, err := strconv.Atoi(char)
		if err != nil {
			return Report{}, err
		}
		levels = append(levels, level)
	}
	return Report{levels}, nil
}

type Report struct {
	levels []int
}

func (r Report) InitialDirection() string {
	if r.levels[0] > r.levels[1] {
		return "down"
	} else if r.levels[0] < r.levels[1] {
		return "up"
	}

	return "level"
}

func (r Report) IsSafe() bool {
	for i, level := range r.levels {
		if i == len(r.levels)-1 {
			return true
		}
		nextLevel := r.levels[i+1]

		if nextLevel == level {
			return false
		}

		if nextLevel > level && r.InitialDirection() == "down" {
			return false
		}

		if nextLevel < level && r.InitialDirection() == "up" {
			return false
		}

		delta := math.Abs(float64(level - nextLevel))

		if delta > 3 {
			return false
		}
	}
	return true
}

func (r Report) isSafeWithTolerance() bool {
	for i := range len(r.levels) {
		levels := make([]int, 0)
		levels = append(levels, r.levels[:i]...)
		levels = append(levels, r.levels[i+1:]...)
		checkReport := Report{levels}
		fmt.Printf("%+v\n", checkReport)
		if checkReport.IsSafe() {
			return true
		}
	}
	return false
}
