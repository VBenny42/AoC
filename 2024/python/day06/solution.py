from enum import Enum, auto
from itertools import cycle
from multiprocessing import Pool
from time import time
from typing import DefaultDict


def timer_func(func):
    # This function shows the execution time of
    # the function object passed
    def wrap_func(*args, **kwargs):
        t1 = time()
        result = func(*args, **kwargs)
        t2 = time()
        print(f"Function {func.__name__!r} executed in {(t2-t1):.4f}s")
        return result

    return wrap_func


class Coordinate:
    __slots__ = ("x", "y")

    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y

    def __repr__(self) -> str:
        return f"Coordinate({self.x}, {self.y})"

    def __eq__(self, other) -> bool:
        return self.x == other.x and self.y == other.y

    def __hash__(self) -> int:
        return hash((self.x, self.y))


class Directions(Enum):
    UP = auto()
    RIGHT = auto()
    DOWN = auto()
    LEFT = auto()


def get_starting_position(grid: list[list[str]]) -> Coordinate:
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "^":
                return Coordinate(x, y)
    raise ValueError("No starting position found")


def get_next_position(
    grid_bounds: tuple[int, int], position: Coordinate, direction: Directions
) -> Coordinate:
    m, n = grid_bounds
    match direction:
        case Directions.UP:
            if position.y == 0:
                raise IndexError
            return Coordinate(position.x, position.y - 1)
        case Directions.LEFT:
            if position.x == 0:
                raise IndexError
            return Coordinate(position.x - 1, position.y)
        case Directions.DOWN:
            if position.y == n - 1:
                raise IndexError
            return Coordinate(position.x, position.y + 1)
        case Directions.RIGHT:
            if position.x == m - 1:
                raise IndexError
            return Coordinate(position.x + 1, position.y)


def mark_guard_path(grid: list[list[str]], position: Coordinate) -> list[list[str]]:
    directions = cycle(Directions)
    current_direction = next(directions)
    next_position = None
    grid_bounds = len(grid[0]), len(grid)
    try:
        while True:
            next_position = get_next_position(grid_bounds, position, current_direction)
            if grid[next_position.y][next_position.x] == "#":
                current_direction = next(directions)
                continue
            grid[position.y][position.x] = "X"
            position = next_position
    # IndexError is raised when we go out of the grid
    # This is the signal that we have reached the end of the path
    except IndexError:
        grid[position.y][position.x] = "X"
        return grid


def does_induce_loop(
    grid: list[list[str]], possible_obstruction: Coordinate, position: Coordinate
) -> bool:
    visited_positions = DefaultDict(set)
    directions = cycle(Directions)
    current_direction = next(directions)
    next_position = None
    grid_bounds = len(grid[0]), len(grid)
    try:
        while True:
            next_position = get_next_position(grid_bounds, position, current_direction)
            if (
                grid[next_position.y][next_position.x] == "#"
                or next_position == possible_obstruction
            ):
                current_direction = next(directions)
                continue
            if (
                position in visited_positions
                and current_direction in visited_positions[position]
            ):
                return True
            visited_positions[position].add(current_direction)
            position = next_position
    # Reached an edge of the grid, no loop found
    except IndexError:
        return False


def find_loops(grid: list[list[str]], position: Coordinate) -> int:
    return sum(
        does_induce_loop(grid, Coordinate(x, y), position)
        for y, row in enumerate(grid)
        for x, cell in enumerate(row)
        if cell == "X"
    )


def worker(task):
    grid, cell, position = task
    return does_induce_loop(grid, cell, position)


def find_loops_multiprocessing(grid: list[list[str]], position: Coordinate) -> int:
    tasks = (
        (grid, Coordinate(x, y), position)
        for y, row in enumerate(grid)
        for x, cell in enumerate(row)
        if cell == "X"
    )

    with Pool() as pool:
        results = pool.map(worker, tasks)

    return sum(results)


def print_grid(grid):
    for row in grid:
        print("".join(row))


# @timer_func
def main1():
    with open("input.txt") as f:
        grid = [list(line.strip()) for line in f]
    starting_position = get_starting_position(grid)
    marked_grid = mark_guard_path(grid, starting_position)
    print(f"ANSWER1: distinct positions = {sum(row.count('X') for row in marked_grid)}")


# @timer_func
def main2():
    with open("input.txt") as f:
        grid = [list(line.strip()) for line in f]
    starting_position = get_starting_position(grid)
    marked_grid = mark_guard_path(grid, starting_position)
    print(
        f"ANSWER2: ways to induce a loop = {find_loops(marked_grid, starting_position)}"
    )


# @timer_func
def main3():
    with open("input.txt") as f:
        grid = [list(line.strip()) for line in f]
    starting_position = get_starting_position(grid)
    marked_grid = mark_guard_path(grid, starting_position)
    print(
        f"ANSWER2: ways to induce a loop = {find_loops_multiprocessing(marked_grid, starting_position)}"
    )


if __name__ == "__main__":
    main1()
    # main2()
    main3()
