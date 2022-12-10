package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, _ := os.Open(inputFile)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	registerX := 1
	xPerCpuCycle := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		cpu_intruction := strings.Split(line, " ")

		switch cpu_intruction[0] {
		case "noop":
			xPerCpuCycle = append(xPerCpuCycle, registerX)
		case "addx":
			// it takes 2 cycless to finish addx instruction
			for i := 0; i < 2; i++ {
				xPerCpuCycle = append(xPerCpuCycle, registerX)
			}
			numberToAdd, _ := strconv.Atoi(cpu_intruction[1])
			registerX += numberToAdd
		}
	}

	// Part 1
	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	sumSignalStrength := 0
	for _, cycle := range interestingCycles {
		signalStrength := xPerCpuCycle[cycle-1] * cycle
		sumSignalStrength += signalStrength
	}
	fmt.Println("Sum of signal strength: ", sumSignalStrength)

	// Part 2
	lineEndings := []int{40, 80, 120, 160, 200, 240}
	// for every cycle draw a pixel. If sprite is visible in currently drawn pixel, draw "#" otherwise "."
	for i, spriteX := range xPerCpuCycle {

		pixelX := i % 40
		if Abs(spriteX-pixelX) > 1 {
			fmt.Print(".")
		} else {
			fmt.Print("#")
		}

		if i+1 == lineEndings[0] {
			fmt.Println()
			lineEndings = lineEndings[1:]
		}
	}

}
