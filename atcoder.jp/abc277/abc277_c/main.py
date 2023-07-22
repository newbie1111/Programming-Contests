#!/usr/bin/env python3
# from typing import *
from collections import deque, defaultdict


class Node():

    def __init__(self, number: int) -> None:
        self.number = number
        self.nears = []
        self.visited = False


# def solve(N: int, A: List[int], B: List[int]) -> int:
def solve(N, A, B):
    ans = 1

    # graph init
    nodes = {i: Node(i) for i in set(A + B + [1])}

    for a, b in zip(A, B):
        nodes[a].nears.append(b)
        nodes[b].nears.append(a)

    # BFS
    que = deque([1])
    while que:
        current_id = que.popleft()
        nodes[current_id].visited = True
        for nears_id in nodes[current_id].nears:
            if not nodes[nears_id].visited:
                que.append(nears_id)
                ans = max(ans, nears_id)

    return ans


def solve2(N, A, B):
    # graph init
    nodes = defaultdict(list)

    for a, b in zip(A, B):
        nodes[a].append(b)
        nodes[b].append(a)

    # BFS
    que = deque([1])
    visited = {1}
    while que:
        for nears_id in nodes[que.popleft()]:
            if nears_id not in visited:
                que.append(nears_id)
                visited.add(nears_id)

    return max(visited)


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    N = int(input())
    A = [None for _ in range(N)]
    B = [None for _ in range(N)]
    for i in range(N):
        A[i], B[i] = map(int, input().split())
    a = solve2(N, A, B)
    print(a)


if __name__ == '__main__':
    main()