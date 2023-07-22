#!/usr/bin/env python3
# from typing import *
import pprint
import itertools
from collections import deque


def is_edge(x, y, left, right, bottom, top):
    neighbor = [(1, 0), (0, 1), (-1, 0), (0, -1)]
    for dx, dy in neighbor:
        if left <= x + dx < right and top <= y + dy < bottom:
            continue
        else:
            return True
    else:
        return False

# def solve(N: str, A: List[List[str]]) -> List[List[str]]:


def counterclockwise(y, x, left, right, bottom, top):
    dx, dy = 0, 0

    if y == top:
        if x == left:
            dx = 1
        else:
            dy = -1
    if y == (bottom-1):
        if x == (right-1):
            dx = -1
        else:
            dy = 1

    if x == left:
        if not y == (bottom-1):
            dx = 1
    if x == (right-1):
        if not y == top:
            dy = -1

    return y + dx, x + dy


def solve(N, A):
    d = {(i, j): counterclockwise(i, j, 0, N, N, 0)
         for i, j in itertools.product(range(N), repeat=2)}
    ans = []
    for i in range(N):
        line = []
        for j in range(N):
            x, y = d[i, j]
            line.append(A[x][y])
        ans.append("".join(line))
    return "\n".join(ans)


def main():
    N = int(input())
    A = []
    for _ in range(N):
        A.append(input())
    ans = solve(N, A)
    print(ans)


if __name__ == '__main__':
    main()
