#!/usr/bin/env python3
import functools
import pathlib
import re
from collections import defaultdict


def parse_input(iterable):
    parse_re = re.compile(r'\[.*:(?P<minute>\d{2})\] (?:Guard #(?P<guard_num>\d+))?')
    guards = defaultdict(lambda: {
        "id": None,
        "mins": 0,  # total time sleeping
        "freq": [0] * 60,
    })

    current_guard = start_sleep = None
    for entry in sorted(iterable, key=lambda x: x[1:17]):
        match = parse_re.match(entry)
        if match.group("guard_num") is not None:
            current_guard = int(match.group("guard_num"))
            guards[current_guard]["id"] = current_guard
        elif "falls asleep" in entry:
            start_sleep = int(match.group("minute"))
        else:
            end_sleep = int(match.group("minute"))
            guards[current_guard]["mins"] += end_sleep - start_sleep
            for i in range(start_sleep, end_sleep):
                guards[current_guard]["freq"][i] += 1
    return guards


def part1(iterable):
    guards = parse_input(iterable)
    longest_sleeper = functools.reduce(lambda x, y: x if x["mins"] > y["mins"] else y, guards.values())
    sleepiest_minute = longest_sleeper["freq"].index(max(longest_sleeper["freq"]))
    return longest_sleeper["id"] * sleepiest_minute


def part2(iterable):
    guards = parse_input(iterable)
    most_common_sleeper = functools.reduce(lambda x, y: x if max(x["freq"]) > max(y["freq"]) else y, guards.values())
    return most_common_sleeper["id"] * most_common_sleeper["freq"].index(max(most_common_sleeper["freq"]))


if __name__ == "__main__":
    input_file = pathlib.Path(__file__).parent / "day_04_input.txt"
    with open(input_file) as f:
        print("part1:", part1(f))
    with open(input_file) as f:
        print("part1:", part2(f))
