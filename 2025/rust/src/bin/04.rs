advent_of_code::solution!(4);
use image::{ImageBuffer, Rgb};
use std::str::FromStr;

enum Spot {
    Empty,
    Paper,
}

type Position = (usize, usize);

struct Day04 {
    grid: Vec<Vec<Spot>>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Invalid Character in input: {0}")]
    InvalidCharacter(char),
}

impl FromStr for Day04 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let grid = input
            .lines()
            .map(|line| {
                line.chars()
                    .map(|ch| match ch {
                        '.' => Ok(Spot::Empty),
                        '@' => Ok(Spot::Paper),
                        _ => Err(ParseError::InvalidCharacter(ch)),
                    })
                    .collect::<Result<Vec<Spot>, Self::Err>>()
            })
            .collect::<Result<Vec<Vec<Spot>>, Self::Err>>()?;

        Ok(Day04 { grid })
    }
}

impl Day04 {
    fn correct_papers(&self, position: Position) -> bool {
        let (x, y) = position;
        let mut paper_neighbors = 0;

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
                && matches!(self.grid[new_y as usize][new_x as usize], Spot::Paper)
            {
                paper_neighbors += 1;
            }
        }

        paper_neighbors < 4
    }

    #[allow(dead_code)]
    fn to_image(&self) -> ImageBuffer<Rgb<u8>, Vec<u8>> {
        let width = self.grid[0].len() as u32;
        let height = self.grid.len() as u32;
        let mut img = ImageBuffer::new(width, height);

        for (i, row) in self.grid.iter().enumerate() {
            for (j, spot) in row.iter().enumerate() {
                let pixel = match spot {
                    Spot::Empty => Rgb([0, 0, 0]),
                    Spot::Paper => Rgb([255, 255, 255]),
                };
                img.put_pixel(j as u32, i as u32, pixel);
            }
        }

        img
    }

    fn remove_papers(&mut self) -> usize {
        let mut accessible_positions = Vec::<Position>::new();

        for (i, row) in self.grid.iter().enumerate() {
            for (j, spot) in row.iter().enumerate() {
                if let Spot::Paper = spot && self.correct_papers((j, i)) {
                        accessible_positions.push((j, i));
                }
            }
        }

        for (x, y) in accessible_positions.iter() {
            self.grid[*y][*x] = Spot::Empty;
        }

        accessible_positions.len()
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let mut day = match Day04::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    // // Uncomment for image
    // let img = day.to_image();
    // img.save("day04_part1.png").unwrap();

    Some(day.remove_papers() as u64)
}

pub fn part_two(input: &str) -> Option<u64> {
    let mut day = match Day04::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let mut total_papers_removed = 0;

    loop {
        let papers_removed = day.remove_papers();
        if papers_removed == 0 {
            break;
        }
        total_papers_removed += papers_removed;
    }

    // // Uncomment for image
    // let img = day.to_image();
    // img.save("day04_part2.png").unwrap();

    Some(total_papers_removed as u64)
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
        assert_eq!(result, Some(43));
    }
}
