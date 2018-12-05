#!/usr/bin/env python3
import itertools
import pathlib


def part1(iterable) -> int:
    """
    >>> part1(["+1", "+1", "+1"])
    3
    >>> part1(["+1", "+1", "-2"])
    0
    >>> part1(["-1", "-2", "-3"])
    -6
    """
    return sum(int(x) for x in iterable)


def part2(iterable) -> int:
    """
    >>> part2([1, -2, 3, 1])
    2
    >>> part2(["+1", "-1"])
    0
    >>> part2(["+3", "+3", "+4", "-2", "-4"])
    10
    >>> part2(["-6", "3", "8", "5", "-6"])
    5
    >>> part2(["7", "7", "-2", "-7", "-4"])
    14
    """
    current_position = 0
    positions = set([current_position])
    for item in itertools.cycle(iterable):
        current_position += int(item)
        if current_position in positions:
            return current_position
        positions.add(current_position)


if __name__ == "__main__":
    path = pathlib.Path(__file__).parent
    input_file = path / "day_01_input.txt"

    with open(input_file) as f:
        print("part 1:", part1(f))

    with open(input_file) as f:
        print("part 2:", part2(f))
