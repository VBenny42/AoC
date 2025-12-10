advent_of_code::solution!(9);

use std::str::FromStr;

use itertools::Itertools;

type Point = (i64, i64);

struct Day09 {
    points: Vec<Point>,
}

trait PointExt {
    fn area(&self, other: &Self) -> u64;
    fn manhattan_distance(&self, other: &Self) -> u64;
}

impl PointExt for Point {
    fn area(&self, other: &Self) -> u64 {
        let width = (other.0 - self.0).unsigned_abs() + 1;
        let height = (other.1 - self.1).unsigned_abs() + 1;
        width * height
    }

    fn manhattan_distance(&self, other: &Self) -> u64 {
        (self.0 - other.0).unsigned_abs() + (self.1 - other.1).unsigned_abs()
    }
}

#[derive(Debug, Clone)]
struct Edge {
    x1: i64,
    y1: i64,
    x2: i64,
    y2: i64,
}

impl Edge {
    fn new(p1: Point, p2: Point) -> Self {
        Edge {
            x1: p1.0,
            y1: p1.1,
            x2: p2.0,
            y2: p2.1,
        }
    }

    fn intersects_rectangle(&self, min_x: i64, min_y: i64, max_x: i64, max_y: i64) -> bool {
        let edge_min_x = self.x1.min(self.x2);
        let edge_max_x = self.x1.max(self.x2);
        let edge_min_y = self.y1.min(self.y2);
        let edge_max_y = self.y1.max(self.y2);

        // Check if rectangles overlap
        min_x < edge_max_x && max_x > edge_min_x && min_y < edge_max_y && max_y > edge_min_y
    }
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Parse int error: {0}")]
    ParseIntError(#[from] std::num::ParseIntError),
}

impl FromStr for Day09 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let points = input
            .lines()
            .map(|line| {
                let mut parts = line.split(',');
                let x = parts
                    .next()
                    .and_then(|s| s.trim().parse::<i64>().map_err(ParseError::from).ok())
                    .unwrap_or(0);
                let y = parts
                    .next()
                    .and_then(|s| s.trim().parse::<i64>().map_err(ParseError::from).ok())
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
        .max();

    Some(sorted_areas.unwrap())
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day09::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    if day.points.is_empty() {
        return Some(0);
    }

    // Build polygon edges
    let mut edges = Vec::new();
    for i in 0..day.points.len() - 1 {
        edges.push(Edge::new(day.points[i], day.points[i + 1]));
    }
    // Close the polygon
    if let (Some(&first), Some(&last)) = (day.points.first(), day.points.last()) {
        edges.push(Edge::new(first, last));
    }

    let mut max_area = 0u64;

    // Check all combinations of points
    for (i, &point_a) in day.points.iter().enumerate() {
        for &point_b in day.points.iter().skip(i) {
            let min_x = point_a.0.min(point_b.0);
            let max_x = point_a.0.max(point_b.0);
            let min_y = point_a.1.min(point_b.1);
            let max_y = point_a.1.max(point_b.1);

            // Early skip
            let manhattan = point_a.manhattan_distance(&point_b);
            if manhattan * manhattan <= max_area {
                continue;
            }

            // Check if rectangle intersects with any edge
            let intersects = edges
                .iter()
                .any(|edge| edge.intersects_rectangle(min_x, min_y, max_x, max_y));

            if !intersects {
                let area = point_a.area(&point_b);
                max_area = max_area.max(area);
            }
        }
    }

    Some(max_area)
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
