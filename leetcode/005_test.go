package leetcode

import (
	"fmt"
	"testing"
)

//最长回文字符串 ababa acad
// 双指针 left right


//动态规划
func MaxhuiStr(s string) string {
	lenS :=len(s)
	dp :=make([][]bool,lenS)
	res := ""
	//初始化
	for i:=0;i<lenS;i++ {
		dp[i] = make([]bool,lenS)
	}

	for i:=0;i<lenS;i++ {
		for j:=0;j<i;j++ {
			dp[j][i] =(s[i]==s[j]) &&(i-j<2 || dp[j+1][i-1])
			if dp[j][i] && i-j>len(res){
				res = s[j:i+1]
			}
		}
		dp[i][i] = true
	}

	return res
}

func TestDemo(t *testing.T) {
	s :="ababa"
	fmt.Println(MaxhuiStr(s))
}