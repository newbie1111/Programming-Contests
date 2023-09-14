#!/usr/bin/env python3
# from typing import *
from icecream import ic

# def solve(N: int, M: int, A: List[int], B: List[int]) -> Tuple[List[str], List[str]]:


def solve(N, M, A, B):
    print(N, M, A, B)


def main():
    N, M = map(int, input().split())
    A = []
    B = []
    Upland = []

    for i in range(2):
        if i == 0:
            A = list(map(int, input().split()))
        else:
            B = list(map(int, input().split()))

    ans = solve(N, M, A, B)

    ic(Upland)
    print(ans)


if __name__ == '__main__':
    main()
