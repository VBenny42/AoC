advent_of_code::solution!(11);

use std::{collections::HashMap, str::FromStr};

struct Day11 {
    devices: HashMap<String, Vec<String>>,
    cache: HashMap<(String, String), u64>,
}

impl Day11 {
    fn new() -> Self {
        Self {
            devices: HashMap::new(),
            cache: HashMap::new(),
        }
    }

    fn add_device(&mut self, name: String, connections: Vec<String>) {
        self.devices.insert(name, connections.into_iter().collect());
    }

    fn count_paths(&mut self, start: &str, end: &str) -> u64 {
        if let Some(&cached) = self.cache.get(&(start.to_string(), end.to_string())) {
            return cached;
        }

        if start == end {
            self.cache.insert((start.to_string(), end.to_string()), 1);
            return 1;
        }

        let neighbors = self.devices.get(start).cloned();

        let mut total = 0;
        if let Some(neighbors) = neighbors {
            for neighbor in neighbors {
                total += self.count_paths(&neighbor, end);
            }
        }

        self.cache
            .insert((start.to_string(), end.to_string()), total);
        total
    }
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("No colon found in line: {0}")]
    NoColon(String),
}

impl FromStr for Day11 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let mut day = Day11::new();

        for line in input.lines() {
            let parts: Vec<&str> = line.split(": ").collect();
            if parts.len() != 2 {
                return Err(ParseError::NoColon(line.to_string()));
            }
            let name = parts[0].to_string();
            let connections: Vec<String> =
                parts[1].split_whitespace().map(|s| s.to_string()).collect();
            day.add_device(name, connections);
        }

        Ok(day)
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let mut day = match Day11::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(day.count_paths("you", "out"))
}

pub fn part_two(input: &str) -> Option<u64> {
    let mut day = match Day11::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    // From looking at input, "fft" is always before "dac"
    // "svr" -> "fft" -> "dac" -> "out" is equivalent to "svr" -> "out"
    Some(
        day.count_paths("svr", "fft")
            * day.count_paths("fft", "dac")
            * day.count_paths("dac", "out"),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(5));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file_part(
            "examples", DAY, 2,
        ));
        assert_eq!(result, Some(2));
    }
}
