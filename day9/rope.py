import math
from itertools import pairwise

import matplotlib.pyplot as plt
import numpy as np

KNOT_NUM = 10


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


# TODO: try with list instead of tuple for rope_stack
def do_direction_p2(direction: list[str], rope_stack: list[tuple[int, int]]):
    curr_head = tuple(rope_stack[0])
    temp = curr_head
    curr_tail = rope_stack[-1]
    match direction:
        case ["R", num]:
            new_head = (rope_stack[0][0] + int(num), rope_stack[0][1])
            print(new_head)
            while rope_stack[0] != new_head:
                print(rope_stack)
                temp = rope_stack[0]
                rope_stack[0] = (rope_stack[0][0] + 1, rope_stack[0][1])
                for knot, next_knot in pairwise(range(KNOT_NUM)):
                    if math.dist(rope_stack[knot], rope_stack[next_knot]) >= 2:
                        # print("in")
                        rope_stack[next_knot], temp = temp, rope_stack[next_knot]
                # if math.dist(curr_head, curr_tail) >= 2:
                #     # Essentially go to the old head
                #     curr_tail = temp_head
                #     tail_past.add(curr_tail)
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
    tail_past.add(rope_stack[-1])
    return rope_stack


with open("small_input_p2.txt", "r", encoding="utf-8") as file:
    directions: list[list[str]] = [line.strip().split() for line in file]
    # use dictionary with tuples for checking, add tail pos to dict if it's not there
    head_pos = (0, 0)
    tail_pos = head_pos
    tail_past = {tail_pos}
    rope_stack = [(0, 0) for _ in range(KNOT_NUM)]
    for direction in directions[:1]:
        rope_stack = do_direction_p2(direction, rope_stack)
    print(len(tail_past))
    print(tail_past)

    # plt.rcParams["figure.autolayout"] = True

    # tail_list = list(tail_past)

    # x = [point[0] for point in tail_list]
    # y = [point[1] for point in tail_list]

    # plt.plot(x, y, "s")
    # plt.yticks(np.arange(min(y), max(y) + 1, 1.0))
    # plt.show()
