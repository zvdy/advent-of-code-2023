package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Game struct holds the game ID and the cube data.
type Game struct {
	ID    int
	Cubes []map[string]int
}

// parseGames function parses the input and returns a slice of Game.
func parseGames(input string) []Game {
	var games []Game
	gameRegex := regexp.MustCompile(`Game (\d+): (.*)`)
	cubeRegex := regexp.MustCompile(`(\d+) (red|green|blue)`)

	// Iterate over each line in the input
	for _, line := range strings.Split(input, "\n") {
		matches := gameRegex.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		// Parse the game ID
		id, _ := strconv.Atoi(matches[1])
		game := Game{ID: id}

		// Parse the cube data for each game
		for _, cubeStr := range strings.Split(matches[2], ";") {
			cubes := make(map[string]int)
			for _, cubeMatch := range cubeRegex.FindAllStringSubmatch(cubeStr, -1) {
				count, _ := strconv.Atoi(cubeMatch[1])
				cubes[cubeMatch[2]] = count
			}
			game.Cubes = append(game.Cubes, cubes)
		}

		games = append(games, game)
	}

	return games
}

// isGamePossible function checks if a game is possible with a given number of red, green, and blue cubes.
func isGamePossible(game Game, red, green, blue int) bool {
	// Iterate over each set of cubes in the game
	for _, cubes := range game.Cubes {
		// If any color has more cubes than the given number, the game is not possible
		if cubes["red"] > red || cubes["green"] > green || cubes["blue"] > blue {
			return false
		}
	}
	return true
}

// minimumCubesNeeded function calculates the minimum number of red, green, and blue cubes needed for a game.
func minimumCubesNeeded(game Game) (int, int, int) {
	red, green, blue := 0, 0, 0
	// Iterate over each set of cubes in the game
	for _, cubes := range game.Cubes {
		// Keep track of the maximum number of cubes of each color needed at any point
		if cubes["red"] > red {
			red = cubes["red"]
		}
		if cubes["green"] > green {
			green = cubes["green"]
		}
		if cubes["blue"] > blue {
			blue = cubes["blue"]
		}
	}
	return red, green, blue
}

// main function reads the input from a file, parses the games, and calculates the sum of the game IDs and the sum of the powers of the minimum sets of cubes.
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	games := parseGames(input)

	totalIDs := 0
	totalPower := 0
	for _, game := range games {
		// Check if the game is possible with 12 red, 13 green, and 14 blue cubes
		if isGamePossible(game, 12, 13, 14) {
			// If so, add the game ID to a running total
			totalIDs += game.ID
		}
		// Calculate the minimum number of cubes needed for the game
		red, green, blue := minimumCubesNeeded(game)
		// Calculate the power of this set (which is the product of the numbers of red, green, and blue cubes)
		power := red * green * blue
		// Add this power to another running total
		totalPower += power
	}

	fmt.Println("Part 1:", totalIDs)
	fmt.Println("Part 2:", totalPower)
}
