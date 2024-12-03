def main1():
    diff_sum = 0
    with open("input1-sorted.txt", "r") as i1:
        with open("input2-sorted.txt", "r") as i2:
            for _ in range(1000):
                line1 = int(i1.readline().strip())
                line2 = int(i2.readline().strip())
                diff = abs(line1 - line2)
                diff_sum += diff
    print(diff_sum)


def main2():
    from collections import Counter

    i1_counter = Counter()
    with open("input1-sorted.txt", "r") as i1:
        for line in i1:
            # Read and process each line
            line = int(line.strip())
            i1_counter.update([line])
    i2_counter = Counter()
    with open("input2-sorted.txt", "r") as i2:
        for line in i2:
            # Read and process each line
            line = int(line.strip())
            i2_counter.update([line])
    similarity_sum = 0
    for key in i1_counter.keys():
        similarity_sum += (key * i1_counter[key]) * i2_counter.get(key, 0)
    print(similarity_sum)


if __name__ == "__main__":
    main1()
    main2()
