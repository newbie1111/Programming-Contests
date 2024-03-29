#!/usr/bin/env python3
# from typing import *


# def solve(A: int, B: int, C: int) -> str:
def solve(A, B, C):
    # TLE
    if A**C < B**C:
        return "<"
    elif A**C == B**C:
        return "="
    else:
        return ">"


def solve2(A, B, C):
    ans = "="
    a, b = (abs(A), abs(B)) if C % 2 == 0 else (A, B)

    if a < b:
        ans = "<"
    elif a == b:
        ans = "="
    else:
        ans = ">"

    return ans


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    A, B, C = map(int, input().split())
    # a = solve(A, B, C)
    a = solve2(A, B, C)
    print(a)


if __name__ == '__main__':
    main()
