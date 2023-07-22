#include <iostream>
#include <vector>
#include <string>

auto solve(int n, std::vector<int> a)
{
    std::vector<int> ans(n);

    for (int i = 0; i < n; i++)
    {
        ans[a[i] - 1] = i + 1;
    }

    return ans;
}

int main()
{
    int n;
    std::cin >> n;
    std::vector<int> a(n);

    for (int i = 0; i < n; i++)
    {
        std::cin >> a[i];
    }

    auto ans = solve(n, a);

    for (int i = 0; i < n; i++)
    {
        std::cout << ans[i];

        if (i != n - 1)
        {
            std::cout << " ";
        }
        else
        {
            std::cout << std::endl;
        }
    }
}