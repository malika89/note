

"""
2n-2个字符串作为一个整体,row=idx%(2n-2)为行数
if row>n,则row=2n-2-row

"""

import datetime
from functools import wraps


def print_time(func):
    @wraps(func)
    def wrapper(*args,**kwargs):
        start_time = datetime.datetime.now()
        try:
            res = func(*args, **kwargs)
        except Exception as e:
            print(e)
            res = "not found"
        print(f"{func.__qualname__}程序耗时:{(datetime.datetime.now()-start_time).microseconds}")
        # return (datetime.datetime.now()-start_time).microseconds
        return res
    return wrapper

class Solution2(object):

    @print_time
    def convert_n_times(self, data, rows, n):
        for i in range(n):
            self.convert(data, rows)

    # @print_time
    def convert(self, strings, rows):
        if rows >= len(strings):
            return strings
        results = [""]*rows
        for i in range(len(strings)):
            row = i % (2*rows-2)
            if row >= rows:
                row = 2 * rows - 2 - row
            results[row] += strings[i]
        return "".join(results)


class Solution(object):

    @print_time
    def convert_n_times(self, data, rows, n):
        for i in range(n):
            self.convert(data, rows)

    # @print_time
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows <= 1:
            return s
        res = ['' for i in range(numRows)]  # 记录各行的字符串
        nrow = 0  # 行号
        d = False  # 方向：False代表向上，True代表向下
        # 从左到右读取Z字形字符串
        for i in s:
            res[nrow] += i  # 同一行的字符join起来
            # 到达首行或者尾行时转向
            if nrow == 0 or nrow == numRows - 1:
                d = (not d)
            step = 1 if d else -1  # 行号变化的方向
            nrow += step
        # 读出转换后的Z字形字符串
        return ''.join(res)


if __name__ == '__main__':
    # begin
    import random

    H = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
    strings = random.sample(H, 50) + random.sample(H,50)
    random.shuffle(strings)
    strings = "".join(strings)
    print("input strings to be converted:::",strings)
    # print(Solution().convert(strings, 5))
    # print(Solution2().convert(strings, 5))

    # 比较两种方法耗时
    Solution().convert_n_times(strings, 100, 50)
    Solution2().convert_n_times(strings, 100, 50)


