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

    # Find the cheapest combination of A and B to get to the prize location

    x, y = machine["prize"]
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

    result = dp(x, y)
    return result if result < float("inf") else None


def main1():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    machines = [read_machine_info(lines[i : i + 3]) for i in range(0, len(lines), 4)]
    min_tokens = 0
    for machine in machines:
        cheapest_combination = find_cheapest_combination(machine)
        if cheapest_combination is not None:
            min_tokens += cheapest_combination
    print(f"LOGF: { min_tokens = }")


def main2():
    pass


if __name__ == "__main__":
    main1()
    main2()
