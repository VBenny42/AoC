from enum import Enum, auto
import heapq
from typing import NamedTuple
from collections.abc import Callable


class Direction(Enum):
    UP = auto()
    RIGHT = auto()
    DOWN = auto()
    LEFT = auto()

    def __hash__(self) -> int:
        return hash(self.value)

    def __lt__(self, other) -> bool:
        return self.value < other.value


State = NamedTuple("State", [("x", int), ("y", int), ("direction", Direction)])
Grid = list[list[str]]


class Dijkstra:
    def __init__(
        self,
        neighbors_fn: Callable[[State], list[State]],
        cost_fn: Callable[[State, State], int],
        min_cost: float,
        max_cost: float,
    ):
        self.cost_function = cost_fn
        self.neighbors_function = neighbors_fn
        self.previous = {}
        self.costs = {}
        self.min_cost = min_cost
        self.max_cost = max_cost

    def find_path(self, start: State):
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

    def get_cost(self, end: State) -> float:
        if end not in self.costs:
            return self.max_cost

        return self.costs[end]

    def get_paths(self, end: State) -> list[State]:
        path = []
        stack = [end]

        while stack:
            current = stack.pop()
            path.append(current)
            for previous in self.previous[current]:
                stack.append(State(*previous))

        return path


def neighbors_fn(
    cell: State,
    grid: Grid,
    width: int,
    height: int,
) -> list[State]:
    neighbors = []
    for direction in Direction:
        if direction == cell.direction:
            continue
        neighbors.append(State(cell.x, cell.y, direction))
    match cell.direction:
        case Direction.UP:
            if cell.y > 0 and grid[cell.y - 1][cell.x] != "#":
                neighbors.append(State(cell.x, cell.y - 1, cell[2]))
        case Direction.RIGHT:
            if cell.x < width - 1 and grid[cell.y][cell.x + 1] != "#":
                neighbors.append(State(cell.x + 1, cell.y, cell[2]))
        case Direction.DOWN:
            if cell.y < height - 1 and grid[cell.y + 1][cell.x] != "#":
                neighbors.append(State(cell.x, cell.y + 1, cell[2]))
        case Direction.LEFT:
            if cell.x > 0 and grid[cell.y][cell.x - 1] != "#":
                neighbors.append(State(cell.x - 1, cell.y, cell[2]))
    return neighbors


def cost_fn(cell1: State, cell2: State) -> int:
    if cell1.x == cell2.x and cell1.y == cell2.y:
        if cell1.direction == cell2.direction:
            return 0
        if cell1.direction in (Direction.UP, Direction.DOWN):
            if cell2.direction in (Direction.UP, Direction.DOWN):
                return 2000
            return 1000
        if cell2.direction in (Direction.UP, Direction.DOWN):
            return 1000
        return 2000
    return 1


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f]

    height, width = len(grid), len(grid[0])

    start: State | None = None
    ends: list[State] | None = None

    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "S":
                start = State(x, y, Direction.RIGHT)
            if cell == "E":
                ends = [
                    State(x, y, Direction.UP),
                    State(x, y, Direction.RIGHT),
                    State(x, y, Direction.DOWN),
                    State(x, y, Direction.LEFT),
                ]

    assert start is not None
    assert ends is not None

    dijkstra = Dijkstra(
        lambda cell: neighbors_fn(cell, grid, width, height),
        cost_fn,
        0.0,
        float("inf"),
    )

    min_cost = float("inf")
    min_end = None
    dijkstra.find_path(start)
    for end in ends:
        min_cost = min(min_cost, dijkstra.get_cost(end))
        min_end = end

    assert min_end is not None
    print(f"LOG: { min_end = }")

    print(f"ANSWER1: Least cost path { int(min_cost) }")


def main2():
    with open("input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f]

    height, width = len(grid), len(grid[0])

    start: State | None = None
    ends: list[State] | None = None

    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "S":
                start = State(x, y, Direction.RIGHT)
            if cell == "E":
                ends = [
                    State(x, y, Direction.UP),
                    State(x, y, Direction.RIGHT),
                    State(x, y, Direction.DOWN),
                    State(x, y, Direction.LEFT),
                ]

    assert start is not None
    assert ends is not None

    dijkstra = Dijkstra(
        lambda cell: neighbors_fn(cell, grid, width, height),
        cost_fn,
        0.0,
        float("inf"),
    )

    min_cost = float("inf")
    end_state: State | None = None
    dijkstra.find_path(start)
    for end in ends:
        cost = dijkstra.get_cost(end)
        if cost < min_cost:
            min_cost = cost
            end_state = State(*end)

    assert end_state is not None

    tiles_on_all_paths = set()
    for node in dijkstra.get_paths(end_state):
        tiles_on_all_paths.add((node.x, node.y))

    print(f"ANSWER2: Tiles on all paths {len(tiles_on_all_paths)}")


if __name__ == "__main__":
    main1()
    main2()
