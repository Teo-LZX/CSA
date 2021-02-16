//@迷宫
//@Author liZhenXin
//@2021-2-4

//注：输入数据时请严格按照格式输入
/*格式例：
请输入M（行），N（列），T（障碍物总数）5 5 4
请输入起点：0 0
请输入终点：4 4
请输入障碍物坐标：1 0 1 1 1 2 1 3
*/
package main

import(
	"fmt"
)
const M, N = 20, 20      //定义地图最大长度,宽度

func StartMap(Map *[M][N]int, x0, y0, x1, y1 *int, m, n, t *int){   //函数功能---初始化地图
	//0表示没走过的路
	//1表示障碍物
	//3表示走过但是走不通的路
	var x, y int      //障碍物的中间变量
	fmt.Printf("请输入M（行），N（列），T（障碍物总数）")
	fmt.Scanf("%d %d %d", m, n, t)
	fmt.Scanln()      //把缓冲区的回车键吸收掉
	fmt.Printf("请输入起点：")
	fmt.Scanln(x0, y0)
	fmt.Printf("请输入终点：")
	fmt.Scanln(x1, y1)
	fmt.Printf("请输入障碍物坐标：")
	for i:=0;i < *t;i++{
		fmt.Scanf("%d %d", &x, &y)
		Map[x][y] = 1
	}
	//下面两个循环进行地图初始化
	for i:=0;i <= *n;i++{
		Map[*m][i] = 1
	}
	for j:=0;j <= *m;j++{
		Map[j][*n] = 1
	}

	fmt.Println("初始地图：")
	for i:=0;i <= *m;i++{
		for j:=0;j <= *n;j++{
			fmt.Printf("%d  ", Map[i][j])
		}
		fmt.Println()
	}
}

func FindWay(Map *[M][N]int, x0, y0, x1, y1, m, n int) bool {   //函数功能---寻找路线
	if x0>=0 && y0>=0 && x0<m && y0<n{
		if Map[x1][y1] == 2{
			return true
		}else{
			if Map[x0][y0] == 0{
				Map[x0][y0] = 2
				if FindWay(Map, x0 + 1, y0, x1, y1, m, n){       //先往下走一格
					return true
				}else if FindWay(Map, x0, y0 + 1, x1, y1, m, n){  //如果往下不行，就退回来往右一格
					return true
				}else if FindWay(Map, x0 - 1, y0, x1, y1, m, n){  //往右不行就退回来往上走一格
					return true
				}else if FindWay(Map, x0, y0 - 1, x1, y1, m, n){  //往左不行就退回来往左走一格
					return true
				}else{
					Map[x0][y0] = 3
					return false
				}
			}else{
				return false
			}
		}
	}

	return false
}
func OutPut(Map *[M][N]int, x0, y0, x1, y1, m, n int) {
	//将刚才地图上为2的路线打印出来
	//每打印一次，就将该处的数字改为3，避免死循环
	Map[x0][y0] = 3               //把起点打印出来并改为3
	fmt.Println(x0, "  ", y0)

	for x0!=x1 || y0!=y1{         //到达终点结束循环
		//打印路径--下右左上
		if Map[x0+1][y0] == 2{
			x0++
			fmt.Println(x0,"  ",y0)
			Map[x0][y0] = 3
		}else if Map[x0][y0+1] == 2{
			y0++
			fmt.Println(x0,"  ",y0)
			Map[x0][y0] = 3
		}else if x0>0 && Map[x0-1][y0] == 2{
			x0--
			fmt.Println(x0,"  ",y0)
			Map[x0][y0] = 3
		}else if y0>0 && Map[x0][y0-1] == 2{
			y0--
			fmt.Println(x0,"  ",y0)
			Map[x0][y0] = 3
		}else{
			fmt.Println("    ")
		}
	}
}

func main() {
	var Map [M][N]int       //地图
	var m,n,t int        //m行n列的迷宫，t个障碍物
	var x0, y0, x1, y1 int //分别为起点，终点
	StartMap(&Map, &x0, &y0, &x1, &y1, &m, &n, &t)		//初始化地图
	err := FindWay(&Map, x0, y0, x1, y1, m, n)        //用err返回值判断迷宫是否可走出
	if err == false{
		fmt.Println("无法走出！程序终止")
	}else{
		fmt.Println("结果地图：")
		for i:=0;i <= m;i++{
			for j:=0;j <= n;j++{
				fmt.Printf("%d  ", Map[i][j])
			}
			fmt.Println()
		}
		fmt.Println("行走路线：")
		OutPut(&Map, x0, y0, x1, y1, m, n)
	}

}

