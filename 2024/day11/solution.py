from typing import List


def apply_rules(stone: int) -> List[int]:
    if stone == 0:
        return [1]

    length = 0
    temp = stone

    while temp > 0:
        length += 1
        temp //= 10

    if length % 2 == 0:
        split_point = length // 2
        first_half = stone // 10 ** split_point
        second_half = stone % 10 ** split_point
        return [first_half, second_half]

    return [stone * 2024]


def blink(stones: List[int]) -> list[int]:
    # new_stones = []
    # for stone in stones:
    #     new_stones.extend(apply_rules(stone))
    # return new_stones
    return [stone for stone in stones for stone in apply_rules(stone)]

def main1():
    with open("input.txt", "r") as f:
        stones = list(map(int, f.read().split()))
    # print(f"LOG: { stones = }")
    for _ in range(25):
        stones = blink(stones)
        # print(f"LOG: { stones = }")
    print(f"LOG: { len(stones) = }")


if __name__ == "__main__":
    main1()
