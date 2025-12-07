use image::ImageBuffer;

pub mod template;

// Use this file to add helper functions and additional modules.

pub type Grid<T> = Vec<Vec<T>>;

pub trait GridExt<T> {
    fn in_bounds(&self, point: &Point) -> bool;
    fn get(&self, point: &Point) -> Option<&T>;
    fn set(&mut self, point: &Point, value: T) -> Result<(), String>;

    fn get_neighbors(&self, point: &Point, directions: &[Point]) -> Vec<(&T, Point)> {
        let mut neighbors = Vec::new();
        for dir in directions.iter() {
            let neighbor_point = *point + *dir;
            if let Some(value) = self.get(&neighbor_point) {
                neighbors.push((value, neighbor_point));
            }
        }
        neighbors
    }
    fn get_cardinal_neighbors(&self, point: &Point) -> Vec<(&T, Point)> {
        self.get_neighbors(point, &DIRECTIONS_CARDINAL)
    }
    fn get_all_neighbors(&self, point: &Point) -> Vec<(&T, Point)> {
        self.get_neighbors(point, &DIRECTIONS_ALL)
    }

    fn to_image<F>(&self, f: F) -> ImageBuffer<image::Rgb<u8>, Vec<u8>>
    where
        F: Fn(&T) -> [u8; 3];
}

pub fn print_grid<T: std::fmt::Display>(grid: &Grid<T>) {
    for row in grid {
        for item in row {
            print!("{}", item);
        }
        println!();
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash)]
pub struct Point {
    pub x: isize,
    pub y: isize,
}

impl Point {
    pub fn new(x: isize, y: isize) -> Self {
        Point { x, y }
    }
}

impl std::ops::Add for Point {
    type Output = Point;

    fn add(self, other: Point) -> Point {
        Point {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}
impl std::ops::Sub for Point {
    type Output = Point;

    fn sub(self, other: Point) -> Point {
        Point {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

pub const UP: Point = Point { x: 0, y: -1 };
pub const DOWN: Point = Point { x: 0, y: 1 };
pub const LEFT: Point = Point { x: -1, y: 0 };
pub const RIGHT: Point = Point { x: 1, y: 0 };
pub const UP_LEFT: Point = Point { x: -1, y: -1 };
pub const UP_RIGHT: Point = Point { x: 1, y: -1 };
pub const DOWN_LEFT: Point = Point { x: -1, y: 1 };
pub const DOWN_RIGHT: Point = Point { x: 1, y: 1 };

pub const DIRECTIONS_CARDINAL: [Point; 4] = [UP, DOWN, LEFT, RIGHT];
pub const DIRECTIONS_ALL: [Point; 8] = [
    UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT,
];

impl<T> GridExt<T> for Grid<T> {
    fn in_bounds(&self, point: &Point) -> bool {
        point.x >= 0
            && point.y >= 0
            && (point.y as usize) < self.len()
            && (point.x as usize) < self[point.y as usize].len()
    }

    fn get(&self, point: &Point) -> Option<&T> {
        if self.in_bounds(point) {
            Some(&self[point.y as usize][point.x as usize])
        } else {
            None
        }
    }

    fn set(&mut self, point: &Point, value: T) -> Result<(), String> {
        if self.in_bounds(point) {
            self[point.y as usize][point.x as usize] = value;
            Ok(())
        } else {
            Err(format!("Point {:?} out of bounds", point))
        }
    }

    fn to_image<F>(&self, f: F) -> ImageBuffer<image::Rgb<u8>, Vec<u8>>
    where
        F: Fn(&T) -> [u8; 3],
    {
        let height = self.len() as u32;
        let width = if height > 0 { self[0].len() as u32 } else { 0 };
        let mut img = ImageBuffer::new(width, height);
        for (y, row) in self.iter().enumerate() {
            for (x, item) in row.iter().enumerate() {
                let pixel = image::Rgb(f(item));
                img.put_pixel(x as u32, y as u32, pixel);
            }
        }
        img
    }
}
