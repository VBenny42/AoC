import operator
from math import ceil, log10
from time import time
from typing import List, Tuple

Equation = Tuple[int, List[int]]


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


def is_valid_equation(equation: Equation, operators) -> bool:
    desired_value, numbers = equation
    if numbers[0] > desired_value:
        return False
    if len(numbers) == 2:
        return any(desired_value == op(*numbers) for op in operators)
    # More than 2 remaining_numbers
    # Solve for the first two numbers and recurse
    first, second, *remaining_numbers = numbers
    return any(
        is_valid_equation(
            (desired_value, [op(first, second)] + remaining_numbers), operators
        )
        for op in operators
    )


def concat_op(a, b):
    return a * (10 ** ceil(log10(b + 1))) + b


def ends_with(a: int, b: int) -> int:
    str_a, str_b = str(a), str(b)
    if str_a.endswith(str_b):
        remaining = str_a[: -len(str_b)] if len(str_a) > len(str_b) else "0"
        return int(remaining)
    return 0


def is_valid_equation_p1(equation: Equation) -> bool:
    desired_value, numbers = equation
    if len(numbers) == 2:
        return (
            desired_value == numbers[0] + numbers[1]
            or desired_value == numbers[0] * numbers[1]
        )
    last = numbers[-1]
    mult, add = False, False
    if desired_value % last == 0:
        mult = is_valid_equation_p1((desired_value // last, numbers[:-1]))
    if desired_value - last >= 0:
        add = is_valid_equation_p1((desired_value - last, numbers[:-1]))
    return any([mult, add])


def is_valid_equation_p2(equation: Equation) -> bool:
    desired_value, numbers = equation
    if len(numbers) == 2:
        return (
            desired_value == numbers[0] + numbers[1]
            or desired_value == numbers[0] * numbers[1]
            or desired_value == concat_op(numbers[0], numbers[1])
        )
    last = numbers[-1]
    mult, add, concat = False, False, False
    if desired_value % last == 0:
        mult = is_valid_equation_p2((desired_value // last, numbers[:-1]))
    if desired_value - last >= 0:
        add = is_valid_equation_p2((desired_value - last, numbers[:-1]))
    if concat := ends_with(desired_value, last):
        concat = is_valid_equation_p2((concat, numbers[:-1]))
    return any([mult, add, concat])


def print_equation(equation: Equation) -> None:
    desired_value, numbers = equation
    print(desired_value, end=": ")
    print(" ".join(map(str, numbers)))


@timer_func
def main1_forward():
    with open("input.txt", "r") as f:
        equations = [
            (int(desired_value), list(map(int, remaining_numbers.split())))
            for line in f
            for desired_value, remaining_numbers in [line.split(":")]
        ]

    print(
        f"ANSWER1: Sum of true equations: {sum(equation[0] for equation in equations if is_valid_equation(equation, {operator.add, operator.mul}))}"
    )


# @timer_func
def main1_backward():
    with open("input.txt", "r") as f:
        equations = [
            (int(desired_value), list(map(int, remaining_numbers.split())))
            for line in f
            for desired_value, remaining_numbers in [line.split(":")]
        ]

    print(
        f"ANSWER1: Sum of true equations: {sum(equation[0] for equation in equations if is_valid_equation_p1(equation))}"
    )


@timer_func
def main2_forward():
    with open("input.txt", "r") as f:
        equations = [
            (int(desired_value), list(map(int, remaining_numbers.split())))
            for line in f
            for desired_value, remaining_numbers in [line.split(":")]
        ]

    print(
        f"ANSWER2: Sum of true equations with concat: {sum(equation[0] for equation in equations if is_valid_equation(equation, {operator.add, operator.mul, concat_op}))}"
    )


# @timer_func
def main2_backward():
    with open("input.txt", "r") as f:
        equations = [
            (int(desired_value), list(map(int, remaining_numbers.split())))
            for line in f
            for desired_value, remaining_numbers in [line.split(":")]
        ]

    print(
        f"ANSWER2: Sum of true equations with concat: {sum(equation[0] for equation in equations if is_valid_equation_p2(equation))}"
    )


if __name__ == "__main__":
    # main1_forward()
    main1_backward()
    # main2_forward()
    main2_backward()
