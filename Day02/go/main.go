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

	for _, line := range strings.Split(input, "\n") {
		matches := gameRegex.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		id, _ := strconv.Atoi(matches[1])
		game := Game{ID: id}

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

// isGamePossible function checks if a game is possible given a certain number of red, green, and blue cubes.
func isGamePossible(game Game, red, green, blue int) bool {
	for _, cubes := range game.Cubes {
		if cubes["red"] > red || cubes["green"] > green || cubes["blue"] > blue {
			return false
		}
	}
	return true
}

// main function reads the input from a file, parses the games, and calculates the sum of the IDs of the possible games.
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

	sum := 0
	for _, game := range games {
		if isGamePossible(game, 12, 13, 14) {
			sum += game.ID
		}
	}

	fmt.Println(sum)
}
