import time,random
def f():
    ts = time.time()
    while True:
        if time.time() - ts>1:
            print "f"
        ts = time.time()
        time.sleep(0.6)
        yield
def ff():
    while True:
        time.sleep(random.random())
        yield

fx = f()
ffx = ff()
while True:
    fx.next()
    print "-----"
    ffx.next()