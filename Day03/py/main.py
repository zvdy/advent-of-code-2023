import sys
import re
from collections import defaultdict

# Function to read input from a file and return a 2D list of characters
def read_input(file_path):
    D = open(file_path).read().strip()
    lines = D.split('\n')
    G = [[c for c in line] for line in lines]
    return G

# Function to calculate p1 and nums based on the grid G
def calculate_p1_and_nums(G):
    R = len(G)
    C = len(G[0])
    p1 = 0
    nums = defaultdict(list)
    for r in range(len(G)):
        gears = set() # positions of '*' characters next to the current number
        n = 0
        has_part = False
        for c in range(len(G[r])+1):
            if c<C and G[r][c].isdigit():
                n = n*10+int(G[r][c])
                for rr in [-1,0,1]:
                    for cc in [-1,0,1]:
                        if 0<=r+rr<R and 0<=c+cc<C:
                            ch = G[r+rr][c+cc]
                            if not ch.isdigit() and ch != '.':
                                has_part = True
                            if ch=='*':
                                gears.add((r+rr, c+cc))
            elif n>0:
                for gear in gears:
                    nums[gear].append(n)
                if has_part:
                    p1 += n
                n = 0
                has_part = False
                gears = set()
    return p1, nums

# Function to calculate p2 based on nums
def calculate_p2(nums):
    p2 = 0
    for k,v in nums.items():
        if len(v)==2:
            p2 += v[0]*v[1]
    return p2

# Main function to call the other functions and print the results
def main():
    G = read_input("input.txt") # Read input from "input.txt"
    p1, nums = calculate_p1_and_nums(G) # Calculate p1 and nums
    print(p1) # Print p1
    p2 = calculate_p2(nums) # Calculate p2
    print(p2) # Print p2

if __name__ == "__main__":
    main()