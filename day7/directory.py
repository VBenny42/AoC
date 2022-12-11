def build_directory_structure(parsed_input: list[list[str]]) -> dict[str, int]:
    directories: dict[str, int] = {}
    cur_size: int = 0

    for line in parsed_input:
        match line[0]:
            case "ls":
                continue
            case "cd":
                match line[1]:
                    case "/":
                        path = [line[1]]
                        cur_path = line[1]
                        directories.setdefault(cur_path, 0)
                    case "..":
                        cur_size = directories.get(cur_path)
                        path.pop()
                        cur_path = "".join(path)
                        directories[cur_path] += cur_size
                    case _:
                        path.append(f"{line[1]}/")
                        cur_path = "".join(path)
            case "dir":
                directories.setdefault(f"{cur_path}{line[1]}/", 0)
            case _:
                directories[cur_path] += int(line[0])

    for _ in range(len(path) - 1):
        cur_size = directories.get(cur_path)
        path.pop()
        cur_path = "".join(path)
        directories[cur_path] += cur_size

    return directories


with open("input.txt", "r", encoding="utf-8") as file:
    lines: list[str] = file.readlines()
    parsed_input: list[list[str]] = [
        line.strip().replace("$ ", "").split() for line in lines
    ]
    dir_structure: dict[str, int] = build_directory_structure(parsed_input)

    print(sum([size for size in dir_structure.values() if size <= 100000]))

    used_space = dir_structure["/"]
    unused_space = 70000000 - used_space
    if unused_space < 30000000:
        print(
            min(
                [
                    size
                    for size in dir_structure.values()
                    if size + unused_space >= 30000000
                ]
            )
        )
