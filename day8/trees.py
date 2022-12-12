DIM_X = 99
DIM_Y = 99


def get_adjacent_cells(cell: int, trees: list[int]) -> list[list[int]]:
    x = cell // DIM_X
    y = cell % DIM_Y
    return [
        [
            tree
            for tree in range(DIM_X * DIM_Y)
            if (tree // DIM_X == x) and tree != cell
        ],
        [
            tree
            # for tree in range(y, DIMENSIONX * DIMENSIONY, DIMENSIONY)
            for tree in range(DIM_X * DIM_Y)
            if (tree % DIM_Y == y) and tree != cell
        ],
    ]


def is_on_edge(cell: int) -> bool:
    # on west  edge if cell %  99 == 0
    # on east  edge if cell %  99 == 98 or 99-1
    # on south edge if cell // 99 == 98 or 99-1
    # on north edge if cell // 99 == 0
    return (cell % DIM_Y in {0, DIM_Y - 1}) or (cell // DIM_X in {0, DIM_X - 1})

# split into two lists for columb and route, check if max of each is tree for visibility

def is_visible(cell: int, trees: list[int], adjacent_trees: list[list[int]]):
    height_to_check = trees[cell]
    if is_on_edge(cell):
        # cell is on edge and is therefore visible
        return True
    # get edge cells
    n_s_edges = filter(is_on_edge, adjacent_trees[0])
    e_w_edges = filter(is_on_edge, adjacent_trees[1])


with open("input.txt", "r", encoding="utf-8") as file:
    trees: list[int] = [
        tree
        for line in file
        for tree in list(map(int, [char for char in line.strip()]))
    ]
    print(get_adjacent_cells(0, trees))
    print(is_visible(0, trees, get_adjacent_cells(0, trees)))
