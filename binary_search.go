package leetcode

import "testing"

func BinarySearch(nums []int,target int) bool {
	if len(nums) <=1 {
		return nums[0]==target
	}
	n :=len(nums)/2
	if nums[n] >target {
		return BinarySearch(nums[:n],target)
	}else if nums[n] == target {
		return true
	}else {
		return BinarySearch(nums[n+1:],target)
	}

}

//非递归方法 1,2,3,4,5
func BinarySearch2(nums []int,target int) bool {
	if len(nums) <=1 {
		return nums[0]==target
	}
	low,high :=0,len(nums)-1
	for low<=high {
		mid :=(high+low)/2
		if nums[mid] ==target{
			return true
		}else if nums[mid] <target{
			low = mid +1
		}else {
			high = mid -1
		}

	}
	return false

}

func TestBinarySearch(t *testing.T) {
	nums :=[]int{1,2,3,4,5,6,7,8,9,12,14,18,30,32}
	target :=9
	t.Log(BinarySearch(nums,target))
	t.Log(BinarySearch2(nums,target))
}