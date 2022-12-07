package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var sizeOfFolders = map[string]int{}

func addFileSizeToAllFolders(folderStack *[]string, size int) {
	folderPath := ""
	for _, folder := range *folderStack {
		folderPath += folder + "/"
		sizeOfFolders[folderPath] += size
	}
}

func getSumOfDirectories(maxSize int) int {
	sum := 0
	for folder, size := range sizeOfFolders {
		fmt.Printf("%s: %d\n", folder, size)
		if size <= maxSize {
			sum += size
		}
	}
	return sum
}

func parseLine(line string, folderStack *[]string, scanner *bufio.Scanner) {
	switch {
	case strings.HasPrefix(line, "$ cd .."):
		*folderStack = (*folderStack)[:len(*folderStack)-1]
	case strings.HasPrefix(line, "$ cd "):
		folderName := strings.Split(line, " ")[2]
		*folderStack = append(*folderStack, folderName)
	default:
		sizeString := strings.Split(line, " ")[0]
		size, err := strconv.Atoi(sizeString)
		if err == nil {
			addFileSizeToAllFolders(folderStack, size)
		}
	}
}

func main() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	folderStack := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line, &folderStack, scanner)
	}

	// Sum all folders with size <= 100 000
	sum := getSumOfDirectories(100000)
	fmt.Printf("Sum of folders with size <= 100 000: %d\n", sum)

	// Find smallest folder with size bigger then sizeToFree
	sizeToFree := 30000000 - (70000000 - sizeOfFolders["//"])
	smallestFolder := ""
	for folder, size := range sizeOfFolders {
		if size > sizeToFree && (smallestFolder == "" || size < sizeOfFolders[smallestFolder]) {
			smallestFolder = folder
		}
	}
	fmt.Printf("Smallest folder with size bigger then %d: %s (%d)\n", sizeToFree, smallestFolder, sizeOfFolders[smallestFolder])

}
