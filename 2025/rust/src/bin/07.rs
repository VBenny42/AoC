advent_of_code::solution!(7);

use std::str::FromStr;

use advent_of_code::{DOWN, Grid, GridExt, LEFT, Point, RIGHT, UP};

#[derive(Clone)]
enum Spot {
    Empty,
    Source,
    Splitter,
    Ray(u64),
}

struct Day07 {
    manifold: Grid<Spot>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Invalid Character in input: {0}")]
    InvalidCharacter(char),
}

impl FromStr for Day07 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let manifold = input
            .lines()
            .map(|line| {
                line.chars()
                    .map(|ch| match ch {
                        '.' => Ok(Spot::Empty),
                        '^' => Ok(Spot::Splitter),
                        'S' => Ok(Spot::Source),
                        _ => Err(ParseError::InvalidCharacter(ch)),
                    })
                    .collect::<Result<Vec<Spot>, Self::Err>>()
            })
            .collect::<Result<Vec<Vec<Spot>>, Self::Err>>()?;

        Ok(Day07 { manifold })
    }
}

fn traverse_manifold(manifold: &mut Grid<Spot>) -> Result<u64, String> {
    let height = manifold.len();
    let width = manifold[0].len();

    let mut total_split = 0;

    for y in 0..height {
        for x in 0..width {
            let spot = manifold[y][x].clone();
            let point = Point::new(x as isize, y as isize);

            match spot {
                Spot::Source => {
                    let down_pos = point + DOWN;
                    manifold.set(&down_pos, Spot::Ray(1))?;
                }
                // Copy down any rays coming from above
                Spot::Empty => {
                    let up_pos = point + UP;
                    if let Some(Spot::Ray(beam_num)) = manifold.get(&up_pos) {
                        manifold.set(&point, Spot::Ray(*beam_num))?;
                    }
                }
                Spot::Splitter => {
                    let up_pos = point + UP;
                    let mut actions = Vec::<(Spot, Point)>::new();
                    if let Some(Spot::Ray(beam_num)) = manifold.get(&up_pos) {
                        let left_pos = point + LEFT;
                        match manifold.get(&left_pos) {
                            Some(Spot::Empty) => {
                                actions.push((Spot::Ray(*beam_num), left_pos));
                            }
                            Some(Spot::Ray(existing_num)) => {
                                actions.push((Spot::Ray(existing_num + beam_num), left_pos));
                            }
                            _ => unreachable!(),
                        }
                        let right_pos = point + RIGHT;
                        match manifold.get(&right_pos) {
                            Some(Spot::Empty) => {
                                actions.push((Spot::Ray(*beam_num), right_pos));
                            }
                            Some(Spot::Ray(existing_num)) => {
                                actions.push((Spot::Ray(existing_num + beam_num), right_pos));
                            }
                            _ => unreachable!(),
                        }

                        total_split += 1; // ← Move it here, inside the if

                        for (new_spot, pos) in actions {
                            manifold.set(&pos, new_spot)?;
                        }
                    }
                }
                Spot::Ray(beam_num) => {
                    let down_pos = point + DOWN;
                    if let Some(Spot::Empty) = manifold.get(&down_pos) {
                        manifold.set(&down_pos, Spot::Ray(beam_num))?;
                    }
                    // Splitter case handled above
                }
            }
        }
    }
    Ok(total_split)
}

pub fn part_one(input: &str) -> Option<u64> {
    let mut day = match Day07::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let split = match traverse_manifold(&mut day.manifold) {
        Ok(split) => split,
        Err(error) => {
            eprintln!("Error traversing manifold: {:?}", error);
            return None;
        }
    };

    // let img = day.manifold.to_image(|spot| match spot {
    //     Spot::Empty => [0, 0, 0],
    //     Spot::Source => [0, 255, 0],
    //     Spot::Splitter => [0, 0, 255],
    //     Spot::Ray(_) => {
    //         // let intensity = ((*count as f64).ln() * 20.0).min(255.0) as u8;
    //         let intensity = 100;
    //         [intensity, intensity, intensity]
    //     }
    // });
    // img.save("day07_part1.png").unwrap();

    Some(split as u64)
}

fn count_timelines(manifold: &Grid<Spot>) -> u64 {
    let width = manifold[0].len();

    // Initialize first row count
    let mut current_row_counts = manifold[0]
        .iter()
        .map(|spot| match spot {
            Spot::Source => 1u64,
            _ => 0u64,
        })
        .collect::<Vec<u64>>();

    // Iterate through the rest
    for row in manifold.iter().skip(1) {
        let mut next_row_counts = vec![0u64; width];

        for x in 0..width {
            let n = current_row_counts[x];
            if n == 0 {
                continue; // No tachyons in this column
            }

            match &row[x] {
                Spot::Empty | Spot::Source => {
                    // n tachyons continue straight down
                    next_row_counts[x] += n;
                }
                Spot::Splitter => {
                    // Split left (x-1)
                    if x > 0 {
                        next_row_counts[x - 1] += n;
                    }
                    // Split right (x+1)
                    if x + 1 < width {
                        next_row_counts[x + 1] += n;
                    }
                }
                // Rays aren't actually added for part two
                _ => unreachable!(),
            }
        }

        current_row_counts = next_row_counts;
    }

    current_row_counts.iter().sum()
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day07::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(count_timelines(&day.manifold))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(21));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(40));
    }
}
