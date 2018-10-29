A={"a":1,"b":2,"c":3}
C={"a":1,"b":1,"c":1}
B=["a","b","c"]
def sums(x):
    C[x]+=A[x]
map(sums,B)
print(C)
import itertools
import findarrary
