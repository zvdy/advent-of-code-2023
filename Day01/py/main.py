import re

# Open the input file
with open('input.txt', 'r') as file:
    lines = file.readlines()

sum = 0
regex = re.compile('\d')

# Loop over all lines in the file
for line in lines:
    digits = regex.findall(line)
    # Check if there are any digits
    if len(digits) > 0:
        # Get first and last digit and sum them
        first_digit = digits[0]
        last_digit = digits[-1]
        two_digit_number = int(first_digit + last_digit)
        sum += two_digit_number

# Print the total sum
print(sum)