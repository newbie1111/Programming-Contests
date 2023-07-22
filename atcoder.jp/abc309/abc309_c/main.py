#!/usr/bin/env python3
# from typing import *

if not __debug__:
    from icecream import ic
    def debug(*x): ic(*x)
else:
    def debug(*x): pass
# def solve(N: int, K: int, a: List[int], b: List[int]) -> int:


def solve(N, K, A, B):
    tup = sorted((a, b) for a, b in zip(A, B))
    medicine = sum(B)

    if medicine <= K:
        return 1
    else:
        for t in tup:
            medicine -= t[1]
            if medicine <= K:
                return t[0] + 1
        else:
            return tup[-1][0] + 1


def solve2(N, K, A, B):
    tup = sorted((a, b) for a, b in zip(A, B))
    medicine = sum(B)
    ans = 1

    if medicine > K:
        for t in tup:
            medicine -= t[1]
            if medicine <= K:
                ans = t[0] + 1
                break

    return ans

# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)


def main():
    N, K = map(int, input().split())
    a = [None for _ in range(N)]
    b = [None for _ in range(N)]
    for i in range(N):
        a[i], b[i] = map(int, input().split())
    a1 = solve2(N, K, a, b)
    print(a1)


if __name__ == '__main__':
    main()