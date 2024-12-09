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
	input, err := os.ReadFile("./cmd/day9a/input.txt")
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
	fmt.Println()
	for emptyBlockIndex, b := range blocks {
		if b.Full() {
			continue
		}
		for !b.Full() {
			lastBlockIndex, lastBlock := blocks.LastNonEmptyBlock()
			if lastBlockIndex <= emptyBlockIndex {
				break
			}
			lastValue := lastBlock.Entries[len(lastBlock.Entries)-1]
			b.Entries = append(b.Entries, lastValue)
			lastBlock.Entries = lastBlock.Entries[:len(lastBlock.Entries)-1]
		}
	}

	var i int
	var total int
	for _, b := range blocks {
		for _, f := range b.Entries {
			total += f * i
			i++
		}
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
