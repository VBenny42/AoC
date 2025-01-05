from itertools import pairwise
import numpy as np


def update_head(head: np.ndarray, direction: str) -> np.array:
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


def update_tail(head: np.array, tail: np.ndarray) -> np.array:
    head_tail_diff = {
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
    return tail + np.array(head_tail_diff.get(tuple(head - tail), (0, 0)))


with open("input.txt", "r") as file:
    directions: list[list[str]] = [line.strip().split() for line in file]


def calculate_tail_pos(directions: list[list[str]], knot_num: int):
    rope_stack = [np.array([0, 0]) for _ in range(knot_num)]
    tail_past = {(0, 0)}
    for direction, num in directions:
        for _ in range(int(num)):
            rope_stack[0] = update_head(rope_stack[0], direction)
            for knot, next_knot in pairwise(range(knot_num)):
                rope_stack[next_knot] = update_tail(
                    rope_stack[knot], rope_stack[next_knot]
                )
            tail_past.add(tuple(rope_stack[-1]))
    print(len(tail_past))


calculate_tail_pos(directions, 2)
calculate_tail_pos(directions, 10)


# plt.rcParams["figure.autolayout"] = True

# tail_list = list(tail_past)

# x = [point[0] for point in tail_list]
# y = [point[1] for point in tail_list]

# plt.plot(x, y, "s")
# plt.yticks(np.arange(min(y), max(y) + 1, 1.0))
# plt.show()
