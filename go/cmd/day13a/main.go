package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

func main() {
	//input, err := os.ReadFile("./cmd/day13a/sample.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, machineInput := range strings.Split(string(input), "\n\n") {
	//	m := NewMachine(machineInput)
	//	fmt.Printf("%+v\n", *m)
	//	for {
	//		if m.ButtonsPressed >= 100 {
	//			break
	//		}
	//
	//		if m.Won() {
	//			break
	//		}
	//	}
	//}
	m := &Machine{
		Claw:    image.Point{},
		ButtonA: image.Point{X: 94, Y: 34},
		ButtonB: image.Point{X: 22, Y: 67},
		Prize:   image.Point{X: 8400, Y: 5400},
	}

	aCount := m.Prize.X / m.ButtonA.X
	bCount := 0
	for aCount > 0 && bCount < 100 && !m.IsWinningCombo(aCount, bCount) {
		fmt.Printf("a: %d, b: %d\n", aCount, bCount)
		aCount--
		bCount++
	}
	fmt.Println(m.IsWinningCombo(80, 40))
}

func newPoint(input string) image.Point {
	coords := strings.Split(input, ",")
	x, _ := strconv.Atoi(coords[0][3:])
	y, _ := strconv.Atoi(coords[1][3:])
	return image.Point{X: x, Y: y}
}

func NewMachine(input string) *Machine {
	lines := strings.Split(input, "\n")
	a := newPoint(strings.Split(lines[0], ":")[1])
	b := newPoint(strings.Split(lines[1], ":")[1])
	p := newPoint(strings.Split(lines[2], ":")[1])

	return &Machine{ButtonA: a, ButtonB: b, Prize: p}
}

type Machine struct {
	Claw       image.Point
	ButtonA    image.Point
	ButtonB    image.Point
	Prize      image.Point
	TokensUsed int
}

func (m *Machine) IsWinningCombo(aCount, bCount int) bool {
	newClaw := m.Claw.Add(image.Point{X: (m.ButtonA.X * aCount) + (m.ButtonB.X * bCount), Y: (m.ButtonA.Y * aCount) + (m.ButtonB.Y * bCount)})
	return newClaw == m.Prize
}

func (m *Machine) PressA() {
	m.Claw = m.Claw.Add(m.ButtonA)
	m.TokensUsed += 3
}

func (m *Machine) PressB() {
	m.Claw = m.Claw.Add(m.ButtonB)
	m.TokensUsed++
}

func (m *Machine) Won() bool {
	return m.Claw == m.Prize
}
