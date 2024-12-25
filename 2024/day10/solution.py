from enum import Enum, auto
from typing import Dict

Coordinate = tuple[int, int]
Grid = list[list[int]]


class Directions(Enum):
    UP = auto()
    RIGHT = auto()
    DOWN = auto()
    LEFT = auto()


class CoordinateError(IndexError):
    pass


def get_next_position(
    grid: Grid, position: Coordinate, direction: Directions
) -> Coordinate:
    # NOTE: See if it makes a difference calling m,n = len(grid[0]), len(grid) here
    # or outside the function
    m, n = len(grid[0]), len(grid)
    value = grid[position[1]][position[0]]
    next_position = None
    match direction:
        case Directions.UP:
            if position[1] == 0:
                raise CoordinateError
            next_position = (position[0], position[1] - 1)
        case Directions.LEFT:
            if position[0] == 0:
                raise CoordinateError
            next_position = (position[0] - 1, position[1])
        case Directions.DOWN:
            if position[1] == n - 1:
                raise CoordinateError
            next_position = (position[0], position[1] + 1)
        case Directions.RIGHT:
            if position[0] == m - 1:
                raise CoordinateError
            next_position = (position[0] + 1, position[1])
    if value - grid[next_position[1]][next_position[0]] != 1:
        raise CoordinateError
    return next_position


def find_paths_to_zero_one(
    position: Coordinate,
    grid: Grid,
    trailheads: Dict[Coordinate, int],
    visited: set[Coordinate],
) -> None:
    if position in visited:
        return
    visited.add(position)
    if position in trailheads:
        trailheads[position] += 1
        return
    for direction in Directions:
        try:
            next_position = get_next_position(grid, position, direction)
            find_paths_to_zero_one(next_position, grid, trailheads, visited)
        except CoordinateError:
            pass
    return


def find_paths_to_zero_all(
    position: Coordinate,
    grid: Grid,
    trailheads: Dict[Coordinate, int],
) -> None:
    if position in trailheads:
        trailheads[position] += 1
        return
    for direction in Directions:
        try:
            next_position = get_next_position(grid, position, direction)
            find_paths_to_zero_all(next_position, grid, trailheads)
        except CoordinateError:
            pass
    return


def print_grid(grid):
    for row in grid:
        print("".join(str(cell) for cell in row))


def main1():
    with open("input.txt") as f:
        grid = (list(line.strip()) for line in f)
        grid = [list(map(int, line)) for line in grid]
    trailheads: dict[Coordinate, int] = {
        (x, y): 0
        for y, row in enumerate(grid)
        for x, cell in enumerate(row)
        if cell == 0
    }
    nine_positions: list[Coordinate] = [
        (x, y) for y, row in enumerate(grid) for x, cell in enumerate(row) if cell == 9
    ]
    for position in nine_positions:
        find_paths_to_zero_one(position, grid, trailheads, set())
    print(f"ANSWER1: score = { sum(trailheads.values()) }")


def main2():
    with open("input.txt") as f:
        grid = (list(line.strip()) for line in f)
        grid = [list(map(int, line)) for line in grid]
    trailheads: dict[Coordinate, int] = {
        (x, y): 0
        for y, row in enumerate(grid)
        for x, cell in enumerate(row)
        if cell == 0
    }
    nine_positions: list[Coordinate] = [
        (x, y) for y, row in enumerate(grid) for x, cell in enumerate(row) if cell == 9
    ]
    for position in nine_positions:
        find_paths_to_zero_all(position, grid, trailheads)
    print(f"ANSWER2: score = { sum(trailheads.values()) }")


if __name__ == "__main__":
    main1()
    main2()
