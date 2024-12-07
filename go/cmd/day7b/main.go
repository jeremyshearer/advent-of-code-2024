package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day7b/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var equations []equation
	var total int
	for _, row := range strings.Split(string(input), "\n") {
		parts := strings.Split(row, ":")
		answer, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		var numbers []int
		for _, rawNum := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			num, err := strconv.Atoi(rawNum)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, num)
		}

		e := equation{answer: answer, input: numbers}
		if e.Possible() {
			total += e.answer
			equations = append(equations, e)
		}
	}
	fmt.Printf("%d\n", total)
}

type equation struct {
	answer int
	input  []int
}

func (e equation) Possible() bool {
	permutations := generatePermutations(len(e.input) - 1)
	for _, perm := range permutations {
		left := e.input[0]
		for idx, op := range perm {
			right := e.input[idx+1]
			if op == "*" {
				left = left * right
			} else if op == "+" {
				left = left + right
			} else if op == "||" {
				left, _ = strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(right))
			}
		}
		if left == e.answer {
			return true
		}
	}

	return false
}

func generatePermutations(size int) [][]string {
	permutations := [][]string{}
	chars := []string{"*", "+", "||"}

	var backtrack func(int, []string)
	backtrack = func(idx int, current []string) {
		if idx == size {
			permutations = append(permutations, append([]string{}, current...))
			return
		}

		for _, char := range chars {
			current = append(current, char)
			backtrack(idx+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, []string{})
	return permutations
}
