#!/usr/bin/env python3
# from typing import *

# from icecream import ic

# def solve(N: int, X: List[int], Y: List[int]) -> int:


def solve(N, X, Y):
    # (おいしさ, 状態)
    dp = [[0, 0], [0, 0]] * (N+1)

    for i, xy in enumerate(zip(X, Y), 1):
        x, y = xy

        # 食べる場合
        if x:
            # 毒入りの皿を出される => お腹を壊す =>  食べてお腹を壊す
            dp[i][1] = dp[i-1][0] + y
        else:
            # 解毒剤入り => そのまま食べる => 壊さないまま食べる or 食べて治す
            dp[i][0] = max(dp[i-1][0], dp[i-1][1]) + y

        # 食べない
        dp[i][1] = max(dp[i][1], dp[i-1][1])
        dp[i][0] = max(dp[i][0], dp[i-1][0])

    return max(dp[N][0], dp[N][1])


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    N = int(input())
    X = [None for _ in range(N)]
    Y = [None for _ in range(N)]
    for i in range(N):
        X[i], Y[i] = map(int, input().split())
    a = solve(N, X, Y)
    print(a)


if __name__ == '__main__':
    main()
