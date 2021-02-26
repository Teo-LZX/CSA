package main

import (
	"fmt"
	"math"
	"time"
)

func WriteChan(intchan chan int) { //将五万个函数写入管道
	for i := 0; i < 50000; i++ {
		intchan <- i + 1
	}
	close(intchan)
}

func ReadChan(intchan chan int, primechan chan int, exitchan chan bool) { //从管道中读取数据并将素数放至primechan管道
	var flag bool
	for {
		num, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
		}
	}
	fmt.Println("读取完成")
	exitchan <- true
}

func main() {
	time1 := time.Now()
	var intchan chan int = make(chan int, 50000)   //写入数据的管道
	var primechan chan int = make(chan int, 30000) //存放数据的管道
	var exitchan chan bool = make(chan bool, 12)   //判断是否退出
	go WriteChan(intchan)                          //写入数据
	for i := 0; i < 12; i++ {
		go ReadChan(intchan, primechan, exitchan) //读取数据并将素数放入primechan中
	}
	go func() {
		for i := 0; i < 12; i++ {
			<-exitchan
		}
		close(primechan) //读取完毕后关闭管道
	}()
	for v := range primechan {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("运行时间：%v", time.Now().Sub(time1))
}

//func ReadChan(intchan chan int, primechan chan int) {
//	for {
//		num, ok := <-intchan
//		if !ok {
//			break
//		}
//		if IsPrime(num) {
//			primechan <- num
//		}
//	}
//	close(primechan)
//	fmt.Println("完毕")
//}
//func main() {
//	time1 := time.Now()
//	intchan := make(chan int, 50000) //写入数据的管道
//	primechan := make(chan int, 30000)
//	WriteChan(intchan)
//	ReadChan(intchan, primechan)
//	for v := range primechan {
//		fmt.Println(v)
//	}
//	fmt.Println("运行时间：", time.Now().Sub(time1))
//}
