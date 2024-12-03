from functools import reduce


with open("input.txt", "r") as file:
    notes = [line.strip() for line in file]


def build_monkey_list():
    monkeys = []
    notes_iter = iter(notes)
    while True:
        try:
            monkey = {"times": 0}
            while (line := next(notes_iter)) != "":
                monkey = monkey_init(line, monkey)
            monkeys.append(monkey)
        except StopIteration:
            monkeys.append(monkey)
            break
    return monkeys


def monkey_init(line: str, monkey: dict):
    match line.split():
        case ["Monkey", monkey_num]:
            monkey["num"] = monkey_num.replace(":", "")
        case ["Starting", "items:", *_]:
            monkey["items"] = list(map(int, line.replace(",", "").split()[2:]))
        case ["Operation:", *_]:
            monkey["operation"] = line.split()[3:]
        case ["Test:", _, _, num]:
            monkey["test"] = int(num)
        case ["If", "true:", _, _, _, monkey_num]:
            monkey["true_throw"] = int(monkey_num)
        case ["If", "false:", _, _, _, monkey_num]:
            monkey["false_throw"] = int(monkey_num)
    return monkey


def round(monkeys: list[dict], level_reduce: int):
    mod_by = reduce(lambda mod_by, monkey: monkey["test"] * mod_by, monkeys, 1)
    for monkey in monkeys:
        monkey["times"] += len(monkey["items"])
        for item in monkey["items"]:
            worry_level = (
                eval(
                    "".join(
                        list(
                            map(
                                lambda x: str(item) if x == "old" else x,
                                monkey["operation"],
                            )
                        )
                    )
                )
            ) // level_reduce
            if level_reduce == 1:
                worry_level %= mod_by
            if worry_level % monkey["test"] == 0:
                monkeys[monkey["true_throw"]]["items"].append(worry_level)
            else:
                monkeys[monkey["false_throw"]]["items"].append(worry_level)
        monkey["items"] = []


# Ik it's read in order, just making sure
# monkeys = sorted(monkeys, key=lambda monkey: monkey["num"])
def calc_rounds(times: int, round_num: int):
    monkeys = build_monkey_list()

    for _ in range(times):
        round(monkeys, round_num)

    print(
        reduce(
            lambda monkey_business, monkey: monkey["times"] * monkey_business,
            sorted(monkeys, key=lambda monkey: monkey["times"], reverse=True)[:2],
            1,
        )
    )


# Part 1
calc_rounds(20, 3)

# Part 2
calc_rounds(10000, 1)
