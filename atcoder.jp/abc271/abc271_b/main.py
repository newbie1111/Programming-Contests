#!/usr/bin/env python3
# from typing import *

from collections import namedtuple

Query = namedtuple('Query', ['s', 't'])


def solve(ary, query) -> list:
    res = []
    for q in query:
        res.append(ary[q.s][q.t])
    return res


def main():
    n, q = map(int, input().split())
    ary, query = [[]], []
    for i in range(n):
        ary.append(list(map(int, input().split())))
    for i in range(q):
        query.append(Query(*map(int, input().split())))
    print(*solve(ary, query), sep="\n")


if __name__ == '__main__':
    main()
