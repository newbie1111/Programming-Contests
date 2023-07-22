#include <iostream>

auto solve(int n)
{
    return (n % 10 < 7) ? 100 * (n / 10) + 15 * (n % 10) : 100 * (n / 10 + 1);
}

int main()
{
    int n;
    std::cin >> n;
    std::cout << solve(n) << std::endl;
}