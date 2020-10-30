package main

import (
	"fmt"
	_ "strings"
)

func main() {
	//arr:=[...]int{1,3,5,8}
	//sum:=0
	////求和
	//for _,v:=range arr{
	//sum+=v
	//}
	//fmt.Println(sum)
	//
	//s1:=[]string{"上海","北京","广州"}
	////s1[3]="深圳"
	//s1 = append(s1, "深圳")
	//fmt.Println(s1)
	//s1:=make([]int,5,10)
	//for i:=0;i<10;i++{
	//	s1=append(s1[:6],i)
	//}
	//fmt.Println(s1,cap(s1))
	//
	//m:=7
	//b:=&m
	//println(m,b)

	//var a *int
	//println(a)
	//var b=new(int)
	//println(b)
	//println(*b)
	//*b=100
	//println(*b)

	//var m=make(map[string]int,3)
	//m["一个"]=1
	//m["二个"]=2
	//m["三个"]=3
	//println(m)
	//println(m["一个"])
	//value,ok:=m["五个"]
	//if !ok{
	//	println("没有")
	//}else {
	//	println(value)
	//}//
	//元素类型为map的切片

	//var s=make([]map[int]string,0,10)
	//写一个程序，统计一个字符串中每个单词出现的次数
	//str:="how do you do you"
	//sstr:=strings.Split(str," ")
	//mstr:=make(map[string]int,10)
	//for i:=0;i<len(sstr);i++{
	////	for j:=i+1;j<len(sstr);j++{
	////		if sstr[i]==sstr[j]{
	////			mstr[sstr[i]]+=1
	////		}else {
	////			mstr[sstr[i]]=1
	////		}
	////
	////	}
	////
	////}
	//for _,w:=range sstr{
	//	//如果map中没有，加入；如果有，加一
	//	if _,ok:=mstr[w];!ok{
	//		mstr[w]=1
	//	}else {
	//		mstr[w]++
	//	}
	//}
	//for k,v:=range mstr{
	//	fmt.Println(k,v)
	//}
	//xstr:="上海自来水来自海上"
	//sm:=make([]rune,0,10)
	////fmt.Println(xstr[0])
	////fmt.Println(len(xstr))
	////把字符串转换成字符切片
	//for _,c:=range xstr{
	//	sm = append(sm, c)
	//	fmt.Println(sm)
	//}
	//copy(sm,xstr...)
	//计算数量
	//fmt.Println(len(sm))
	//从中间开始向两边数，对称相等
	deferdemo()
	fmt.Println(f1()) //5
	fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
}
func deferdemo() {
	fmt.Println("第一句")
	defer fmt.Println("第二句")
	defer fmt.Println("第三句")
	fmt.Println("第四句")
}

//1返回值赋值
//2运行defer
//3ret指令
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//查找单词是否出现过
//func isHas(key string,m map[string]int)(ret bool){
//
//	return true
//}
