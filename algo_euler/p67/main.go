package main

// https://projecteuler.net/problem=67

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Read a file line by line and returns an array of lines in string format
 * https://stackoverflow.com/a/16615559/4906586
 * http://networkbit.ch/read-text-file-in-go/
 */
func readFile(path string) ([]string, error) {
	// absPath, _ := filepath.Abs(path)
	// file, err := os.Open(absPath)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// readLines splits each line by number and parse it into int64
func readLines(path string) ([][]int64, error) {
	lines, err := readFile(path)
	if err != nil {
		return nil, err
	}

	var entriesMatrix [][]int64

	for _, line := range lines {
		split := strings.Fields(line)

		var entries []int64

		for _, nbInTxt := range split {
			parsedInt, _ := strconv.ParseInt(nbInTxt, 10, 64)
			entries = append(entries, parsedInt)
		}

		entriesMatrix = append(entriesMatrix, entries)
	}

	return entriesMatrix, nil
}

// Move to next row
func buildMaxSum(path string) ([][]int64, error) {
	entries, err := readLines(path)
	if err != nil {
		return nil, err
	}

	// initialise with first value
	var bestSums [][]int64
	bestSums = append(bestSums, entries[0])
	// fmt.Println("Init BestSum", bestSums)

	for i := 1; i < len(entries); i++ {
		// get entries for a given line
		entry := entries[i]
		prevSum := bestSums[i-1]
		var bestSum []int64

		// Handle top left value
		bestSum = append(bestSum, prevSum[0]+entry[0])

		for j := 1; j < len(prevSum); j++ {
			leftParent := entry[j] + prevSum[j-1]
			rightParent := entry[j] + prevSum[j]

			if leftParent > rightParent {
				bestSum = append(bestSum, leftParent)
			} else {
				bestSum = append(bestSum, rightParent)
			}
		}

		// top right entry
		bestSum = append(bestSum, prevSum[len(prevSum)-1]+entry[len(entry)-1])

		bestSums = append(bestSums, bestSum)
		// fmt.Println("Got Entries", entry)
		// fmt.Println("New BestSum", bestSums)
	}

	return bestSums, nil
}

func main() {
	bestSum, _ := buildMaxSum("algo_euler/p67/p67_triangle.txt")
	// fmt.Println(bestSum)

	max := int64(0)
	for _, val := range bestSum[len(bestSum)-1] {
		if val > max {
			max = val
		}
	}

	fmt.Println("The Max: ", max)
}
