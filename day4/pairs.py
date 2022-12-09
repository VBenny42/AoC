from functools import reduce


def parse_int(line: str) -> list[list[int]]:
    return [tuple(map(int, x.split("-"))) for x in line.split(",")]


def determine_subset(count: int, elf_pair: list[tuple[int, int]]):
    if set(range(elf_pair[0][0], elf_pair[0][1] + 1)).issubset(
        set(range(elf_pair[1][0], elf_pair[1][1] + 1))
    ) or set(range(elf_pair[1][0], elf_pair[1][1] + 1)).issubset(
        set(range(elf_pair[0][0], elf_pair[0][1] + 1))
    ):
        return count + 1
    return count


def determine_intersect(count: int, elf_pair: list[tuple[int, int]]):
    if set(range(elf_pair[0][0], elf_pair[0][1] + 1)).intersection(
        set(range(elf_pair[1][0], elf_pair[1][1] + 1))
    ):
        return count + 1
    return count


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    elf_pairs = [parse_int(line) for line in lines]
    print(reduce(determine_subset, elf_pairs, 0))
    print(reduce(determine_intersect, elf_pairs, 0))
