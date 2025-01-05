from typing import List
from enum import Enum, auto


class Quadrants(Enum):
    TOP_LEFT = auto()
    TOP_RIGHT = auto()
    BOTTOM_LEFT = auto()
    BOTTOM_RIGHT = auto()


def parse_robot(line: List[str]):
    position = list(map(int, line[0][2:].split(",")))
    velocity = tuple(map(int, line[1][2:].split(",")))
    return [position, velocity]


def print_room(room):
    for row in room:
        print("".join(str(cell) if cell > 0 else "." for cell in row))
    print()


def initialize_robot(room, robot):
    x, y = robot[0]
    room[y][x] += 1


def print_bitmap(room, rows, cols, filename):
    with open(filename, "w") as f:
        f.write(f"P1\n{cols} {rows}\n")
        for row in range(rows):
            for col in range(cols):
                f.write(str(room[row][col]))
            f.write("\n")


def invert_room(room):
    for row in room:
        for i, cell in enumerate(row):
            row[i] = 1 if cell == 0 else 0


def print_scaled_bitmap(room, rows, cols, filename, scale_factor=5):
    with open(filename, "w") as f:
        f.write(f"P1\n{cols*scale_factor} {rows*scale_factor}\n")
        for row in range(rows):
            for _ in range(scale_factor):
                for col in range(cols):
                    for _ in range(scale_factor):
                        f.write(str(room[row][col]))
                f.write("\n")


def move_robot(robot, room, rows, cols):
    v_x, v_y = robot[1]
    p_x, p_y = robot[0]
    new_x = (p_x + v_x) % cols
    new_y = (p_y + v_y) % rows
    room[p_y][p_x] -= 1
    room[new_y][new_x] += 1
    robot[0] = [new_x, new_y]


def get_quadrant(room, quadrant):
    rows, cols = len(room), len(room[0])
    match quadrant:
        case quadrant.TOP_LEFT:
            return [[room[y][x] for x in range(cols // 2)] for y in range(rows // 2)]
        case quadrant.TOP_RIGHT:
            return [
                [room[y][x] for x in range(cols // 2 + 1, cols)]
                for y in range(rows // 2)
            ]
        case quadrant.BOTTOM_LEFT:
            return [
                [room[y][x] for x in range(cols // 2)]
                for y in range(rows // 2 + 1, rows)
            ]
        case quadrant.BOTTOM_RIGHT:
            return [
                [room[y][x] for x in range(cols // 2 + 1, cols)]
                for y in range(rows // 2 + 1, rows)
            ]


def main1():
    with open("input.txt", "r") as f:
        robots = [parse_robot(line.strip().split()) for line in f]
    rows = 103
    cols = 101
    room = [[0 for _ in range(cols)] for _ in range(rows)]
    for robot in robots:
        initialize_robot(room, robot)
        for _ in range(100):
            move_robot(robot, room, rows, cols)
    safety_factor = 1
    for quadrant in Quadrants:
        room_quadrant = get_quadrant(room, quadrant)
        num_robots = sum(sum(row) for row in room_quadrant)
        safety_factor *= num_robots
    print(f"ANSWER: { safety_factor = }")


def main2():
    with open("input.txt", "r") as f:
        robots = [parse_robot(line.strip().split()) for line in f]
    rows = 103
    cols = 101
    room = [[0 for _ in range(cols)] for _ in range(rows)]
    for robot in robots:
        initialize_robot(room, robot)
    for i in range(10000):
        for robot in robots:
            move_robot(robot, room, rows, cols)
        # My solution was at 7753 seconds
        # print_bitmap(room, rows, cols, f"pics/output_{i+1}.pbm")
        if i == 7752:
            # invert_room(room)
            print_scaled_bitmap(room, rows, cols, f"output_{i+1}_scaled.pbm", 1)


if __name__ == "__main__":
    main1()
    main2()
