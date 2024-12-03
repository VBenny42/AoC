from functools import reduce


def get_priority(character: str) -> int:
    priority = ord(character) - ord("a") + 1
    if character.isupper():
        priority += 58
    return priority


def line_parse(count: int, line: str) -> int:
    first, second = set(line[: len(line) // 2]), set(line[len(line) // 2 :])
    common: str = first.intersection(second).pop()
    priority = get_priority(common)
    return count + priority


def three_line_parse(count: int, lines: tuple[str, str, str]) -> int:
    first, second, third = (
        set(lines[0]),
        set(lines[1]),
        set(lines[2]),
    )
    common = first.intersection(second).intersection(third).pop()
    priority = get_priority(common)
    return count + priority


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()

    list_iter = iter(lines)
    trio_list = []
    for line in list_iter:
        elf_trio = (line.strip(), next(list_iter).strip(), next(list_iter).strip())
        trio_list.append(elf_trio)

    print(reduce(line_parse, lines, 0))
    print(reduce(three_line_parse, trio_list, 0))
