package main

import "fmt"

func main() {
	/*a, b := 1, "你？"
	c, d := 3, 7
	println(a, b, c, d)

	var s,sep string
	for i :=1;i<len(os.Args);i++{
		s+=sep +os.Args[i]
		sep=" "
	}
	fmt.Println(s)
	*/
	/*x:=1
	p:=&x
	fmt.Println(p,*p)
	*p=2
	fmt.Println(x)*/

	/*var x3, y3 int
	fmt.Println(&x3 == &x3, &x3 == &y3, &x3 == nil)
	*/
	var p4 = f()
	fmt.Println(p4, *p4)
	x5 := 1
	fmt.Println(padd(&x5))
}
func f() *int {
	x4 := 1
	return &x4
}
func padd(v *int) int {
	*v++
	return *v
}
