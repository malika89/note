#!/usr/bin/python
# coding:utf-8
"https://github.com/wawacode/python_miaosha/blob/main/app.py"

from redis import Redis
import uuid
import time
from threading import Thread,Lock
import pymysql
import random


class RedisX(Redis):

    def __init__(self, host='localhost', port=6379, db=0, password=None, username=None, decode_responses=True,**kwargs):
        super().__init__(host, port, db, password, username=username,decode_responses=decode_responses, **kwargs)

    def set(self, k, value, expire=0):
        super().set(k, value)

    # 获取一个锁
    def acquire_lock(self, lock_name, acquire_time=10):
        identifier = str(uuid.uuid4())
        end_time = time.time()+acquire_time
        while time.time() < end_time:
            if self.setnx(lock_name, identifier):
                return identifier
            time.sleep(0.001)
        return False

    # 释放一个锁
    def release_lock(self, lock_name, identifier):
        if self.get(lock_name) == identifier:
            self.delete(lock_name)
            return True
        return False


class ThreadWithReturnValue(Thread):
    def __init__(self, group=None, target=None, name=None,args=(), kwargs={}, Verbose=None):
        Thread.__init__(self, group, target, name, args, kwargs)
        self._return = None

    def run(self):
        if self._target is not None:
            self._return = self._target(*self._args, **self._kwargs)

    def join(self, *args):
        Thread.join(self, *args)
        return self._return


def deal_with_mysql(m):
    start_time = time.time()
    lock = Lock()
    threads = []

    def deal_stock_db(n, user):
        lock.acquire()
        db = pymysql.connect(host="127.0.0.1", port=3306, user="root", password="testpwd123", database="test")
        cursor = db.cursor()
        db.begin()
        sql = 'select no from resource where name="ticket" for update'
        cursor.execute(sql)
        ret = cursor.fetchone()
        if ret[0] < n:
            cursor.close()
            db.close()
            lock.release()
            return "no stock"
        try:
            update_sql = "update resource set no={} where name='ticket'".format(ret[0] - n)
            cursor.execute(update_sql)
            sql = 'insert into records(user,no) values("{}",{})'.format(user, n)
            cursor.execute(sql)
            db.commit()
            lock.release()
        except Exception as e:
            print(e)
            pass
    for i in range(m):
        per_ticket = random.randint(1, 2)
        t = ThreadWithReturnValue(target=deal_stock_db, args=(per_ticket,f'user_{i}'))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()

    print("票已售完:mysql 完成耗时：", time.time()-start_time)


def deal_with_redis(m):
    start_time = time.time()
    threads = []

    def deal_stock_redis(m, user):
        client = RedisX(host="127.0.0.1", port=6379, db=1, password="testpwd123")
        with client.pipeline() as p:
            while True:
                lock = client.acquire_lock("lock:stock")
                if not lock:
                    continue
                p.watch(lock)
                try:
                    p.multi()
                    count = int(client.get("stock"))
                    if count>0 and count>=m:
                       p.set("stock", count-m)
                       p.execute()
                    break
                except Exception:
                    pass
                finally:
                    res = client.release_lock("lock:stock",lock)
                    if not res:
                        client.delete("lock:stock")

    for i in range(0, m):
        per_ticket = random.randint(1, 2)
        t = ThreadWithReturnValue(target=deal_stock_redis, args=(per_ticket, f'user_{i}'))
        t.start()
    print("票已售完:redis 完成耗时：", time.time()-start_time)


if __name__ == '__main__':
    totalNum = 500
    # 100个人抢30张票
    deal_with_mysql(1000)
    deal_with_redis(1000)



