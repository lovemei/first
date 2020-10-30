package main

import "fmt"

func main() {
	//f1(f2)
	//ret:=f3(f2,5,8)
	//f1(ret)

	//addfunc,subfunc:=calc(9)
	//fmt.Println(addfunc(5))
	//fmt.Println(subfunc(10))
	//fmt.Println(subfunc(10))

}
func f1(f func()) {
	fmt.Println("这是函数f1")
	f()
}
func f2(x int, y int) int {
	fmt.Println("这是函数f2")
	return x + y
}

//闭包就是 函数+外部变量的引用
func f3(f func(x int, y int) int, m int, n int) func() {
	tmp := func() {
		fmt.Println(f(m, n))
	}
	return tmp
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
