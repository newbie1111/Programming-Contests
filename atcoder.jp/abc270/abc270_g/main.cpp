#include <bits/stdc++.h>
#define REP(i, n) for (int i = 0; (i) < (int)(n); ++(i))
#define REP3(i, m, n) for (int i = (m); (i) < (int)(n); ++(i))
#define REP_R(i, n) for (int i = (int)(n)-1; (i) >= 0; --(i))
#define REP3R(i, m, n) for (int i = (int)(n)-1; (i) >= (int)(m); --(i))
#define ALL(x) ::std::begin(x), ::std::end(x)
using namespace std;

auto solve(int a, const std::vector<int64_t> &b, const std::vector<int64_t> &c, const std::vector<int64_t> &d, const std::vector<int64_t> &e, const std::vector<int64_t> &f) {
    // TODO: edit here
}

// generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    int a;
    std::cin >> a;
    std::vector<int64_t> b(a), c(a), d(a), e(a), f(a);
    REP (i, a) { std::cin >> b[i] >> c[i] >> d[i] >> e[i] >> f[i]; }
    auto ans = solve(a, b, c, d, e, f);
    std::cout << m << '\n';
    std::cout << n << '\n';
    std::cout << o << '\n';
    return 0;
}
