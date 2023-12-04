import re

# Dictionary to match string representation of numbers to their integer values
matchmap = {
    "one":   1,
    "two":   2,
    "three": 3,
    "four":  4,
    "five":  5,
    "six":   6,
    "seven": 7,
    "eight": 8,
    "nine":  9,
    "0":     0,
    "1":     1,
    "2":     2,
    "3":     3,
    "4":     4,
    "5":     5,
    "6":     6,
    "7":     7,
    "8":     8,
    "9":     9,
}

def part1():
    # Open the input file
    with open("input.txt", "r") as file:
        lines = file.readlines()

    sum = 0
    regex = re.compile(r'\d')  # Regular expression to match digits

    # Loop over all lines in the file
    for line in lines:
        digits = regex.findall(line)  # Find all digits in the line
        # Check if there are any digits
        if len(digits) > 0:
            # Get first and last digit and sum them
            first_digit = digits[0]
            last_digit = digits[-1]
            two_digit_number = int(first_digit + last_digit)  # Convert string to integer
            sum += two_digit_number

    # Print the total sum
    print("Part one:", sum)

def part2():
    # Open the input file
    with open("input.txt", "r") as file:
        lines = file.readlines()

    sum = 0

    # Loop over all lines in the input
    for line in lines:
        first = -1
        last = -1
        minpos = len(line) + 1
        maxpos = -1

        # Loop over all entries in the matchmap
        for k, v in matchmap.items():
            p = line.find(k)  # Find the position of the key in the line
            if p == -1:
                # substring not found
                continue
            if p < minpos:
                first = v
                minpos = p

            p = line.rfind(k)  # Find the last position of the key in the line
            if p > maxpos:
                last = v
                maxpos = p

        sum += 10*first + last

    print("Part two:", sum)

def main():
    part1()
    part2()

if __name__ == "__main__":
    main()