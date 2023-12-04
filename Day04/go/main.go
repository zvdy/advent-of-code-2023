package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// calculatePoints calculates the points for each card
func calculatePoints(cards []string) int {
	totalPoints := 0

	for _, card := range cards {
		// split the card into winning numbers and your numbers
		sides := strings.Split(card, "|")
		winningNumbers := strings.Fields(sides[0])
		yourNumbers := strings.Fields(sides[1])

		points := 0
		for _, yourNumber := range yourNumbers {
			for _, winningNumber := range winningNumbers {
				if yourNumber == winningNumber {
					// if a number matches, double the points (or set to 1 if this is the first match)
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break
				}
			}
		}
		totalPoints += points
	}

	return totalPoints
}

// part1 reads the cards from the file and calculates the total points
func part1() string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []string
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	result := calculatePoints(cards)
	return fmt.Sprintf("part one: %d", result)
}

// calculateMatches calculates the number of matches for each card
func calculateMatches(cards []string) []int {
	matches := make([]int, len(cards))

	for i, card := range cards {
		// split the card into winning numbers and your numbers
		sides := strings.Split(card, "|")
		winningNumbers := strings.Fields(sides[0])
		yourNumbers := strings.Fields(sides[1])

		for _, yourNumber := range yourNumbers {
			for _, winningNumber := range winningNumbers {
				if yourNumber == winningNumber {
					// if a number matches, increment the match count for this card
					matches[i]++
					break
				}
			}
		}
	}

	return matches
}

// part2 reads the cards from the file and calculates the total number of cards won
func part2() string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []string
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	matches := calculateMatches(cards)
	cardsWon := make([]int, len(cards))
	for i := range cardsWon {
		cardsWon[i] = 1
	}

	for i, match := range matches {
		// for each match, add the number of cards won to the next cards
		for j := 1; j <= match && i+j < len(cards); j++ {
			cardsWon[i+j] += cardsWon[i]
		}
	}

	total := 0
	for _, won := range cardsWon {
		total += won
	}

	return fmt.Sprintf("Part two: %d", total)
}

// main function prints the results of part 1 and part 2
func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
