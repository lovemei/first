package main

import "fmt"

func main() {
	str1 := "你好北京，hello go on"
	fmt.Println(str1)
	//str[0]='a'不能再赋值
	str2 := "abc\ngogogo"
	fmt.Println(str2)
	//``所见即所得
	str3 := `
	func main()  {
		str1:="你好北京，hello go on"
		fmt.Println(str1)
		//str[0]='a'不能再赋值
		str2:="abc\ngogogo"
		fmt.Println(str2)
		str3:=""` + str2 + `}""""`

	fmt.Println(str3)

	//基础数据类型和string类型的转换
	var num1 int = 99
	var num2 float64 = 23.456

	fmt.Println(num1, num2)

}
