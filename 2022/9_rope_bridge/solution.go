package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func Sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Motion struct {
	direction string
	steps     int
}

type Position struct {
	x int
	y int
}

func (p *Position) move(direction string) {
	switch direction {
	case "U":
		p.y++
	case "D":
		p.y--
	case "L":
		p.x--
	case "R":
		p.x++
	}
}

func (p *Position) follow(position Position) {
	// Calculate the difference between the two positions
	x_diff := position.x - p.x
	y_diff := position.y - p.y

	// Move the tail in the direction of the head in case its 2 steps away
	if Abs(x_diff)+Abs(y_diff) > 2 {
		p.x += Sign(x_diff)
		p.y += Sign(y_diff)
		return
	}
	if Abs(x_diff) > 1 {
		p.x += Sign(x_diff)
		return
	}
	if Abs(y_diff) > 1 {
		p.y += Sign(y_diff)
		return
	}

}

func loadInput() []Motion {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	motions := []Motion{}
	for scanner.Scan() {
		line := scanner.Text()
		motions_intructions := strings.Split(line, " ")
		steps, _ := strconv.Atoi(motions_intructions[1])
		motion := Motion{motions_intructions[0], steps}
		motions = append(motions, motion)
	}
	return motions
}

func ropeMovement(motions []Motion, numberOfKnots int) map[Position]bool {
	knots := []Position{}
	for i := 0; i < numberOfKnots; i++ {
		knots = append(knots, Position{0, 0})
	}

	tailPositionHistory := map[Position]bool{}

	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			knots[0].move(motion.direction)
			for j := 1; j < len(knots); j++ {
				knots[j].follow(knots[j-1])
			}
			tailPositionHistory[knots[len(knots)-1]] = true
		}
	}
	return tailPositionHistory
}

func main() {
	// Read input and store to 2D array of trees
	motions := loadInput()
	fmt.Println(motions)

	// First part
	tail := ropeMovement(motions, 2)
	fmt.Println("Number of unique positions in first part:", len(tail))

	// Second part
	tail = ropeMovement(motions, 10)
	fmt.Println("Number of unique positions in second part:", len(tail))

}
