package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	sum := 0
	regex := regexp.MustCompile(`\d`)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		digits := regex.FindAllString(line, -1)
		// Check if there are any digits
		if len(digits) > 0 {
			// Get first and last digit and sum them
			firstDigit := digits[0]
			lastDigit := digits[len(digits)-1]
			twoDigitNumber, _ := strconv.Atoi(firstDigit + lastDigit)
			sum += twoDigitNumber
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Print the total sum
	fmt.Println(sum)
}
