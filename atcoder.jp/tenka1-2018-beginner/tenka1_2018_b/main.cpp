#include <iostream>

auto solve(int a, int b, int k)
{
    int taka = a, aoki = b;
    for (int i = 0; i < k; i++)
    {
        if (i % 2 == 0)
        {
            taka /= 2;
            aoki += taka;
        }
        else
        {
            aoki /= 2;
            taka += aoki;
        }
    }

    return std::to_string(taka) + " " + std::to_string(aoki);
}

int main()
{
    int a, b, k;
    std::cin >> a >> b >> k;
    std::cout << solve(a, b, k) << std::endl;
}