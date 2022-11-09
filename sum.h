#include <stdio.h>
__attribute__((weak))
int sum(int c, int d)
{
  return c+d;
}
int delfunc(int c, int d);
