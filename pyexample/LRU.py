from collections import OrderedDict


class LRUCache(object):
    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.cache = OrderedDict()
        self.max_size = capacity
        self.size = 0

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        return self.cache.get(key) or -1

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        if key in self.cache:
            self.cache.pop(key)
        else:
            if self.size >= self.max_size:
                self.cache.popitem(last=True)
            self.size += 1
        self.cache[key] = value
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None