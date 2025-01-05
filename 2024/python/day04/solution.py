from enum import Enum


class Coordinate:
    __slots__ = ("x", "y")

    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y


class Directions(Enum):
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
) -> dict[Directions, tuple[str, Coordinate]]:
    x, y = coord.x, coord.y
    adjacent = {}
    if x > 0:
        adjacent[Directions.LEFT] = (grid[y][x - 1], Coordinate(x - 1, y))
    if x < len(grid[y]) - 1:
        adjacent[Directions.RIGHT] = (grid[y][x + 1], Coordinate(x + 1, y))
    if y > 0:
        adjacent[Directions.UP] = (grid[y - 1][x], Coordinate(x, y - 1))
    if y < len(grid) - 1:
        adjacent[Directions.DOWN] = (grid[y + 1][x], Coordinate(x, y + 1))
    if x > 0 and y > 0:
        adjacent[Directions.UP_LEFT] = (grid[y - 1][x - 1], Coordinate(x - 1, y - 1))
    if x < len(grid[y]) - 1 and y > 0:
        adjacent[Directions.UP_RIGHT] = (grid[y - 1][x + 1], Coordinate(x + 1, y - 1))
    if x > 0 and y < len(grid) - 1:
        adjacent[Directions.DOWN_LEFT] = (grid[y + 1][x - 1], Coordinate(x - 1, y + 1))
    if x < len(grid[y]) - 1 and y < len(grid) - 1:
        adjacent[Directions.DOWN_RIGHT] = (grid[y + 1][x + 1], Coordinate(x + 1, y + 1))
    return adjacent


def is_xmas_match(
    grid: list[list[str]],
    coord: Coordinate,
    current_match: str,
    direction: Directions,
) -> bool:
    xmas = "XMAS"

    adjacent = get_adjacent_letters(coord, grid)
    adjacent_letter = adjacent.get(direction, None)
    if adjacent_letter is None:
        return False

    letter = grid[coord.y][coord.x]

    potential_xmas = current_match + letter + adjacent_letter[0]

    if xmas.startswith(potential_xmas):
        if len(potential_xmas) == len(xmas):
            return True
        return is_xmas_match(
            grid,
            adjacent_letter[1],
            current_match + letter,
            direction,
        )
    return False


# For each letter in the grid, check if the adjacent letters can continue xmas
def xmas_matches1(grid: list[list[str]]) -> int:
    matches = 0

    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if grid[y][x] == "X":
                for direction in Directions:
                    matches += (
                        1 if is_xmas_match(grid, Coordinate(x, y), "", direction) else 0
                    )
    return matches


def is_x_mas_match(grid: list[list[str]], coord: Coordinate) -> bool:
    corners = {
        Directions.UP_LEFT,
        Directions.UP_RIGHT,
        Directions.DOWN_LEFT,
        Directions.DOWN_RIGHT,
    }
    adjacent = get_adjacent_letters(coord, grid)
    # For an X-MAS match, A must have letters in all 4 corners
    if (corners).intersection(adjacent.keys()) != corners:
        return False
    # For an X-MAS match, There must be an M and an S in each of the diagonals
    return (
        {adjacent[Directions.UP_LEFT][0], adjacent[Directions.DOWN_RIGHT][0]}
        == {"M", "S"}
    ) and (
        {adjacent[Directions.UP_RIGHT][0], adjacent[Directions.DOWN_LEFT][0]}
        == {"M", "S"}
    )


def xmas_matches2(grid: list[list[str]]) -> int:
    matches = 0
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if grid[y][x] == "A":
                match = is_x_mas_match(grid, Coordinate(x, y))
                if match:
                    matches += 1
    return matches


def main1():
    with open("input.txt") as f:
        lines = f.readlines()
    grid = [list(line.strip()) for line in lines]
    print("ANSWER1: xmas matches", xmas_matches1(grid))


def main2():
    with open("input.txt") as f:
        lines = f.readlines()
    grid = [list(line.strip()) for line in lines]
    print("ANSWER2: x-mas matches", xmas_matches2(grid))


if __name__ == "__main__":
    main1()
    main2()
