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

type Scratchcard struct {
	cardID         int
	winningNumbers []int
	myNumbers      []int
}

type computedScratchcard struct {
	cardID              int
	matchingNumberCount int
}

func main() {
	f, err := os.Open(common.InputFile)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}
	defer f.Close()

	cards, err := getCards(f)
	if err != nil {
		fmt.Println("could not get cards from file", err)

		return
	}

	// day 4 part 1
	fmt.Printf("Points for all scratchcards: %d\n", getPointsForAllGames(cards))

	// day 4 part 2
	idToScratchCards := make(map[int][]computedScratchcard, len(cards))
	for i := range cards {
		idToScratchCards[cards[i].cardID] = []computedScratchcard{getComputedScratchCard(cards[i])}
	}

	scratchcardSum := getScratchCardCount(idToScratchCards)
	fmt.Printf("Sum of scratchcards: %d\n", scratchcardSum)
}

// getComputedScratchCard computes the winningNumber count for each scratchcard.
func getComputedScratchCard(card Scratchcard) computedScratchcard {
	return computedScratchcard{
		cardID:              card.cardID,
		matchingNumberCount: len(common.HashIntersect[int](card.winningNumbers, card.myNumbers)),
	}
}

// getScratchCardCount returns the count of scratchcards after the copying of winning numbers.
//
// Precomputing the scratchCard matchingNumbers already improved the function performance from 9.7 seconds to ~500ms.
func getScratchCardCount(idToScratchCards map[int][]computedScratchcard) int {
	// implements copying logik for winning numbers.
	// cannot loop directly over the map because the execution order is random
	for cardID := 1; cardID <= len(idToScratchCards); cardID++ {
		cards := idToScratchCards[cardID]
		for _, card := range cards {
			for i := 1; i <= card.matchingNumberCount; i++ {
				idToScratchCards[cardID+i] = append(idToScratchCards[cardID+i], idToScratchCards[cardID+i][0]) // add a single instance for every matching number of the appropriate card.
			}
		}
	}

	sum := 0
	for _, cards := range idToScratchCards {
		// fmt.Printf("CardID %d contains %d instances of the card\n", cardID, len(cards))
		sum += len(cards)
	}

	return sum
}

// getPointsForAllGames sums up the scratchcards winning numbers for part1.
func getPointsForAllGames(cards []Scratchcard) int {
	res := 0
	for _, card := range cards {
		matchingNumbers := common.HashIntersect[int](card.winningNumbers, card.myNumbers)
		singleGamePoints := 0
		for i := 0; i < len(matchingNumbers); i++ {
			if singleGamePoints == 0 {
				singleGamePoints = 1
			} else {
				singleGamePoints *= 2 // double the value
			}
		}
		res += singleGamePoints
	}

	return res
}

func getCards(in io.Reader) ([]Scratchcard, error) {
	scanner := bufio.NewScanner(in)

	cards := make([]Scratchcard, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line found: %s", line)
		}
		gameID, err := getGameID(parts[0])
		if err != nil {
			return nil, err
		}
		winning, matching, err := getWinningAndMatchingNumbers(parts[1])
		if err != nil {
			return nil, err
		}

		cards = append(cards, Scratchcard{
			cardID:         gameID,
			winningNumbers: winning,
			myNumbers:      matching,
		})
	}

	return cards, nil
}

func getGameID(gameID string) (int, error) {
	_, numStr, found := strings.Cut(gameID, " ")
	if found {
		return strconv.Atoi(strings.TrimSpace(numStr))
	}

	return 0, fmt.Errorf("invalid gameID string: %s", gameID)
}

// getWinningAndMatchingNumbers from a string like this '41 48 83 86 17 | 83 86  6 31 17  9 48 53'.
func getWinningAndMatchingNumbers(numLine string) (winning []int, mine []int, err error) {
	parts := strings.Split(numLine, "|")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("getWinningAndMatchingNumbers invalid input: %s", numLine)
	}
	winning, err = getNumbers(parts[0])
	if err != nil {
		return
	}

	mine, err = getNumbers(parts[1])

	return
}

// getNumbers from a string like this: '41 48 83 86 17 '.
func getNumbers(numbers string) ([]int, error) {
	numbers = strings.TrimSpace(numbers)

	parts := strings.Split(numbers, " ")
	res := make([]int, 0, len(parts))
	for _, singleNum := range parts {
		singleNum = strings.TrimSpace(singleNum)
		if len(singleNum) > 0 {
			val, err := strconv.Atoi(singleNum)
			if err != nil {
				return nil, fmt.Errorf("getNumbers: %w", err)
			}
			res = append(res, val)
		}
	}

	return res, nil
}
