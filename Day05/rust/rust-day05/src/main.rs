use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

// Define the Map struct
struct Map {
    destination_start: i32,
    source_start: i32,
    length: i32,
}

// Function to parse the input file
fn parse_input(filename: &str) -> (Vec<i32>, std::collections::HashMap<String, Vec<Map>>) {
    // Create a path from the filename
    let path = Path::new(filename);
    // Open the file
    let file = File::open(&path).unwrap();
    // Create a buffered reader
    let reader = io::BufReader::new(file);

    // Initialize the maps and seeds
    let mut maps: std::collections::HashMap<String, Vec<Map>> = std::collections::HashMap::new();
    let mut seeds: Vec<i32> = Vec::new();

    // Initialize the current map
    let mut current_map: String = String::new();

    // Loop through each line in the file
    for line in reader.lines() {
        let line = line.unwrap();
        // Check if the line contains seeds
        if line.contains("seeds:") {
            seeds = line.split_whitespace().skip(1).map(|s| s.parse().unwrap()).collect();
        // Check if the line contains a map
        } else if line.contains("map:") {
            current_map = line.split_whitespace().next().unwrap().to_string();
        // Check if the line is not empty
        } else if !line.is_empty() {
            // Parse the numbers from the line
            let nums: Vec<i32> = line.split_whitespace().map(|s| s.parse().unwrap()).collect();
            // Create a new map from the numbers
            let map = Map {
                destination_start: nums[0],
                source_start: nums[1],
                length: nums[2],
            };
            // Add the map to the maps
            maps.entry(current_map.clone()).or_insert(Vec::new()).push(map);
        }
    }

    // Return the seeds and maps
    (seeds, maps)
}

// Function to create a map
fn create_map(maps: &Vec<Map>, num: i32) -> i32 {
    // Loop through each map
    for m in maps {
        // Check if the number is within the map's range
        if num >= m.source_start && num < m.source_start + m.length {
            return m.destination_start + (num - m.source_start);
        }
    }
    // Return the number if it is not in any map's range
    num
}

// Function to find the lowest location
fn find_lowest_location(seeds: Vec<i32>, maps: std::collections::HashMap<String, Vec<Map>>) -> i32 {
    // Initialize the lowest location
    let mut lowest_location = -1;
    // Loop through each seed
    for seed in seeds {
        // Create the soil, fertilizer, water, light, temperature, humidity, and location from the maps
        let soil = create_map(&maps["seed-to-soil"], seed);
        let fertilizer = create_map(&maps["soil-to-fertilizer"], soil);
        let water = create_map(&maps["fertilizer-to-water"], fertilizer);
        let light = create_map(&maps["water-to-light"], water);
        let temperature = create_map(&maps["light-to-temperature"], light);
        let humidity = create_map(&maps["temperature-to-humidity"], temperature);
        let location = create_map(&maps["humidity-to-location"], humidity);
        // Check if the location is the lowest found so far
        if lowest_location == -1 || location < lowest_location {
            lowest_location = location;
        }
    }
    // Return the lowest location
    lowest_location
}

// Main function
fn main() {
    // Parse the input file
    let (seeds, maps) = parse_input("input.txt");
    // Print the lowest location
    println!("{}", find_lowest_location(seeds, maps));
}