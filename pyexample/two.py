
a= 1
def f():
    global a
    if a!=2:
        a= 2
def c():
    f()


def d():
    f()