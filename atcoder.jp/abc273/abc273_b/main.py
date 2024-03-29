#!/usr/bin/env python3
# from typing import *

from decimal import Decimal, ROUND_HALF_UP
import sys

# def solve(X: int, K: int) -> int:


def solve(X, K):
    ans = X

    for k in range(K):
        ans = int(Decimal(ans).quantize(
            Decimal(f"1E{k + 1}"), rounding=ROUND_HALF_UP))

    return ans


def solve2(X, K):
    x = X

    for k in range(K):
        x //= 10 ** k
        m = x % 10
        print(x, k, m, file=sys.stderr)
        if m < 5:
            x -= m
        else:
            x += (10-m)

        x *= 10 ** k
        print(x, "\n", file=sys.stderr)

    return x


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)


def main():
    X, K = map(int, input().split())
    # a = solve(X, K)
    a = solve2(X, K)
    print(a)


if __name__ == '__main__':
    main()
