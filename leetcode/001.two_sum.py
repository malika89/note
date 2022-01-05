# coding:utf-8

import datetime
from functools import wraps
import random

"""
一个数组，还有一个目标数target，让我们找到两个数字，使其和为targe
仅有一个解
数据比较大的时候双端指针solution方法效率更高
"""

# 方法1，遍历数组存储在map里面，target-值判断是否存在
# 方法2：排序后进行遍历

def generate_data(n):
    z = random.choices(range(0, 1000000), k=n)
    target = z[random.randint(0,n-1)] + z[random.randint(0,n-1)]
    return z,target


def print_time(func):
    @wraps(func)
    def wrapper(*args,**kwargs):
        start_time = datetime.datetime.now()
        try:
            res = func(*args, **kwargs)
        except Exception as e:
            res = "not found"
        #print(f"{func.__qualname__}程序耗时:{(datetime.datetime.now()-start_time).microseconds}")
        # return res
        return (datetime.datetime.now()-start_time).microseconds
    return wrapper


class Solution2(object):

    @print_time
    def two_sum(self, lst, target):
        map_dict = dict()
        for idx, v in enumerate(lst):
            if target - v not in map_dict:
                map_dict[v] = idx
            else:
                return target - v, v
        return "not found"


class Solution(object):

    # def twoSum(self, nums, target):
    #     # hash 2
    #     hash_nums = {}
    #     for index, num in enumerate(nums):
    #         another = target - num
    #         try:
    #             hash_nums[another]
    #             return [hash_nums[another], index]
    #         except KeyError:
    #             hash_nums[num] = index

    @print_time
    def two_sum(self, nums, target):
        # two point
        nums_index = [(v, index) for index, v in enumerate(nums)]
        nums_index.sort()
        begin, end = 0, len(nums) - 1
        while begin < end:
            curr = nums_index[begin][0] + nums_index[end][0]
            if curr == target:
                return [nums_index[begin][0], nums_index[end][0]]
            elif curr < target:
                begin += 1
            else:
                end -= 1
        return "not found"


if __name__ == '__main__':
    # begin
    score = {"solution":0,"solution2":0}

    for i in range(300):
        data = generate_data(5000)
        # print("生成的数据target：",data[1])
        if Solution().two_sum(data[0], data[1]) >Solution2().two_sum(data[0],data[1]):
            score["solution"] +=1
        else:
            score["solution2"]+=1

    print(score)

