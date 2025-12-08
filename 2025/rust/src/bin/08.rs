advent_of_code::solution!(8);

use std::str::FromStr;

use itertools::Itertools;

type Point = [u64; 3];

struct Day08 {
    points: Vec<Point>,
}

trait PointExt {
    fn euclidean_distance(&self, other: &Self) -> u64;
}

impl PointExt for Point {
    fn euclidean_distance(&self, other: &Self) -> u64 {
        (self[0] as i64 - other[0] as i64).pow(2) as u64
            + (self[1] as i64 - other[1] as i64).pow(2) as u64
            + (self[2] as i64 - other[2] as i64).pow(2) as u64
    }
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Non-digit character found in line: {0}")]
    NonDigitLine(String),
    #[error("Parse int error: {0}")]
    ParseIntError(#[from] std::num::ParseIntError),
}

impl FromStr for Day08 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        Ok(Day08 {
            points: input
                .lines()
                .map(|line| {
                    let bytes: Vec<u64> = line
                        .splitn(3, ',')
                        .map(|c| c.trim().parse::<u64>().map_err(ParseError::from))
                        .collect::<Result<Vec<u64>, ParseError>>()?;

                    if bytes.len() != 3 {
                        return Err(ParseError::NonDigitLine(line.to_string()));
                    }

                    Ok([bytes[0], bytes[1], bytes[2]])
                })
                .collect::<Result<Vec<Point>, ParseError>>()?,
        })
    }
}

fn get_sorted_edges(points: &[Point]) -> Vec<(usize, usize)> {
    (0..points.len())
        .tuple_combinations()
        .sorted_by_key(|&(i, j)| points[i].euclidean_distance(&points[j]))
        .collect()
}

fn find(parents: &mut [usize], id: usize) -> usize {
    if id != parents[id] {
        parents[id] = find(parents, parents[id]);
    }
    parents[id]
}

fn merge(parents: &mut [usize], i: usize, j: usize) {
    let i_root = find(parents, i);
    let j_root = find(parents, j);
    if i_root != j_root {
        parents[j_root] = i_root;
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day08::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let mut parents: Vec<usize> = (0..day.points.len()).collect();

    #[cfg(test)]
    const NUMBER_OF_TRIES: usize = 10;

    #[cfg(not(test))]
    const NUMBER_OF_TRIES: usize = 1000;

    for &(i, j) in get_sorted_edges(&day.points).iter().take(NUMBER_OF_TRIES) {
        merge(&mut parents, i, j);
    }

    let mut sizes = vec![0u64; day.points.len()];
    for i in 0..day.points.len() {
        let root = find(&mut parents, i);
        sizes[root] += 1;
    }

    Some(sizes.iter().sorted().rev().take(3).product())
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day08::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let mut parents: Vec<usize> = (0..day.points.len()).collect();
    let mut count = 0;
    let len = day.points.len();

    for (i, j) in get_sorted_edges(&day.points) {
        if find(&mut parents, i) != find(&mut parents, j) {
            count += 1;
            if count == len - 1 {
                return Some(day.points[i][0] * day.points[j][0]);
            }
            merge(&mut parents, i, j);
        }
    }

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(40));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(25272));
    }
}
