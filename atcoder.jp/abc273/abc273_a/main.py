#!/usr/bin/env python3
# from typing import *


# def solve(N: int) -> int:
def solve(N):
    return 1 if N == 0 else N * solve(N - 1)


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    N = int(input())
    a = solve(N)
    print(a)


if __name__ == '__main__':
    main()
