def is_possible(towels: set[str], design: str) -> bool:
    n = len(design)
    dp = [False] * (n + 1)
    dp[0] = True

    for i in range(1, n + 1):
        for towel in towels:
            if design.startswith(towel, i - len(towel)) and dp[i - len(towel)]:
                dp[i] = True
                break

    return dp[n]


def different_combos(towels: set[str], design: str) -> int:
    n = len(design)
    dp = [0] * (n + 1)
    dp[0] = 1

    for i in range(1, n + 1):
        for towel in towels:
            if design.startswith(towel, i - len(towel)):
                dp[i] += dp[i - len(towel)]

    return dp[n]


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        part1, part2 = f.read().split("\n\n")
        towels = set(part1.split(", "))

        designs = part2.splitlines()

    possible_designs = sum(is_possible(towels, design) for design in designs)

    print(f"ANSWER1: { possible_designs = }")


def main2():
    with open("input.txt", "r", encoding="utf-8") as f:
        part1, part2 = f.read().split("\n\n")
        towels = set(part1.split(", "))

        designs = part2.splitlines()

    different_possible_designs = sum(
        different_combos(towels, design) for design in designs
    )

    print(f"ANSWER2: { different_possible_designs = }")


if __name__ == "__main__":
    main1()
    main2()
