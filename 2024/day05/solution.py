from functools import cmp_to_key
from time import time
from typing import DefaultDict, Iterable


def timer_func(func):
    # This function shows the execution time of
    # the function object passed
    def wrap_func(*args, **kwargs):
        t1 = time()
        result = func(*args, **kwargs)
        t2 = time()
        print(f"Function {func.__name__!r} executed in {(t2-t1):.4f}s")
        return result

    return wrap_func


def build_ruleset(rules: Iterable[list[int]]) -> dict[int, set[int]]:
    ruleset = DefaultDict(set)
    for rule in rules:
        ruleset[rule[0]].add(rule[1])
    return ruleset


def is_valid1(update: tuple[int, ...], ruleset: dict[int, set[int]]) -> bool:
    for i in range(len(update)):
        before = update[:i]
        page = update[i]
        if page in ruleset:
            if any(dep in before for dep in ruleset[page]):
                return False
    return True


def is_valid2(update: tuple[int, ...], ruleset: dict[int, set[int]]) -> bool:
    compare_with_ruleset = lambda a, b: compare(ruleset, a, b)
    return sorted(update, key=cmp_to_key(compare_with_ruleset)) == list(update)


def compare(ruleset, a, b):
    if a in ruleset:
        if b in ruleset[a]:
            return -1
    return 0


def reordering(
    update: tuple[int, ...], ruleset: dict[int, set[int]]
) -> tuple[int, ...]:
    compare_with_ruleset = lambda a, b: compare(ruleset, a, b)
    return tuple(sorted(update, key=cmp_to_key(compare_with_ruleset)))


@timer_func
def main1():
    with open("input-rules.txt", "r") as f:
        rules = (
            [int(value) for value in line.strip().split("|")] for line in f.readlines()
        )
    with open("input-updates.txt", "r") as f:
        updates = (
            tuple(int(value) for value in line.strip().split(","))
            for line in f.readlines()
        )
    ruleset = build_ruleset(rules)

    sum = 0
    for update in updates:
        if is_valid2(update, ruleset):
            sum += update[(len(update) - 1) // 2]
    print(f"LOGF: {sum = }")


@timer_func
def main3():
    with open("input-rules.txt", "r") as f:
        rules = (
            [int(value) for value in line.strip().split("|")] for line in f.readlines()
        )
    with open("input-updates.txt", "r") as f:
        updates = (
            tuple(int(value) for value in line.strip().split(","))
            for line in f.readlines()
        )
    ruleset = build_ruleset(rules)

    sum = 0
    for update in updates:
        if is_valid2(update, ruleset):
            sum += update[(len(update) - 1) // 2]
    print(f"LOGF: {sum = }")


@timer_func
def main2():
    with open("input-rules.txt", "r") as f:
        rules = (
            [int(value) for value in line.strip().split("|")] for line in f.readlines()
        )
    with open("input-updates.txt", "r") as f:
        updates = (
            tuple(int(value) for value in line.strip().split(","))
            for line in f.readlines()
        )
    ruleset = build_ruleset(rules)

    sum = 0
    for update in updates:
        if not is_valid2(update, ruleset):
            valid_ordering = reordering(update, ruleset)
            sum += valid_ordering[(len(valid_ordering) - 1) // 2]
    print(f"LOGF: {sum = }")


if __name__ == "__main__":
    main1()
    main2()
    main3()
