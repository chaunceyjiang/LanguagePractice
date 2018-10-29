import time
from functools import wraps
from functools import lru_cache
registry = []

def register(func):
    print("running register(%s)" % func)
    registry.append(func)
    return func
@register
def f1():
    print("runnging f1()")

@register
def f2():
    print("runing f2()")

def f3():

    print("running f3()")

def main():
    print("running main()")
    print("registry -> ",registry)
    f1()
    f2()
    f3()
   
    print('*'*40 ,"Calling snooze(.123)")
    snooze(.123)
    print('*'*40 ,"Calling factorial(6)")
    print('10! = ',factorial(3))
    P()
def clock(func):
    def clocked(*args):
        t0=time.perf_counter()
        result = func(*args)
        elapsed = time.perf_counter() - t0
        name = func.__name__
        arg_str = ','.join(repr(arg) for arg in args)
        print('[%0.8fs] %s(%s) -> %r' % (elapsed,name,arg_str,result))
        return result
    return clocked
@clock
def snooze(seconds):
    time.sleep(seconds)

@lru_cache()
@clock
def factorial(n):
    return 1 if n< 2 else n * factorial(n-1)

def log(text):
    print("log --> ",text)
    def addlog(func):
        print("func before")
        func()
        print("func after")
        return func
    return addlog
@log("aaa")
def P():
    print("func")


if __name__=="__main__":
    main()