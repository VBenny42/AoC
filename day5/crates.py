import re
import itertools


def split_crates(line: str) -> list[str]:
    step = 4
    return [
        re.sub("\[|\]", "", line[i : i + step].strip())
        for i in range(0, len(line), step)
    ]


def do_instruction_p1(line: str, stack: list[list[str]]):
    split_line = line.strip().split()
    number, src, dst = (
        int(split_line[1]),
        int(split_line[3]) - 1,
        int(split_line[5]) - 1,
    )
    for _ in range(number):
        crate = stack[src].pop()
        stack[dst].append(crate)


def do_instruction_p2(line: str, stack: list[list[str]]):
    split_line = line.strip().split()
    number, src, dst = (
        int(split_line[1]),
        int(split_line[3]) - 1,
        int(split_line[5]) - 1,
    )
    temp_stack = []
    for _ in range(number):
        temp_stack.append(stack[src].pop())
    for _ in range(len(temp_stack)):
        stack[dst].append(temp_stack.pop())


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    crates = list(
        map(
            split_crates,
            itertools.takewhile(lambda x: not re.sub("\s*", "", x).isdigit(), lines),
        )
    )

    stack = [[] for _ in range(len(crates[0]))]

    crate_rev_iter = reversed(crates)
    for _ in range(0, len(crates)):
        row = next(crate_rev_iter)
        for i in range(0, len(row)):
            if row[i] != "":
                stack[i].append(row[i])
    # print(stack)

    instructions = lines[len(crates) + 2 :]
    for instruction in instructions:
        # do_instruction_p1(instruction, stack)
        do_instruction_p2(instruction, stack)
    # print(stack)
    for elem in stack:
        print(elem[-1], end="")
    print("")
