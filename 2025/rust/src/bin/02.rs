use std::str::FromStr;

advent_of_code::solution!(2);

// Brute-force solution, can't be bothered to optimize right now

struct Day02 {
    ids: Vec<String>,
}

impl FromStr for Day02 {
    type Err = ();

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let ids = input
            .lines()
            .next()
            .map(|line| line.split(",").map(|s| s.to_string()).collect())
            .unwrap();

        Ok(Day02 { ids })
    }
}

#[allow(dead_code)]
fn is_invalid(id: &str, multiple: usize) -> bool {
    let length = id.len();
    if !length.is_multiple_of(multiple) {
        return false;
    }

    let pattern_len = length / multiple;
    let pattern = &id[..pattern_len];

    // Check chunk by chunk
    for chunk_start in (pattern_len..length).step_by(pattern_len) {
        if &id[chunk_start..chunk_start + pattern_len] != pattern {
            return false;
        }
    }

    true
}

fn is_invalid_int(id: u64, multiple: usize) -> bool {
    if id == 0 {
        return multiple == 1;
    }

    let length = id.ilog10() as usize + 1;

    if !length.is_multiple_of(multiple) {
        return false;
    }

    let pattern_len = length / multiple;
    let chunk_divisor = 10u64.pow(pattern_len as u32);

    // Extract pattern (leftmost chunk)
    let pattern = id / 10u64.pow((length - pattern_len) as u32);

    // Check remaining chunks from right to left
    let mut remaining = id;
    for _ in 0..multiple {
        let chunk = remaining % chunk_divisor;
        if chunk != pattern {
            return false;
        }
        remaining /= chunk_divisor;
    }

    true
}

fn pattern_is_repeated(id: u64) -> bool {
    let length = if id == 0 { 1 } else { id.ilog10() as usize + 1 };
    if id == 0 {
        return false;
    }

    for p in 1..=length / 2 {
        if is_invalid_int(id, length / p) {
            return true;
        }
    }
    false
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day02::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input {error:?}");
            return None;
        }
    };

    let mut sum = 0;

    for pair in day.ids {
        let pairs = pair.splitn(2, "-").collect::<Vec<&str>>();
        if pairs.len() != 2 {
            println!("Invalid pair: {}", pair);
            return None;
        }

        let start: u64 = pairs[0].parse().unwrap();
        let end: u64 = pairs[1].parse().unwrap();

        for id_num in start..=end {
            if is_invalid_int(id_num, 2) {
                sum += id_num;
            }
        }
    }

    Some(sum)
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day02::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input {error:?}");
            return None;
        }
    };

    let mut sum = 0;

    for pair in day.ids {
        let pairs = pair.splitn(2, "-").collect::<Vec<&str>>();
        if pairs.len() != 2 {
            println!("Invalid pair: {}", pair);
            return None;
        }

        let start: u64 = pairs[0].parse().unwrap();
        let end: u64 = pairs[1].parse().unwrap();

        for id_num in start..=end {
            if pattern_is_repeated(id_num) {
                sum += id_num;
            }
        }
    }

    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1227775554));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(4174379265));
    }
}
