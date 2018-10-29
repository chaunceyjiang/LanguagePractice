EXAMPLE = ["a", "b", "b", "b", "c", "c"]
from collections import defaultdict
CharClass = defaultdict(list)
i = 0
for k, v in enumerate(EXAMPLE):
    if v not in CharClass.keys():
        i = 0
    CharClass[v].append({"Oindex": k, "index": i})
    i += 1
print(CharClass)
