use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

// Function to read lines from a file
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    // Dictionary to match string representation of numbers to their integer values
    let mut matchmap = HashMap::new();
    matchmap.insert("one", 1);
    matchmap.insert("two", 2);
    matchmap.insert("three", 3);
    matchmap.insert("four", 4);
    matchmap.insert("five", 5);
    matchmap.insert("six", 6);
    matchmap.insert("seven", 7);
    matchmap.insert("eight", 8);
    matchmap.insert("nine", 9);
    matchmap.insert("0", 0);
    matchmap.insert("1", 1);
    matchmap.insert("2", 2);
    matchmap.insert("3", 3);
    matchmap.insert("4", 4);
    matchmap.insert("5", 5);
    matchmap.insert("6", 6);
    matchmap.insert("7", 7);
    matchmap.insert("8", 8);
    matchmap.insert("9", 9);

    // Read lines from the input file
    let lines = read_lines("input.txt").unwrap().collect::<Result<Vec<_>, _>>().unwrap();

    part1(&lines);
    part2(&lines, &matchmap);
}

fn part1(lines: &[String]) {
    let mut sum = 0;

    // Loop over all lines in the file
    for line in lines {
        let digits: Vec<_> = line.chars().filter(|c| c.is_digit(10)).collect();
        // Check if there are any digits
        if !digits.is_empty() {
            // Get first and last digit and sum them
            let two_digit_number = format!("{}{}", digits[0], digits[digits.len() - 1])
                .parse::<i32>()
                .unwrap();
            sum += two_digit_number;
        }
    }

    // Print the total sum
    println!("Part one: {}", sum);
}

fn part2(lines: &[String], matchmap: &HashMap<&str, i32>) {
    let mut sum = 0;

    // Loop over all lines in the input
    for line in lines {
        let mut first = -1;
        let mut last = -1;
        let mut minpos = line.len() + 1;
        let mut maxpos = -1;

        // Loop over all entries in the matchmap
        for (&k, &v) in matchmap {
            if let Some(p) = line.find(k) {
                if p < minpos {
                    first = v;
                    minpos = p;
                }
            }

            if let Some(p) = line.rfind(k) {
                if p > maxpos {
                    last = v;
                    maxpos = p;
                }
            }
        }

        sum += 10 * first + last;
    }

    println!("Part two: {}", sum);
}