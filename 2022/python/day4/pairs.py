from functools import reduce


def determine_subset(count: int, elf_pair: list[int, int, int, int]):
    a, b, x, y = elf_pair
    if ((a <= x) and (b >= y)) or ((a >= x) and (b <= y)):
        return count + 1
    return count


def determine_intersect(count: int, elf_pair: list[int, int, int, int]):
    a, b, x, y = elf_pair
    if ((a <= x) and (b >= x)) or ((a <= y) and ((b >= y))) or ((a >= x) and (b <= y)):
        return count + 1
    return count


from re import split


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    elf_pairs = [[int(x) for x in split("-|,", line)] for line in lines]
    print(reduce(determine_subset, elf_pairs, 0))
    print(reduce(determine_intersect, elf_pairs, 0))
