'''
一个简单的二维向量类
'''
from math import hypot
class Vector:
    def __init__(self,x=0,y=0):
        self.x = x
        self.y = y
    def __repr__(self):
    #把一个对象用字符串的形式表达出来以便辨认__repr__ 和 __str__ 的区别在于，
    # 后者是在 str() 函数被使用，或是在用 print 函数打印一个对象的时候才被调用的，
    # 并且它返回的字符串对终端用户更友好
        return 'Vector(%r,%r)' % (self.x,self.y)
    def __abs__(self):
        return hypot(self.x,self.y)  #hypot(x,y) == sqrt(x*x + y*y)
    def __bool__(self):
        return bool(abs(self))    #abs() 会自动调用 __abs__方法
    def __add__(self, other):     # + 运算符
        x = self.x + other.x
        y = self.y + other.y
        return Vector(x,y)
    def __mul__(self, other):    # * 运算符
        return Vector(self.x * other,self.y * other)

''' 解释器输出
>>> Vector(3,5)+Vector(6,8)
Vector(9,13)
>>> abs(Vector(3,4))
5.0
>>> Vector(3,4)
Vector(3,4)
>>> Vector(3,4)*5
Vector(15,20)
>>> abs(Vector(3,4)*5)
25.0
'''