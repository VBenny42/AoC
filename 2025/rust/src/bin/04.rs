advent_of_code::solution!(4);
use image::{ImageBuffer, Rgb};
use std::str::FromStr;

enum Spot {
    Empty,
    Paper,
    Accessible,
}

type Position = (usize, usize);

struct Day04 {
    grid: Vec<Vec<Spot>>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {}

impl FromStr for Day04 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let grid = input
            .lines()
            .map(|line| {
                line.chars()
                    .map(|ch| match ch {
                        '.' => Spot::Empty,
                        '@' => Spot::Paper,
                        _ => Spot::Empty, // Shouldn't happen
                    })
                    .collect()
            })
            .collect();

        Ok(Day04 { grid })
    }
}

impl Day04 {
    fn get_neighbors(&self, position: Position) -> Vec<Position> {
        let (x, y) = position;
        let mut neighbors = Vec::new();

        let directions = [
            (-1, 0),
            (1, 0),
            (0, -1),
            (0, 1),
            (-1, -1),
            (1, 1),
            (-1, 1),
            (1, -1),
        ];

        for (dx, dy) in directions.iter() {
            let new_x = x as isize + dx;
            let new_y = y as isize + dy;

            if new_x >= 0
                && new_x < self.grid[0].len() as isize
                && new_y >= 0
                && new_y < self.grid.len() as isize
            {
                neighbors.push((new_x as usize, new_y as usize));
            }
        }

        neighbors
    }

    #[allow(dead_code)]
    fn to_image(&self) -> ImageBuffer<Rgb<u8>, Vec<u8>> {
        let width = self.grid[0].len() as u32;
        let height = self.grid.len() as u32;
        let mut img = ImageBuffer::new(width, height);

        for (i, row) in self.grid.iter().enumerate() {
            for (j, spot) in row.iter().enumerate() {
                let pixel = match spot {
                    Spot::Empty => Rgb([255, 255, 255]),
                    Spot::Paper => Rgb([0, 0, 0]),
                    Spot::Accessible => Rgb([255, 0, 0]),
                };
                img.put_pixel(j as u32, i as u32, pixel);
            }
        }

        img
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day04::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let mut accessible_spots = 0;

    // let mut accessible_positions = Vec::<Position>::new();

    for (i, row) in day.grid.iter().enumerate() {
        for (j, spot) in row.iter().enumerate() {
            if let Spot::Paper = spot {
                let neighbors = day
                    .get_neighbors((j, i))
                    .iter()
                    .filter(|(x, y)| matches!(day.grid[*y][*x], Spot::Paper))
                    .count();
                if neighbors < 4 {
                    accessible_spots += 1;
                    // accessible_positions.push((j, i));
                }
            }
        }
    }

    // for (x, y) in accessible_positions.iter() {
    //     day.grid[*y][*x] = Spot::Accessible;
    // }

    // Uncomment for image
    // let img = day.to_image();
    // img.save("day04_part1.png").unwrap();

    Some(accessible_spots)
}

pub fn part_two(input: &str) -> Option<u64> {
    let _day = match Day04::from_str(input) {
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
        assert_eq!(result, Some(13));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
