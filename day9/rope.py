import math
from itertools import pairwise
import numpy as np


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


def update_head(head, direction: str):
    # print(direction)
    match direction:
        case "R":
            head[1] += 1
        case "L":
            head[1] -= 1
        case "U":
            head[0] += 1
        case "D":
            head[0] -= 1
    return head


def update_tail(head, tail):
    dirset = {
        (2, 1): (1, 1),
        (1, 2): (1, 1),
        (2, 0): (1, 0),
        (2, -1): (1, -1),
        (1, -2): (1, -1),
        (0, -2): (0, -1),
        (-1, -2): (-1, -1),
        (-2, -1): (-1, -1),
        (-2, 0): (-1, 0),
        (-2, 1): (-1, 1),
        (-1, 2): (-1, 1),
        (0, 2): (0, 1),
        (2, 2): (1, 1),
        (-2, -2): (-1, -1),
        (-2, 2): (-1, 1),
        (2, -2): (1, -1),
    }
    return tail + np.array(dirset.get(tuple(head - tail), (0, 0)))


# change to 2 for p1 results
KNOT_NUM = 10

with open("input.txt", "r") as file:
    directions: list[list[str]] = [line.strip().split() for line in file]
    rope_stack = [np.array([0, 0]) for _ in range(KNOT_NUM)]
    tail_past = {tuple(rope_stack[0])}
    for dir, num in directions:
        for _ in range(int(num)):
            rope_stack[0] = update_head(rope_stack[0], dir)
            for knot, next_knot in pairwise(range(KNOT_NUM)):
                rope_stack[next_knot] = update_tail(
                    rope_stack[knot], rope_stack[next_knot]
                )
            tail_past.add(tuple(rope_stack[-1]))
    print(len(tail_past))
    # print(tail_past)

    # plt.rcParams["figure.autolayout"] = True

    # tail_list = list(tail_past)

    # x = [point[0] for point in tail_list]
    # y = [point[1] for point in tail_list]

    # plt.plot(x, y, "s")
    # plt.yticks(np.arange(min(y), max(y) + 1, 1.0))
    # plt.show()
