import operator
from typing import List, Tuple

Equation = Tuple[int, List[int]]


def is_valid_equation1(equation: Equation) -> bool:
    operators = {operator.add, operator.mul}
    test_value, numbers = equation
    if len(numbers) == 2:
        return any(test_value == op(*numbers) for op in operators)
    # More than 2 remaining_numbers
    # Solve for the first two numbers and recurse
    first, second, *remaining_numbers = numbers
    return any(
        is_valid_equation1((test_value, [op(first, second)] + remaining_numbers))
        for op in operators
    )


def concat_op(a, b):
    return int(str(a) + str(b))


def is_valid_equation2(equation: Equation) -> bool:
    operators = {operator.add, operator.mul, concat_op}
    test_value, numbers = equation
    if len(numbers) == 2:
        return any(test_value == op(*numbers) for op in operators)
    # More than 2 remaining_numbers
    # Solve for the first two numbers and recurse
    first, second, *remaining_numbers = numbers
    return any(
        is_valid_equation2((test_value, [op(first, second)] + remaining_numbers))
        for op in operators
    )


def print_equation(equation: Equation) -> None:
    test_value, numbers = equation
    print(test_value, end=": ")
    print(" ".join(map(str, numbers)), end="")
    print()


def main1():
    with open("input.txt", "r") as f:
        equations = [
            (int(test_value), list(map(int, remaining_numbers.split())))
            for line in f
            for test_value, remaining_numbers in [line.split(":")]
        ]

    print(
        f"LOGF: Sum of true equations: {sum(equation[0] for equation in equations if is_valid_equation1(equation))}"
    )


def main2():
    with open("input.txt", "r") as f:
        equations = [
            (int(test_value), list(map(int, remaining_numbers.split())))
            for line in f
            for test_value, remaining_numbers in [line.split(":")]
        ]

    # for equation in equations:
    #     if is_valid_equation1(equation):
    #         print_equation(equation)
    #     if is_valid_equation2(equation):
    #         print_equation(equation)

    print(
        f"LOGF: Sum of true equations: {sum(equation[0] for equation in equations if is_valid_equation2(equation))}"
    )


if __name__ == "__main__":
    # main1()
    main2()
