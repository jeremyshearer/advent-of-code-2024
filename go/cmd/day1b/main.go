package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day1b/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var leftList []int
	rightMap := map[int]int{}
	for _, row := range strings.Split(string(input), "\n") {
		points := strings.Split(row, "   ")
		left, err := strconv.Atoi(points[0])
		if err != nil {
			log.Fatal(err)
		}
		leftList = append(leftList, left)
		right, err := strconv.Atoi(points[1])
		if err != nil {
			log.Fatal(err)
		}
		rightMap[right] += 1
	}

	var total int
	for _, val := range leftList {
		total += val * rightMap[val]
	}
	fmt.Println(total)
}
