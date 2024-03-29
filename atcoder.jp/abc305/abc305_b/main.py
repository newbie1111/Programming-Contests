#!/usr/bin/env python3
# from typing import *

def cusum(itr):
    res = [itr[0]]
    for v in itr[1:]:
        res.append(res[-1] + v)
    return res

# def solve(p: str, q: str) -> int:


def solve(p, q):
    d = {chr(c): i for c, i in zip(
        range(ord("A"), ord("G")+1), cusum([0, 3, 1, 4, 1, 5, 9]))}
    return abs(d[p]-d[q])


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    p, q = input().split()
    a = solve(p, q)
    print(a)


if __name__ == '__main__':
    main()
