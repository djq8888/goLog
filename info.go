package goLog

import (
	"fmt"
	"runtime"
	"time"
)

func logInfo()  {
	for {
		createTime := time.Now().Format("2006-01-02_15_04_05")
		infoLogName := "log/INFO" + createTime + ".txt"
		infoLog = new(secureFile)
		infoLog.create(infoLogName)
		var timeout <-chan time.Time
		if createInterval > 0 {
			timeout = time.After(createInterval * time.Minute)
		}
		loop:
		for {
			select {
			case buf := <- infoQueue:
				infoLog.write(buf)
				wg.Done()
			case <-timeout:
				break loop
			}
		}
	}
}

func LogInfo(format string, args... interface{})  {
	wg.Add(1)
	buf := fmt.Sprintf(format, args...)
	//获取调用处函数名，文件名，行数
	pc, callFile, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()
	buf = fmt.Sprintf("INFO %s %s %d %s", callFile, function, line, buf)
	infoQueue <- buf
}