from typing import DefaultDict, Dict, Set, List
from itertools import chain
from enum import Enum, auto


class Directions(Enum):
    UP = auto()
    RIGHT = auto()
    DOWN = auto()
    LEFT = auto()
    UP_RIGHT = auto()
    UP_LEFT = auto()
    DOWN_RIGHT = auto()
    DOWN_LEFT = auto()


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
        case Directions.UP_LEFT:
            if position[1] == 0 or position[0] == 0:
                raise CoordinateError
            next_position = (position[0] - 1, position[1] - 1)
        case Directions.UP_RIGHT:
            if position[1] == 0 or position[0] == m - 1:
                raise CoordinateError
            next_position = (position[0] + 1, position[1] - 1)
        case Directions.DOWN_LEFT:
            if position[1] == n - 1 or position[0] == 0:
                raise CoordinateError
            next_position = (position[0] - 1, position[1] + 1)
        case Directions.DOWN_RIGHT:
            if position[1] == n - 1 or position[0] == m - 1:
                raise CoordinateError
            next_position = (position[0] + 1, position[1] + 1)
    if (
        direction
        in {
            Directions.UP,
            Directions.RIGHT,
            Directions.DOWN,
            Directions.LEFT,
        }
        and value != grid[next_position[1]][next_position[0]]
    ):
        raise CoordinateError
    return next_position


def build_regions(
    grid: Grid,
) -> tuple[dict[str, list[set[Coordinate]]], list[list[set]]]:
    m, n = len(grid[0]), len(grid)
    regions = DefaultDict(list)
    visited = [[False for _ in range(m)] for _ in range(n)]
    not_neighbors = [[set() for _ in range(m)] for _ in range(n)]
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
                for direction in (
                    Directions.UP,
                    Directions.RIGHT,
                    Directions.DOWN,
                    Directions.LEFT,
                ):
                    try:
                        next_position = get_next_position(grid, position, direction)
                        if not visited[next_position[1]][next_position[0]]:
                            stack.append(next_position)
                    except CoordinateError:
                        not_neighbors[position[1]][position[0]].add(direction)
            regions[grid[i][j]].append(region)
    return regions, not_neighbors


def mark_perimeter(
    grid_with_adjacent_spaces: Grid, region: Set[Coordinate], grid: Grid
) -> set:
    changes = set()
    for x, y in region:
        # Calculate corresponding coordinates in grid_with_adjacent_spaces
        adj_x, adj_y = 2 * x + 1, 2 * y + 1
        # Check all four directions
        for direction, dx, dy, marker in [
            (Directions.LEFT, -1, 0, "|"),
            (Directions.RIGHT, 1, 0, "|"),
            (Directions.UP, 0, -1, "-"),
            (Directions.DOWN, 0, 1, "-"),
        ]:
            nx, ny = x + dx, y + dy
            adj_nx, adj_ny = adj_x + dx, adj_y + dy
            # Check if the neighboring cell is out of bounds or not part of the region
            if (
                nx < 0
                or nx >= len(grid[0])
                or ny < 0
                or ny >= len(grid)
                or (nx, ny) not in region
            ):
                grid_with_adjacent_spaces[adj_ny][adj_nx] = marker
                changes.add((adj_nx, adj_ny, marker))
    return changes


