from typing import DefaultDict, Set
from enum import Enum, auto


class Directions(Enum):
    UP = auto()
    RIGHT = auto()
    DOWN = auto()
    LEFT = auto()


Coordinate = tuple[int, int]
Grid = list[list[str]]


class CoordinateError(IndexError):
    pass


def get_next_position(
    grid: Grid, position: Coordinate, direction: Directions
) -> Coordinate:
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
    if value != grid[next_position[1]][next_position[0]]:
        raise CoordinateError
    return next_position


def build_regions(grid: Grid):
    m, n = len(grid[0]), len(grid)
    regions = DefaultDict(list)
    visited = [[False for _ in range(m)] for _ in range(n)]
    not_neighbors_count = [[set() for _ in range(m)] for _ in range(n)]
    for i in range(n):
        for j in range(m):
            if visited[i][j]:
                continue
            region = set()
            stack = [(j, i)]
            while stack:
                position = stack.pop()
                region.add(position)
                visited[position[1]][position[0]] = True
                for direction in Directions:
                    try:
                        next_position = get_next_position(grid, position, direction)
                        if not visited[next_position[1]][next_position[0]]:
                            stack.append(next_position)
                    except CoordinateError:
                        not_neighbors_count[position[1]][position[0]].add(direction)
            regions[grid[i][j]].append(region)
    return regions, not_neighbors_count


def print_grid(grid: Grid) -> None:
    for row in grid:
        print("".join(row))


def count_perimeter_neighbors(region: Set[Coordinate], not_neighbors_count) -> int:
    perimeter = 0
    for coordinate in region:
        perimeter += len(not_neighbors_count[coordinate[1]][coordinate[0]])
    return perimeter


def main1():
    with open("input.txt", "r") as f:
        grid = [list(line.strip()) for line in f]

    regions, not_neighbors_count = build_regions(grid)
    price = 0
    for region in regions:
        for r in regions[region]:
            price += count_perimeter_neighbors(r, not_neighbors_count) * len(r)

    print(f"LOG: { price = }")


if __name__ == "__main__":
    main1()
