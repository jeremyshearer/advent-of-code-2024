package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day1a/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var listA []int
	var listB []int
	for _, row := range strings.Split(string(input), "\n") {
		points := strings.Split(row, "   ")
		pointA, err := strconv.Atoi(points[0])
		if err != nil {
			log.Fatal(err)
		}
		listA = append(listA, pointA)
		pointB, err := strconv.Atoi(points[1])
		if err != nil {
			log.Fatal(err)
		}
		listB = append(listB, pointB)
	}

	sort.Ints(listA)
	sort.Ints(listB)
	var total int
	for i := range len(listA) {
		delta := math.Abs(float64(listA[i] - listB[i]))
		total += int(delta)
	}
	fmt.Println(total)
}
