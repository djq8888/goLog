package goLog

import (
	"os"
	"sync"
	"time"
)

var (
	infoLog, warnLog, errorLog *secureFile
	infoQueue, warnQueue, errorQueue chan string
	wg sync.WaitGroup
	createInterval time.Duration
	onlyInfo bool
	maxFile, maxLine int
)

//createinterval：切分日志文件的时间间隔（分钟），如果设为0，则会将运行期间所有日志写入同一文件
//onlyinfo：如果为true，则只创建INFO文件
//以下配置仅在onlyInfo为true时生效
//maxfile：log文件夹中可以保留的最大文件个数，当文件个数到达maxFile时，清除最早创建的文件，如果设为0，则没有限制
//maxline：log文件中可以打印的最大行数，当日志行数到达maxLine时，将新建日志文件，如果设为0，则无限制
func Init(createinterval time.Duration, onlyinfo bool, maxfile, maxline int) {
	createInterval = createinterval
	onlyInfo = onlyinfo
	if onlyInfo {
		maxFile = maxfile
		maxLine = maxline + 1
	}
	//创建日志文件夹
	os.Mkdir("./log",0777)
	//创建队列
	infoQueue = make(chan string)
	warnQueue = make(chan string)
	errorQueue = make(chan string)
	//开启日志写入协程
	go logInfo()
	if !onlyInfo {
		go logWarn()
		go logError()
	}
}

func Wait() {
	wg.Wait()
}