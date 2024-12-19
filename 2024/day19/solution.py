from collections.abc import Iterable


def is_possible2(towels: Iterable[str], design: str) -> bool:
    n = len(design)
    dp = [False] * (n + 1)
    dp[0] = True

    for i in range(1, n + 1):
        for towel in towels:
            if design.startswith(towel, i - len(towel)) and dp[i - len(towel)]:
                dp[i] = True
                break

    return dp[n]


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        part1, part2 = f.read().split("\n\n")
        towels = set(part1.split(", "))

        designs = part2.splitlines()

    possible_designs = sum(is_possible2(towels, design) for design in designs)

    print(f"LOGF: { possible_designs = }")


def main2():
    pass


if __name__ == "__main__":
    main1()
    main2()
