advent_of_code::solution!(10);

use std::{
    collections::{HashSet, VecDeque},
    str::FromStr,
};

use good_lp::{Expression, Solution, SolverModel, Variable, default_solver, variable, variables};

struct Day10 {
    machines: Vec<Machine>,
}

struct Machine {
    lights: Vec<bool>,
    buttons: Vec<Vec<usize>>,
    joltages: Vec<u64>,
}

#[derive(Debug, thiserror::Error)]
pub enum ParseError {
    #[error("Parse int error: {0}")]
    ParseIntError(#[from] std::num::ParseIntError),
}

impl FromStr for Day10 {
    type Err = ParseError;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let machines: Result<Vec<Machine>, ParseError> = input
            .lines()
            .map(|line| {
                let (lights_str, rest) = line.split_once(" ").unwrap();
                let (buttons_str, joltages_str) = rest.rsplit_once(" ").unwrap();

                let lights = lights_str
                    .trim_matches(['[', ']'])
                    .chars()
                    .map(|c| c == '#')
                    .collect();

                let buttons: Vec<Vec<_>> = buttons_str
                    .split(' ')
                    .map(|s| {
                        s.trim_matches(['(', ')'])
                            .split(',')
                            .map(|s| s.parse().map_err(ParseError::from))
                            .collect::<Result<Vec<_>, _>>()
                    })
                    .collect::<Result<Vec<_>, _>>()?;

                let joltages = joltages_str
                    .trim_matches(['{', '}'])
                    .split(',')
                    .map(|s| s.trim().parse::<u64>().unwrap())
                    .collect();

                Ok(Machine {
                    lights,
                    buttons,
                    joltages,
                })
            })
            .collect();

        Ok(Day10 {
            machines: machines?,
        })
    }
}

impl Machine {
    fn min_needed_button_presses(&self) -> usize {
        let goal =
            self.lights.iter().enumerate().fold(
                0u32,
                |acc, (i, &lit)| {
                    if lit { acc | (1 << i) } else { acc }
                },
            );

        let buttons = self
            .buttons
            .iter()
            .map(|btn| btn.iter().fold(0u32, |acc, &idx| acc | (1 << idx)))
            .collect::<Vec<u32>>();

        // BFS until first goal state is found
        let mut queue = VecDeque::from([(0u32, 0usize)]);
        let mut visited = HashSet::from([0u32]);

        while let Some((state, presses)) = queue.pop_front() {
            if state == goal {
                return presses;
            }

            for &btn in &buttons {
                let next_state = state ^ btn;
                if visited.insert(next_state) {
                    queue.push_back((next_state, presses + 1));
                }
            }
        }

        unreachable!("There should always be a solution");
    }

    fn min_needed_joltage_presses(&self) -> usize {
        let mut vars = variables!();

        let presses: Vec<Variable> = (0..self.buttons.len())
            .map(|_| vars.add(variable().integer().min(0)))
            .collect();

        let total_presses: Expression = presses.iter().sum();
        let mut problem = vars.minimise(total_presses).using(default_solver);

        for (jolt_i, &target) in self.joltages.iter().enumerate() {
            let mut expr = Expression::from(0.0);

            for (btn_i, needed_is) in self.buttons.iter().enumerate() {
                if needed_is.contains(&jolt_i) {
                    expr += presses[btn_i];
                }
            }

            problem.add_constraint(expr.eq(target as f64));
        }

        let solution = problem.solve().unwrap();

        presses
            .iter()
            .map(|var| solution.value(*var).round() as usize)
            .sum()
    }
}

pub fn part_one(input: &str) -> Option<u64> {
    let day = match Day10::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(
        day.machines
            .iter()
            .map(|machine| machine.min_needed_button_presses() as u64)
            .sum(),
    )
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day10::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    Some(
        day.machines
            .iter()
            .map(|machine| machine.min_needed_joltage_presses() as u64)
            .sum(),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(7));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(33));
    }
}
