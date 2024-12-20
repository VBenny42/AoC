from collections import deque

Grid = list[list[str]]
Coord = tuple[int, int]


def bfs(
    grid: Grid, start: Coord, end: Coord, can_pass_walls=False
) -> tuple[int | float, list[Coord]]:
    q = deque([(start, 0, [start])])  # (coord, distance, path)
    visited = set()

    while q:
        (x, y), distance, path = q.popleft()
        if (x, y) == end:
            return distance, path

        if (x, y) in visited:
            continue
        visited.add((x, y))

        for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            new_y, new_x = y + dy, x + dx
            if 0 <= new_y < len(grid) and 0 <= new_x < len(grid[0]):
                if grid[new_y][new_x] != "#" or can_pass_walls:
                    q.append(((new_x, new_y), distance + 1, path + [(new_x, new_y)]))

    return float("inf"), []


def find_cheats(grid, path):
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

    def in_bounds(x, y):
        return 0 <= y < len(grid) and 0 <= x < len(grid[0])

    cheats = []
    for cell in path:
        x, y = cell

        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            if in_bounds(nx, ny) and grid[ny][nx] == "#":
                nx, ny = nx + dx, ny + dy
                if (
                    in_bounds(nx, ny)
                    and grid[ny][nx] == "."
                    and path.index((nx, ny)) > path.index(cell)
                ):
                    cheats.append((cell, (nx, ny)))

    return cheats


def calculate_savings(path: list[Coord], cheat: tuple[Coord, Coord]):
    return path.index(cheat[1]) - path.index(cheat[0]) - 2


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f.readlines()]

    get_point = lambda value: [
        (x, y)
        for y in range(len(grid))
        for x in range(len(grid[0]))
        if grid[y][x] == value
    ][0]

    start = get_point("S")
    end = get_point("E")

    grid[start[1]][start[0]] = "."
    grid[end[1]][end[0]] = "."

    _, normal_path = bfs(grid, start, end)

    cheats = find_cheats(grid, normal_path)

    threshold = 100

    savings = sum(
        1 for cheat in cheats if calculate_savings(normal_path, cheat) >= threshold
    )

    print(f"ANSWER: { savings = }")


def main2():
    pass


if __name__ == "__main__":
    main1()
    main2()
