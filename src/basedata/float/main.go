package main

import "fmt"

func main() {

	var f1 float32 = 123.0000901
	var f2 float64 = 123.0000901
	fmt.Println(f1, f2)
	f3 := 4.5
	fmt.Printf("f3的数据类型是%T", f3)
	//基本数据类型的相互转换
	var i int32 = 300
	var f4 float32 = float32(i)
	var f5 float64 = float64(f4)
	var i8 int8 = int8(i)
	fmt.Println(i, f4, f5, i8)
}
