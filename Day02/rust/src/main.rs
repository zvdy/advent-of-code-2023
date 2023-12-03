use regex::Regex;
use std::collections::HashMap;
use std::fs;

// Define the Game struct
struct Game {
    id: i32,
    cubes: Vec<HashMap<String, i32>>,
}

// Function to parse games from the input string
fn parse_games(input: &str) -> Vec<Game> {
    let mut games = Vec::new();
    let game_regex = Regex::new(r"Game (\d+): (.*)").unwrap();
    let cube_regex = Regex::new(r"(\d+) (red|green|blue)").unwrap();

    for line in input.lines() {
        if let Some(game_match) = game_regex.captures(line) {
            let id = game_match[1].parse::<i32>().unwrap();
            let mut cubes = Vec::new();

            for cube_str in game_match[2].split(';') {
                let mut cube = HashMap::new();
                for cube_match in cube_regex.captures_iter(cube_str) {
                    let count = cube_match[1].parse::<i32>().unwrap();
                    let color = cube_match[2].to_string();
                    cube.insert(color, count);
                }
                cubes.push(cube);
            }

            let game = Game { id, cubes };
            games.push(game);
        }
    }

    games
}

// Function to check if a game is possible
fn is_game_possible(game: &Game, red: i32, green: i32, blue: i32) -> bool {
    for cube in &game.cubes {
        if *cube.get("red").unwrap_or(&0) > red || *cube.get("green").unwrap_or(&0) > green || *cube.get("blue").unwrap_or(&0) > blue {
            return false;
        }
    }
    true
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("Failed to read file");
    let games = parse_games(&input);

    let mut sum = 0;
    for game in &games {
        if is_game_possible(game, 12, 13, 14) {
            sum += game.id;
        }
    }

    println!("{}", sum);
}