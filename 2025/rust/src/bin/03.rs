use std::str::FromStr;

advent_of_code::solution!(3);

struct Day03 {
    banks: Vec<Vec<u8>>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {}

impl FromStr for Day03 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let banks = input
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(|c| c.to_digit(10).unwrap() as u8)
                    .collect::<Vec<u8>>()
            })
            .collect::<Vec<Vec<u8>>>();

        Ok(Day03 { banks })
    }
}

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
        .map(|bank| get_largest_combo(bank) as u64)
        .sum::<u64>()
        .into()
}

pub fn part_two(input: &str) -> Option<u64> {
    None
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
        assert_eq!(result, None);
    }
}
