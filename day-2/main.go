package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/klaus112/advent-of-code-2023/common"
)

// game contains information about a single game recording it's highest values
type game struct {
	id               int
	red, blue, green int
}

func main() {
	f, err := os.Open(common.InputFile)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}
	defer f.Close()

	sum, powerOfGames, err := getSumFromIOReader(f)
	if err != nil {
		fmt.Println("calculating sum failed", err)

		return
	}

	fmt.Printf("Sum of valid games is: %d\n", sum)
	fmt.Printf("Power of all games is: %d\n", powerOfGames)
}

func getSumFromIOReader(reader io.Reader) (sum int, powerOfGames int, err error) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		game, err := lineToGame(line)
		if err != nil {
			return 0, 0, err
		}

		powerOfGames += game.blue * game.red * game.green

		if game.isValid() {
			sum += game.id
		}
	}

	return sum, powerOfGames, nil
}

func lineToGame(line string) (game, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return game{}, fmt.Errorf("no game in row: %s", line)
	}

	gameID, err := getGameID(parts)
	if err != nil {
		return game{}, err
	}

	red, green, blue := getHighestColorValuesForGame(parts[1])

	return game{
		id:    gameID,
		red:   red,
		blue:  blue,
		green: green,
	}, nil
}

func getGameID(parts []string) (int, error) {
	if gameIDParts := strings.Split(parts[0], " "); len(gameIDParts) == 2 {
		gameID, err := strconv.Atoi(gameIDParts[1])
		if err != nil {
			return 0, fmt.Errorf("could not convert gameID to integer for line %s: %w", parts[0], err)
		}

		return gameID, nil
	}

	return 0, fmt.Errorf("gameID part can not be split, %s", parts[0])
}

// getHighestColorValuesForGame gameOutput should only be the part behind the colon (:).
//
//	eg.: " 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue"
func getHighestColorValuesForGame(revealedValues string) (red int, green int, blue int) {
	draws := strings.Split(revealedValues, ";")

	for _, singleDraw := range draws {
		colors := strings.Split(singleDraw, ",")
		colors = removeEmpty(colors)
		for _, colorStr := range colors {
			s := strings.Split(colorStr, " ")
			val, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}
			color := s[1]
			switch color {
			case "blue":
				if val > blue {
					blue = val
				}
			case "red":
				if val > red {
					red = val
				}
			case "green":
				if val > green {
					green = val
				}
			}
		}
	}

	return
}

func removeEmpty(parts []string) []string {
	res := make([]string, 0, len(parts))
	for i := range parts {
		parts[i], _ = strings.CutPrefix(parts[i], " ")
		parts[i], _ = strings.CutSuffix(parts[i], " ")
		res = append(res, parts[i])
	}

	return res
}

func (g game) isValid() bool {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	if g.red <= maxRed && g.blue <= maxBlue && g.green <= maxGreen {
		return true
	}

	return false
}
