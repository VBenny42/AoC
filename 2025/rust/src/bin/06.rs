advent_of_code::solution!(6);

use std::str::FromStr;

struct Problem {
    operator: char,
    numbers: Vec<u64>,
}

struct Day06 {
    problems: Vec<Problem>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Operator not allowed: {0}")]
    InvalidOperator(char),
    #[error("Failed to parse number: {0}")]
    InvalidNumber(#[from] std::num::ParseIntError),
}

impl FromStr for Day06 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let array = input
            .lines()
            .map(|line| line.split_whitespace().collect::<Vec<_>>())
            .collect::<Vec<_>>();

        let mut transposed: Vec<Vec<&str>> = Vec::new();
        for i in 0..array[0].len() {
            let mut column = Vec::new();
            for row in &array {
                column.push(row[i]);
            }
            transposed.push(column);
        }

        let problems = transposed
            .into_iter()
            .map(|mut items| {
                let operator = items.remove(items.len() - 1).chars().next().unwrap();
                if operator != '+' && operator != '*' {
                    return Err(ParseError::InvalidOperator(operator));
                }

                let numbers: Result<Vec<u64>, _> = items
                    .into_iter()
                    .map(|num_str| num_str.parse::<u64>().map_err(ParseError::from))
                    .collect();

                Ok(Problem {
                    operator,
                    numbers: numbers?,
                })
            })
            .collect::<Result<Vec<Problem>, ParseError>>()?;

        Ok(Day06 { problems })
    }
}

impl Problem {
    fn evaluate(&self) -> u64 {
        match self.operator {
            '+' => self.numbers.iter().sum(),
            '*' => self.numbers.iter().product(),
            _ => unreachable!(),
        }
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day06::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(day.problems.iter().map(|p| p.evaluate()).sum())
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day06::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(4277556));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
