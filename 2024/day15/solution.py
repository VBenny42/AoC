from enum import Enum
from typing import List


class Directions(Enum):
    UP = (0, -1)
    DOWN = (0, 1)
    LEFT = (-1, 0)
    RIGHT = (1, 0)

    def __repr__(self) -> str:
        match self:
            case Directions.UP:
                return "^"
            case Directions.DOWN:
                return "v"
            case Directions.LEFT:
                return "<"
            case Directions.RIGHT:
                return ">"


Map = List[List[str]]
Coordinate = tuple[int, int]


def parse_movement(movement: str) -> Directions:
    match movement:
        case "^":
            return Directions.UP
        case "v":
            return Directions.DOWN
        case "<":
            return Directions.LEFT
        case ">":
            return Directions.RIGHT
        case _:
            raise ValueError(f"Invalid movement: {movement}")


def parse_input(lines: List[str]) -> tuple[Map, List[Directions]]:
    i = 0
    warehouse_map = []
    while lines[i] != "\n":
        warehouse_map.append(list(lines[i].strip()))
        i += 1

    movements = []
    for j in range(i + 1, len(lines)):
        for movement in lines[j].strip():
            movements.append(parse_movement(movement))
    return warehouse_map, movements


def print_map(map):
    for row in map:
        print("".join(row))
    print()


def boxes_to_move(map: Map, box: Coordinate, direction: Directions) -> list[Coordinate]:
    x, y = box
    dx, dy = direction.value
    new_x, new_y = x + dx, y + dy
    # No need to check for further boxes if there is a wall or empty space
    # Return the current box
    if map[new_y][new_x] == "#" or map[new_y][new_x] == ".":
        return [box]
    if map[new_y][new_x] == "O":
        return boxes_to_move(map, (new_x, new_y), direction) + [box]
    # Shouldn't reach here
    return []


def move_robot(map: Map, robot: Coordinate, direction: Directions) -> Coordinate:
    x, y = robot
    dx, dy = direction.value
    new_x, new_y = x + dx, y + dy
    # Can't move if there is a wall, return Early
    if map[new_y][new_x] == "#":
        return robot
    # Can move freely, update the map
    if map[new_y][new_x] == ".":
        map[y][x] = "."
        map[new_y][new_x] = "@"
        return (new_x, new_y)
    # Box is in new spot, check if box can be moved
    if map[new_y][new_x] == "O":
        boxes = boxes_to_move(map, (new_x, new_y), direction)
        for box in boxes:
            box_x, box_y = box
            new_box_x, new_box_y = box_x + dx, box_y + dy
            # Furthest box can't be moved, return Early
            if map[new_box_y][new_box_x] == "#":
                return robot
            else:
                map[box_y][box_x] = "."
                map[new_box_y][new_box_x] = "O"
        map[y][x] = "."
        map[new_y][new_x] = "@"
        return (new_x, new_y)
    # Shouldn't reach here
    raise ValueError("Spot is invalid")


def find_robot(map: Map) -> Coordinate:
    for i in range(len(map)):
        for j in range(len(map[i])):
            if map[i][j] == "@":
                return (i, j)
    raise ValueError("Robot not found")


def find_boxes(map: Map) -> list[Coordinate]:
    boxes = []
    for i in range(len(map)):
        for j in range(len(map[i])):
            if map[i][j] == "O":
                boxes.append((j, i))
    return boxes


def calculate_gps_coordinate(box: Coordinate) -> int:
    return box[1] * 100 + box[0]


def main1():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    warehouse_map, movements = parse_input(lines)
    # print_map(warehouse_map)
    robot = find_robot(warehouse_map)
    for movement in movements:
        robot = move_robot(warehouse_map, robot, movement)
    # print_map(warehouse_map)

    boxes = find_boxes(warehouse_map)
    coordinates = sum(map(calculate_gps_coordinate, boxes))
    print(f"LOGF: { coordinates = }")


def main2():
    pass


if __name__ == "__main__":
    main1()
    main2()
