'''
二分法查找
'''

import bisect,sys
HAYSTACK = [1, 4, 5, 6, 8, 12, 15, 20, 21, 23, 23, 26, 29, 30]
NEEDLES = [0, 1, 2, 5, 8, 10, 22, 23, 29, 30, 31]
print(bisect.bisect(HAYSTACK,22))#查找

print(id(HAYSTACK[0]),HAYSTACK[0])

bisect.insort(HAYSTACK,1)#插入

print(id(HAYSTACK[0]),HAYSTACK[1])
print(id(HAYSTACK[0]),HAYSTACK[1],id(1))