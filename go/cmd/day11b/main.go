package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day11b/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	hallway := map[int]int{}
	for _, v := range strings.Fields(string(input)) {
		stoneValue, _ := strconv.Atoi(v)
		hallway[stoneValue] = 1
	}
	for range 75 {
		newHallway := map[int]int{}
		for s, c := range hallway {
			newStones := blink(s)
			for _, n := range newStones {
				newHallway[n] += c
			}
		}
		hallway = newHallway
	}
	var total int
	for _, count := range hallway {
		total += count
	}
	fmt.Printf("%d\n", total)
}

func blink(val int) []int {
	if val == 0 {
		return []int{1}
	} else if evenDigitCount(val) {
		return split(val)
	} else {
		return multiply(val)
	}
}

func split(val int) []int {
	digitStr := strconv.Itoa(val)
	digits := []string{digitStr[:len(digitStr)/2], digitStr[len(digitStr)/2:]}
	firstStoneVal, _ := strconv.Atoi(digits[0])
	secondStoneVal, _ := strconv.Atoi(digits[1])
	return []int{firstStoneVal, secondStoneVal}
}

func multiply(val int) []int {
	return []int{val * 2024}
}

func evenDigitCount(val int) bool {
	valString := strconv.Itoa(val)

	return len(valString)%2 == 0
}
