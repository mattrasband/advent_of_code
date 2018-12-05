#!/usr/bin/env python3
import difflib
import itertools
import pathlib
from collections import Counter
from typing import Iterable


def part1(iterable: Iterable[str]) -> int:
    """
    >>> part1(["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"])
    12
    """
    count_3 = count_2 = 0
    for line in iterable:
        counter = Counter(line)
        counter = {v: k for k, v in counter.items()}
        if 2 in counter:
            count_2 += 1
        if 3 in counter:
            count_3 += 1
    return count_3 * count_2

def part2(iterable: Iterable[str]) -> str:
    """
    >>> part2(["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"])
    'fgij'
    """
    for a, b in itertools.combinations(iterable, r=2):
        diff = list(difflib.ndiff(a.strip(), b.strip()))
        if sum(x.startswith("-") or x.startswith("+") for x in diff) == 2:
            return ''.join(x.strip() for x in diff if not (x.startswith("+") or x.startswith("-")))


if __name__ == "__main__":
    year_root = pathlib.Path(__file__).parent
    input_file = year_root / "day_02_input.txt"
    with open(input_file) as f:
        print("part 1:", part1(f))

    with open(input_file) as f:
        print("part 2:", part2(f))