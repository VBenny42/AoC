from typing import Counter


def find_secret_number(n: int) -> int:
    mod = 16777216
    n = ((n * 64) ^ n) % mod
    n = ((n // 32) ^ n) % mod
    n = ((n * 2048) ^ n) % mod
    return n


def get_prices_and_changes(secret_number: int) -> tuple[list[int], list[int]]:
    last = secret_number
    prices = []
    changes = []
    for _ in range(2000):
        secret_number = find_secret_number(secret_number)
        # Only need to keep last digits
        prices.append(secret_number % 10)
        changes.append(secret_number % 10 - last % 10)
        last = secret_number
    return prices, changes


def get_banana_sequences(
    prices: list[int], changes: list[int]
) -> dict[tuple[int, ...], int]:
    sequences = {}
    for i in range(3, len(changes)):
        seq = tuple(changes[i - 3 : i + 1])
        if sum(seq) > 0 and seq not in sequences:
            # Add price of the last banana in the sequence
            sequences[seq] = prices[i]
    return sequences


def main1():
    with open("input.txt", "r", encoding="utf-8") as f:
        secret_numbers = map(int, f.read().split())

    sum_secret_numbers = 0
    for secret_number in secret_numbers:
        for _ in range(2000):
            secret_number = find_secret_number(secret_number)
        sum_secret_numbers += secret_number

    print(f"ANSWER1: { sum_secret_numbers }")


def main2():
    with open("input.txt", "r", encoding="utf-8") as f:
        secret_numbers = map(int, f.read().split())

    banana_sequences = Counter()
    for secret_number in secret_numbers:
        prices, changes = get_prices_and_changes(secret_number)
        sequences = get_banana_sequences(prices, changes)
        banana_sequences.update(sequences)

    max_sequence = max(banana_sequences.values())
    print(f"ANSWER2: { max_sequence }")


if __name__ == "__main__":
    main1()
    main2()
