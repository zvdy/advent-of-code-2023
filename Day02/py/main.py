import re

# Define the Game class
class Game:
    def __init__(self, id, cubes):
        self.id = id
        self.cubes = cubes

# Function to parse games from the input string
def parse_games(input):
    games = []
    game_regex = re.compile(r'Game (\d+): (.*)')
    cube_regex = re.compile(r'(\d+) (red|green|blue)')

    for line in input.split('\n'):
        match = game_regex.search(line)
        if match is None:
            continue

        id = int(match.group(1))
        cubes = []

        for cube_str in match.group(2).split(';'):
            cube = {}
            for cube_match in cube_regex.findall(cube_str):
                count = int(cube_match[0])
                color = cube_match[1]
                cube[color] = count
            cubes.append(cube)

        game = Game(id, cubes)
        games.append(game)

    return games

# Function to check if a game is possible
def is_game_possible(game, red, green, blue):
    for cube in game.cubes:
        if cube.get('red', 0) > red or cube.get('green', 0) > green or cube.get('blue', 0) > blue:
            return False
    return True

# Main function
def main():
    with open('input.txt', 'r') as file:
        input = file.read()

    games = parse_games(input)

    sum = 0
    for game in games:
        if is_game_possible(game, 12, 13, 14):
            sum += game.id

    print(sum)

if __name__ == "__main__":
    main()