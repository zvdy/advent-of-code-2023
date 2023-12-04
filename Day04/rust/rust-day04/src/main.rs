use std::fs;

// calculate_points calculates the points for each card
fn calculate_points(cards: &Vec<String>) -> i32 {
    let mut total_points = 0;

    for card in cards {
        // split the card into winning numbers and your numbers
        let sides: Vec<&str> = card.split("|").collect();
        let winning_numbers: Vec<&str> = sides[0].split_whitespace().collect();
        let your_numbers: Vec<&str> = sides[1].split_whitespace().collect();

        let mut points = 0;
        for your_number in &your_numbers {
            if winning_numbers.contains(your_number) {
                // if a number matches, double the points (or set to 1 if this is the first match)
                points = if points == 0 { 1 } else { points * 2 };
            }
        }

        total_points += points;
    }

    total_points
}

// part1 reads the cards from the file and calculates the total points
fn part1() -> String {
    let cards = fs::read_to_string("input.txt").unwrap();
    let cards: Vec<String> = cards.lines().map(|s| s.to_string()).collect();

    let result = calculate_points(&cards);
    format!("part one: {}", result)
}

// calculate_matches calculates the number of matches for each card
fn calculate_matches(cards: &Vec<String>) -> Vec<i32> {
    let mut matches = vec![0; cards.len()];

    for (i, card) in cards.iter().enumerate() {
        // split the card into winning numbers and your numbers
        let sides: Vec<&str> = card.split("|").collect();
        let winning_numbers: Vec<&str> = sides[0].split_whitespace().collect();
        let your_numbers: Vec<&str> = sides[1].split_whitespace().collect();

        for your_number in &your_numbers {
            if winning_numbers.contains(your_number) {
                // if a number matches, increment the match count for this card
                matches[i] += 1;
            }
        }
    }

    matches
}

// part2 reads the cards from the file and calculates the total number of cards won
fn part2() -> String {
    let cards = fs::read_to_string("input.txt").unwrap();
    let cards: Vec<String> = cards.lines().map(|s| s.to_string()).collect();

    let matches = calculate_matches(&cards);
    let mut cards_won = vec![1; cards.len()];

    for (i, &match_) in matches.iter().enumerate() {
        // for each match, add the number of cards won to the next cards
        for j in 1..=match_ {
            if i + j < cards.len() {
                cards_won[i + j] += cards_won[i];
            }
        }
    }

    let total: i32 = cards_won.iter().sum();
    format!("Part two: {}", total)
}

// main function prints the results of part 1 and part 2
fn main() {
    println!("{}", part1());
    println!("{}", part2());
}