package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func newRule(input string) (rule, error) {
	parts := strings.Split(input, "|")
	before, err := strconv.Atoi(parts[0])
	if err != nil {
		return rule{}, err
	}
	after, err := strconv.Atoi(parts[1])
	if err != nil {
		return rule{}, err
	}

	return rule{before: before, after: after}, nil
}

func main() {
	input, err := os.ReadFile("./cmd/day5a/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(string(input), "\n\n")
	orderingRules := strings.Split(sections[0], "\n")
	var rules []rule
	for _, ruleInput := range orderingRules {
		r, err := newRule(ruleInput)
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, r)
	}
	updatesInput := strings.Split(sections[1], "\n")
	var updates [][]int
	for _, updateInput := range updatesInput {
		var update []int
		for _, pageInput := range strings.Split(updateInput, ",") {
			pageNumber, err := strconv.Atoi(pageInput)
			if err != nil {
				log.Fatal(err)
			}
			update = append(update, pageNumber)
		}
		updates = append(updates, update)
	}

	var total int

	for _, update := range updates {
		isValid := true
		for _, r := range rules {
			beforeIndex := slices.Index(update, r.before)
			afterIndex := slices.Index(update, r.after)
			if beforeIndex == -1 || afterIndex == -1 {
				continue
			}

			if beforeIndex > afterIndex {
				isValid = false
				break
			}
		}
		if isValid {
			// todo find middle number
			middle := update[len(update)/2]
			total += middle
		}
	}
	fmt.Println(total)
}
