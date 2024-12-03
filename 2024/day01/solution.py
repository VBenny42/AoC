def main1():
    with open("input1-sorted.txt", "rb") as i1:
        i1_lines = i1.readlines()
        i1_lines = map(int, i1_lines)
    with open("input2-sorted.txt", "rb") as i2:
        i2_lines = i2.readlines()
        i2_lines = map(int, i2_lines)

    diff_sum = sum(abs(line1 - line2) for line1, line2 in zip(i1_lines, i2_lines))
    print(diff_sum)


def main2():
    from collections import Counter

    with open("input1-sorted.txt", "rb") as i1:
        lines = i1.readlines()
        lines = map(int, lines)
        i1_counter = Counter(lines)
    with open("input2-sorted.txt", "rb") as i2:
        lines = i2.readlines()
        lines = map(int, lines)
        i2_counter = Counter(lines)

    similarity_sum = 0
    for key in i1_counter.keys():
        similarity_sum += (key * i1_counter[key]) * i2_counter.get(key, 0)
    print(similarity_sum)


if __name__ == "__main__":
    main1()
    main2()
