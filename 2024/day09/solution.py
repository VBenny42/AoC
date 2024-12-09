from itertools import takewhile

FREE_SPACE = "."


def convert(line: list) -> list:
    is_freespace = False
    id = 0
    diskmap = []
    for i in line:
        if is_freespace:
            for _ in range(i):
                diskmap.append(FREE_SPACE)
        else:
            for _ in range(i):
                diskmap.append(id)
            id += 1
        is_freespace = not is_freespace
    return diskmap


def convert_with_size(line: list) -> list:
    is_freespace = False
    id = 0
    diskmap = []
    for i in line:
        if is_freespace:
            diskmap.append((FREE_SPACE, i))
        else:
            diskmap.append((id, i))
            id += 1
        is_freespace = not is_freespace
    return diskmap


def make_contiguous1(diskmap: list) -> list:
    first_free_block = 0
    last_file_block = len(diskmap) - 1

    while first_free_block < last_file_block:
        # Find the first free block
        while (
            first_free_block < len(diskmap) and diskmap[first_free_block] != FREE_SPACE
        ):
            first_free_block += 1

        # Find the last file block
        while last_file_block >= 0 and diskmap[last_file_block] == FREE_SPACE:
            last_file_block -= 1

        # Swap if the indices are still valid
        if first_free_block < last_file_block:
            diskmap[first_free_block], diskmap[last_file_block] = (
                diskmap[last_file_block],
                diskmap[first_free_block],
            )

    return diskmap


def make_contiguous2(diskmap: list) -> list:
    files = list(reversed([x for x in (diskmap) if x[0] != FREE_SPACE]))
    for file in files:
        swap_file(diskmap, file)
    return diskmap


def swap_file(diskmap: list, file: tuple):
    file_index = diskmap.index(file)
    free_space_index = -1

    for i in range(len(diskmap)):
        if diskmap[i][0] == FREE_SPACE and diskmap[i][1] >= file[1]:
            free_space_index = i
            break

    if free_space_index == -1:
        return

    if free_space_index > file_index:
        return

    if diskmap[free_space_index][1] > file[1]:
        diskmap[free_space_index] = (FREE_SPACE, diskmap[free_space_index][1] - file[1])
        diskmap[file_index] = (FREE_SPACE, file[1])
        diskmap.insert(free_space_index, file)

    elif diskmap[free_space_index][1] == file[1]:
        diskmap[free_space_index], diskmap[file_index] = (
            diskmap[file_index],
            diskmap[free_space_index],
        )

    return


def print_diskmap(diskmap: list):
    representation = ""
    for x in diskmap:
        if x[0] == FREE_SPACE:
            representation += "." * x[1]
        else:
            representation += str(x[0]) * x[1]
    return representation


def checksum(diskmap: list):
    index = 0
    checksum = 0
    for x in diskmap:
        if x[0] != FREE_SPACE:
            for i in range(x[1]):
                checksum += (index + i) * x[0]
        index += x[1]
    return checksum


def main1():
    with open("input.txt", "r") as f:
        line = f.read().strip()
        int_line = [int(x) for x in line]
    diskmap = convert(int_line)
    contiguous_diskmap = make_contiguous1(diskmap)
    files_only = takewhile(lambda x: x != FREE_SPACE, contiguous_diskmap)
    checksum = sum(i * id for i, id in enumerate(files_only))
    print(f"LOGF: {checksum = }")


def main2():
    with open("input.txt", "r") as f:
        line = f.read().strip()
        int_line = [int(x) for x in line]
    diskmap = convert_with_size(int_line)
    contiguous_diskmap = make_contiguous2(diskmap)
    print(f"LOGF: checksum for contiguous files {checksum(contiguous_diskmap)}")


if __name__ == "__main__":
    main1()
    main2()
