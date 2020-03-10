package goLog

import (
	"fmt"
	"runtime"
	"time"
)

func logInfo()  {
	createTime := time.Now().Format("2006-01-02_15_04_05")
	infoLogName = "log/INFO" + createTime + ".txt"
	createLog(infoLogName)
	for {
		select {
		case buf := <- infoQueue:
			log(infoLogName, buf)
		}
	}
}

func LogInfo(format string, args... interface{})  {
	buf := fmt.Sprintf(format, args...)
	//获取调用处函数名，文件名，行数
	pc, callFile, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()
	buf = fmt.Sprintf("INFO %s %s %d %s", callFile, function, line, buf)
	infoQueue <- buf
}