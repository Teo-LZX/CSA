//第一次作业

//@简易计算器
//@Author LiZhenXin
//@2021-1-31

package main

import("fmt")

func main(){
	var a,b int
	var ch string
	fmt.Println("input:")
	fmt.Scan(&a,&ch,&b)      //输入格式：以空格为分隔符，例如：9 + 9
	switch ch{
	case "+":fmt.Println(a,ch,b,"=",a + b)
	case "-":fmt.Println(a,ch,b,"=",a - b)
	case "*":fmt.Println(a,ch,b,"=",a * b)
	case "/":fmt.Println(a,ch,b,"=",a / b)
	default:
		fmt.Println("error!")

	}
}
