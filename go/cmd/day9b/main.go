package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./cmd/day9b/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var values []int
	var fileID int
	var blocks Blocks
	for i, val := range strings.Split(string(input), "") {
		isFreeSpace := i%2 == 1
		blockSize, _ := strconv.Atoi(val)
		values = append(values, blockSize)
		if isFreeSpace {
			blocks = append(blocks, &Block{
				Capacity: blockSize,
			})
		} else {
			blocks = append(blocks, &Block{
				Entries:  slices.Repeat([]int{fileID}, blockSize),
				Capacity: blockSize,
			})
			fileID++
		}
	}
	for blockIndex := len(blocks) - 1; blockIndex >= 0; blockIndex-- {
		b := blocks[blockIndex]
		if len(b.Entries) == 0 {
			continue
		}

		for freeIndex, fb := range blocks {
			if fb.Availability() < b.Capacity || freeIndex > blockIndex {
				continue
			}
			fb.Entries = append(fb.Entries, b.Entries...)
			b.Entries = []int{}
			break
		}
	}

	var i int
	var total int
	for _, b := range blocks {
		for _, f := range b.Entries {
			total += f * i
			i++
		}
		i += b.Availability()
	}
	fmt.Println(total)
}

type Block struct {
	Capacity int
	Entries  []int
}

type Blocks []*Block

func (b Blocks) Print() {
	for _, val := range b {
		for i := range val.Capacity {
			if i < len(val.Entries) {
				fmt.Print(val.Entries[i])
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}

func (b Blocks) LastNonEmptyBlock() (int, *Block) {
	i := len(b) - 1
	for i >= 0 {
		block := b[i]
		if len(block.Entries) != 0 {
			return i, block
		}
		i--
	}
	return 0, nil
}

func (b Block) Checksum(i int) int {
	if !b.Full() {
		return 0
	}
	var total int
	for _, v := range b.Entries {
		total += i * v
	}
	return total
}

func (b Block) Full() bool {
	return b.Availability() == 0
}

func (b Block) Availability() int {
	return b.Capacity - len(b.Entries)
}
