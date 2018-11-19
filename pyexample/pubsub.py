import time
from redis import StrictRedis
from redisQueue import RedisQueue
qq = RedisQueue('rq', host='192.168.1.119', port=6379, db=1)
redis = StrictRedis(host='192.168.1.119', port=6379, db=2)

pubsub = redis.pubsub()


def event_handlerx(msg):
    print ('Handler', qq.exists(msg['data']))
    qq.delete(msg['data'])
    print ('Handler', qq.exists(msg['data']))


pubsub.psubscribe(**{'__keyevent@1__:*': event_handlerx})

print('Starting message loop')
while True:
    message = pubsub.get_message()
    if message:
        print(message)
    else:
        time.sleep(4)
