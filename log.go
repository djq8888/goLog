package goLog

import (
	"os"
	"sync"
	"time"
)

var infoLog, warnLog, errorLog *secureFile
var infoQueue, warnQueue, errorQueue chan string
var wg sync.WaitGroup
var createInterval time.Duration
var onlyInfo bool
var maxFile int

//createinterval：切分日志文件的时间间隔（分钟），如果设为0，则会将运行期间所有日志写入同一文件
//onlyInfo：如果为true，则只创建INFO文件
//以下配置仅在onlyInfo为true时生效
//maxFile：log文件夹中可以保留的最大文件个数，当文件个数到达maxFile时，清除最早创建的文件，如果设为0，则没有限制
func Init(createinterval time.Duration, onlyinfo bool) {
	createInterval = createinterval
	onlyInfo = onlyinfo
	//if onlyInfo {
	//	maxFile = maxfile
	//}
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