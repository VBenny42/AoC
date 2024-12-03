from functools import reduce
from itertools import pairwise
from typing import Iterable


def validate_report_1(report: Iterable[int]) -> bool:
    direction = ""

    for prev, curr in pairwise(report):
        if direction == "":
            if prev > curr:
                direction = "decreasing"
            else:
                direction = "increasing"
        # Adjacent levels can only differ by at least one and at most three
        if (abs(prev - curr) < 1) or (abs(prev - curr) > 3):
            return False
        if direction == "increasing" and prev > curr:
            return False
        if direction == "decreasing" and prev < curr:
            return False
    return True


def validate_report_2(report: list[int]) -> bool:
    # Generate all reports with one level removed
    possible_reports = (report[:i] + report[i + 1 :] for i in range(len(report)))
    # If any of the reports with one level removed is valid,
    # then the original report is considered safe
    return any(
        validate_report_1(possible_report) for possible_report in possible_reports
    )


def main1():
    with open("input.txt", "rb") as f:
        lines = f.readlines()

    converted_lines = ((map(int, line.split())) for line in lines)

    sum_of_valid_reports_part_1 = reduce(
        lambda sum, report: sum + (1 if validate_report_1(report) else 0),
        converted_lines,
        0,
    )
    print(f"LOG: {sum_of_valid_reports_part_1 = }")


def main2():
    with open("input.txt", "rb") as f:
        lines = f.readlines()

    # Need to be able to index the levels for part 2,
    # so we must convert the map object to a list
    converted_lines = ((list(map(int, line.split()))) for line in lines)

    sum_of_valid_reports_part_2 = reduce(
        lambda sum, report: sum + (1 if validate_report_2(report) else 0),
        converted_lines,
        0,
    )
    print(f"LOG: {sum_of_valid_reports_part_2 = }")


if __name__ == "__main__":
    main1()
    main2()
