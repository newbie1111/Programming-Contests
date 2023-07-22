#include <iostream>

int solve(int a, int b)
{
    return (a - 1) * (b - 1);
}

int main()
{
    int a, b;
    std::cin >> a >> b;
    std::cout << solve(a, b) << std::endl;
}