#include <bits/stdc++.h>
#define REP(i, n) for (int i = 0; (i) < (int)(n); ++ (i))
#define REP3(i, m, n) for (int i = (m); (i) < (int)(n); ++ (i))
#define REP_R(i, n) for (int i = (int)(n) - 1; (i) >= 0; -- (i))
#define REP3R(i, m, n) for (int i = (int)(n) - 1; (i) >= (int)(m); -- (i))
#define ALL(x) ::std::begin(x), ::std::end(x)
using namespace std;


auto solve(int64_t N, int M, const std::vector<int64_t> &s, const std::vector<int64_t> &t) {
    // TODO: edit here
}

// generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    int64_t N;
    int M;
    std::cin >> N >> M;
    std::vector<int64_t> s(M), t(M);
    REP (i, M) {
        std::cin >> s[i] >> t[i];
    }
    auto ans = solve(N, M, s, t);
    REP (i, M) {
        std::cout << ans[i] << '\n';
    }
    return 0;
}