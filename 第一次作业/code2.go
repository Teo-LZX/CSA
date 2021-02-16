//@成语接龙
//@Author LiZhenXin
//@2021-1-31
package main
import(
	"fmt"
)

func Sort(str []string,ch byte){    //函数功能--排序接龙并输出
	for j:=0;j<len(str)-1;j++{     //外层循环控制输出次数，避免重复
		for i:=0;i<len(str);i++{   //内层循环寻找首字母为ch的单词
			if str[i][0] == ch{
				fmt.Printf("%s  ", str[i])
				ch = str[i][len(str[i])-1]  //ch更改为已输出单词的尾字母
				break
			}
		}
	}
}

func main()  {
	var str []string   //存放输入的字符串
	var s string		//中间变量
	var ch byte = 'c'
	fmt.Println("请输入")
	for {
		_,err := fmt.Scanf("%s", &s)
		str = append(str, s)
		if err != nil{
			break
		}
	}
	fmt.Println("请输入起始字母")
	fmt.Scanf("%c",&ch)
	fmt.Printf("成语接龙结果：")
	Sort(str, ch)		         //排序并打印
}

