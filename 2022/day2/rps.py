from functools import reduce

OUTCOMES = [
    {
        "A": {"X": 4, "Y": 8, "Z": 3},
        "B": {"X": 1, "Y": 5, "Z": 9},
        "C": {"X": 7, "Y": 2, "Z": 6},
    },
    {
        "A": {"X": 3, "Y": 4, "Z": 8},
        "B": {"X": 1, "Y": 5, "Z": 9},
        "C": {"X": 2, "Y": 6, "Z": 7},
    },
]


def get_outcome(score: int, line: str) -> int:
    opponent, me = line.split(maxsplit=2)
    return score + OUTCOMES[0][opponent][me]


def get_score(score: int, line: str) -> int:
    opponent, me = line.split(maxsplit=2)
    return score + OUTCOMES[1][opponent][me]


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    print(reduce(get_outcome, lines, 0))
    print(reduce(get_score, lines, 0))
