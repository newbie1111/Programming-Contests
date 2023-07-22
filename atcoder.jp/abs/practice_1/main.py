#!/usr/bin/env python3
# from typing import *


# def solve(n: int, a: List[int]) -> int:
def solve(a, b, c, s):
    return f"{sum([a,b,c])} {s}"


def main():
    a = int(input())
    b, c = map(int, input().split())
    s = input()
    ans = solve(a, b, c, s)
    print(ans)


if __name__ == '__main__':
    main()
