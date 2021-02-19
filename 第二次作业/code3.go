package main

import "fmt"

var count int = 1

type AuSlice []Author
type ViSlice []Video
type Video struct {
	AuthorName string //作者名
	VideoName  string //视频名
	Fabulous   int    //点赞
	Collection int    //收藏
	Coins      int    //投币
	Three      int    //一键三连
	Au         Author //作者
}
type Author struct {
	Name      string //名字
	VIP       bool   //是否是高贵的带会员
	Signature string //签名
	Focus     int    //关注人数
	Vios      int    //视频数量
}

func (this *Video) Fabu() {
	this.Fabulous++
	fmt.Println("点赞次数+1：", this.Fabulous)
}
func (this *Video) Collect() {
	this.Collection++
	fmt.Println("收藏量+1：", this.Collection)
}
func (this *Video) Coin() {
	this.Coins++
	fmt.Println("投币数+1：", this.Coins)
}
func (this *Video) Thr() {
	this.Three++
	fmt.Println("三连成功+1")
}
func Publish(Auth Author, VideoName string) Video { //发布视频
	fmt.Printf("\n~~~~第%d个视频发布成功~~~~\n", count)
	count++
	v := Video{VideoName: VideoName, AuthorName: Auth.Name, Fabulous: 0, Coins: 0, Collection: 0, Three: 0, Au: Author{Auth.Name, Auth.VIP, Auth.Signature, Auth.Focus, Auth.Vios}}
	return v
}

func Menu() { //菜单
	var (
		m    int
		n    int = 1
		flag string
		//AuthorName string
		VideoName string
		Aus       AuSlice //所有用户组成的切片
		Vis       ViSlice //所有视频的组成切片
	)
	Aus = AuSlice{
		Author{"吃花椒的喵酱", true, "国民老婆", 4280000, 8},
		Author{"罗翔", true, "张三老师", 2000000, 50},
		Author{"宋浩老师官方", true, "期末不挂科", 10234546, 123},
	}
	fmt.Printf("~~~~~~~~~~~~~~~~~~欢迎来到bilibili~~~~~~~~~~~~~~~~~~\n发布视频（Y/N）?")
	fmt.Scanln(&flag)
	for flag == "Y" || flag == "y" {
		fmt.Println("请选择作者")
		for i, k := range Aus {
			fmt.Print(i+1, ".", k.Name, "  ")
		}
		fmt.Println()
		fmt.Scanln(&m)
		fmt.Print("请输入要发布的视频名字：")
		fmt.Scanln(&VideoName)
		fmt.Println()
		v := Publish(Aus[m-1], VideoName)
		Vis = append(Vis, v)
		fmt.Print("继续发布（Y/N）?")
		fmt.Scanln(&flag)
	}
	fmt.Printf("\n浏览视频（Y/N）?")
	fmt.Scanln(&flag)
	//浏览视频
	if flag == "Y" || flag == "y" {
		for n != 0 {
			fmt.Println()
			fmt.Println("---------当前所有视频--------")
			if len(Vis) == 0 {
				fmt.Println("当前无视频")
			}
			for i, k := range Vis {
				fmt.Print(i+1, ".", k.VideoName, "  ")
			}
			fmt.Println()
			m = -1
			for m < 0 || m > 6 {
				fmt.Printf("\n请选择视频:")
				fmt.Scanln(&m)
			}
			fmt.Println()
			fmt.Println("~~~~~~~~~~~~~~~~~~~", Vis[m-1].VideoName, "~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("请选择操作：1.点赞  2.收藏  3.投币  4.一键三连  5.查看作者信息  6.查看视频详情  0.退出")
			fmt.Scanln(&n)
			switch n {
			case 1:
				Vis[m-1].Fabu()
			case 2:
				Vis[m-1].Collect()
			case 3:
				Vis[m-1].Coin()
			case 4:
				Vis[m-1].Thr()
			case 5:
				DisplayAuthor(&Aus[m-1])
			case 6:
				DisplayVideo(&Vis[m-1])
			case 0:
				break
			default:
				fmt.Println("输入错误")
			}
		}
	}
}
func DisplayAuthor(author *Author) { //打印作者信息
	fmt.Printf("名字:%s  VIP:%t  签名:%s  关注人数:%d  视频数量:%d\n", author.Name, author.VIP, author.Signature, author.Focus, author.Vios)
}
func DisplayVideo(FirstVideo *Video) { //打印视频信息
	fmt.Printf("视频详情：\n 视频名：%-15s\n 作者名：%-15s\n 点赞数：%-15d\n 收藏数：%-15d\n 投币数：%-15d\n 一键三连：%-13d\n", FirstVideo.VideoName, FirstVideo.AuthorName, FirstVideo.Fabulous, FirstVideo.Collection, FirstVideo.Coins, FirstVideo.Three)
}
func main() {
	Menu()
}
