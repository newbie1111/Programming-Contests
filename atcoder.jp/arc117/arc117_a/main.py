#!/usr/bin/env python3
# from typing import *

import sys

# def solve(A: int, B: int) -> List[str]:


def solve(A, B):
    p, n = list(range(1, A+1)), list(range(-1, -(B+1), -1))
    print(p, n, file=sys.stderr)

    if A < B:
        tail = sum(n[len(p):])
        p[-1] += -tail
    elif A > B:
        tail = sum(p[len(n):])
        n[-1] += -tail

    print(sum(p+n), file=sys.stderr)
    return p + n


# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
def main():
    A, B = map(int, input().split())
    E = solve(A, B)
    print(*[E[i] for i in range(A + B)])


if __name__ == '__main__':
    main()