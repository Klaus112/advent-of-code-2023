package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "files/advent_of_code_input_1.txt"
	testFile = "files/advent_of_code_check_file_1.txt"
)

func main() {
	// arg this does not work.
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}
	defer f.Close()

	sum, err := sumFromIoReader(f)
	if err != nil {
		fmt.Println("sum failed with err", err)

		return
	}

	fmt.Println("sum of file is: ", sum)
}

func sumFromIoReader(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Print(line)
		onlyNumbers := replace(line)
		if len(onlyNumbers) > 0 {
			num, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(onlyNumbers[0]), string(onlyNumbers[len(onlyNumbers)-1])))
			// fmt.Println(":", num)
			// fmt.Println(num)
			sum = sum + num
		} else {
			fmt.Println("no numbers: ", line)
		}
	}

	return sum, nil
}

func replace(line string) string {
	line = replaceWrittenNumbersWithLetters(line)
	onlyNumbers := strings.Map(mapOnlyNumbers, line)

	return onlyNumbers
}

func replaceWrittenNumbersWithLetters(line string) string {
	valMaps := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	lowestIndex := math.MaxInt64
	highestIndex := -1
	var lowestFoundSearchTerm, highestFoundSearchTerm string

	for searchTerm := range valMaps {
		i1 := strings.Index(line, searchTerm)
		i2 := strings.LastIndex(line, searchTerm)

		// lowest
		if i1 < lowestIndex && i1 != -1 {
			lowestIndex = i1
			lowestFoundSearchTerm = searchTerm
		}

		// highest
		if i2 > highestIndex {
			highestIndex = i2
			highestFoundSearchTerm = searchTerm
		}
	}

	if highestIndex == -1 {
		// no replace
		return line
	}

	line = line[0:highestIndex] + valMaps[highestFoundSearchTerm] + line[highestIndex+len(highestFoundSearchTerm):]
	if highestIndex == lowestIndex {
		return line
	}

	// replace first
	line = line[0:lowestIndex] + valMaps[lowestFoundSearchTerm] + line[lowestIndex+len(lowestFoundSearchTerm)-1:] // i forgot the -1 here to so it stripped the 2 in the following string: 88eigh2ffg

	return line
}

func mapOnlyNumbers(r rune) rune {
	if strings.IndexRune("0123456789", r) >= 0 {
		return r
	}

	return -1
}
