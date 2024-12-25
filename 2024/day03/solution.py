import re


def get_result1(line: str) -> int:
    # Search for mul(a,b) in the line,
    # Group 1: a, Group 2: b
    pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
    matches = re.finditer(pattern, line)

    return sum(int(match.group(1)) * int(match.group(2)) for match in matches)


def get_result2(line: str) -> int:
    # Search for mul(a,b), do() or don't() in the line
    # Group 1: a,b, Group 2: do(), Group 3: don't()
    pattern = r"mul\((\d{1,3},\d{1,3})\)|(do\(\))|(don't\(\))"
    matches = re.finditer(pattern, line)

    sum = 0
    enabled = True
    for match in matches:
        # Perform the multiplication only if enabled true, continue otherwise
        if match.group(1) and enabled:
            a, b = map(int, match.group(1).split(","))
            sum += a * b
        # Match is do(), enable future multiplication
        elif match.group(2):
            enabled = True
        # Match is don't(), disable future multiplication
        elif match.group(3):
            enabled = False
    return sum


def main1():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    print(f"ANSWER1: {sum(get_result1(line) for line in lines)}")


def main2():
    with open("input.txt", "r") as f:
        lines = f.readlines()
    # Need to join the lines, as do/don't state carries over to the next line
    joined = "".join(lines)
    print(f"ANSWER2: {get_result2(joined)}")


if __name__ == "__main__":
    main1()
    main2()
