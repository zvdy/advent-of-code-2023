def calculate_points(cards):
    total_points = 0

    for card in cards:
        # split the card into winning numbers and your numbers
        sides = card.split("|")
        winning_numbers = sides[0].split()
        your_numbers = sides[1].split()

        points = 0
        for your_number in your_numbers:
            # if a number matches, double the points (or set to 1 if this is the first match)
            if your_number in winning_numbers:
                points = 1 if points == 0 else points * 2

        total_points += points

    return total_points

def part1():
    # read the cards from the file
    with open('input.txt', 'r') as file:
        cards = file.read().splitlines()

    # calculate the total points
    result = calculate_points(cards)
    return f'part one: {result}'

def calculate_matches(cards):
    matches = [0] * len(cards)

    for i, card in enumerate(cards):
        # split the card into winning numbers and your numbers
        sides = card.split("|")
        winning_numbers = sides[0].split()
        your_numbers = sides[1].split()

        for your_number in your_numbers:
            # if a number matches, increment the match count for this card
            if your_number in winning_numbers:
                matches[i] += 1

    return matches

def part2():
    # read the cards from the file
    with open('input.txt', 'r') as file:
        cards = file.read().splitlines()

    # calculate the number of matches for each card
    matches = calculate_matches(cards)
    cards_won = [1] * len(cards)

    for i, match in enumerate(matches):
        # for each match, add the number of cards won to the next cards
        for j in range(1, match + 1):
            if i + j < len(cards):
                cards_won[i + j] += cards_won[i]

    # calculate the total number of cards won
    total = sum(cards_won)

    return f'Part two: {total}'

def main():
    # print the results of part 1 and part 2
    print(part1())
    print(part2())

if __name__ == "__main__":
    main()