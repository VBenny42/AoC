from functools import cache

DIRECTIONAL_KEYPAD_TRANSITIONS = {
    "^": {
        "<": ["v<"],
        ">": [">v", "v>"],
        "A": [">"],
        "^": [""],
        "v": ["v"],
    },
    "v": {
        "<": ["<"],
        ">": [">"],
        "A": [">^", "^>"],
        "^": ["^"],
        "v": [""],
    },
    ">": {
        "<": ["<<"],
        ">": [""],
        "A": ["^"],
        "^": ["<^", "^<"],
        "v": ["<"],
    },
    "<": {
        "<": [""],
        ">": [">>"],
        "A": [">>^", ">^>"],
        "^": [">^"],
        "v": [">"],
    },
    "A": {
        "<": ["<v<", "v<<"],
        ">": ["v"],
        "A": [""],
        "^": ["<"],
        "v": ["<v", "v<"],
    },
}


NUMERIC_KEYPAD_TRANSITIONS = {
    "A": {
        "0": ["<"],
        "1": ["<^<", "^<<"],
        "2": ["<^", "^<"],
        "3": ["^"],
        "4": ["<^<^", "^<<^", "<^^<", "^<^<", "^^<<"],
        "5": ["<^^", "^<^", "^^<"],
        "6": ["^^"],
        "7": [
            "<^<^^",
            "^<<^^",
            "<^^<^",
            "^<^<^",
            "^^<<^",
            "<^^^<",
            "^<^^<",
            "^^<^<",
            "^^^<<",
        ],
        "8": ["<^^^", "^<^^", "^^<^", "^^^<"],
        "9": ["^^^"],
        "A": [""],
    },
    "0": {
        "0": [""],
        "1": ["^<"],
        "2": ["^"],
        "3": [">^", "^>"],
        "4": ["^<^", "^^<"],
        "5": ["^^"],
        "6": [">^^", "^>^", "^^>"],
        "7": ["^<^^", "^^<^", "^^^<"],
        "8": ["^^^"],
        "9": [">^^^", "^>^^", "^^>^", "^^^>"],
        "A": [">"],
    },
    "1": {
        "0": [">v"],
        "1": [""],
        "2": [">"],
        "3": [">>"],
        "4": ["^"],
        "5": [">^", "^>"],
        "6": [">>^", ">^>", "^>>"],
        "7": ["^^"],
        "8": [">^^", "^>^", "^^>"],
        "9": [">>^^", ">^>^", "^>>^", ">^^>", "^>^>", "^^>>"],
        "A": [">>v", ">v>"],
    },
    "2": {
        "0": ["v"],
        "1": ["<"],
        "2": [""],
        "3": [">"],
        "4": ["<^", "^<"],
        "5": ["^"],
        "6": [">^", "^>"],
        "7": ["<^^", "^<^", "^^<"],
        "8": ["^^"],
        "9": [">^^", "^>^", "^^>"],
        "A": [">v", "v>"],
    },
    "3": {
        "0": ["<v", "v<"],
        "1": ["<<"],
        "2": ["<"],
        "3": [""],
        "4": ["<<^", "<^<", "^<<"],
        "5": ["<^", "^<"],
        "6": ["^"],
        "7": ["<<^^", "<^<^", "^<<^", "<^^<", "^<^<", "^^<<"],
        "8": ["<^^", "^<^", "^^<"],
        "9": ["^^"],
        "A": ["v"],
    },
    "4": {
        "0": [">vv", "v>v"],
        "1": ["v"],
        "2": [">v", "v>"],
        "3": [">>v", ">v>", "v>>"],
        "4": [""],
        "5": [">"],
        "6": [">>"],
        "7": ["^"],
        "8": [">^", "^>"],
        "9": [">>^", ">^>", "^>>"],
        "A": [">>vv", ">v>v", "v>>v", ">vv>", "v>v>"],
    },
    "5": {
        "0": ["vv"],
        "1": ["<v", "v<"],
        "2": ["v"],
        "3": [">v", "v>"],
        "4": ["<"],
        "5": [""],
        "6": [">"],
        "7": ["<^", "^<"],
        "8": ["^"],
        "9": [">^", "^>"],
        "A": [">vv", "v>v", "vv>"],
    },
    "6": {
        "0": ["<vv", "v<v", "vv<"],
        "1": ["<<v", "<v<", "v<<"],
        "2": ["<v", "v<"],
        "3": ["v"],
        "4": ["<<"],
        "5": ["<"],
        "6": [""],
        "7": ["<<^", "<^<", "^<<"],
        "8": ["<^", "^<"],
        "9": ["^"],
        "A": ["vv"],
    },
    "7": {
        "0": [">vvv", "v>vv", "vv>v"],
        "1": ["vv"],
        "2": [">vv", "v>v", "vv>"],
        "3": [">>vv", ">v>v", "v>>v", ">vv>", "v>v>", "vv>>"],
        "4": ["v"],
        "5": [">v", "v>"],
        "6": [">>v", ">v>", "v>>"],
        "7": [""],
        "8": [">"],
        "9": [">>"],
        "A": [
            ">>vvv",
            ">v>vv",
            "v>>vv",
            ">vv>v",
            "v>v>v",
            "vv>>v",
            ">vvv>",
            "v>vv>",
            "vv>v>",
        ],
    },
    "8": {
        "0": ["vvv"],
        "1": ["<vv", "v<v", "vv<"],
        "2": ["vv"],
        "3": [">vv", "v>v", "vv>"],
        "4": ["<v", "v<"],
        "5": ["v"],
        "6": [">v", "v>"],
        "7": ["<"],
        "8": [""],
        "9": [">"],
        "A": [">vvv", "v>vv", "vv>v", "vvv>"],
    },
    "9": {
        "0": ["<vvv", "v<vv", "vv<v", "vvv<"],
        "1": ["<<vv", "<v<v", "v<<v", "<vv<", "v<v<", "vv<<"],
        "2": ["<vv", "v<v", "vv<"],
        "3": ["vv"],
        "4": ["<<v", "<v<", "v<<"],
        "5": ["<v", "v<"],
        "6": ["v"],
        "7": ["<<"],
        "8": ["<"],
        "9": [""],
        "A": ["vvv"],
    },
}


