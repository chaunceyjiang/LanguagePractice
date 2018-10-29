#include <iostream>
#include <stack>
#include <string>
using namespace std;

int main()
{
    unsigned int a, b, c, d;
    char ch;
    while (cin >> a >> ch >> b >> ch >> c >> ch >> d)
    {
        cout << ((a << 24) | (b << 16) | (c << 8) | d) << endl;
        cin >> a;
        cout << ((a & 0xff000000) >> 24) << "." << ((a & 0x00ff0000) >> 16) << "." << ((a & 0x0000ff00) >> 8) << "." << (a & 0x000000ff) << endl;
    }
    return 0;
}