package main

import "fmt"

//包内部变量
var aa = 3
var (
	bb = true
	cc = "kkk"
)

// 函数的外面定义变量不可以省略var
//bb := true

func varibleZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varibleInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func varibleTypeDefuction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func varibleShorter() {
	// := 就是定义一个变量，最好使用这种形式
	a, b, c, s := 3, 5, true, "def"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello world")
	varibleZeroValue()
	varibleInitialValue()
	varibleTypeDefuction()
	varibleShorter()
	fmt.Println(aa, bb, cc)
}
