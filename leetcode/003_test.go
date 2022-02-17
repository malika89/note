package leetcode

import (
	"testing"
)


//最长不重复子串
//abcdbef=>cdb
//滑动窗口：左指针k表示起始位置，右指针表示当前最大未重复位置rk;左指针移动一步；移动右指针
func MaxStr(str string) int {
	setDict :=map[byte]int{}
	rk,ans :=-1,0 //右指针，初始-1表示没移动
	for i:=0;i<len(str);i++{ //i为左指针
		if i!=0 {
			delete(setDict,str[i-1])
		}
		for rk<len(str)-1 && setDict[str[rk+1]] ==0 {
			setDict[str[rk+1]] ++
			rk +=1
		}
		if rk-i+1 >ans {
			ans = rk-i+1
		}
	}
	return ans
}

//双指针
func MaxStr2(str string) int {
	l,r :=0,0
	res :=0
	window :=map[string]int{}
	for r<len(str) {
		c :=string(str[r])
		r +=1
		window[c] +=1
		for window[c]>1{
			//2 6
			d :=string(str[l])
			l+=1
			window[d] -=1
		}
		if r-l>res {
			res = r-l
		}
	}
	return res
}

//dp解题
func MaxStr3(str string) int {
	l,r :=0,0
	res :=0
	window :=map[string]int{}
	for r<len(str) {
		c :=string(str[r])
		r +=1
		window[c] +=1
		for window[c]>1{
			//2 6
			d :=string(str[l])
			l+=1
			window[d] -=1
		}
		if r-l>res {
			res = r-l
		}
	}
	return res
}

func TestMaxStr(t *testing.T) {
	str :="abdcodefjbjicwulkxchjk"
	t.Log(MaxStr(str))
	t.Log(MaxStr2(str))
}
