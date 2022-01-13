import redis_lock
from redis import Redis
import random
from threading import Thread
import time


def order(m):
    conn = Redis(host="127.0.0.1", port=6379, db=1, password="test123")
    with redis_lock.Lock(conn, "stock") as l:
        now_stock = conn.get("stock")
        if now_stock >= m:
            conn.set("stock", now_stock-m)
            print("order placed:",m)


if __name__ == '__main__':
    start_time = time.time()
    for i in range(0, 30):
        per_ticket = random.randint(1, 2)
        t = Thread(target=order, args=(per_ticket,))
        t.start()
    print("票已售完:redis 完成耗时：", time.time()-start_time)