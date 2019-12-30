package main

import "fmt"

func main() {
	str1 := "你好北京，hello go on"
	fmt.Println(str1)
	//str[0]='a'不能再赋值
	str2 := "abc\ngogogo"
	fmt.Println(str2)
	//``所见即所得
	str3 := `package main

	import "fmt"

	func main()  {
		str1:="你好北京，hello go on"
		fmt.Println(str1)
		//str[0]='a'不能再赋值
		str2:="abc\ngogogo"
		fmt.Println(str2)
		str3:=""` + str2 + `}""""`

	fmt.Println(str3)
}
