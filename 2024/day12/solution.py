from typing import DefaultDict, List, Set
from itertools import chain
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
                        not_neighbors_count[position[1]][position[0]].add(
                            direction
                        )
            regions[grid[i][j]].append(region)
    return regions, not_neighbors_count


def calculate_perimeter(region: List[Coordinate], grid: Grid) -> int:
    perimeter = 0
    m, n = len(grid[0]), len(grid)
    region_set = set(region)

    for x, y in region:
        for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            nx, ny = x + dx, y + dy

            # Check if neighbor is out of bounds (outer boundary)
            if nx < 0 or nx >= m or ny < 0 or ny >= n:
                perimeter += 1
            # Check if neighbor is not part of the same region (internal boundary)
            elif (nx, ny) not in region_set:
                perimeter += 1

    return perimeter


def print_grid(grid: Grid) -> None:
    for row in grid:
        print("".join(row))


def mark_perimeter(
    grid_with_adjacent_spaces: Grid, region: List[Coordinate], grid: Grid
) -> None:
    region_set = set(region)
    for x, y in region:
        # Calculate corresponding coordinates in grid_with_adjacent_spaces
        adj_x, adj_y = 2 * x + 1, 2 * y + 1

        # Check all four directions
        for dx, dy, marker in [(-1, 0, "|"), (1, 0, "|"), (0, -1, "-"), (0, 1, "-")]:
            nx, ny = x + dx, y + dy
            adj_nx, adj_ny = adj_x + dx, adj_y + dy

            # Check if the neighboring cell is out of bounds or not part of the region
            if (
                nx < 0
                or nx >= len(grid[0])
                or ny < 0
                or ny >= len(grid)
                or (nx, ny) not in region_set
            ):
                grid_with_adjacent_spaces[adj_ny][adj_nx] = marker


def count_perimeter(grid_with_adjacent_spaces: Grid, region: List[Coordinate]) -> int:
    perimeter = 0
    seen = set()
    for x, y in region:
        seen.add((x, y))
        for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            nx, ny = x + dx, y + dy
            if (nx, ny) in seen:
                continue
            if grid_with_adjacent_spaces[2 * y + dy + 1][2 * x + dx + 1] in "-|":
                perimeter += 1
    return perimeter


def add_corners(grid_with_adjacent_spaces: Grid) -> None:
    rows = len(grid_with_adjacent_spaces)
    cols = len(grid_with_adjacent_spaces[0])

    for y in range(rows):
        for x in range(cols):
            # Check if this position should be a corner
            if grid_with_adjacent_spaces[y][x] == " ":
                if (y > 0 and grid_with_adjacent_spaces[y - 1][x] == "|") and (
                    x > 0 and grid_with_adjacent_spaces[y][x - 1] == "-"
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                elif (y > 0 and grid_with_adjacent_spaces[y - 1][x] == "|") and (
                    x < cols - 1 and grid_with_adjacent_spaces[y][x + 1] == "-"
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                elif (y < rows - 1 and grid_with_adjacent_spaces[y + 1][x] == "|") and (
                    x > 0 and grid_with_adjacent_spaces[y][x - 1] == "-"
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                elif (y < rows - 1 and grid_with_adjacent_spaces[y + 1][x] == "|") and (
                    x < cols - 1 and grid_with_adjacent_spaces[y][x + 1] == "-"
                ):
                    grid_with_adjacent_spaces[y][x] = "+"


def count_perimeter_neighbors(
     region: List[Coordinate], not_neighbors_count
) -> int:
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
