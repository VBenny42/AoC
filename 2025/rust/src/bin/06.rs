advent_of_code::solution!(6);

use std::str::FromStr;

#[derive(Clone, Debug)]
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

fn transpose(input: &str) -> Vec<String> {
    let lines = input.lines().collect::<Vec<&str>>();
    let longest_line_length = lines.iter().map(|line| line.len()).max().unwrap_or(0);

    let padded_lines: Vec<String> = lines
        .iter()
        .map(|line| {
            if line.len() < longest_line_length {
                format!("{}{}", line, " ".repeat(longest_line_length - line.len()))
            } else {
                line.to_string()
            }
        })
        .collect();

    let mut transposed = vec![String::new(); longest_line_length];
    for line in padded_lines {
        for (i, ch) in line.chars().enumerate() {
            transposed[i].push(ch);
        }
    }
    transposed
}

pub fn part_two(input: &str) -> Option<u64> {
    let transposed_lines = transpose(input);

    let mut problems = Vec::<Problem>::new();
    let mut on_new_problem = true;
    let mut current_problem = Problem {
        operator: '+',
        numbers: Vec::new(),
    };

    for line in transposed_lines {
        let trimmed = line.trim();
        let len = trimmed.len();

        match on_new_problem {
            true => {
                let operator = trimmed[len - 1..].chars().next().unwrap();
                if operator != '+' && operator != '*' {
                    eprintln!("Error parsing input: Invalid operator {}", operator);
                    return None;
                }

                let first_number = match &trimmed[..len - 1].trim().parse::<u64>() {
                    Ok(num) => *num,
                    Err(e) => {
                        eprintln!("Error parsing input: Invalid number {}", e);
                        return None;
                    }
                };

                current_problem = Problem {
                    operator,
                    numbers: vec![first_number],
                };
                on_new_problem = false;
            }
            false => match len {
                0 => {
                    problems.push(current_problem.clone());
                    on_new_problem = true;
                }
                _ => match trimmed.parse::<u64>() {
                    Ok(num) => current_problem.numbers.push(num),
                    Err(e) => {
                        eprintln!("Error parsing input: Invalid number {}", e);
                        return None;
                    }
                },
            },
        }
    }

    // Push last problem if not already pushed
    if !on_new_problem {
        problems.push(current_problem);
    }

    Some(problems.iter().map(|p| p.evaluate()).sum())
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
        assert_eq!(result, Some(3263827));
    }
}
