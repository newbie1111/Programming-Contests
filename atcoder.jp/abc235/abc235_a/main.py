#!/usr/bin/env python3
# from typing import *
import itertools

# def solve(abc: int) -> int:


def solve(abc):
    return 111 * sum([int(c) for c in abc])


def solve2(abc):
    a, b, c = abc[0], abc[1], abc[2]
    return int(a+b+c) + int(b+c+a) + int(c+a+b)


def main():
    abc = input()
    # a = solve(abc)
    a = solve2(abc)
    print(a)


if __name__ == '__main__':
    main()
