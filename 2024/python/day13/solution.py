from typing import List
from re import search


def read_machine_info(lines: List[str]) -> dict:
    pattern_behavior = r"([A|B]):.*X\+(\d+).*Y\+(\d+)"
    res = {}
    for line in lines[:2]:
        matches = search(pattern_behavior, line)
        assert matches is not None
        res[matches.group(1)] = (int(matches.group(2)), int(matches.group(3)))
    pattern_prize = r"X=(\d+), Y=(\d+)"
    matches = search(pattern_prize, lines[2])
    assert matches is not None
    res["prize"] = tuple(map(int, matches.groups()))
    return res


def find_cheapest_combination(machine: dict):
    a_cost = 3
    b_cost = 1

    p_x, p_y = machine["prize"]
    a_x, a_y = machine["A"]
    b_x, b_y = machine["B"]

    memo = {}

    def dp(x, y):
        if x < 0 or y < 0:
            return float("inf")
        if x == 0 and y == 0:
            return 0
        if (x, y) in memo:
            return memo[(x, y)]
        cost = min(a_cost + dp(x - a_x, y - a_y), b_cost + dp(x - b_x, y - b_y))
        memo[(x, y)] = cost
        return cost

    result = dp(p_x, p_y)
    return result if result < float("inf") else None


def find_cheapest_combination2(machine: dict):
    p_x, p_y = machine["prize"]
    a_x, a_y = machine["A"]
    b_x, b_y = machine["B"]

    determinant = a_x * b_y - a_y * b_x
    if determinant == 0:
        return None  # No unique solution

    a = (p_x * b_y - p_y * b_x) // determinant
    b = (a_x * p_y - a_y * p_x) // determinant

    # Validate the solution
    if (a_x * a + b_x * b) != p_x or (a_y * a + b_y * b) != p_y:
        return None

    return 3 * a + b  # Return the cost


def main1():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    machines = [read_machine_info(lines[i : i + 3]) for i in range(0, len(lines), 4)]
    min_tokens = 0
    for machine in machines:
        cheapest_combination = find_cheapest_combination2(machine)
        if cheapest_combination is not None:
            min_tokens += cheapest_combination
    print(f"ANSWER1: { min_tokens = }")


def main2():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    machines = [read_machine_info(lines[i : i + 3]) for i in range(0, len(lines), 4)]
    addition = 10000000000000
    min_tokens = 0
    for machine in machines:
        machine["prize"] = (
            machine["prize"][0] + addition,
            machine["prize"][1] + addition,
        )
        cheapest_combination = find_cheapest_combination2(machine)
        if cheapest_combination is not None:
            min_tokens += cheapest_combination
    print(f"ANSWER2: { min_tokens = }")


if __name__ == "__main__":
    main1()
    main2()
