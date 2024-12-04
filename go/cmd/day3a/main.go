package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section string

const (
	sectionInstruction  section = "instruction"
	sectionFirstNumber  section = "firstNumber"
	sectionDelimiter    section = "delimiter"
	sectionSecondNumber section = "secondNumber"
)

// mul(123,123)
func main() {
	input, err := os.ReadFile("./cmd/day3a/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	maxLen := len(input) - 1
	var total int
	for i, _ := range input {
		var buffer []byte
		maxIndex := min(i+4, maxLen)
		if string(input[i:maxIndex]) == "mul(" {
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
					fmt.Printf("%d*%d=%d\n", firstNum, secondNum, firstNum*secondNum)
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
