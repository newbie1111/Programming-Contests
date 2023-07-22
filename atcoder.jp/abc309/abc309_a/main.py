#!/usr/bin/env python3
# from typing import *

YES = 'Yes'
NO = 'No'


# def solve(A: int, B: int) -> str:
def solve(A, B):
    corr = {(i, i+1) for i in range(9) if i % 3 != 0}
    return YES if (A, B) in corr else NO


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    A, B = map(int, input().split())
    a = solve(A, B)
    print(a)


if __name__ == '__main__':
    main()