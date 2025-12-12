advent_of_code::solution!(3);

use std::str::FromStr;

struct Day03 {
    banks: Vec<Vec<u8>>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("0 in input file")]
    ZeroEncountered,
    #[error("Invalid character found: {0}")]
    InvalidChar(char),
}

impl FromStr for Day03 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let banks = input
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(|c| match c.to_digit(10) {
                        None => Err(ParseError::InvalidChar(c)),
                        Some(0) => Err(ParseError::ZeroEncountered),
                        Some(d) => Ok(d as u8),
                    })
                    .collect::<Result<Vec<u8>, ParseError>>()
            })
            .collect::<Result<Vec<Vec<u8>>, ParseError>>()?;

        Ok(Day03 { banks })
    }
}

#[allow(dead_code)]
// Works for part 1, but needed to use stack solution for part 2
fn get_largest_combo(bank: &[u8]) -> u8 {
    let mut largest_pick: [u8; 2] = [0; 2];

    largest_pick.copy_from_slice(&bank[0..2]);

    /*
     * If current pick is [a, b] and next two numbers are [c, d]:
     * Valid picks:
     * - [a,b]
     * - [a,c]
     * - [a,d]
     * - [b,c]
     * - [b,d]
     * - [c,d]
     * Assuming b and c are not the same index
     *
     */

    // Use sliding window approach
    for i in 1..=bank.len() - 2 {
        let a = largest_pick[0];
        let b = largest_pick[1];
        let c = bank[i];
        let d = bank[i + 1];

        let picks = [
            ((10 * a) + b, [a, b]),
            ((10 * a) + c, [a, c]),
            ((10 * a) + d, [a, d]),
            // ((10 * b) + c, [b, c]),
            ((10 * b) + d, [b, d]),
            ((10 * c) + d, [c, d]),
        ];

        let (_, max_pick) = picks.iter().max_by_key(|(value, _)| *value).unwrap();
        largest_pick = *max_pick;
    }

    (10 * largest_pick[0]) + largest_pick[1]
}

fn get_largest_combo_k(bank: &[u8], k: usize) -> u64 {
    let mut stack: Vec<(u8, usize)> = Vec::with_capacity(k); // (digit, original_index)
    let n = bank.len();

    for (i, &digit) in bank.iter().enumerate() {
        // Remove smaller digits from stack if we have room to pick more digits
        while !stack.is_empty() && stack.last().unwrap().0 < digit && stack.len() + (n - i) > k {
            stack.pop();
        }

        if stack.len() < k {
            stack.push((digit, i));
        }
    }

    stack
        .into_iter()
        .map(|(digit, _)| digit)
        .fold(0, |acc, d| acc * 10 + d as u64)
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day03::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {}", error);
            return None;
        }
    };

    day.banks
        .iter()
        .map(|bank| get_largest_combo_k(bank, 2))
        .sum::<u64>()
        .into()
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day03::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {}", error);
            return None;
        }
    };

    day.banks
        .iter()
        .map(|bank| get_largest_combo_k(bank, 12))
        .sum::<u64>()
        .into()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(357));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(3121910778619));
    }
}
