# from icecream import ic
from collections import deque, defaultdict
import sys

sys.setrecursionlimit(10000000)
"""


class Edge:

    def __init__(self, to, weight=1) -> None:
        self.to = to
        self.weight = weight

    def __str__(self) -> str:
        return f"to : {self.to} | weight : {self.weight} | id : {id(self)}"

    def __repr__(self) -> str:
        return self.__str__()


class Graph:

    def __init__(self) -> None:
        self.edges = dict()

    def add_vertex(self, key):
        self.edges[key] = list()

    def add_edge(self, key, to, weight=1):
        self.edges[key].append(Edge(to, weight))

    # def show(self):
    #     for key in self.edges.keys():
    #         ic(key, self.edges[key])

    def dfs(self, start, end=None):
        seen = defaultdict(bool)
        stack = []
        prev = dict()

        stack.append(start)

        while stack:
            head = stack.pop()

            # ic("now", head, id(self.edges[head]))

            for edge in self.edges[head]:

                if not seen[edge.to]:
                    stack.append(edge.to)
                    prev[edge.to] = head

            seen[head] = True

            if head and head == end:
                break

        return stack, seen, prev

    def bfs(self, start, end=None, seen=None):
        seen = defaultdict(bool) if not seen else seen
        deq = deque()
        prev = dict()

        deq.append(start)

        while deq:
            head = deq.popleft()

            # ic("now", head, id(self.edges[head]))

            for edge in self.edges[head]:

                if not seen[edge.to]:
                    deq.append(edge.to)
                    prev[edge.to] = head

            seen[head] = True

            if end and head == end:
                break

        return deq, seen, prev

    def get_path(self, end, prev):
        path = [end]

        while path[-1] in prev:
            path.append(prev[path[-1]])
        return list(reversed(path))


g = Graph()
seen = defaultdict(bool)
r, c = map(int, input().split())
sy, sx = map(int, input().split())
gy, gx = map(int, input().split())

s = []
for _ in range(r):
    s.append(input())

for row in range(r):
    for col in range(c):
        g.add_vertex((row, col))

        if s[row][col] == "#":
            seen[(row, col)] = True

for row in range(r):
    for col in range(c):
        left, right = (0, 1, 0, -1), (-1, 0, 1, 0)

        for i, j in zip(left, right):
            if 0 < row + i < r and 0 < col + j < c:
                g.add_edge((row, col), (row + i, col + j))
                # ic(f"connected {row , col} -> {row + i, col + j}")

res = g.bfs((sy - 1, sx - 1), (gy - 1, gx - 1), seen)
# ic(res)
# ic(g.get_path((gy - 1, gx - 1), res[2]))
print(len(g.get_path((gy - 1, gx - 1), res[2])) - 1)
"""


class Node:
    """

    ノード情報の管理

    Attributes:
        index (int): ノード番号
        nears (list): 隣接リスト（隣接するノード番号を格納）
        arrival (bool): 探索済フラグ（到達済の場合Trueが返される）
    """

    # def __init__(self, index):
    #     self.index = index
    #     self.nears = []
    #     self.parent = -1
    #     self.arrival = False

    def __init__(self, row, col):
        self.row = row
        self.col = col
        self.nears = []
        self.parent = -1
        self.arrival = False

    def __repr__(self):
        return f"(index:{self.index}, nears:{self.nears}, arrival:{self.arrival})"


r, c = map(int, input().split())

sy, sx = map(int, input().split())
sy, sx = sy - 1, sx - 1

gy, gx = map(int, input().split())
gy, gx = gy - 1, gx - 1

s = []
for _ in range(r):
    s.append(input())

nodes = [[Node(i, j) for j in range(c)] for i in range(r)]
adjacents = [
    (0, 1),
    (0, -1),
    (1, 0),
    (-1, 0),
]

for i in range(r):
    for j in range(c):
        for dx, dy in adjacents:
            nx, ny = i + dx, j + dy

            if 0 <= nx < r and 0 <= ny < c:
                nodes[i][j].nears.append([nx, ny])

        if s[i][j] == "#":
            nodes[i][j].arrival = True

queue = deque([nodes[sy][sx]])

while queue:
    node = queue.popleft()
    nears = node.nears

    for i, j in nears:
        if not nodes[i][j].arrival:
            queue.append(nodes[i][j])
            nodes[i][j].parent = node.row, node.col
            nodes[i][j].arrival = True

ans = [nodes[gy][gx]]
while ans[-1] != -1:
    ans.append(ans[-1])
print(ans)
