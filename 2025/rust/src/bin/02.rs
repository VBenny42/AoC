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

fn is_repeated_substring(s: &str) -> bool {
    let length = s.len();
    if length == 0 {
        return false;
    }

    for p in 1..=length / 2 {
        if is_invalid(s, length / p) {
            return true;
        }
    }
    false
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day02::from_str(input) {
        Ok(day) => day,
        Err(_) => {
            eprintln!("Error parsing input");
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
            let id = id_num.to_string();
            if is_invalid(&id, 2) {
                sum += id_num;
            }
        }
    }

    Some(sum)
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day02::from_str(input) {
        Ok(day) => day,
        Err(_) => {
            eprintln!("Error parsing input");
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
            let id = id_num.to_string();
            if is_repeated_substring(&id) {
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
