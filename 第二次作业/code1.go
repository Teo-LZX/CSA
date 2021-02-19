package main

import "fmt"

type person struct { //人
	Name   string
	Age    int
	Gender string
}
type dove interface { //鸽子
	gugugu()
}
type repeater interface { //复读机
	repeat(words string)
}
type lemon interface { //柠檬精
	say()
}
type delicious interface { //真香怪
	zhenxiang(object string)
}

func (p *person) gugugu() {
	fmt.Println(p.Name, "又鸽了")
}
func (p *person) repeat(words string) {
	fmt.Println(p.Name, ": ", words)
}
func (p *person) say() {
	fmt.Println(p.Name, ": 我酸了")
}
func (p *person) zhenxiang(object string) {
	fmt.Println(object, "真香!")
}

func main() {
	per := &person{"linlin", 19, "girl"}
	var a dove = per
	var b repeater = per
	var c lemon = per
	var d delicious = per
	a.gugugu()
	b.repeat("我爱你")
	c.say()
	d.zhenxiang("goland")
}
