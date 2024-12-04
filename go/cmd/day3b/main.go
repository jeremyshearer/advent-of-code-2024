package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section string

// mul(123,123)
func main() {
	input, err := os.ReadFile("./cmd/day3b/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var total int

	maxLen := len(input) - 1
	enabled := true

	for i, _ := range input {
		var buffer []byte
		maxIndex := min(i+4, maxLen)
		maxDont := min(i+7, maxLen)
		maxDo := min(i+4, maxLen)

		if string(input[i:maxDont]) == "don't()" {
			enabled = false
		}

		if string(input[i:maxDo]) == "do()" {
			enabled = true
		}

		if enabled && string(input[i:maxIndex]) == "mul(" {
			for innerIndex, innerC := range input[maxIndex:] {
				if string(innerC) == ")" {
					parts := strings.Split(string(buffer), ",")
					firstNum, err := strconv.Atoi(parts[0])
					if err != nil {
						log.Printf("%+v", err)
						break
					}
					secondNum, err := strconv.Atoi(parts[1])
					if err != nil {
						log.Printf("%+v", err)
						break
					}
					total += firstNum * secondNum
					break
				}
				if innerIndex > 11 {
					break
				}
				buffer = append(buffer, innerC)
			}
		}
	}

	fmt.Println(total)
}
