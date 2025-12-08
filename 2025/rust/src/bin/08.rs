advent_of_code::solution!(8);

use std::str::FromStr;

type JBox = [u64; 3];

struct Day08 {
    jboxes: Vec<JBox>,
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
        let jboxes = input
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
            .collect::<Result<Vec<JBox>, ParseError>>()?;

        Ok(Day08 { jboxes })
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

    let n = day.jboxes.len();

    // Create all possible edges with their distances
    let mut edges = Vec::new();
    for i in 0..n {
        for j in (i + 1)..n {
            let dist_squared = (day.jboxes[i][0] as i64 - day.jboxes[j][0] as i64).pow(2)
                + (day.jboxes[i][1] as i64 - day.jboxes[j][1] as i64).pow(2)
                + (day.jboxes[i][2] as i64 - day.jboxes[j][2] as i64).pow(2);
            let dist = (dist_squared as f64).sqrt();
            edges.push((i, j, dist));
        }
    }

    // Sort edges by distance
    edges.sort_by(|a, b| a.2.partial_cmp(&b.2).unwrap());

    // Union-Find
    let mut parent: Vec<usize> = (0..n).collect();

    fn find(parent: &mut Vec<usize>, x: usize) -> usize {
        if parent[x] != x {
            parent[x] = find(parent, parent[x]);
        }
        parent[x]
    }

    fn union(parent: &mut Vec<usize>, x: usize, y: usize) -> bool {
        let root_x = find(parent, x);
        let root_y = find(parent, y);
        if root_x == root_y {
            return false;
        }
        parent[root_x] = root_y;
        true
    }

    #[cfg(test)]
    const NUMBER_OF_TRIES: usize = 10;

    #[cfg(not(test))]
    const NUMBER_OF_TRIES: usize = 1000;

    // Try the 10 shortest edges (whether they connect or not)
    for &(i, j, _dist) in edges.iter().take(NUMBER_OF_TRIES) {
        union(&mut parent, i, j);
    }

    // Count component sizes
    let mut component_sizes = std::collections::HashMap::new();
    for i in 0..n {
        let root = find(&mut parent, i);
        *component_sizes.entry(root).or_insert(0) += 1;
    }

    // Get the three largest component sizes
    let mut sizes: Vec<u64> = component_sizes.values().copied().collect();
    sizes.sort_by(|a, b| b.cmp(a)); // Sort descending

    if sizes.len() < 3 {
        return None;
    }

    Some(sizes[0] * sizes[1] * sizes[2])
}

pub fn part_two(input: &str) -> Option<u64> {
    let day = match Day08::from_str(input) {
        Ok(day) => day,
        Err(error) => {
            eprintln!("Error parsing input: {:?}", error);
            return None;
        }
    };

    let n = day.jboxes.len();

    // Create all possible edges with their distances
    let mut edges = Vec::new();
    for i in 0..n {
        for j in (i + 1)..n {
            let dist_squared = (day.jboxes[i][0] as i64 - day.jboxes[j][0] as i64).pow(2)
                + (day.jboxes[i][1] as i64 - day.jboxes[j][1] as i64).pow(2)
                + (day.jboxes[i][2] as i64 - day.jboxes[j][2] as i64).pow(2);
            let dist = (dist_squared as f64).sqrt();
            edges.push((i, j, dist));
        }
    }

    // Sort edges by distance
    edges.sort_by(|a, b| a.2.partial_cmp(&b.2).unwrap());

    // Union-Find
    let mut parent: Vec<usize> = (0..n).collect();
    let mut num_components = n;

    fn find(parent: &mut Vec<usize>, x: usize) -> usize {
        if parent[x] != x {
            parent[x] = find(parent, parent[x]);
        }
        parent[x]
    }

    fn union(parent: &mut Vec<usize>, x: usize, y: usize) -> bool {
        let root_x = find(parent, x);
        let root_y = find(parent, y);
        if root_x == root_y {
            return false;
        }
        parent[root_x] = root_y;
        true
    }

    // Keep connecting until all are in one circuit
    let mut last_connection = None;
    for &(i, j, _dist) in &edges {
        if union(&mut parent, i, j) {
            num_components -= 1;
            last_connection = Some((i, j));

            if num_components == 1 {
                // All connected!
                break;
            }
        }
    }

    // Get the X coordinates of the last two junction boxes connected
    let (i, j) = last_connection?;
    let x1 = day.jboxes[i][0];
    let x2 = day.jboxes[j][0];

    Some(x1 * x2)
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
