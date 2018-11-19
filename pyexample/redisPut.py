from redisQueue import RedisQueue,StrictRedis
import time,datetime,os,sys,redis
import msgpack
# qq=RedisQueue('rq',host='192.168.1.119',port=6379,db=2)
# q=RedisQueue('rq',host='192.168.1.119',port=6379,db=1)
qq=redis.Redis(host="192.168.1.119",db=5)
# qq.set("xxx",["222",2,3])
qq.lpush('lis1t',"1")
qq.lpush('lis1t',"2")
qq.lpush('lis1t',"2")
print type(qq.lrange('lis1t',0,-1))
# q.set("xxx","xxx",expire=3)

# qq.set("yyyy",["222",2,3])
# q.set("yyyy","yyyy",expire=6)
# print q.setnx("xxx","2222",expire=3)
# print q.setnx("xxx","2222",expire=3)
# print q.get("sss")
# # e=StrictRedis(host='192.168.1.119',port=6379,db=1)
# # while True:  
# #     message = e.get_messaget()
# #     if message:
# #         print message,"222222"
# #     else:
# #         time.sleep(0.01)

# import time  
# from redis import StrictRedis

# redisc = StrictRedis(host='192.168.1.119', port=6379,db=1)

# def event_handler(msg):
#     print ("thread",  qq.exists(msg['data']))
#     qq.delete(msg['data'])
#     print ("thread", qq.exists(msg['data']) )

# pubsubs = redisc.pubsub()  
# pubsubs.psubscribe(**{'__keyevent@1__:expired': event_handler})  
# thread = pubsubs.run_in_thread(sleep_time=0.01,daemon=True)

# print "daemon=True"
# # import time  
# # from redis import StrictRedis

# redis = StrictRedis(host='192.168.1.119', port=6379,db=1)

# pubsub = redis.pubsub()

# def event_handlerx(msg):  
#     print ('Handler', qq.exists(msg['data']))
#     qq.delete(msg['data'])
#     print ('Handler', qq.exists(msg['data']) )

# pubsub.psubscribe(**{'__keyevent@1__:*': event_handlerx})

# print('Starting message loop')  
# while True:  
#     message = pubsub.get_message()
#     if message:
#         print(message)
#     else:
#         time.sleep(0.01)