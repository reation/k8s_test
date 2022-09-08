package main

import (
	"fmt"
	"os"
)

func main() {
	file6, error := os.OpenFile("./4.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if error != nil {
		fmt.Println(error)
	}
	data := "我00\r\n"
	for i := 0; i < 10; i++ {
		//写入byte的slice数据
		file6.Write([]byte(data))
		//写入字符串
		file6.WriteString(data)
	}
	file6.Close()
}
