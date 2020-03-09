package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func logInfo()  {
	
}

func init()  {
	//创建日志文件夹
	os.Mkdir("./log",0777)
	//创建日志文件
	createTime := time.Now().Format("2006-01-02_15_04_05")
	filename := "log/" + createTime + ".txt"
	os.Create(filename)
	//写入初始日志
	initInfo := "logfile is created at " + createTime
	ioutil.WriteFile(filename, []byte(initInfo), 0666)
}

func main()  {
	//初始化
	fmt.Println(1)
}