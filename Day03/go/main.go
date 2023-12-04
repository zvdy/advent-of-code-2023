package main

import (
	"bufio"
	"os"
)

// Function to read input from a file and return a 2D list of characters
func readInput(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

// Function to calculate p1 and nums based on the grid G
func calculateP1AndNums(grid [][]rune) (int, map[[2]int][]int) {
	R := len(grid)
	C := len(grid[0])
	p1 := 0
	nums := make(map[[2]int][]int)

	for r := 0; r < R; r++ {
		gears := make(map[[2]int]bool) // positions of '*' characters next to the current number
		n := 0
		hasPart := false

		// Loop through the characters in the current row
		for c := 0; c <= C; c++ {
			if c < C && grid[r][c] >= '0' && grid[r][c] <= '9' {
				n = n*10 + int(grid[r][c]-'0')
				// Check if the current number has a part next to it
				for _, dr := range []int{-1, 0, 1} {
					for _, dc := range []int{-1, 0, 1} {
						nr, nc := r+dr, c+dc
						// Check if the current position is valid
						if nr >= 0 && nr < R && nc >= 0 && nc < C {
							ch := grid[nr][nc]
							// Check if the current character is a part of the current number
							if ch != '.' && (ch < '0' || ch > '9') {
								hasPart = true
							}
							if ch == '*' {
								gears[[2]int{nr, nc}] = true
							}
						}
					}
				}
				// Check if the current number is the last number in the current row
			} else if n > 0 {
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				if hasPart {
					p1 += n
				}
				n = 0
				hasPart = false
				gears = make(map[[2]int]bool)
			}
		}
	}

	return p1, nums
}

// Function to calculate p2 based on the grid G and nums
func calculateP2(grid [][]rune, nums map[[2]int][]int) int {
	R := len(grid)
	C := len(grid[0])
	p2 := 0

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == '*' {
				if len(nums[[2]int{r, c}]) == 2 {
					p2 += nums[[2]int{r, c}][0] * nums[[2]int{r, c}][1]
				}
			}
		}
	}

	return p2
}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	p1, nums := calculateP1AndNums(grid)
	p2 := calculateP2(grid, nums)

	println("p1:", p1)
	println("p2:", p2)
}
