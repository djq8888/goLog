package goLog

import (
	"fmt"
	"os"
	"time"
)

var infoLogName, warnLogName, errorLogName string
var infoQueue, warnQueue, errorQueue chan string
var lock map[string]chan int8

func init()  {
	//创建日志文件夹
	os.Mkdir("./log",0777)
	//创建队列
	infoQueue = make(chan string)
	warnQueue = make(chan string)
	errorQueue = make(chan string)
	//初始化文件锁
	lock = make(map[string]chan int8)
	//开启日志写入协程
	go logInfo()
	go logWarn()
	go logError()
}

func log(filename, buf string)  {
	//添加日期和时间戳（精确到ms）
	buf = fmt.Sprintf("\n%s(%v) %s", time.Now().Format("2006-01-02_15_04_05"), time.Now().UnixNano() / 1e6, buf)
	lock[filename] <- 1
	file, err := os.OpenFile(filename, os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
		<- lock[filename]
	} else {
		return
	}
	file.Write([]byte(buf))
}

func createLog(filename string)  {
	//创建日志文件
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()
	} else {
		return
	}
	//写入初始日志
	initInfo := "logfile is created"
	file.Write([]byte(initInfo))
	//创建文件锁
	lock[filename] = make(chan int8, 1)
}