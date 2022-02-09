package main

import (
	"fmt"
	"strings"
	"testing"
)

//栈的应用(顶端栈、尾端栈)：先进后出特性

//stack 构建--version1

type Stack []interface{}

func (s Stack ) Push(item interface{}) Stack  {
	s =append(s,item)
	return s
}

func (s Stack ) Pop() (Stack,interface{})  {
	l :=len(s)
	if l <=0 {
		return s,0
	}
	return s[:l-1],s[l-1]
}

func (s Stack ) Len() int  {
	return len(s)
}

func (s Stack ) Peek() interface{}  {
	if s.Len() <=0 {
		return 0
	}
	return s[s.Len()-1]
}


func TestStack(t *testing.T) {
	s :=make(Stack,0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	fmt.Println("the top value:",s.Peek())
	s, v :=s.Pop()
	fmt.Printf("poped item:%v,rest stack values:%v",v,s)

	fmt.Println("the top value:",s.Peek())
}

//stack version2 加入length 和链

type (
	Stack2 struct {
		top    *stackNode
		length int
	}
	stackNode struct {
		pre *stackNode
		value interface{}
	}
)

func NewStack() *Stack2 {
	return &Stack2{
		top:    nil,
		length: 0,
	}
}

func (s *Stack2) Push(item interface{}) {
	s.top = &stackNode{
		pre:   s.top,
		value: item,
	}
	s.length +=1
}

func (s *Stack2) Pop() interface{} {
	if s.length ==0 {
		return nil
	}
	n :=s.top
	s.top = n.pre
	s.length -=1
	return n.value
}

func (s *Stack2) Len() int {
	return s.length
}

func (s *Stack2) Peek() interface{} {
	if s.length ==0 {
		return nil
	}
	return s.top.value
}

func TestStack2(t *testing.T)  {
	stackNode :=NewStack()
	stackNode.Push(1)
	stackNode.Push(3)
	stackNode.Push(5)

	fmt.Println("the top value:",stackNode.Peek())

	fmt.Printf("poped item:%v \n",stackNode.Pop())

	fmt.Println("the top value:",stackNode.Peek())

}
//case1 : 括号匹配 example: ({}) true
/*==============================================================*/
func bracketMatch(str string) bool {
	left :="({["
	right := ")}]"
	//matched :=true

	stack :=NewStack()
	for _,s := range str {
		strS :=string(s)
		if strings.Contains(left,strS) {
			stack.Push(strS)
		}else {
			z :=stack.Pop()
			if  strings.Index(left,z.(string)) != strings.Index(right,strS) {
				return false
			}
		}
	}
	if stack.Len() ==0 {
		return true
	}
	return false

}

func TestBracketMatch(t *testing.T) {
	testStr :="{}([])"
	t.Log(bracketMatch(testStr))

	testStr ="{}([)]"
	t.Log(bracketMatch(testStr))
}

//case2 : 进制转换 n禁止转换(位运算符二进制)
/*==============================================================*/



//case3: 中缀表达式转为后缀表达式
/*==============================================================*/