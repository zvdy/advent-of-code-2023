use regex::Regex;
use std::collections::HashMap;
use std::fs;

#[derive(Debug)]
struct Game {
    id: i32,
    cubes: Vec<HashMap<String, i32>>,
}

fn parse_games(input: &str) -> Vec<Game> {
    let mut games = Vec::new();
    let game_regex = Regex::new(r"Game (\d+): (.*)").unwrap();
    let cube_regex = Regex::new(r"(\d+) (red|green|blue)").unwrap();

    for line in input.lines() {
        if let Some(captures) = game_regex.captures(line) {
            let id = captures[1].parse::<i32>().unwrap();
            let mut cubes = Vec::new();

            for cube_str in captures[2].split(';') {
                let mut cube = HashMap::new();
                for cube_match in cube_regex.captures_iter(cube_str) {
                    let count = cube_match[1].parse::<i32>().unwrap();
                    let color = cube_match[2].to_string();
                    cube.insert(color, count);
                }
                cubes.push(cube);
            }

            games.push(Game { id, cubes });
        }
    }

    games
}

fn is_game_possible(game: &Game, red: i32, green: i32, blue: i32) -> bool {
    for cubes in &game.cubes {
        if *cubes.get("red").unwrap_or(&0) > red
            || *cubes.get("green").unwrap_or(&0) > green
            || *cubes.get("blue").unwrap_or(&0) > blue
        {
            return false;
        }
    }
    true
}

fn minimum_cubes_needed(game: &Game) -> (i32, i32, i32) {
    let (mut red, mut green, mut blue) = (0, 0, 0);
    for cubes in &game.cubes {
        red = red.max(*cubes.get("red").unwrap_or(&0));
        green = green.max(*cubes.get("green").unwrap_or(&0));
        blue = blue.max(*cubes.get("blue").unwrap_or(&0));
    }
    (red, green, blue)
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let games = parse_games(&input);

    let (mut total_ids, mut total_power) = (0, 0);
    for game in &games {
        if is_game_possible(game, 12, 13, 14) {
            total_ids += game.id;
        }
        let (red, green, blue) = minimum_cubes_needed(game);
        total_power += red * green * blue;
    }

    println!("Part 1: {}", total_ids);
    println!("Part 2: {}", total_power);
}