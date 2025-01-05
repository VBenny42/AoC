from functools import cache
from math import floor, log10
from typing import List


def apply_rules(stone: int) -> List[int]:
    if stone == 0:
        return [1]

    length = floor(log10(stone)) + 1

    if length % 2 == 0:
        split_point = length // 2
        first_half = stone // 10**split_point
        second_half = stone % 10**split_point
        return [first_half, second_half]

    return [stone * 2024]


@cache
def apply_rules_recursive(stone: int, blinks: int) -> int:
    if blinks == 0:
        return 1
    if stone == 0:
        return apply_rules_recursive(1, blinks - 1)
    length = floor(log10(stone)) + 1
    if length % 2 == 0:
        split_point = length // 2
        first_half = stone // 10**split_point
        second_half = stone % 10**split_point
        return apply_rules_recursive(first_half, blinks - 1) + apply_rules_recursive(
            second_half, blinks - 1
        )
    return apply_rules_recursive(stone * 2024, blinks - 1)


def blink(stones: List[int]) -> list[int]:
    return [new_stone for stone in stones for new_stone in apply_rules(stone)]


def blink_recursive(stones: List[int], iterations: int) -> int:
    return sum(apply_rules_recursive(stone, iterations) for stone in stones)


def main1():
    with open("input.txt", "r") as f:
        stones = list(map(int, f.read().split()))
    for _ in range(25):
        stones = blink(stones)
    print(f"ANSWER1: { len(stones) = }")


def main3():
    with open("input.txt", "r") as f:
        stones = list(map(int, f.read().split()))
    stones = blink_recursive(stones, 25)
    print(f"ANSWER1: { stones = }")


def main2():
    with open("input.txt", "r") as f:
        stones = list(map(int, f.read().split()))
    stones = blink_recursive(stones, 75)
    print(f"ANSWER2: { stones = }")


if __name__ == "__main__":
    # main1()
    main3()
    main2()
