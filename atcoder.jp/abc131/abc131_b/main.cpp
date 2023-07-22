#include <iostream>
#include <vector>
#include <cmath>

auto solve(int N, int L)
{
    std::vector<int> v(N);
    auto abs_min = 0;
    auto sum = 0;

    for (int i = 0; i < N; i++)
    {
        v[i] = L + (i + 1) - 1;

        if (std::abs(v[abs_min]) > std::abs(v[i]))
        {
            abs_min = i;
        }
    }

    for (int i = 0; i < N; i++)
    {
        std::cerr << v[i] << " ";
        if (abs_min != i)
        {
            sum += v[i];
        }
    }

    std::cerr << abs_min << std::endl;
    return sum;
}

int main()
{
    int N, L;
    std::cin >> N >> L;
    std::cout << solve(N, L) << std::endl;
}