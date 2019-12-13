import json
import time

# from pytz import utc
# from apscheduler.executors.pool import ThreadPoolExecutor, ProcessPoolExecutor
# from apscheduler.schedulers.background import BackgroundScheduler
# from apscheduler.jobstores.sqlalchemy import SQLAlchemyJobStore
#
# scheduler = BackgroundScheduler()
# executors = {
#     'default': ThreadPoolExecutor(20),
#     'processpool': ProcessPoolExecutor(5)
# }
#
# scheduler = BackgroundScheduler(executors=executors, timezone=utc)
# scheduler.start()
# def p():
#     with open("./a.txt","a") as f:
#         f.write(str(time.time()))
# scheduler.add_job(p,'interval',minutes=1,id="dsdsdsd")
#
#
#
#
# while True:
#     time.sleep(10)

import stomp
from stomp import *
c = Connection([('localhost', 61613)])


class MyListener(ConnectionListener):
    def on_before_message(self, headers, body):
        """
        :param dict headers:
        :param body:
        """
        print('on_before_message %s %s' % (headers, body))
        return headers, body

    def on_message(self, headers, body):
        """
        :param dict headers:
        :param body:
        """
        print('on_message %s %s' % (headers, type(json.loads(body))))


c.set_listener('', MyListener())
c.start()
c.connect('admin', 'admin', wait=True)


c.subscribe('/topic/simo-event-MODULE_ONLINE', '1')

time.sleep(100)