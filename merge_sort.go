package main

import "fmt"

//归并排序

func MergeSort(s []int) []int {
	if len(s)<=1 {
		return s
	}
	mid :=len(s)/2
	left := MergeSort(s[:mid])
	right := MergeSort(s[mid:])
	return Merge(left,right)
}

func Merge(s1,s2 []int) []int {
	i:=0
	j:=0
	res :=make([]int,0)
	for i<len(s1) && j<len(s2) {
		if  s1[i]<s2[j] {
			res = append(res,s1[i])
			i ++
		}else {
			res = append(res,s2[j])
			j ++
		}
	}
	if i<len(s1)-1 {
		res = append(res,s1[i:]...)
	}
	if j<len(s2)-1 {
		res = append(res,s2[j:]...)
	}
	return res
}

func main() {
	s :=[]int{1,7,3,5,8,4}
	fmt.Println(MergeSort(s))
}