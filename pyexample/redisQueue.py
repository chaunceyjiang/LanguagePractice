# coding: utf-8
import redis
class RedisQueue(object):
    def __init__(self, name, namespace='queue', **redis_kwargs):
       # redis的默认参数为：host='localhost', port=6379, db=0， 其中db为定义redis database的数量
       self.__db= redis.Redis(**redis_kwargs)
       self.key = '%s:%s' %(namespace, name)
    def qsize(self):
        return self.__db.llen(self.key)  # 返回队列里面list内元素的数量

    def put(self, item):
        self.__db.rpush(self.key, item)  # 添加新元素到队列最右方

    def get_wait(self, timeout=None):
        # 返回队列第一个元素，如果为空则等待至有元素被加入队列（超时时间阈值为timeout，如果为None则一直等待）
        item = self.__db.blpop(self.key, timeout=timeout)
        # if item:
        #     item = item[1]  # 返回值为一个tuple
        return item

    def get_nowait(self):
        # 直接返回队列第一个元素，如果队列为空返回的是None
        item = self.__db.lpop(self.key)  
        return item
    def get_rwait(self):
        item = self.__db.rpop(self.key)
        return item

    def set(self, key, value,expire=None):
        self.__db.set(key, value)
        if expire:
            self.__db.expire(key,expire)

    def get(self, key):
        return self.__db.get(key)
    
    def delete(self, key):
        self.__db.delete(key)

    def exists(self,key):
        return self.__db.exists(key)

    def setnx(self,key,value,expire=None):
        n = self.__db.setnx(key,value)
        if expire and n:
            self.__db.expire(key,expire)
        return n

class StrictRedis(object):

    def __init__(self,**redis_kwargs):
        self.__db= redis.StrictRedis(**redis_kwargs)
        self.pubsub = self.__db.pubsub()
        self.pubsub.psubscribe('__keyevent@1__:*')
        
    def get_messaget(self):
        return self.pubsub.get_message()