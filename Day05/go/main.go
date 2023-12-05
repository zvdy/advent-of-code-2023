package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct to hold the map data
type Map struct {
	destinationStart int
	sourceStart      int
	length           int
}

// Parse the input file into a map of maps
func parseInput(filename string) ([]int, map[string][]Map) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maps := make(map[string][]Map)
	var seeds []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds:") {
			seedsStr := strings.Split(line, " ")[1:]
			for _, seedStr := range seedsStr {
				seed, _ := strconv.Atoi(seedStr)
				seeds = append(seeds, seed)
			}
			// Parse the map data into a map of maps
		} else if strings.Contains(line, "map:") {
			mapName := strings.Split(line, " ")[0]
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}
				nums := strings.Split(line, " ")
				destinationStart, _ := strconv.Atoi(nums[0])
				sourceStart, _ := strconv.Atoi(nums[1])
				length, _ := strconv.Atoi(nums[2])
				maps[mapName] = append(maps[mapName], Map{destinationStart, sourceStart, length})
			}
		}
	}

	return seeds, maps
}

// Create a map from the given maps
func createMap(maps []Map, num int) int {
	for _, m := range maps {
		if num >= m.sourceStart && num < m.sourceStart+m.length {
			return m.destinationStart + (num - m.sourceStart)
		}
	}
	return num
}

// Find the lowest location from the given seeds and maps
func findLowestLocation(seeds []int, maps map[string][]Map) int {
	lowestLocation := -1
	for _, seed := range seeds {
		soil := createMap(maps["seed-to-soil"], seed)
		fertilizer := createMap(maps["soil-to-fertilizer"], soil)
		water := createMap(maps["fertilizer-to-water"], fertilizer)
		light := createMap(maps["water-to-light"], water)
		temperature := createMap(maps["light-to-temperature"], light)
		humidity := createMap(maps["temperature-to-humidity"], temperature)
		location := createMap(maps["humidity-to-location"], humidity)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

/*
--- Part Two ---

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?
*/

func main() {
	seeds, maps := parseInput("input.txt")
	fmt.Println(findLowestLocation(seeds, maps))
}
