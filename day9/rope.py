import math

# import matplotlib.pyplot as plt
# import numpy as np


def do_direction_p1(
    direction: list[str],
    head_pos: tuple[int, int],
    tail_pos: tuple[int, int],
):
    curr_head = head_pos
    temp_head = curr_head
    curr_tail = tail_pos
    match direction:
        case ["R", num]:
            new_head = (head_pos[0] + int(num), head_pos[1])
            while curr_head != new_head:
                temp_head = curr_head
                curr_head = (curr_head[0] + 1, curr_head[1])
                if math.dist(curr_head, curr_tail) >= 2:
                    # Essentially go to the old head
                    curr_tail = temp_head
                    tail_past.add(curr_tail)
        case ["U", num]:
            new_head = (head_pos[0], head_pos[1] + int(num))
            while curr_head != new_head:
                temp_head = curr_head
                curr_head = (curr_head[0], curr_head[1] + 1)
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = temp_head
                    tail_past.add(curr_tail)
        case ["L", num]:
            new_head = (head_pos[0] - int(num), head_pos[1])
            while curr_head != new_head:
                temp_head = curr_head
                curr_head = (curr_head[0] - 1, curr_head[1])
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = temp_head
                    tail_past.add(curr_tail)
        case ["D", num]:
            new_head = (head_pos[0], head_pos[1] - int(num))
            while curr_head != new_head:
                temp_head = curr_head
                curr_head = (curr_head[0], curr_head[1] - 1)
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = temp_head
                    tail_past.add(curr_tail)
    return (new_head, curr_tail)


def do_direction_p2(
    direction: list[str],
    head_pos: tuple[int, int],
    tail_pos: tuple[int, int],
):
    curr_head = head_pos
    curr_tail = tail_pos
    match direction:
        case ["R", num]:
            new_head = (head_pos[0] + int(num), head_pos[1])
            while curr_head != new_head:
                curr_head = (curr_head[0] + 1, curr_head[1])
                if math.dist(curr_head, curr_tail) >= 2:
                    # Essentially go to the old head
                    curr_tail = (curr_head[0] - 1, curr_head[1])
                    tail_past.add(curr_tail)
        case ["U", num]:
            new_head = (head_pos[0], head_pos[1] + int(num))
            while curr_head != new_head:
                curr_head = (curr_head[0], curr_head[1] + 1)
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = (curr_head[0], curr_head[1] - 1)
                    tail_past.add(curr_tail)
        case ["L", num]:
            new_head = (head_pos[0] - int(num), head_pos[1])
            while curr_head != new_head:
                curr_head = (curr_head[0] - 1, curr_head[1])
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = (curr_head[0] + 1, curr_head[1])
                    tail_past.add(curr_tail)
        case ["D", num]:
            new_head = (head_pos[0], head_pos[1] - int(num))
            while curr_head != new_head:
                curr_head = (curr_head[0], curr_head[1] - 1)
                if math.dist(curr_head, curr_tail) >= 2:
                    curr_tail = (curr_head[0], curr_head[1] + 1)
                    tail_past.add(curr_tail)
    return (new_head, curr_tail)


with open("input.txt", "r", encoding="utf-8") as file:
    directions: list[list[str]] = [line.strip().split() for line in file]
    # use dictionary with tuples for checking, add tail pos to dict if it's not there
    head_pos = (0, 0)
    tail_pos = head_pos
    tail_past = {tail_pos}
    for direction in directions:
        head_pos, tail_pos = do_direction_p1(direction, head_pos, tail_pos)
    print(len(tail_past))

    # plt.rcParams["figure.autolayout"] = True

    # tail_list = list(tail_past)

    # x = [point[0] for point in tail_list]
    # y = [point[1] for point in tail_list]

    # plt.plot(x, y, "s")
    # plt.yticks(np.arange(min(y), max(y) + 1, 1.0))
    # plt.show()
