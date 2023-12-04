import re

# Game struct holds the game ID and the cube data.
class Game:
    def __init__(self, id, cubes):
        self.id = id
        self.cubes = cubes

# parse_games function parses the input and returns a list of Game.
def parse_games(input):
    games = []
    game_regex = re.compile(r"Game (\d+): (.*)")
    cube_regex = re.compile(r"(\d+) (red|green|blue)")

    # Iterate over each line in the input
    for line in input.split("\n"):
        matches = game_regex.search(line)
        if matches is None:
            continue

        # Parse the game ID
        id = int(matches.group(1))
        cubes = []

        # Parse the cube data for each game
        for cube_str in matches.group(2).split(";"):
            cube = {}
            for cube_match in cube_regex.findall(cube_str):
                count = int(cube_match[0])
                color = cube_match[1]
                cube[color] = count
            cubes.append(cube)

        game = Game(id, cubes)
        games.append(game)

    return games

# is_game_possible function checks if a game is possible with a given number of red, green, and blue cubes.
def is_game_possible(game, red, green, blue):
    # Iterate over each set of cubes in the game
    for cubes in game.cubes:
        # If any color has more cubes than the given number, the game is not possible
        if cubes.get("red", 0) > red or cubes.get("green", 0) > green or cubes.get("blue", 0) > blue:
            return False
    return True

# minimum_cubes_needed function calculates the minimum number of red, green, and blue cubes needed for a game.
def minimum_cubes_needed(game):
    red, green, blue = 0, 0, 0
    # Iterate over each set of cubes in the game
    for cubes in game.cubes:
        # Keep track of the maximum number of cubes of each color needed at any point
        red = max(red, cubes.get("red", 0))
        green = max(green, cubes.get("green", 0))
        blue = max(blue, cubes.get("blue", 0))
    return red, green, blue

# main function reads the input from a file, parses the games, and calculates the sum of the game IDs and the sum of the powers of the minimum sets of cubes.
def main():
    with open("input.txt", "r") as file:
        input = file.read()

    games = parse_games(input)

    total_ids = 0
    total_power = 0
    for game in games:
        # Check if the game is possible with 12 red, 13 green, and 14 blue cubes
        if is_game_possible(game, 12, 13, 14):
            # If so, add the game ID to a running total
            total_ids += game.id
        # Calculate the minimum number of cubes needed for the game
        red, green, blue = minimum_cubes_needed(game)
        # Calculate the power of this set (which is the product of the numbers of red, green, and blue cubes)
        power = red * green * blue
        # Add this power to another running total
        total_power += power

    print("Part 1:", total_ids)
    print("Part 2:", total_power)

if __name__ == "__main__":
    main()