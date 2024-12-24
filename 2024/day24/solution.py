def run_gate(gate: str, inputs: dict[str, int]):
    left_input, operator, right_input, _, output = gate.split(" ")
    if left_input not in inputs or right_input not in inputs:
        return False
    match operator:
        case "AND":
            inputs[output] = inputs[left_input] & inputs[right_input]
        case "OR":
            inputs[output] = inputs[left_input] | inputs[right_input]
        case "XOR":
            inputs[output] = inputs[left_input] ^ inputs[right_input]
        case _:
            raise NotImplementedError(f"Operator {operator} not implemented")
    return True


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        inputs, gates = f.read().split("\n\n")

    inputs = inputs.splitlines()
    inputs = {input.split(": ")[0]: int(input.split(": ")[1]) for input in inputs}
    gates = gates.splitlines()

    while gates:
        gate = gates.pop(0)
        if run_gate(gate, inputs):
            continue
        gates.append(gate)

    z_wires = ((k, v) for k, v in inputs.items() if k.startswith("z"))
    z_wires = sorted(z_wires, key=lambda x: x[0], reverse=True)

    binary_number = int("".join(str(x[1]) for x in z_wires), 2)
    print(f"ANSWER1: { binary_number = }")


def main2():
    pass


if __name__ == "__main__":
    main1()
    main2()
