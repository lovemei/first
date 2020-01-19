package main

import "fmt"

func main() {
	var n1 int32 = 10
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20
	n3 = int8(n1 + 20)
	fmt.Println("n2=", n2, "n3=", n3)

	var n4 int8
	var n5 int8

	n4 = int8(n1) + 127 //不提示报错
	n5 = int8(n1 + 128) //这样才不报错
	fmt.Println(n4, n5)
}
