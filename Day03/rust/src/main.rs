use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;

// Function to read input from a file and return a 2D list of characters
fn read_input(file_path: &str) -> io::Result<Vec<Vec<char>>> {
    let file = File::open(&Path::new(file_path))?;
    let reader = io::BufReader::new(file);
    let grid: Vec<Vec<char>> = reader.lines()
        .map(|line| line.unwrap().chars().collect())
        .collect();
    Ok(grid)
}

// Function to calculate p1 and nums based on the grid G
fn calculate_p1_and_nums(grid: &Vec<Vec<char>>) -> (i32, HashMap<(usize, usize), Vec<i32>>) {
    let R = grid.len();
    let C = grid[0].len();
    let mut p1 = 0;
    let mut nums: HashMap<(usize, usize), Vec<i32>> = HashMap::new();

    for r in 0..R {
        let mut gears: HashMap<(usize, usize), bool> = HashMap::new(); // positions of '*' characters next to the current number
        let mut n = 0;
        let mut has_part = false;

        for c in 0..=C {
            if c < C && grid[r][c].is_digit(10) {
                n = n * 10 + grid[r][c].to_digit(10).unwrap() as i32;
                for dr in [-1, 0, 1].iter() {
                    for dc in [-1, 0, 1].iter() {
                        let nr = (r as i32) + dr;
                        let nc = (c as i32) + dc;
                        if nr >= 0 && (nr as usize) < R && nc >= 0 && (nc as usize) < C {
                            let ch = grid[nr as usize][nc as usize];
                            if ch != '.' && !ch.is_digit(10) {
                                has_part = true;
                            }
                            if ch == '*' {
                                gears.insert((nr as usize, nc as usize), true);
                            }
                        }
                    }
                }
            } else if n > 0 {
                for gear in gears.keys() {
                    nums.entry(*gear).or_insert(Vec::new()).push(n);
                }
                if has_part {
                    p1 += n;
                }
                n = 0;
                has_part = false;
                gears.clear();
            }
        }
    }

    (p1, nums)
}

fn main() {
    let grid = read_input("input.txt").unwrap();
    let (p1, nums) = calculate_p1_and_nums(&grid);
    println!("p1: {}", p1);
}