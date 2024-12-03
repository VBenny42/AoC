def read_line(line: str, distinct_num: int) -> int:
    for i in range(len(line) - distinct_num):
        if len(set(line[i : i + distinct_num])) == distinct_num:
            return i


with open("input.txt", "r", encoding="utf-8") as file:
    line: str = file.readline()
    distinct_num = 14
    print(read_line(line, distinct_num) + distinct_num)
