package main

import "fmt"

func Receiver(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("这个是string")
	case int:
		fmt.Println("这个是int")
	case bool:
		fmt.Println("这个是bool")
	case float64:
		fmt.Println("这个是float64")
	case float32:
		fmt.Println("这个是float32")
	default:
		fmt.Println("未知类型")

	}
}
func main() {
	Receiver("hello")
	Receiver(false)
	Receiver(23.3)
	Receiver(1)
}
