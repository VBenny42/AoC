import heapq
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


def convert_with_heaps(line: list) -> tuple[list, list]:
    is_freespace = False
    id = 0
    diskmap = []
    heaps = [[] for _ in range(10)]
    for i in line:
        if is_freespace:
            # Push index of free space to the heap of the size of the free space
            heapq.heappush(heaps[i], len(diskmap))
            for _ in range(i):
                diskmap.append(FREE_SPACE)
        else:
            for _ in range(i):
                diskmap.append(id)
            id += 1
        is_freespace = not is_freespace
    return diskmap, heaps


def make_contiguous2_heap(diskmap: list, heaps: list) -> list:
    index = len(diskmap) - 1
    while index >= 0:
        if diskmap[index] == FREE_SPACE:
            index -= 1
            continue

        id = diskmap[index]
        file_width = 0
        # Get the width of the file
        while index >= 0 and diskmap[index] == id:
            file_width += 1
            index -= 1

        best_width = -1
        smallest_index = len(diskmap)
        # Find the leftmost index of free space that can fit the file
        for width in range(file_width, 10):
            if heaps[width]:
                if smallest_index > heaps[width][0]:
                    smallest_index = heaps[width][0]
                    best_width = width

        if smallest_index == len(diskmap):
            continue
        if smallest_index > index:
            continue

        # Remove the smallest index from the heap
        # In-place swap the file with the free space
        heapq.heappop(heaps[best_width])
        for j in range(file_width):
            diskmap[smallest_index + j] = id
            diskmap[index + j + 1] = FREE_SPACE
        # Push the new smaller free space to the heap
        heapq.heappush(heaps[best_width - file_width], smallest_index + file_width)

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

    for i in range(file_index):
        if diskmap[i][0] == FREE_SPACE and diskmap[i][1] >= file[1]:
            free_space_index = i
            break

    if free_space_index == -1:
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
    print(f"ANSWER1: {checksum = }")


def main2():
    with open("input.txt", "r") as f:
        line = f.read().strip()
        int_line = [int(x) for x in line]
    diskmap = convert_with_size(int_line)
    contiguous_diskmap = make_contiguous2(diskmap)
    print(f"ANSWER2: checksum for contiguous files {checksum(contiguous_diskmap)}")


def main3():
    with open("input.txt", "r") as f:
        line = f.read().strip()
        int_line = [int(x) for x in line]
    diskmap, heaps = convert_with_heaps(int_line)
    contiguous_diskmap = make_contiguous2_heap(diskmap, heaps)
    checksum = sum(
        i * id for i, id in enumerate(contiguous_diskmap) if id != FREE_SPACE
    )
    print(f"ANSWER2: {checksum = }")


if __name__ == "__main__":
    main1()
    # main2()
    main3()
