#!/usr/bin/env python3
# from typing import *


# def solve(N: int, H: List[int]) -> int:
def solve(N, H):
    pass  # TODO: edit here


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    import sys
    tokens = iter(sys.stdin.read().split())
    N = int(next(tokens))
    H = [None for _ in range(N)]
    for i in range(N):
        H[i] = int(next(tokens))
    assert next(tokens, None) is None
    a = solve(N, H)
    print(a)


if __name__ == '__main__':
    main()
