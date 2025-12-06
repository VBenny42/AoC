advent_of_code::solution!(5);

use std::{cmp, num::ParseIntError, ops::Range, str::FromStr};

struct Day05 {
    ranges: Vec<Range<u64>>,
    items: Vec<u64>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("No digit found: {0}")]
    NonDigitFound(String),
    #[error("Parse int error: {0}")]
    ParseIntError(#[from] ParseIntError),
}

impl FromStr for Day05 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let mut items_reached = false;

        let mut ranges = Vec::<Range<u64>>::new();
        let mut items = Vec::<u64>::new();

        for line in input.lines() {
            if line.is_empty() {
                items_reached = true;
                continue;
            }
            match items_reached {
                false => match line.split_once("-") {
                    None => {
                        return Err(ParseError::NonDigitFound(line.to_string()));
                    }
                    Some((left, right)) => {
                        ranges.push(Range {
                            start: left.parse::<u64>()?,
                            end: right.parse::<u64>()? + 1,
                        });
                    }
                },
                true => {
                    items.push(line.parse::<u64>()?);
                }
            }
        }

        let sorted_ranges = merge_overlapping_ranges(&mut ranges);

        Ok(Day05 {
            ranges: sorted_ranges,
            items,
        })
    }
}

fn merge_overlapping_ranges(ranges: &mut [Range<u64>]) -> Vec<Range<u64>> {
    ranges.sort_by(|a, b| a.start.cmp(&b.start));

    let mut result = Vec::<Range<u64>>::new();
    result.push(ranges[0].clone());

    for current in ranges.iter().skip(1) {
        let j = result.len() - 1;

        if current.start >= result[j].start && current.start < result[j].end {
            result[j].end = cmp::max(current.end, result[j].end);
        } else {
            result.push(current.clone())
        }
    }

    result
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day05::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(
        day.items
            .iter()
            .filter(|item| day.ranges.iter().any(|range| range.contains(item)))
            .count() as u64,
    )
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day05::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(
        day.ranges
            .iter()
            .map(|range| range.end - range.start)
            .sum::<u64>(),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(3));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(14));
    }
}
