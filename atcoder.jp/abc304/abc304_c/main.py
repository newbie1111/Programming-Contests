#!/usr/bin/env python3
# from typing import *

YES = 'Yes'
NO = 'No'


# def solve(N: int, D: int, X: List[int], Y: List[int]) -> List[str]:
def solve(N, D, X, Y):
    pass  # TODO: edit here


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    N, D = map(int, input().split())
    X = [None for _ in range(N)]
    Y = [None for _ in range(N)]
    for i in range(N):
        X[i], Y[i] = map(int, input().split())
    ans = solve(N, D, X, Y)
    for i in range(N):
        print(ans[i])


if __name__ == '__main__':
    main()
