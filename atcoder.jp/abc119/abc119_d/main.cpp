#include <bits/stdc++.h>
#define REP(i, n) for (int i = 0; (i) < (int)(n); ++ (i))
#define REP3(i, m, n) for (int i = (m); (i) < (int)(n); ++ (i))
#define REP_R(i, n) for (int i = (int)(n) - 1; (i) >= 0; -- (i))
#define REP3R(i, m, n) for (int i = (int)(n) - 1; (i) >= (int)(m); -- (i))
#define ALL(x) ::std::begin(x), ::std::end(x)
using namespace std;


auto solve(int A, int B, int Q, const std::vector<int64_t> &s, const std::vector<int64_t> &t, const std::vector<int64_t> &x) {
    // TODO: edit here
}

// generated by oj-template v4.8.1 (https://github.com/online-judge-tools/template-generator)
int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    int A, B, Q;
    std::cin >> A;
    std::vector<int64_t> s(A);
    std::cin >> B;
    std::vector<int64_t> t(B);
    std::cin >> Q;
    std::vector<int64_t> x(Q);
    REP (i, A) {
        std::cin >> s[i];
    }
    REP (i, B) {
        std::cin >> t[i];
    }
    REP (i, Q) {
        std::cin >> x[i];
    }
    auto ans = solve(A, B, Q, s, t, x);
    REP (i, Q) {
        std::cout << a[i] << '\n';
    }
    return 0;
}