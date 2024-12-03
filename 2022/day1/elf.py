elves: list[int] = []

line: str = ""

with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    line_iter = iter(lines)
    while True:
        try:
            elf = 0
            while (line := next(line_iter).strip()) != "":
                elf += int(line)
            elves.append(elf)
        except StopIteration:
            print("Breaking")
            break

sumt = max(elves)
elves.remove(max(elves))

sumt += max(elves)
elves.remove(max(elves))

sumt += max(elves)
print(sumt)
