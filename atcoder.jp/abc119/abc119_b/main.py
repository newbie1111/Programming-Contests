#!/usr/bin/env python3
# from typing import *


# def solve(N: int, x: List[float], u: List[str]) -> float:
def solve(N, x, u):
    pass  # TODO: edit here


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    import sys
    tokens = iter(sys.stdin.read().split())
    N = int(next(tokens))
    x = [None for _ in range(N)]
    u = [None for _ in range(N)]
    for i in range(N):
        x[i] = float(next(tokens))
        u[i] = next(tokens)
    assert next(tokens, None) is None
    a = solve(N, x, u)
    print(a)


if __name__ == '__main__':
    main()