def add_corners(grid_with_adjacent_spaces: Grid, perimeter_edges: Set) -> set:
    rows = len(grid_with_adjacent_spaces)
    cols = len(grid_with_adjacent_spaces[0])
    corners = set()

    for y in range(rows):
        for x in range(cols):
            # Check if this position should be a corner
            if grid_with_adjacent_spaces[y][x] == " ":
                if (y > 0 and (x, y - 1, "|") in perimeter_edges) and (
                    x > 0 and (x - 1, y, "-") in perimeter_edges
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                    corners.add((x, y))
                elif (y > 0 and (x, y - 1, "|") in perimeter_edges) and (
                    x < cols - 1 and (x + 1, y, "-") in perimeter_edges
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                    corners.add((x, y))
                elif (y < rows - 1 and (x, y + 1, "|") in perimeter_edges) and (
                    x > 0 and (x - 1, y, "-") in perimeter_edges
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                    corners.add((x, y))
                elif (y < rows - 1 and (x, y + 1, "|")) and (
                    x < cols - 1 and (x + 1, y, "-") in perimeter_edges
                ):
                    grid_with_adjacent_spaces[y][x] = "+"
                    corners.add((x, y))
    return corners


def print_grid(grid: Grid) -> None:
    for row in grid:
        print("".join(row))


def calculate_perimeter(region: Set[Coordinate], not_neighbors: List[List[Set]]) -> int:
    perimeter = 0
    for coordinate in region:
        perimeter += len(not_neighbors[coordinate[1]][coordinate[0]])
    return perimeter


def is_corner_of_region(c: Coordinate, direction: Directions, not_neighbors) -> bool:
    try:
        match direction:
            case Directions.UP_RIGHT:
                return (
                    Directions.UP in not_neighbors[c[1]][c[0]]
                    and Directions.RIGHT in not_neighbors[c[1]][c[0]]
                ) or (
                    # ( (c[0] + 1, c[1] - 1), Directions.DOWN_LEFT, not_neighbors)
                    Directions.DOWN in not_neighbors[c[1] - 1][c[0] + 1]
                    and Directions.LEFT in not_neighbors[c[1] - 1][c[0] + 1]
                )
            case Directions.UP_LEFT:
                return (
                    Directions.UP in not_neighbors[c[1]][c[0]]
                    and Directions.LEFT in not_neighbors[c[1]][c[0]]
                ) or (
                    # ( (c[0] - 1, c[1] - 1), Directions.DOWN_RIGHT, not_neighbors)
                    Directions.DOWN in not_neighbors[c[1] - 1][c[0] - 1]
                    and Directions.RIGHT in not_neighbors[c[1] - 1][c[0] - 1]
                )
            case Directions.DOWN_RIGHT:
                return (
                    Directions.DOWN in not_neighbors[c[1]][c[0]]
                    and Directions.RIGHT in not_neighbors[c[1]][c[0]]
                ) or (
                    # ( (c[0] + 1, c[1] + 1), Directions.UP_LEFT, not_neighbors)
                    Directions.UP in not_neighbors[c[1] + 1][c[0] + 1]
                    and Directions.LEFT in not_neighbors[c[1] + 1][c[0] + 1]
                )

            case Directions.DOWN_LEFT:
                return (
                    Directions.DOWN in not_neighbors[c[1]][c[0]]
                    and Directions.LEFT in not_neighbors[c[1]][c[0]]
                ) or (
                    # ( (c[0] - 1, c[1] + 1), Directions.UP_RIGHT, not_neighbors)
                    Directions.UP in not_neighbors[c[1] + 1][c[0] - 1]
                    and Directions.RIGHT in not_neighbors[c[1] + 1][c[0] - 1]
                )
    except IndexError:
        return True
    return False


def calculate_corners(
    region: Set[Coordinate],
    grid_with_adjacent_spaces: Grid,
    not_neighbors,
    corners_map: Set,
) -> int:
    corners = set()
    for coordinate in region:
        x, y = coordinate
        adj_x, adj_y = 2 * x + 1, 2 * y + 1
        for direction in (
            Directions.UP_RIGHT,
            Directions.UP_LEFT,
            Directions.DOWN_RIGHT,
            Directions.DOWN_LEFT,
        ):
            try:
                next_position = get_next_position(
                    grid_with_adjacent_spaces, (adj_x, adj_y), direction
                )
                if (
                    next_position[0],
                    next_position[1],
                ) in corners_map and is_corner_of_region(
                    coordinate, direction, not_neighbors
                ):
                    corners.add((coordinate, next_position))
            except CoordinateError:
                pass
    return len(corners)


def count_sides(region: Set[Coordinate]) -> int:
    side_count = 0
    for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
        visited = set()
        for plot in region:
            if plot in visited:
                continue

            x, y = plot
            # Adjacent plot also in region, no need to count
            if (x + dx, y + dy) in region:
                continue
            # Adjacent cell not in region, count the side
            side_count += 1

            for direction in (-1, 1):
                fx, fy = plot
                while (fx, fy) in region and (fx + dx, fy + dy) not in region:
                    visited.add((fx, fy))
                    fx += direction * dy
                    fy += direction * dx

    return side_count


def main1():
    with open("input.txt", "r") as f:
        grid = [list(line.strip()) for line in f]

    regions, not_neighbors = build_regions(grid)
    price = sum(
        calculate_perimeter(r, not_neighbors) * len(r)
        for region in regions
        for r in regions[region]
    )

    print(f"ANSWER1: { price = }")


def make_adjacent_grid(grid: Grid) -> Grid:
    grid_with_adjacent_spaces = [[" " for _ in range(2 * len(grid[0]) + 1)]]
    for row in grid:
        grid_with_adjacent_spaces.append(
            list(chain.from_iterable([[" ", cell] for cell in row])) + [" "]
        )
        grid_with_adjacent_spaces.append([" " for _ in range(2 * len(grid[0]) + 1)])
    return grid_with_adjacent_spaces


def main2():
    with open("sample-input2.txt", "r") as f:
        grid = [list(line.strip()) for line in f]

    regions, not_neighbors = build_regions(grid)

    price = 0
    g_a_s = make_adjacent_grid(grid)
    # NOTE: Currently does not give correct answer
    for region in regions:
        for r in regions[region]:
            perimeter_lines = mark_perimeter(g_a_s, r, grid)
            corners = add_corners(g_a_s, perimeter_lines)
            cal = calculate_corners(r, g_a_s, not_neighbors, corners)
            price += cal * len(r)
    print_grid(g_a_s)
    print(f"ANSWER2: { price = }")


def main3():
    with open("sample-input3.txt", "r") as f:
        grid = [list(line.strip()) for line in f]

    regions, _ = build_regions(grid)

    price = 0
    for region in regions:
        for r in regions[region]:
            price += count_sides(r) * len(r)
    print(f"ANSWER2: { price = }")


if __name__ == "__main__":
    main1()
    main2()
    main3()
