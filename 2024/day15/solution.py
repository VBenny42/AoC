from enum import Enum
from typing import List, Tuple


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


class EdgeError(Exception):
    pass


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

    movements = [
        parse_movement(movement)
        for j in range(i + 1, len(lines))
        for movement in lines[j].strip()
    ]
    return warehouse_map, movements


def scale_map(map: Map) -> Map:
    new_map = []
    for row in map:
        new_row = []
        for spot in row:
            match spot:
                case "O":
                    new_row.extend(["[", "]"])
                case "@":
                    new_row.extend(["@", "."])
                case "#":
                    new_row.extend(["#", "#"])
                case ".":
                    new_row.extend([".", "."])
        new_map.append(new_row)
    return new_map


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
    if map[new_y][new_x] == "#":
        raise EdgeError("Edge of the map")
    if map[new_y][new_x] == ".":
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
        try:
            boxes = boxes_to_move(map, (new_x, new_y), direction)
        except EdgeError:
            return robot
        for box in boxes:
            assert type(box) == tuple
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


def boxes_to_move2(
    map: Map, box: Tuple[Coordinate, Coordinate], direction: Directions
) -> list:
    if direction == Directions.LEFT or direction == Directions.RIGHT:
        # same code as part 1
        x, y = box[0] if direction == Directions.LEFT else box[1]
        bracket = "]" if direction == Directions.LEFT else "["
        dx, dy = direction.value
        new_x, new_y = x + dx, y + dy
        adjacent_box = (
            ((new_x - 1, new_y), (new_x, new_y))
            if direction == Directions.LEFT
            else ((new_x, new_y), (new_x + 1, new_y))
        )
        if map[new_y][new_x] == "#":
            raise EdgeError("Edge of the map")
        if map[new_y][new_x] == ".":
            return [box]
        if map[new_y][new_x] == bracket:
            return boxes_to_move2(map, adjacent_box, direction) + [box]
    if direction == Directions.UP or direction == Directions.DOWN:
        left, right = box
        left_x, left_y = left
        right_x, right_y = right
        dx, dy = direction.value
        new_left_x, new_left_y = left_x + dx, left_y + dy
        new_right_x, new_right_y = right_x + dx, right_y + dy

        if map[new_left_y][new_left_x] == "#" or map[new_right_y][new_right_x] == "#":
            raise EdgeError("Wall")

        if map[new_left_y][new_left_x] == "." and map[new_right_y][new_right_x] == ".":
            return [box]

        # No need to check right bracket really
        if map[new_left_y][new_left_x] == "[" and map[new_right_y][new_right_x] == "]":
            return boxes_to_move2(
                map, ((new_left_x, new_left_y), (new_right_x, new_right_y)), direction
            ) + [box]

        if map[new_left_y][new_left_x] == "]" and map[new_right_y][new_right_x] == ".":
            return boxes_to_move2(
                map, ((new_left_x - 1, new_left_y), (new_left_x, new_left_y)), direction
            ) + [box]
        if map[new_right_y][new_right_x] == "[" and map[new_left_y][new_left_x] == ".":
            return boxes_to_move2(
                map,
                ((new_right_x, new_right_y), (new_right_x + 1, new_right_y)),
                direction,
            ) + [box]
        else:
            return (
                boxes_to_move2(
                    map,
                    ((new_left_x - 1, new_left_y), (new_left_x, new_left_y)),
                    direction,
                )
                + boxes_to_move2(
                    map,
                    (
                        (new_right_x, new_right_y),
                        (new_right_x + 1, new_right_y),
                    ),
                    direction,
                )
                + [box]
            )

    return []


def move_robot2(map: Map, robot: Coordinate, direction: Directions) -> Coordinate:
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
    if map[new_y][new_x] == "[" or map[new_y][new_x] == "]":
        if map[new_y][new_x] == "[":
            box = ((new_x, new_y), (new_x + 1, new_y))
        else:
            box = ((new_x - 1, new_y), (new_x, new_y))
        try:
            boxes = boxes_to_move2(map, box, direction)
        except EdgeError:
            return robot
        for box in boxes:
            box_l, box_r = box
            new_box_l, new_box_r = (
                (box_l[0] + dx, box_l[1] + dy),
                (box_r[0] + dx, box_r[1] + dy),
            )
            map[box_l[1]][box_l[0]] = "."
            map[box_r[1]][box_r[0]] = "."
            map[new_box_l[1]][new_box_l[0]] = "["
            map[new_box_r[1]][new_box_r[0]] = "]"
        map[y][x] = "."
        map[new_y][new_x] = "@"
        return (new_x, new_y)
    # Shouldn't reach here
    raise ValueError("Spot is invalid")


def find_robot(map: Map) -> Coordinate:
    for y, row in enumerate(map):
        for x, spot in enumerate(row):
            if spot == "@":
                return (x, y)
    raise ValueError("Robot not found")


def find_boxes(map: Map) -> list[Coordinate]:
    return [
        (x, y)
        for y, row in enumerate(map)
        for x, cell in enumerate(row)
        if cell == "[" or cell == "O"
    ]


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
    with open("input.txt", "r") as f:
        lines = f.readlines()
    warehouse_map, movements = parse_input(lines)
    warehouse_map_scaled = scale_map(warehouse_map)
    robot = find_robot(warehouse_map_scaled)
    for movement in movements:
        robot = move_robot2(warehouse_map_scaled, robot, movement)

    boxes = find_boxes(warehouse_map_scaled)
    coordinates = sum(map(calculate_gps_coordinate, boxes))
    print(f"LOGF: { coordinates = }")


if __name__ == "__main__":
    # main1()
    main2()
