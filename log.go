package goLog

import (
	"os"
	"time"
)

var filename string
var logString = make(chan string)

func init()  {
	//创建日志文件夹
	os.Mkdir("./log",0777)
	//创建日志文件
	createTime := time.Now().Format("2006-01-02_15_04_05")
	filename = "log/" + createTime + ".txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()
	} else {
		return
	}
	//写入初始日志
	initInfo := "logfile is created at " + createTime
	file.Write([]byte(initInfo))
	//开启日志写入协程
	go func() {
		for {
			select {
			case buf := <- logString:
				log(buf)
			}
		}
	}()
}

func log(buf string)  {
	tmp := "\n" + time.Now().Format("2006-01-02_15_04_05") + " " + buf
	file, err := os.OpenFile(filename, os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
	} else {
		return
	}
	file.Write([]byte(tmp))
}

func Log(buf string)  {
	logString <- buf
}