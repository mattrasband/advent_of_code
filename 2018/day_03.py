import pathlib
import re
from collections import defaultdict


def part1(iterable, size = 1_000) -> int:
    """
    >>> part1(["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"], 8)
    4
    >>> part1(["#123 @ 3,2: 5x4"], 11)
    0
    """
    fabric = [[0] * size for _ in range(size)]
    parser_re = re.compile(
        r"^#(?P<claim>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<width>\d+)x(?P<height>\d+)$"
    )

    for line in iterable:
        match = parser_re.search(line.strip())
        _, x, y, width, height = (int(x) for x in match.groups())

        for y_offset in range(y, y + height):
            for x_offset in range(x, x + width):
                fabric[y_offset][x_offset] += 1

    return sum(sum(1 for x in y if x > 1) for y in fabric)


def part2(iterable, size = 1_000):
    """
    >>> part2(["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"], 8)
    3
    """
    fabric = [[0] * size for _ in range(size)]
    parser_re = re.compile(
        r"^#(?P<claim>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<width>\d+)x(?P<height>\d+)$"
    )

    matches = []
    for line in iterable:
        match = parser_re.search(line.strip())
        claim, x, y, width, height = (int(x) for x in match.groups())
        matches.append((claim, x, y, width, height))

        for y_offset in range(y, y + height):
            for x_offset in range(x, x + width):
                fabric[y_offset][x_offset] += 1

    for claim, x, y, width, height in matches:
        for y_offset in range(y, y + height):
            for x_offset in range(x, x + width):
                if fabric[y_offset][x_offset] == 1:
                    continue
                break
            else:
                continue
            break
        else:
            return claim


if __name__ == "__main__":
    input_ = pathlib.Path(__file__).parent / "day_03_input.txt"

    with open(input_) as f:
        print("part1:", part1(f))

    with open(input_) as f:
        print("part2:", part2(f))