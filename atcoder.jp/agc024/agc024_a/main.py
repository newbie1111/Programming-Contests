#!/usr/bin/env python3
# from typing import *


# def solve(A: int, B: int, C: int, K: int) -> int:


def solve(A, B, C, K):
    if K == 0:
        return A-B
    else:
        base = K * (A+B+C)
        takahashi = base - A
        nakahashi = base - B

        if abs(takahashi - nakahashi) >= 10**18:
            return "Unfair"
        else:
            return takahashi - nakahashi


def solve2(A, B, C, K):
    if K % 2 == 0:
        return A - B
    else:
        return B - A


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    A, B, C, K = map(int, input().split())
    a = solve2(A, B, C, K)
    print(a)


if __name__ == '__main__':
    main()
