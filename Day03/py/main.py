import re

# Function to read the input file and return a grid
def read_input(file_path):
    with open(file_path, 'r') as file:
        grid = [list(line.strip()) for line in file]
    return grid

# Function to calculate p1 and nums based on the grid
def calculate_p1_and_nums(grid):
    R, C = len(grid), len(grid[0])
    p1 = 0
    nums = {}

    for r in range(R):
        gears = {}
        n = 0
        has_part = False

        for c in range(C+1):
            # If the cell contains a number
            if c < C and '0' <= grid[r][c] <= '9':
                n = n*10 + int(grid[r][c])
                for dr in range(-1, 2):
                    for dc in range(-1, 2):
                        nr, nc = r+dr, c+dc
                        if 0 <= nr < R and 0 <= nc < C:
                            ch = grid[nr][nc]
                            # If the cell is not empty and not a number
                            if ch != '.' and not re.match('[0-9]', ch):
                                has_part = True
                            # If the cell is a gear
                            if ch == '*':
                                gears[(nr, nc)] = True
            # If the cell does not contain a number
            elif n > 0:
                for gear in gears:
                    if gear not in nums:
                        nums[gear] = []
                    nums[gear].append(n)
                if has_part:
                    p1 += n
                n = 0
                has_part = False
                gears = {}

    return p1, nums

# Function to calculate p2 based on the grid and nums
def calculate_p2(grid, nums):
    R, C = len(grid), len(grid[0])
    p2 = 0

    for r in range(R):
        for c in range(C):
            # If the cell is a gear
            if grid[r][c] == '*':
                # If the gear is adjacent to exactly two part numbers
                if (r, c) in nums and len(nums[(r, c)]) == 2:
                    # Add the gear ratio to p2
                    p2 += nums[(r, c)][0] * nums[(r, c)][1]

    return p2

# Main function
def main():
    grid = read_input("input.txt")
    p1, nums = calculate_p1_and_nums(grid)
    p2 = calculate_p2(grid, nums)

    print("p1:", p1)
    print("p2:", p2)

if __name__ == "__main__":
    main()