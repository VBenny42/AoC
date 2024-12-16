from enum import Enum
import heapq

class Directions(Enum):
    NORTH = (0, -1)
    SOUTH = (0, 1)
    WEST = (-1, 0)
    EAST = (1, 0)

    def turn_clockwise(self):
        directions = [Directions.NORTH, Directions.EAST, Directions.SOUTH, Directions.WEST]
        return directions[(directions.index(self) + 1) % 4]

    def turn_counterclockwise(self):
        directions = [Directions.NORTH, Directions.EAST, Directions.SOUTH, Directions.WEST]
        return directions[(directions.index(self) - 1) % 4]

    def __lt__(self, other):
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

def heuristic(state, goal):
    x, y, _ = state
    goal_x, goal_y = goal
    return abs(x - goal_x) + abs(y - goal_y)

def is_valid_move(grid, x, y):
    return 0 <= y < len(grid) and 0 <= x < len(grid[0]) and grid[y][x] != "#"

def aStar(grid, start_state, goal_state) -> tuple:
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
            heapq.heappush(fringe, (cost + 1 + heuristic(new_state, goal_state), new_state, cost + 1, path + [new_state]))

        # Turn clockwise
        new_direction = direction.turn_clockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(fringe, (cost + 1000 + heuristic(new_state, goal_state), new_state, cost + 1000, path + [new_state]))

        # Turn counterclockwise
        new_direction = direction.turn_counterclockwise()
        new_state = (x, y, new_direction)
        heapq.heappush(fringe, (cost + 1000 + heuristic(new_state, goal_state), new_state, cost + 1000, path + [new_state]))
    raise ValueError("No path found")

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

def main():
    with open("input.txt", "r", encoding="utf-8") as f:
        grid = [list(line.strip()) for line in f.readlines()]

    # print_grid(grid)

    start_state = find_start(grid)
    goal_state = find_end(grid)

    cost, path = aStar(grid, start_state, goal_state)

    print(f"Lowest cost: {cost}")
    # mark_path(grid, path)
    # print_grid(grid)
    print(f"{len(path)} steps")

if __name__ == "__main__":
    main()
