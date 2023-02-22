#!/usr/bin/env python3
# from typing import *
import numpy

YES = 'Yes'
NO = 'No'


# def solve(a: int, b: int, c: List[int], d: List[List[int]]) -> str:
def solve(n, m, x):
    ans = numpy.zeros((n, n), numpy.int8)
    for x_val in x:
        for i in range(len(x_val)):
            for j in range(i, len(x_val)):
                ans[x_val[i] - 1][x_val[j] - 1] = 1

    return YES if numpy.sum(ans) == numpy.sum(numpy.tri(n, dtype=numpy.int8)) else NO


def main():
    n, m = map(int, input().split())
    x = [list(map(int, input().split()))[1:] for i in range(m)]
    print(solve(n, m, x))


if __name__ == '__main__':
    main()
