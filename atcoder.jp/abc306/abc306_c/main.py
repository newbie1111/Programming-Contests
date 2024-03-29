#!/usr/bin/env python3
# from typing import *

from collections import defaultdict, Counter
# from icecream import ic
# def solve(N: int, A: List[int]) -> List[str]:


def solve(N, A):
    dd = defaultdict(list)

    for i, a in enumerate(A, 1):
        dd[a].append(i)
    f = tuple((a[1], i) for i, a in dd.items())

    return [i for a, i in sorted(f)]


def solve2(N, A):
    counter = Counter()
    ans = []
    for i, a in enumerate(A):
        counter[a] += 1
        if counter[a] == 2:
            ans.append(a)

    return ans

# generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)


def main():
    import sys
    tokens = iter(sys.stdin.read().split())
    N = int(next(tokens))
    A = [None for _ in range(3 * N)]
    for i in range(3 * N):
        A[i] = int(next(tokens))
    assert next(tokens, None) is None
    ans = solve2(N, A)
    print(*[ans[i] for i in range(N)])


if __name__ == '__main__':
    main()
