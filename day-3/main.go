package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "files/input.txt"
	testFile  = "files/test.txt"
)

const (
	dotVal     = -10
	specialVal = -1
)

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}
	defer f.Close()

	inputArray := inputTo2DArray(f)

	sum := getSumOfValidNumers(inputArray)

	fmt.Println(sum)
}

func getSumOfValidNumers(inputArray [][]int32) int {
	validNumbers := getValidNumbers(inputArray)
	sum := 0
	for _, val := range validNumbers {
		sum += val
	}

	return sum
}

func inputTo2DArray(input io.Reader) [][]int32 {
	scanner := bufio.NewScanner(input)

	var res [][]int32

	for scanner.Scan() {
		line := scanner.Text()
		convertedLine := make([]int32, len(line))
		for i, val := range line {
			if strings.IndexRune("0123456789", val) >= 0 {
				convertedLine[i] = int32(val - '0') // subtract Zero rune from integer runes will lead to the int representation
			} else if strings.IndexRune(".", val) >= 0 {
				convertedLine[i] = dotVal
			} else {
				convertedLine[i] = specialVal
			}
		}
		res = append(res, convertedLine)
	}

	return res
}

// getValidNumbers which are next to a 'specialCharacter' (=everything except numbers and dots).
func getValidNumbers(in [][]int32) []int {
	var res []int
	for rowIndex := 0; rowIndex < len(in); rowIndex++ {
		var minInt, maxInt int
		wasOnNum := false
		for columnIndex := 0; columnIndex < len(in[rowIndex]); columnIndex++ {
			val := in[rowIndex][columnIndex]
			if val >= 0 {
				if !wasOnNum {
					minInt = columnIndex
				}
				wasOnNum = true
			} else if wasOnNum {
				maxInt = columnIndex - 1
				wasOnNum = false
				if shouldNumberBeCounted(in, rowIndex, minInt-1, maxInt+1) {
					val := combineNumbersIntoSingle(in[rowIndex][minInt : maxInt+1])
					res = append(res, val)
				}
			}
		}

		// don't forget  the last one if wasOnNum is still true
		if wasOnNum && shouldNumberBeCounted(in, rowIndex, minInt-1, len(in[rowIndex])) {
			val := combineNumbersIntoSingle(in[rowIndex][minInt:len(in[rowIndex])])
			res = append(res, val)
		}
	}

	return res
}

// combineNumbersIntoSingle combines array numbers next to each other into one.
//
//	e.g.: []int{1,5,7} -> 157
func combineNumbersIntoSingle(in []int32) int {
	combined := ""
	for _, val := range in {
		combined += fmt.Sprint(val)
	}

	res, err := strconv.Atoi(combined)
	if err != nil {
		panic(fmt.Errorf("Invalid number %s: %w", combined, err))
	}

	return res
}

// shouldNumberBeCounted checks if the number with the boundaries matrix[row][minColumn+1] : matrix[row][maxColumn-1] should be counted.
func shouldNumberBeCounted(matrix [][]int32, row, minColumn, maxColumn int) bool {
	// check row above
	if rowVal := row - 1; rowVal >= 0 {
		if rowHasSpecialChar(matrix[rowVal], minColumn, maxColumn) {
			return true
		}
	}

	// check same row
	if rowHasSpecialChar(matrix[row], minColumn, maxColumn) {
		return true
	}

	// check row below
	if rowVal := row + 1; rowVal < len(matrix) {
		if rowHasSpecialChar(matrix[rowVal], minColumn, maxColumn) {
			return true
		}
	}

	return false
}

// rowHasSpecialChar checks if the row between min and max val hasa special character.
func rowHasSpecialChar(row []int32, minVal, maxVal int) bool {
	minVal = clamp(minVal, 0, len(row)-1)
	maxVal = clamp(maxVal, 0, len(row)-1)

	for i := minVal; i <= maxVal; i++ {
		if row[i] == specialVal {
			return true
		}
	}

	return false
}

// clamp val between min and max (inclusive).
func clamp(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}

	return val
}
