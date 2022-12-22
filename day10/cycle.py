from time import sleep


with open("input.txt", "r") as file:
    instructions = [line.strip().split() for line in file]


x = 1
cycle_count = 0
cycle_checks = (20, 60, 100, 140, 180, 220)
index_check = 0
signal_strength_sum = 0


def cycle_counter():
    global cycle_count, index_check, signal_strength_sum
    cycle_count += 1
    if cycle_count == cycle_checks[index_check]:
        index_check += 1
        signal_strength_sum += cycle_count * x


for instruction in instructions:
    try:
        match instruction:
            case ["noop"]:
                cycle_counter()
            case ["addx", num]:
                for _ in range(2):
                    cycle_counter()
                x += int(num)
        # sleep(2)
    except IndexError:
        break

print(f"{signal_strength_sum = }")

ROW_SIZE = 40
PAD_NUM = 3

cycle_count = 0
x = 1


def cycle_drawer():
    global cycle_count
    # sleep(0.01)
    print(
        f"{'#' * PAD_NUM}"
        if cycle_count % ROW_SIZE in set(range(x - 1, x + 2))
        else f"{'.' * PAD_NUM}",
        end="",
        flush=True,
    )
    cycle_count += 1
    if cycle_count % ROW_SIZE == 0:
        # print(f"{cycle_count=}")
        print()


for instruction in instructions:
    match instruction:
        case ["noop"]:
            cycle_drawer()
        case ["addx", num]:
            for _ in range(2):
                cycle_drawer()
            x += int(num)
