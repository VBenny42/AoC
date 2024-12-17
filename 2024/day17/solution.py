Registers = dict[str, int]


def get_combo_value(registers: Registers, operand: int) -> int:
    value = -1
    if operand in {0, 1, 2, 3}:
        value = operand
    match operand:
        case 4:
            value = registers["A"]
        case 5:
            value = registers["B"]
        case 6:
            value = registers["C"]
        case 7:
            raise ValueError("Invalid combo operand")
    return value


def adv(registers: Registers, operand: int):
    numerator = registers["A"]
    divisor = get_combo_value(registers, operand)
    registers["A"] = numerator // pow(2, divisor)


def bxl(registers: Registers, operand: int):
    registers["B"] = registers["B"] ^ operand


def bst(registers: Registers, operand: int):
    value = get_combo_value(registers, operand)
    registers["B"] = value % 8


def jnz(registers: Registers, operand: int):
    if registers["A"] == 0:
        return -1
    return operand


def bxc(registers: Registers, operand: int):
    registers["B"] = registers["B"] ^ registers["C"]


def out(registers: Registers, operand: int):
    value = get_combo_value(registers, operand)
    return value % 8


def bdv(registers: Registers, operand: int):
    numerator = registers["A"]
    divisor = get_combo_value(registers, operand)
    registers["B"] = numerator // pow(2, divisor)


def cdv(registers: Registers, operand: int):
    numerator = registers["A"]
    divisor = get_combo_value(registers, operand)
    registers["C"] = numerator // pow(2, divisor)


def execute_instructions(registers: Registers, program: list[int]):
    ip = 0
    outs = []
    instructions = {
        0: adv,
        1: bxl,
        2: bst,
        3: jnz,
        4: bxc,
        5: out,
        6: bdv,
        7: cdv,
    }
    while ip < len(program):
        try:
            instruction = program[ip]
            operand = program[ip + 1]
        except IndexError:
            break
        print(f"LOG: { instruction = }, { operand = } { ip = }")
        instruction_fn = instructions[instruction]
        if instruction == 3:
            res = jnz(registers, operand)
            if res != -1:
                ip = res
                continue
        else:
            ret = instruction_fn(registers, operand)
            if ret is not None:
                outs.append(ret)
        print(f"LOG: { registers = }")
        print()
        ip += 2

    return outs


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        lines = f.readlines()
        registers = {}
        registers["A"] = int(lines[0].split(":")[1].strip())
        registers["B"] = int(lines[1].split(":")[1].strip())
        registers["C"] = int(lines[2].split(":")[1].strip())

        program = list(map(int, lines[4].split(":")[1].strip().split(",")))

    print(f"LOG: { registers = }")
    print(f"LOG: { program = }")

    out = execute_instructions(registers, program)
    print(f"LOG: { registers = }")
    print(f"LOG: outs", ",".join(map(str, out)))


def main2():
    pass


if __name__ == "__main__":
    main1()
    # main2()
