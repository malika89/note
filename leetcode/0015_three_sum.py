# coding:utf-8


# 三数之和为0
def solution(lst, target):
    res = []
    for i, v in enumerate(lst):
        left = i + 1
        right = len(lst) - 1
        while left <= right and right>i:
            if lst[right] + lst[left] + v == target:
                tmp = [v, lst[left], lst[right]]
                tmp.sort()
                if tmp not in res:
                    res.append(tmp)
                break
            if left == right:
                left += 1
                right = len(lst) - 1
            right -= 1
    return res


# 排序后再进行双指针比较
def solution2(lst, target):
    res = []
    lst.sort()
    for i, v in enumerate(lst):
        if i>=1 and lst[i] == lst[i-1]:
            continue
        if lst[i] >0:
            break
        left = i+1
        right = len(lst)-1
        while left < right and right > i:
            sum_value = lst[left] + lst[right] + v
            if sum_value > target:
                right -= 1
            elif sum_value < target:
                left += 1
            else:
                res.append([v, lst[left], lst[right]])
                while left != right and lst[left] == lst[left + 1]:
                    left += 1
                while left != right and lst[right] == lst[right - 1]:
                    right -= 1
                left += 1
                right -= 1
    return res


if __name__ == '__main__':
    print(solution([-1, 0, 1, 2, -1, -4], 0))
    print(solution2([-1, 0, 1, 2, -1, -4], 0))