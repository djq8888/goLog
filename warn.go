package goLog

import (
	"fmt"
	"runtime"
	"time"
)

func logWarn()  {
	createTime := time.Now().Format("2006-01-02_15_04_05")
	warnLogName = "log/WARN" + createTime + ".txt"
	createLog(warnLogName)
	for {
		select {
		case buf := <- warnQueue:
			log(infoLogName, buf)
			log(warnLogName, buf)
		}
	}
}

func LogWarn(format string, args... interface{})  {
	buf := fmt.Sprintf(format, args...)
	//获取调用处函数名，文件名，行数
	pc, callFile, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()
	buf = fmt.Sprintf("WARN %s %s %d %s", callFile, function, line, buf)
	warnQueue <- buf
}