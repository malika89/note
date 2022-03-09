package leetcode

import (
	"testing"
)

//最适合抢劫银行的日子
//[3,2,1,1,2,3]

func BestDay(n []int,time int) []int {
	var res []int
	if time ==0 {
		for i:=range n {
			res = append(res,i)
		}
	}
	if len(n) <=time {
		return res
	}
	start ,mid,end :=0,time,2*time
	for start <=end && end<len(n){
		if n[start] >=n[mid] &&n[start]<n[start-1] && n[mid] <=n[end] && n[end]>n[end-1]{
			end --
		} else{
			mid ++
			end ++
		}
		start ++
		if start ==end {
			res = append(res,mid)
			mid ++
			end = mid+1
		}
	}
	return res
}

//DP思想；dp[start][mid] = 递增1 dp[mid][end] =-1 递减
// dp[i-1]是，则dp[i] = dp[i-1] &&(n[i] >n[i-1] && n[time+i]<n[time+i-1])
// dp[i-1] 否，则dp[i-1+time:] 否;

func BestDay2(n []int,time int) []int {
	var res []int
	lenN := len(n)
	left := make([]int,lenN)
	right := make([]int,lenN)
	for i:=1;i<lenN;i++ {
		if n[i] <= n[i-1] {
			left[i] = left[i-1] + 1
		}
		if n[lenN-i-1] <= n[lenN-i] {
			right[lenN-i-1] = right[lenN-i] + 1
		}
		if left[i] >= time && right[i] >= time {
			res = append(res,i)
		}
	}
	return res
}

func TestRobo(t *testing.T) {
	//res :=[]int{5,3,3,4,5,6,2} // 5
	//time :=2
	//t.Log(BestDay(res,time))
	//t.Log(BestDay2(res,time))
	//
	//res2 :=[]int{1,1,1,1,1}
	//time2:=0
	//t.Log(BestDay(res2,time2))
	//
	res3 :=[]int{1,2,3,4,5,6}
	time3:=2
	t.Log(BestDay2(res3,time3))
}