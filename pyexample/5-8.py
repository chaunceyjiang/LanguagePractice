import random
class BingoCage:
    def __init__(self,items):
        self._items = list(items)
        random.shuffle(self._items)
    def pick(self):
        try:
            return self._items.pop()
        except IndexError :
            raise LookupError("pick from empty BingoCage")
    def __call__(self):
        return self.pick()


#参数注释，返回值注释
def clip(text:str,max_len:'int > 0'=80)->str:
    return 80
