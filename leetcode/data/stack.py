#!/usr/bin/python
# coding:utf-8

"""
stack:lifo pop push方法。 python中list(单向队列)是stack的一种应用
linked list : insert delete ；工作方式：存储数据和其他的工作节点
deque是Python中stack和queue的通用形式，也就是既能当做栈使用，又能当做双向队列。拥有list所有方法，且额外拥有：
   appendleft(x) 头部添加元素
   extendleft(iterable) 头部添加多个元素
   popleft() 头部返回并删除
   rotate(n=1) 旋转
   maxlen 最大空间，如果是无边界的，返回None
"""


class Stack:
    def __init__(self):
        self.items = []

    def pop(self):
        return self.items.pop()

    def push(self,value):
        self.items.append(value)

    def peek(self):
        self.items[len(self.items)-1]

    def size(self):
        return len(self.items)


# case2: 进制转换
def Hexconvert(number,base):
    dights = [str(i) for i in range(10) ] + [chr(i) for i in range(65,72)]
    remstack = []
    while number >0:
        rem = number%base
        remstack.append(dights[rem])
        number //=base
    res = ""
    while remstack !=[]:
        res +=remstack.pop()

    return res


# case3: 中缀表达式转换为后缀表达式

# 递归法 爬楼梯：有n层，每次可以爬1-2楼
def solutions(n,lst):
    if lst[n] >0:
        return lst[n]
    if n <= 3:
        lst[n] = n
    else:
        lst[n] = solutions(n-1,lst)+solutions(n-2,lst)
    return lst[n]


# 回溯法 爬楼梯：有n层，每次可以爬1-2楼
def solutions2(n):
    pass


if __name__ == '__main__':
    # print(Hexconvert(22, 2))
    print(solutions(6,[0]*7))