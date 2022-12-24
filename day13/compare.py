from itertools import zip_longest

with open("small_input.txt", "r") as file:
    pair_list = []
    pair = []
    for line in file:
        if line.strip() == "":
            pair_list.append(pair)
            pair = []
            continue
        pair.append(eval(line.strip()))
    pair_list.append(pair)

# from pprint import pprint


# def pair_evaluator(pair: list) -> bool:
#     left, right = pair
#     # print(left, right)
#     if isinstance(left, list) and isinstance(right, list):
#         prev = True
#         for l, r in zip(left, right):
#             # if isinstance(l, list) and isinstance(r, list):
#             # return pair_evaluator([l, r])
#             prev = pair_evaluator([l, r])
#             if not prev:
#                 return False
#             # if pair_evaluator([l, r]):
#             # return True
#         # if len(left) > len(right):
#         #     # return pair_evaluator(list(zip(left, right))[-1])
#         return len(left) <= len(right)  # or prev
#     if isinstance(left, int) and isinstance(right, int):
#         # print("in")
#         # print(left <= right)
#         return left <= right
#     if isinstance(left, int) or isinstance(right, int):
#         return pair_evaluator(
#             [[left], right] if isinstance(left, int) else [left, [right]]
#         )


def pair_eval_retry(pair):
    left, right = pair
    for l, r in zip_longest(left, right, fillvalue=None):
        if l == None:
            return True
        if r == None:
            return False
        if isinstance(l, int) and isinstance(r, int):
            if l > r:
                return False
            if l < r:
                return True
        else:
            if isinstance(r, int):
                r = [r]
            if isinstance(l, int):
                l = [l]

            ret = pair_eval_retry([l, r])
            if ret in [True, False]:
                return ret


print(list(map(pair_eval_retry, pair_list)))
print(sum([i for i, pair in enumerate(pair_list, 1) if pair_eval_retry(pair)]))