@cache
def shortest_sequence(level: int, sequence_str: str, num_robots: int):
    if level == num_robots + 1:
        return len(sequence_str)

    transitions = (
        NUMERIC_KEYPAD_TRANSITIONS if level == 0 else DIRECTIONAL_KEYPAD_TRANSITIONS
    )

    sequence = 0
    # Always start at "A"
    for current, target in zip("A" + sequence_str, sequence_str):
        # Must press "A" to enter sequence for next bot
        possible_paths = (
            shortest_sequence(level + 1, path + "A", num_robots)
            for path in transitions[current][target]
        )
        # If no possible paths, then start and end are the same,
        # so only one path is possible
        sequence += min(
            (path for path in possible_paths if path is not None), default=1
        )

    return sequence


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        lines = f.read().splitlines()

    sum_of_shortest_sequences = 0
    for line in lines:
        number = int(line[:-1])
        shortest_sequence_length = shortest_sequence(0, line, 2)
        sum_of_shortest_sequences += shortest_sequence_length * number
    print(f"ANSWER1: { sum_of_shortest_sequences }")


def main2():
    with open("input.txt", "r", encoding="utf-8") as f:
        lines = f.read().splitlines()

    sum_of_shortest_sequences = 0
    for line in lines:
        number = int(line[:-1])
        shortest_sequence_length = shortest_sequence(0, line, 25)
        sum_of_shortest_sequences += shortest_sequence_length * number
    print(f"ANSWER2: { sum_of_shortest_sequences  }")


if __name__ == "__main__":
    main1()
    main2()
