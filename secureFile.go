package goLog

import (
	"fmt"
	"os"
	"time"
)

type secureFile struct {
	filename string
	lock chan int8
}

func (f *secureFile) create(filename string) {
	f.lock = make(chan int8, 1)
	f.filename = filename
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()
	} else {
		return
	}
	//写入初始日志
	initInfo := "logfile is created"
	file.Write([]byte(initInfo))
}

func (f *secureFile) write(buf string) {
	f.lock <- 1
	//添加日期和时间戳（精确到ms）
	buf = fmt.Sprintf("\n%s(%v) %s", time.Now().Format("2006-01-02_15_04_05"), time.Now().UnixNano() / 1e6, buf)
	file, err := os.OpenFile(f.filename, os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
	} else {
		<- f.lock
		return
	}
	file.Write([]byte(buf))
	<- f.lock
}