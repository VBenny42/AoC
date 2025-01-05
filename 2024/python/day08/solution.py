from collections import defaultdict
from itertools import combinations
from typing import Set

Coordinate = tuple[int, int]
Grid = list[list[str]]


def find_frequencies(grid: Grid) -> dict[str, Set[Coordinate]]:
    frequencies = defaultdict(set)
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell != ".":
                frequencies[cell].add((x, y))
    return frequencies


def in_bounds(c: Coordinate, grid: Grid) -> bool:
    return 0 <= c[0] < len(grid[0]) and 0 <= c[1] < len(grid)


def find_antinodes_1(frequency: Set[Coordinate], grid: Grid) -> set[Coordinate]:
    antinodes = set()
    for f1, f2 in combinations(frequency, 2):
        diff_x, diff_y = f2[0] - f1[0], f2[1] - f1[1]
        pos_x, pos_y = f1[0] - diff_x, f1[1] - diff_y
        neg_x, neg_y = f2[0] + diff_x, f2[1] + diff_y
        if in_bounds((pos_x, pos_y), grid):
            antinodes.add((pos_x, pos_y))
        if in_bounds((neg_x, neg_y), grid):
            antinodes.add((neg_x, neg_y))
    return antinodes


def find_antinodes_2(frequency: Set[Coordinate], grid: Grid) -> set[Coordinate]:
    antinodes = set()
    for f1, f2 in combinations(frequency, 2):
        diff_x, diff_y = f2[0] - f1[0], f2[1] - f1[1]
        start_x, start_y = f1[0], f1[1]
        while in_bounds((start_x, start_y), grid):
            antinodes.add((start_x, start_y))
            start_x, start_y = start_x - diff_x, start_y - diff_y
        start_x, start_y = f2[0], f2[1]
        while in_bounds((start_x, start_y), grid):
            antinodes.add((start_x, start_y))
            start_x, start_y = start_x + diff_x, start_y + diff_y
    return antinodes


def print_grid(grid: Grid):
    for row in grid:
        print("".join(row))


def main1():
    with open("input.txt", "r") as f:
        grid = [list(line.strip()) for line in f]
    frequencies = find_frequencies(grid)
    antinodes = set()
    for frequency in frequencies:
        antinodes.update(find_antinodes_1(frequencies[frequency], grid))
    print(f"ANSWER1: Number of unique antinodes within the map {len(antinodes)}")


def main2():
    with open("input.txt", "r") as f:
        grid = [list(line.strip()) for line in f]
    frequencies = find_frequencies(grid)
    antinodes = set()
    for frequency in frequencies:
        antinodes.update(find_antinodes_2(frequencies[frequency], grid))
    print(f"ANSWER2: Number of unique antinodes within the map {len(antinodes)}")


if __name__ == "__main__":
    main1()
    main2()
