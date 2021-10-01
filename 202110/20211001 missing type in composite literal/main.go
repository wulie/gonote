package main

import "fmt"

/*

上面代码在编辑器中会提示这样的错误：Reports composite literals with incompatible types and values。原因在于定义的Person结构体的Favourite字段的类型是字符串切片，但是在main函数中使用时，直接传入了一个结构体对象。两者的类型不相同，所以提示不兼容的类型。

对于一些复合数据类型，在使用时应该声明其类型信息，不然Go无法自动判别。因此上面的错误示例只需要改变person初始化时的类型声明即可。
————————————————
版权声明：本文为CSDN博主「benben_2015」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/benben_2015/article/details/95166239

*/
type Person struct {
	Name      string
	Age       int
	Favourite []string
}

func main() {
	person := Person{
		"benben_2015",
		18,
		[]string{"huoguo", "chuanchuan"}, //正确的格式
		//{"huoguo", "chuanchuan"},//错误的格式
	}
	fmt.Println(person)
}

//output

/*
	错误的输出😠
	GOROOT=/usr/local/go #gosetup
	GOPATH=/Users/wangxuhui/go #gosetup
	/usr/local/go/bin/go build -o /private/var/folders/g5/hb8g81c579lb0kj0szbz7dvw0000gn/T/GoLand/___go_build_github_com_wulie_gonote_202110_20211001 github.com/wulie/gonote/202110/20211001 missing type in composite literal #gosetup
	# github.com/wulie/gonote/202110/20211001 missing type in composite literal
	./main.go:17:3: missing type in composite literal

**********************************************************

	正确的输出😊
	GOROOT=/usr/local/go #gosetup
	GOPATH=/Users/wangxuhui/go #gosetup
	/usr/local/go/bin/go build -o /private/var/folders/g5/hb8g81c579lb0kj0szbz7dvw0000gn/T/GoLand/___go_build_github_com_wulie_gonote_202110_20211001 github.com/wulie/gonote/202110/20211001 missing type in composite literal #gosetup
	/private/var/folders/g5/hb8g81c579lb0kj0szbz7dvw0000gn/T/GoLand/___go_build_github_com_wulie_gonote_202110_20211001
	{benben_2015 18 [huoguo chuanchuan]}

*/
