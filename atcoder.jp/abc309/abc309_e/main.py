#!/usr/bin/env python3
# from typing import *

import sys
from collections import defaultdict

sys.setrecursionlimit(10 ** 8)

if not __debug__:
    from icecream import ic
    def debug(*x): ic(x)
else:
    def debug(*x): pass

# def solve(N: int, M: int, p: List[int], x: List[int], y: List[int]) -> int:


def solve(N, M, p, X, Y):
    CHILD = "child"
    CHECK = "check"
    person_data = {i+1: {CHILD: set(), CHECK: False} for i in range(N)}
    insurance = defaultdict(int)

    ans = 0

    for i, v in enumerate(p, 2):
        person_data[v][CHILD].add(i)

    for x, y in zip(X, Y):
        insurance[x] = max(insurance[x], y)

    def __solve(person_id, depth):
        if depth > 0:
            person_data[person_id][CHECK] = True

            for child_id in person_data[person_id][CHILD]:
                person_data[child_id][CHECK] = True
                debug(person_data[person_id], person_data[child_id])
                __solve(child_id, depth-1)

    debug(person_data, insurance)

    for k, v in insurance.items():
        __solve(k, v)

    debug(person_data)

    for v in person_data.values():
        ans += 1 if v[CHECK] else 0

    return ans


def solve2(N, M, P, X, Y):
    dp = [-1 for _ in range(N)]
    ans = 0

    # for x, y in zip(X, Y):
    #     dp[x] = max(dp[x], y)
    # debug(dp, P)
    # for i in range(1, N):
    #     dp[i] = max(dp[i], dp[P[i]]-1)

    for v in dp:
        ans += v >= 0

    return ans


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)


def main():
    import sys
    tokens = iter(sys.stdin.read().split())
    N = int(next(tokens))
    M = int(next(tokens))
    p = [None for _ in range(N - 1)]
    x = [None for _ in range(M)]
    y = [None for _ in range(M)]
    for i in range(N - 1):
        p[i] = int(next(tokens)) - 1
    for i in range(M):
        x[i] = int(next(tokens)) - 1
        y[i] = int(next(tokens))
    assert next(tokens, None) is None
    a = solve2(N, M, p, x, y)
    print(a)


if __name__ == '__main__':
    main()
