advent_of_code::solution!(9);

use std::str::FromStr;

use geo::{Coord, Covers, LineString, Polygon};
use itertools::Itertools;

type Point = (i64, i64);

struct Day09 {
    points: Vec<Point>,
}

trait PointExt {
    fn area(&self, other: &Self) -> u64;
}

impl PointExt for Point {
    fn area(&self, other: &Self) -> u64 {
        let width = (other.0 - self.0).unsigned_abs() + 1;
        let height = (other.1 - self.1).unsigned_abs() + 1;
        width * height
    }
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {}

impl FromStr for Day09 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let points = input
            .lines()
            .map(|line| {
                let mut parts = line.split(',');
                let x = parts
                    .next()
                    .and_then(|s| s.trim().parse::<i64>().ok())
                    .unwrap_or(0);
                let y = parts
                    .next()
                    .and_then(|s| s.trim().parse::<i64>().ok())
                    .unwrap_or(0);
                (x, y)
            })
            .collect();

        Ok(Day09 { points })
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day09::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let sorted_areas = day
        .points
        .iter()
        .tuple_combinations()
        .map(|(a, b)| a.area(b))
        .sorted();

    Some(sorted_areas.last().unwrap())
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day09::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let polygon = Polygon::new(LineString::from(day.points.clone()), vec![]);

    let valid_points = day
        .points
        .iter()
        .tuple_combinations()
        .filter(|(a, b)| {
            let c = Coord { x: b.0, y: a.1 };
            let d = Coord { x: a.0, y: b.1 };
            polygon.covers(&c) && polygon.covers(&d)
        })
        .map(|(a, b)| a.area(b))
        .sorted();

    Some(valid_points.last().unwrap())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(50));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(24));
    }
}
