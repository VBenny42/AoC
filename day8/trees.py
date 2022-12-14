# split into two lists for column and row, check if max of each is tree for visibility
def get_adjacent_cells(x: int, y: int, trees: list[list[int]]) -> list[list[int]]:
    row = [tree for tree in trees[y]]
    col = [tree[x] for tree in trees]
    return (
        row[:x],
        row[x + 1 :],
        col[:y],
        col[y + 1 :],
    )


def is_visible(x: int, y: int, trees: list[list[int]]) -> bool:
    tree_height = trees[y][x]

    west_side, east_side, north_side, south_side = get_adjacent_cells(x, y, trees)
    if (not west_side) or (not east_side) or (not north_side) or (not south_side):
        return True
    for height in {
        max(west_side),
        max(east_side),
        max(north_side),
        max(south_side),
    }:
        if tree_height > height:
            return True
    return False


def takewhile_inclusive(predicate, it):
    for x in it:
        if predicate(x):
            yield x
        else:
            yield x
            break


def get_score(x: int, y: int, trees: list[list[int]]) -> int:
    height = trees[y][x]
    west_side, east_side, north_side, south_side = get_adjacent_cells(x, y, trees)

    def get_side_score(tree_side: list[int]) -> int:
        return len(list(takewhile_inclusive(lambda tree: tree < height, tree_side)))

    west_score = get_side_score(reversed(west_side))
    east_score = get_side_score(east_side)
    north_score = get_side_score(reversed(north_side))
    south_score = get_side_score(south_side)
    return west_score * east_score * north_score * south_score


with open("input.txt", "r", encoding="utf-8") as file:
    trees: list[list[int]] = [[int(char) for char in line.strip()] for line in file]
    # assuming matrix is square
    dimension: int = len(trees)
    count: int = 0
    max_score: int = 0
    score: int = 0
    for y in range(dimension):
        for x in range(dimension):
            count += int(is_visible(x, y, trees))
            if (score := get_score(x, y, trees)) > max_score:
                max_score = score
    print(count)
    print(max_score)
