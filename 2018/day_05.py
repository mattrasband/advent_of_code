import pathlib
import string


def part1(iterable) -> int:
    content = ''.join(iterable)
    while True:
        start_len = len(content)
        for letter in string.ascii_lowercase:
            content = content.replace(f"{letter}{letter.upper()}", "")
            content = content.replace(f"{letter.upper()}{letter}", "")

        if len(content) == start_len:
            break

    return len(content)


def part2(iterable):
    iterable = ''.join(iterable)
    counts = {}

    for letter in string.ascii_lowercase:
        content = iterable.replace(letter.lower(), "").replace(letter.upper(), "")
        counts[letter] = part1(content)

    return min(counts.values())


if __name__ == "__main__":
    in_file = pathlib.Path(__file__).parent / "day_05_input.txt"
    with open(in_file) as f:
        print("part1:", part1(f))

    with open(in_file) as f:
        print("part2:", part2(f))