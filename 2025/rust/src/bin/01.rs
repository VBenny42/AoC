advent_of_code::solution!(1);
use std::error::Error;

struct Day01 {
    values: Vec<i32>,
}

const MOD: i32 = 100;
const INITIAL_VALUE: i32 = 50;

fn parse(input: &str) -> Result<Day01, Box<dyn Error>> {
    let values = input
        .lines()
        .map(|line| {
            let (sign, number) = line.split_at(1);
            match sign {
                "L" => Ok(-number.parse::<i32>()?),
                "R" => Ok(number.parse::<i32>()?),
                _ => Err("Invalid rotation".into()),
            }
        })
        .collect::<Result<Vec<i32>, Box<dyn Error>>>()?;
    Ok(Day01 { values })
}

pub fn part_one(input: &str) -> Option<u64> {
    let data = parse(input).ok()?;
    let mut actual_password: u64 = 0;

    let mut current_value = INITIAL_VALUE;

    for value in data.values {
        current_value = (current_value + value).rem_euclid(MOD);

        if current_value == 0 {
            actual_password += 1;
        }
    }

    Some(actual_password)
}

pub fn part_two(input: &str) -> Option<u64> {
    let data = parse(input).ok()?;
    let mut actual_password: u64 = 0;

    let mut current_value = INITIAL_VALUE;

    for value in data.values {
        let new_value = (current_value + value).rem_euclid(MOD);

        let (start, end) = match value.signum() {
            -1 => (
                (new_value - 1).rem_euclid(MOD),
                (current_value - 1).rem_euclid(MOD),
            ),
            1 => (current_value, new_value),
            _ => panic!("Rotation will not be zero"),
        };

        current_value = new_value;

        // If end < start, we wrapped around the modulo once
        if end < start {
            actual_password += 1
        }
        // Count how many complete modulo cycles were made
        actual_password += value.abs().div_euclid(MOD) as u64;
    }

    Some(actual_password)
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
        assert_eq!(result, Some(6));

        let result = part_two(&advent_of_code::template::read_file_part(
            "examples", DAY, 2,
        ));
        assert_eq!(result, Some(10));
    }
}
