class memoize(dict):
    def __init__(self,func):
        self.func=func
    def __call__(self,*key):
        return self[key]
    def __missing__(self,key):
        result = self[key] = self.func(*key)
        return result
def f(ids):
    return ids
        
fs=memoize(f)
print(fs(1))

def get(a:int,b:str)->str:
    return a*b
print(get(2,2))