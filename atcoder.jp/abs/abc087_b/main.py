#!/usr/bin/env python3
# from typing import *
import itertools

# def solve(A: int, B: int, C: int, X: int) -> int:


def solve(A, B, C, X):
    return sum([1 for a, b, c in itertools.product(range(A+1), range(B+1), range(C+1)) if sum([a*500, b*100, c*50]) == X])


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    A = int(input())
    B = int(input())
    C = int(input())
    X = int(input())
    a = solve(A, B, C, X)
    print(a)


if __name__ == '__main__':
    main()
