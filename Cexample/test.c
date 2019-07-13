#include <stdio.h>
int main()
{
    int x, y = 0, i, a[10], sum = 0, min, max;
    for (i = 0; i < 10; i++)
    {
        scanf("%d", &a[i]);
        sum += a[i];
    }
    min = max = a[0];
    for (i = 0; i < 10; i++)
    {
        if (max < a[i])
            max = a[i];
        if (min > a[i])
            min = a[i];
        
    }
    printf("%f\n", (sum - min - max) / 8.0);
    return 0;
}