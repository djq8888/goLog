package goLog

import (
	"fmt"
	"runtime"
	"time"
)

func logError()  {
	createTime := time.Now().Format("2006-01-02_15_04_05")
	errorLogName = "log/ERROR" + createTime + ".txt"
	createLog(errorLogName)
	for {
		select {
		case buf := <- errorQueue:
			log(infoLogName, buf)
			log(warnLogName, buf)
			log(errorLogName, buf)
		}
	}
}

func LogError(format string, args... interface{})  {
	buf := fmt.Sprintf(format, args...)
	//获取调用处函数名，文件名，行数
	pc, callFile, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()
	buf = fmt.Sprintf("ERROR %s %s %d %s", callFile, function, line, buf)
	errorQueue <- buf
}