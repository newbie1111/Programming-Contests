#!/usr/bin/env python3
# from typing import *


# def solve(N: int, B: List[int]) -> int:
def solve(N, B):
    A = [0] * 100
    A[0] = B[0]
    A[-1] = B[-1]

    for i in range(1, N-1):
        A[i] = B[i]
        A[i+1] = max(B[i], B[i-1])
    return sum(A[N:])


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    import sys
    tokens = iter(sys.stdin.read().split())
    N = int(next(tokens))
    B = [None for _ in range(N - 1)]
    for i in range(N - 1):
        B[i] = int(next(tokens))
    assert next(tokens, None) is None
    a = solve(N, B)
    print(a)


if __name__ == '__main__':
    main()