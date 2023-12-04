use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;

// Function to read the input file and return a grid
fn read_input(file_path: &str) -> io::Result<Vec<Vec<char>>> {
    let path = Path::new(file_path);
    let file = File::open(&path)?;
    let grid: Vec<Vec<char>> = io::BufReader::new(file).lines()
        .map(|line| line.unwrap().chars().collect())
        .collect();
    Ok(grid)
}

// Function to calculate p1 and nums based on the grid
fn calculate_p1_and_nums(grid: &Vec<Vec<char>>) -> (i32, HashMap<(usize, usize), Vec<i32>>) {
    let (R, C) = (grid.len(), grid[0].len());
    let (mut p1, mut nums) = (0, HashMap::new());

    for r in 0..R {
        let (mut gears, mut n, mut has_part) = (HashMap::new(), 0, false);

        for c in 0..=C {
            if c < C && grid[r][c].is_digit(10) {
                n = n * 10 + grid[r][c].to_digit(10).unwrap() as i32;
                for dr in -1..=1 {
                    for dc in -1..=1 {
                        let (nr, nc) = ((r as i32 + dr) as usize, (c as i32 + dc) as usize);
                        if nr < R && nc < C {
                            let ch = grid[nr][nc];
                            if ch != '.' && !ch.is_digit(10) {
                                has_part = true;
                            }
                            if ch == '*' {
                                gears.insert((nr, nc), true);
                            }
                        }
                    }
                }
            } else if n > 0 {
                for (&gear, _) in &gears {
                    nums.entry(gear).or_insert(Vec::new()).push(n);
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

// Function to calculate p2 based on the grid and nums
fn calculate_p2(grid: &Vec<Vec<char>>, nums: &HashMap<(usize, usize), Vec<i32>>) -> i32 {
    let (R, C) = (grid.len(), grid[0].len());
    let mut p2 = 0;

    for r in 0..R {
        for c in 0..C {
            if grid[r][c] == '*' {
                if let Some(num) = nums.get(&(r, c)) {
                    if num.len() == 2 {
                        p2 += num[0] * num[1];
                    }
                }
            }
        }
    }

    p2
}

// Main function
fn main() -> io::Result<()> {
    let grid = read_input("input.txt")?;
    let (p1, nums) = calculate_p1_and_nums(&grid);
    let p2 = calculate_p2(&grid, &nums);

    println!("p1: {}", p1);
    println!("p2: {}", p2);

    Ok(())
}