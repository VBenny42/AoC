from solution import *
import sys
import time


def visualize(filename: str):
    with open(filename, "r", encoding="utf-8") as f:
        lines = f.readlines()

    warehouse_grid, movements = parse_input(lines)
    warehouse_grid_scaled = scale_grid(warehouse_grid)
    warehouse_grid_scaled = warehouse_grid
    robot = find_robot(warehouse_grid_scaled)

    rows = len(warehouse_grid_scaled)
    movements_len = len(movements)

    for i, movement in enumerate(movements):
        print(f"{movement = } {i+1} of {movements_len}")
        robot = move_robot(warehouse_grid_scaled, robot, movement)
        print_grid(warehouse_grid_scaled)
        time.sleep(0.07)
        sys.stdout.write("\033[F" * (rows + 2))


if __name__ == "__main__":
    visualize("sample-input-smaller.txt")
