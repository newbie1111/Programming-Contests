#!/usr/bin/env python3
# from typing import *


# def solve(N: int) -> int:
def solve(N):
    return 100 * (N // 10) + 15 * (N % 10) if N % 10 <= 6 else 100 * (N // 10 + 1)


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    N = int(input())
    a = solve(N)
    print(a)


if __name__ == '__main__':
    main()
