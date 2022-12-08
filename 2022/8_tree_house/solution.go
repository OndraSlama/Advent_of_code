package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func loadInputTo2DArray() [][]int {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	trees := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for _, char := range line {
			treeLength, _ := strconv.Atoi(string(char))
			row = append(row, treeLength)
		}

		trees = append(trees, row)
	}
	return trees
}

func maxInArray(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func reverseArray(arr []int) []int {
	reversed := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func getTreeVisibility(row, col int, trees [][]int) (bool, int) {
	// set tree to visible if any trees leading to an edge (on any side) are lower than it
	treeHeight := trees[row][col]

	treeRow := trees[row]
	var treeColumn []int
	for _, rowArray := range trees {
		treeColumn = append(treeColumn, rowArray[col])
	}
	directions := [][]int{reverseArray(treeRow[:col]), treeRow[col+1:], reverseArray(treeColumn[:row]), treeColumn[row+1:]}

	isVisible := false
	scenicScore := 1
	for _, direction := range directions {
		isVisibleInDirection, noVisibleTrees := getVisibilityInDirection(treeHeight, direction)
		if isVisibleInDirection {
			isVisible = true
		}
		scenicScore *= noVisibleTrees
	}
	return isVisible, scenicScore
}

func getVisibilityInDirection(treeHeight int, treesInDirection []int) (bool, int) {
	isVisible := true
	noVisibleTrees := 0
	for _, tree := range treesInDirection {
		noVisibleTrees++
		if treeHeight <= tree {
			isVisible = false
			break
		}
	}
	return isVisible, noVisibleTrees
}

func create2DArray(rows, cols int) [][]bool {
	arr := make([][]bool, rows)
	for i := range arr {
		arr[i] = make([]bool, cols)
	}
	return arr
}

func main() {
	// Read input and store to 2D array of trees
	trees := loadInputTo2DArray()
	fmt.Println(trees)

	visibleTrees := create2DArray(len(trees), len(trees[0]))
	noVisibleTrees := 0
	hihestScenicScore := 0
	for columnIndex, row := range trees {
		for rowIndex, _ := range row {
			isVisible, scenicScore := getTreeVisibility(columnIndex, rowIndex, trees)
			visibleTrees[columnIndex][rowIndex] = isVisible
			if isVisible {
				noVisibleTrees++
			}
			if scenicScore > hihestScenicScore {
				hihestScenicScore = scenicScore
			}
		}
	}
	fmt.Println("No visible trees:", noVisibleTrees)
	fmt.Println("Highest scenic score:", hihestScenicScore)

}
