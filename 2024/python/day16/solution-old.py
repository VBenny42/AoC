from enum import Enum
import heapq
from typing import Set


class Directions(Enum):
    NORTH = (0, -1)
    SOUTH = (0, 1)
    WEST = (-1, 0)
    EAST = (1, 0)

    def turn_clockwise(self):
        directions = [
            Directions.NORTH,
            Directions.EAST,
            Directions.SOUTH,
            Directions.WEST,
        ]
        return directions[(directions.index(self) + 1) % 4]

    def turn_counterclockwise(self):
        directions = [
            Directions.NORTH,
            Directions.EAST,
            Directions.SOUTH,
            Directions.WEST,
        ]
        return directions[(directions.index(self) - 1) % 4]

    def __lt__(self, _):
        return 0


def print_grid(grid):
    for row in grid:
        print("".join(row))


def find_start(grid):
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "S":
                return (x, y, Directions.EAST)


def find_end(grid):
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "E":
                return (x, y)


def manhattan_heuristic(state, goal):
    x, y, _ = state
    goal_x, goal_y = goal
    return abs(x - goal_x) + abs(y - goal_y)


def is_valid_move(grid, x, y):
    return 0 <= y < len(grid) and 0 <= x < len(grid[0]) and grid[y][x] != "#"


def aStar(grid, start_state, goal_state, heuristic) -> tuple:
    visited = set()
    fringe = []
    heapq.heappush(fringe, (0, start_state, 0, []))  # (priority, state, cost, path)

    while fringe:
        _, state, cost, path = heapq.heappop(fringe)

        if (state) in visited:
            continue
        visited.add((state))

        x, y, direction = state
        if (x, y) == goal_state:
            return cost, path

        # Move forward
        dx, dy = direction.value
        new_x, new_y = x + dx, y + dy
        if is_valid_move(grid, new_x, new_y):
            new_state = (new_x, new_y, direction)
            heapq.heappush(
                fringe,
                (
                    cost + 1 + heuristic(new_state, goal_state),
                    new_state,
                    cost + 1,
                    path + [new_state],
                ),
            )

        # Turn clockwise
        new_direction = direction.turn_clockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(
            fringe,
            (
                cost + 1000 + heuristic(new_state, goal_state),
                new_state,
                cost + 1000,
                path + [new_state],
            ),
        )

        # Turn counterclockwise
        new_direction = direction.turn_counterclockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(
            fringe,
            (
                cost + 1000 + heuristic(new_state, goal_state),
                new_state,
                cost + 1000,
                path + [new_state],
            ),
        )
    raise ValueError("No path found")


def generate_all_paths(grid, start_state, goal_state, best_cost):
    paths = []  # Store all paths with their costs
    fringe = [(0, start_state, [])]  # (cost, state, path)

    while fringe:
        cost, state, path = fringe.pop()

        if cost > best_cost:
            continue

        x, y, direction = state

        # Check if we've reached the goal
        if (x, y) == goal_state:
            paths.append((cost, path + [(x, y)]))
            continue

        # Move forward
        dx, dy = direction.value
        new_x, new_y = x + dx, y + dy
        if is_valid_move(grid, new_x, new_y):
            new_state = (new_x, new_y, direction)
            fringe.append((cost + 1, new_state, path + [(x, y)]))

        # Turn clockwise
        new_direction = direction.turn_clockwise()
        new_state = (x, y, new_direction)
        fringe.append((cost + 1000, new_state, path + [(x, y)]))

        # Turn counterclockwise
        new_direction = direction.turn_counterclockwise()
        new_state = (x, y, new_direction)
        fringe.append((cost + 1000, new_state, path + [(x, y)]))

    return paths


def aStar2(grid, start_state, goal_state, heuristic=manhattan_heuristic):
    visited = set()
    fringe = []
    heapq.heappush(fringe, (0, start_state, 0, []))  # (priority, state, cost, path)

    tiles_on_optimal_path = set()
    best_cost = None

    while fringe:
        priority, state, cost, path = heapq.heappop(fringe)

        if (state[0], state[1]) in visited and (
            best_cost is not None and cost > best_cost
        ):
            continue

        visited.add((state[0], state[1]))

        x, y, direction = state
        if (x, y) == goal_state:
            if best_cost is None or cost < best_cost:
                best_cost = cost
                tiles_on_optimal_path.update(path)
            elif cost == best_cost:
                tiles_on_optimal_path.update(path)
            continue

        # Move forward
        dx, dy = direction.value
        new_x, new_y = x + dx, y + dy
        if is_valid_move(grid, new_x, new_y):
            new_state = (new_x, new_y, direction)
            heapq.heappush(
                fringe,
                (
                    cost + 1 + heuristic(new_state, goal_state),
                    new_state,
                    cost + 1,
                    path + [(state[0], state[1])],
                ),
            )

        # Turn clockwise
        new_direction = direction.turn_clockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(
            fringe,
            (
                cost + 1000 + heuristic(new_state, goal_state),
                new_state,
                cost + 1000,
                path + [(state[0], state[1])],
            ),
        )

        # Turn counterclockwise
        new_direction = direction.turn_counterclockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(
            fringe,
            (
                cost + 1000 + heuristic(new_state, goal_state),
                new_state,
                cost + 1000,
                path + [(state[0], state[1])],
            ),
        )

    return best_cost, tiles_on_optimal_path


def mark_path(grid, path):
    for x, y, direction in path:
        match direction:
            case Directions.NORTH:
                grid[y][x] = "^"
            case Directions.SOUTH:
                grid[y][x] = "v"
            case Directions.WEST:
                grid[y][x] = "<"
            case Directions.EAST:
                grid[y][x] = ">"


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f.readlines()]

    # print_grid(grid)

    start_state = find_start(grid)
    goal_state = find_end(grid)

    cost, path = aStar(grid, start_state, goal_state, manhattan_heuristic)

    print(f"Lowest cost: {cost}")
    # mark_path(grid, path)
    # print_grid(grid)
    print(f"ANSWER1: {len(path)} steps")


def part2_heuristic(grid, start_state, state, goal_state, on_optimal: Set, cost):
    start_to_state, _ = aStar(
        grid, start_state, (state[0], state[1]), manhattan_heuristic
    )
    state_to_goal, _ = aStar(grid, state, goal_state, manhattan_heuristic)

    if start_to_state + state_to_goal == cost:
        on_optimal.add((state[0], state[1]))

    return start_to_state + state_to_goal


def main2():
    with open("sample-input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f.readlines()]

    start_state = find_start(grid)
    goal_state = find_end(grid)

    # tiles = set()

    cost, _ = aStar(grid, start_state, goal_state, manhattan_heuristic)

    paths = generate_all_paths(grid, start_state, goal_state, cost)
    print(f"ANSWER2: { paths = }")

    # _, tiles = aStar2(grid, start_state, goal_state)
    #
    # for tile in tiles:
    #     grid[tile[1]][tile[0]] = "O"
    # print_grid(grid)
    #

    # print(len(tiles))


if __name__ == "__main__":
    main1()
    # main2()
