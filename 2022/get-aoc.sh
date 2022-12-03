#!/usr/bin/env bash

mkdir -p $(date +'%d')
curl --cookie "session=$AOC_SESSION_COOKIE" https://adventofcode.com/$(date +'%Y')/day/$(date +'%-d'/input) > $(date +'%d')/input.txt
