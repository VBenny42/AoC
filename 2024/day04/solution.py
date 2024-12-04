from collections import Counter
from enum import Enum
from pprint import pprint


class Coordinate:
    __slots__ = ("x", "y")


class Direction(Enum):
    UP = 1
    DOWN = 2
    LEFT = 3
    RIGHT = 4
    UP_LEFT = 5
    UP_RIGHT = 6
    DOWN_LEFT = 7
    DOWN_RIGHT = 8


def get_adjacent_letters(
    coord: Coordinate, grid: list[list[str]]
) -> dict[Direction, tuple[str, Coordinate]]:
    x, y = coord.x, coord.y
    adjacent = {}
    if x > 0:
        c = Coordinate()
        c.x = x - 1
        c.y = y
        adjacent[Direction.LEFT] = (grid[y][x - 1], c)
    if x < len(grid[y]) - 1:
        c = Coordinate()
        c.x = x + 1
        c.y = y
        adjacent[Direction.RIGHT] = (grid[y][x + 1], c)
    if y > 0:
        c = Coordinate()
        c.x = x
        c.y = y - 1
        adjacent[Direction.UP] = (grid[y - 1][x], c)
    if y < len(grid) - 1:
        c = Coordinate()
        c.x = x
        c.y = y + 1
        adjacent[Direction.DOWN] = (grid[y + 1][x], c)
    if x > 0 and y > 0:
        c = Coordinate()
        c.x = x - 1
        c.y = y - 1
        adjacent[Direction.UP_LEFT] = (grid[y - 1][x - 1], c)
    if x < len(grid[y]) - 1 and y > 0:
        c = Coordinate()
        c.x = x + 1
        c.y = y - 1
        adjacent[Direction.UP_RIGHT] = (grid[y - 1][x + 1], c)
    if x > 0 and y < len(grid) - 1:
        c = Coordinate()
        c.x = x - 1
        c.y = y + 1
        adjacent[Direction.DOWN_LEFT] = (grid[y + 1][x - 1], c)
    if x < len(grid[y]) - 1 and y < len(grid) - 1:
        c = Coordinate()
        c.x = x + 1
        c.y = y + 1
        adjacent[Direction.DOWN_RIGHT] = (grid[y + 1][x + 1], c)
    return adjacent


def is_xmas_match(
    grid: list[list[str]],
    coord: Coordinate,
    current_match: list[Coordinate],
    current_direction: Direction,
) -> list[Coordinate]:
    xmas = "XMAS"
    if len(current_match) == len(xmas):
        return current_match

    adjacent = get_adjacent_letters(coord, grid)
    adjacent_letter = adjacent.get(current_direction, None)
    if adjacent_letter is None:
        return []

    potential_xmas = (
        "".join(grid[c.y][c.x] for c in current_match)
        + grid[coord.y][coord.x]
        + adjacent_letter[0]
    )

    if xmas.startswith(potential_xmas):
        if len(potential_xmas) == len(xmas):
            return current_match + [coord, adjacent_letter[1]]
        return is_xmas_match(
            grid,
            adjacent_letter[1],
            current_match + [coord],
            current_direction,
        )
    return []


# For each letter in the grid, check if the adjacent letters can continue xmas
def xmas_matches1(grid: list[list[str]]) -> int:
    matches = set()

    for y in range(len(grid)):
        for x in range(len(grid[y])):
            for direction in Direction:
                c = Coordinate()
                c.x = x
                c.y = y
                match = is_xmas_match(grid, c, [], direction)
                if match:
                    matches.add(tuple(match))

    return len(matches)


def is_x_mas_match(grid: list[list[str]], coord: Coordinate) -> bool:
    corners = {
        Direction.UP_LEFT,
        Direction.UP_RIGHT,
        Direction.DOWN_LEFT,
        Direction.DOWN_RIGHT,
    }
    adjacent = get_adjacent_letters(coord, grid)
    # For an X-MAS match, A must have letters in all 4 corners
    if (corners).intersection(adjacent.keys()) != corners:
        return False
    # For an X-MAS match, There must be an M and an S in each of the diagonals
    return (
        {adjacent[Direction.UP_LEFT][0], adjacent[Direction.DOWN_RIGHT][0]}
        == {"M", "S"}
    ) and (
        {adjacent[Direction.UP_RIGHT][0], adjacent[Direction.DOWN_LEFT][0]}
        == {"M", "S"}
    )


def xmas_matches2(grid: list[list[str]]) -> int:
    matches = 0
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if grid[y][x] == "A":
                c = Coordinate()
                c.x = x
                c.y = y
                match = is_x_mas_match(grid, c)
                if match:
                    matches += 1
    return matches


def main1():
    with open("input.txt") as f:
        lines = f.readlines()
    grid = [list(line.strip()) for line in lines]
    print(xmas_matches1(grid))


def main2():
    with open("input.txt") as f:
        lines = f.readlines()
    grid = [list(line.strip()) for line in lines]
    print(xmas_matches2(grid))


if __name__ == "__main__":
    main1()
    main2()
