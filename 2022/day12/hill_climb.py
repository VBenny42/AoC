with open("small_input.txt", "r") as file:
    height_map = [list(map(ord, [char for char in line.strip()])) for line in file]
