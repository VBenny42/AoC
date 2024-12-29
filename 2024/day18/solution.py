import heapq
from collections.abc import Callable


Grid = list[list[str]]
Coordinate = tuple[int, int]


class Dijkstra:
    def __init__(
        self,
        neighbors_fn: Callable[[Coordinate], list[Coordinate]],
        cost_fn: Callable[[Coordinate, Coordinate], float],
        min_cost: float,
        max_cost: float,
    ):
        self.cost_function = cost_fn
        self.neighbors_function = neighbors_fn
        self.previous = {}
        self.costs = {}
        self.min_cost = min_cost
        self.max_cost = max_cost

    def find_path(self, start: Coordinate):
        queue = []
        queue.append([0, start])
        self.previous = {}
        self.costs = {}
        self.costs[start] = self.min_cost
        self.previous[start] = []

        while queue:
            _, current = heapq.heappop(queue)
            for neighbor in self.neighbors_function(current):
                new_cost = self.costs[current] + self.cost_function(current, neighbor)

                if neighbor not in self.costs or new_cost < self.costs[neighbor]:
                    self.costs[neighbor] = new_cost
                    heapq.heappush(queue, [new_cost, neighbor])
                    self.previous[neighbor] = [current]

                elif new_cost == self.costs[neighbor]:
                    self.previous[neighbor].append(current)

    def get_cost(self, end: Coordinate) -> float:
        if end not in self.costs:
            return self.max_cost

        return self.costs[end]

    def get_paths(self, end: Coordinate) -> list[Coordinate]:
        path = []
        stack = [end]

        while stack:
            current = stack.pop()
            path.append(current)
            for previous in self.previous[current]:
                stack.append((previous))

        return path


def adjacent_neighbors_fn(
    cell: Coordinate,
    grid: Grid,
    width: int,
    height: int,
) -> list[Coordinate]:
    neighbors = []
    if cell[0] > 0 and grid[cell[1]][cell[0] - 1] == ".":
        neighbors.append((cell[0] - 1, cell[1]))
    if cell[0] < width - 1 and grid[cell[1]][cell[0] + 1] == ".":
        neighbors.append((cell[0] + 1, cell[1]))
    if cell[1] > 0 and grid[cell[1] - 1][cell[0]] == ".":
        neighbors.append((cell[0], cell[1] - 1))
    if cell[1] < height - 1 and grid[cell[1] + 1][cell[0]] == ".":
        neighbors.append((cell[0], cell[1] + 1))
    return neighbors


def neutral_cost_fn(cell1: Coordinate, cell2: Coordinate) -> int:
    return 1


def print_grid(grid: Grid) -> None:
    for row in grid:
        print("".join(row))
    print()


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        falling_bytes: list[Coordinate] = [
            (int(x), int(y)) for line in f for x, y in [line.strip().split(",")]
        ]

    width, height = 71, 71

    grid = [["." for _ in range(width)] for _ in range(height)]

    for x, y in falling_bytes[:1024]:
        grid[y][x] = "#"

    # print_grid(grid)

    start: Coordinate = (0, 0)
    end: Coordinate = (width - 1, height - 1)

    dijkstra = Dijkstra(
        lambda cell: adjacent_neighbors_fn(cell, grid, width, height),
        neutral_cost_fn,
        0.0,
        float("inf"),
    )

    min_cost = float("inf")
    dijkstra.find_path(start)
    min_cost = dijkstra.get_cost(end)

    print(f"ANSWER1: Least cost path { int(min_cost) }")


def main2():
    with open("input.txt", "r", encoding="utf-8") as f:
        falling_bytes: list[Coordinate] = [
            (int(x), int(y)) for line in f for x, y in [line.strip().split(",")]
        ]

    width, height = 71, 71

    grid = [["." for _ in range(width)] for _ in range(height)]

    # My value was near 2900
    index = 2913
    for x, y in falling_bytes[:index]:
        grid[y][x] = "#"

    # print_grid(grid)

    start: Coordinate = (0, 0)
    end: Coordinate = (width - 1, height - 1)

    i = index
    while i < len(falling_bytes):
        falling_byte = falling_bytes[i]
        grid[falling_byte[1]][falling_byte[0]] = "#"

        dijkstra = Dijkstra(
            lambda cell: adjacent_neighbors_fn(cell, grid, width, height),
            neutral_cost_fn,
            0.0,
            float("inf"),
        )

        min_cost = float("inf")
        dijkstra.find_path(start)
        min_cost = dijkstra.get_cost(end)
        if min_cost == float("inf"):
            break

        i += 1

    print(f"ANSWER2: Byte that breaks the path { falling_bytes[i] } at index { i }")


if __name__ == "__main__":
    main1()
    main2()
