with open("input.txt", "r") as file:
    instructions = [line.strip().split() for line in file]


x = 1
cycle_count = 0
cycle_checks = (20, 60, 100, 140, 180, 220)
index_check = 0
signal_strength_sum = 0


def cycle_counter():
    global cycle_count
    global index_check
    cycle_count += 1
    if cycle_count == cycle_checks[index_check]:
        global signal_strength_sum
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
    except IndexError:
        break

print(f"{signal_strength_sum = }")

ROW_SIZE = 40

cycle_count = 0
x = 1


def cycle_drawer():
    global cycle_count, x
    print("#" if cycle_count % ROW_SIZE in set(range(x - 1, x + 2)) else ".", end="")
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
