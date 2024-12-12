package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day11a/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var stones []Stone
	for _, v := range strings.Fields(string(input)) {
		stoneValue, _ := strconv.Atoi(v)
		stones = append(stones, Stone{Value: stoneValue})
	}
	for range 25 {
		var tmp []Stone
		for _, s := range stones {
			tmp = append(tmp, s.Blink()...)
		}
		stones = tmp
	}
	fmt.Printf("%d\n", len(stones))
}

type Stone struct {
	Value int
}

func (s *Stone) Blink() []Stone {
	if s.Value == 0 {
		return []Stone{{Value: 1}}
	} else if s.EvenDigitCount() {
		return s.Split()
	} else {
		return s.Multiply()
	}
}

func (s *Stone) Split() []Stone {
	digitStr := strconv.Itoa(s.Value)
	digits := []string{digitStr[:len(digitStr)/2], digitStr[len(digitStr)/2:]}
	firstStoneVal, _ := strconv.Atoi(digits[0])
	secondStoneVal, _ := strconv.Atoi(digits[1])
	return []Stone{{Value: firstStoneVal}, {Value: secondStoneVal}}
}

func (s *Stone) Multiply() []Stone {
	return []Stone{{Value: s.Value * 2024}}
}

func (s *Stone) EvenDigitCount() bool {
	valString := strconv.Itoa(s.Value)

	return len(valString)%2 == 0
}
